import argparse
import itertools
from collections import defaultdict
from enum import Enum


def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    rows = 0
    cols = 0

    frequencies: dict[str, list[tuple[int, int]]] = defaultdict(list)
    antinodes: set[tuple[int, int]] = set()

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            cols = len(line.strip())

            for col, ch in enumerate(line.strip()):
                if ch.isalnum():
                    frequencies[ch].append((rows, col))

            rows += 1

    def validate_bound(x: int, y: int) -> bool:
        return 0 <= x < rows and 0 <= y < cols

    class Direction(Enum):
        POSITIVE = 1
        NEGATIVE = -1

    def find_antinodes(a: tuple[int, int], b: tuple[int, int]) -> set[tuple[int, int]]:
        """Apply general formula in both positive and negative directions"""

        def find_antinode(x: tuple[int, int], y: tuple[int, int], factor: int) -> tuple[int, int]:
            """General formula"""

            return (x[0] + factor * (x[0] - y[0]), x[1] + factor * (x[1] - y[1]))

        def find_antinodes_in_direction(a, b, direction: Direction) -> set[tuple[int, int]]:
            """Apply general formula in just one direction until out of bounds"""

            antis: set[tuple[int, int]] = set()
            factor = 0
            while True:
                antinode = find_antinode(a, b, factor)
                if not validate_bound(*antinode):
                    break
                antis.add(antinode)
                factor += direction.value
            return antis

        return find_antinodes_in_direction(a, b, Direction.POSITIVE) \
               | find_antinodes_in_direction(a, b, Direction.NEGATIVE)

    for positions in frequencies.values():
        for a, b in itertools.combinations(positions, 2):
            antinodes.update(find_antinodes(a, b))

    print(len(antinodes))


if __name__ == '__main__':
    main()

