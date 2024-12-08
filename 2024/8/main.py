import argparse
import itertools
import math
from collections import defaultdict

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
    
    def find_antinodes(a: tuple[int, int], b: tuple[int, int]) -> list[tuple[int, int]]:
        first = (2*a[0] - b[0], 2*a[1] - b[1])
        second = (2*b[0] - a[0], 2*b[1] - a[1])
        return [antinode for antinode in (first, second) if validate_bound(*antinode)]

    for positions in frequencies.values():
        for a, b in itertools.combinations(positions, 2):
            antinodes.update(find_antinodes(a, b))

    print(len(antinodes))


if __name__ == '__main__':
    main()

