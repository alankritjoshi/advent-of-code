import argparse
from enum import Enum, auto


def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i", "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    grid: list[list[int | str]] = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            grid.append([int(ch) if ch.isnumeric() else ch for ch in line.strip()])

    class Direction(Enum):
        LEFT = auto()
        RIGHT = auto()
        UP = auto()
        DOWN = auto()

    def climb(
        curr_x: int,
        curr_y: int,
        path: list[Direction],
        visited: set[tuple[int, int]],
        unique: list[list[Direction]],
    ) -> list[list[Direction]]:
        curr_val = grid[curr_x][curr_y]
        assert type(curr_val) is int

        if curr_val == 9:
            return unique + [path]

        new_unique = []
        for ii, jj, direction in [
            (0, 1, Direction.DOWN),
            (0, -1, Direction.UP),
            (1, 0, Direction.RIGHT),
            (-1, 0, Direction.LEFT),
        ]:
            neigh_x, neigh_y = curr_x + ii, curr_y + jj
            if (
                not (0 <= neigh_x < len(grid) and 0 <= neigh_y < len(grid[0]))
                or (neigh_x, neigh_y) in visited
                or type(grid[neigh_x][neigh_y]) is not int
                or grid[neigh_x][neigh_y] != curr_val + 1
            ):
                continue

            visited.add((neigh_x, neigh_y))
            new_unique.extend(
                climb(
                    neigh_x,
                    neigh_y,
                    path + [direction],
                    visited,
                    unique,
                )
            )
            visited.remove((neigh_x, neigh_y))

        return unique + new_unique

    total = 0
    for x in range(len(grid)):
        for y in range(len(grid)):
            if grid[x][y] == 0:
                paths = climb(x, y, [], {(x, y)}, [])
                unique: set[tuple[Direction, ...]] = set()
                for path in paths:
                    unique.add(tuple(path))
                total += len(unique)

    print(total)


if __name__ == "__main__":
    main()
