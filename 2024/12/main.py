import argparse
from collections import defaultdict
from enum import Enum


class Direction(Enum):
    UP = "up"
    DOWN = "down"
    LEFT = "left"
    RIGHT = "right"

    @property
    def is_vertical(self):
        return self in {Direction.LEFT, Direction.RIGHT}

    @property
    def is_horizontal(self):
        return self in {Direction.UP, Direction.DOWN}


Sides = defaultdict[tuple[Direction, int, int], set[int]] 

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    grid: list[list[str | None]] = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            grid.append(list(line.strip()))


    def search(
        curr_x: int, 
        curr_y: int, 
        char: str, 
        sides: Sides,
        area: int, 
        visited: set[tuple[int, int]]
    ) -> tuple[Sides, int]:
        grid[curr_x][curr_y] = "*"
        visited.add((curr_x, curr_y))

        for i, j, direction in (
            (1, 0, Direction.DOWN), 
            (-1, 0, Direction.UP),
            (0, 1, Direction.RIGHT),
            (0, -1, Direction.LEFT),
        ):
            neigh_x, neigh_y = curr_x + i, curr_y + j
            if not (0 <= neigh_x < len(grid) and 0 <= neigh_y < len(grid[0])):
                sides[
                    (
                        direction,
                        neigh_x if direction.is_horizontal else 0, 
                        neigh_y if direction.is_vertical else 0,
                    )
                ].add(neigh_y if direction.is_horizontal else neigh_x)
            elif (neigh_x, neigh_y) in visited:
                    continue
            elif grid[neigh_x][neigh_y] == char:
                sides, curr_area = search(neigh_x, neigh_y, char, sides, 0, visited)
                area += curr_area
            else:
                sides[
                    (
                        direction,
                        neigh_x if direction.is_horizontal else 0, 
                        neigh_y if direction.is_vertical else 0,
                    )
                ].add(neigh_y if direction.is_horizontal else neigh_x)

        grid[curr_x][curr_y] = None

        return sides, area + 1

    total = 0
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            symbol = grid[i][j]
            if symbol is None:
                continue

            s: Sides = defaultdict(set)
            s, area = search(i, j, symbol, s, 0, set())
            sides_count = 0

            for edges in s.values():
                edges = sorted(edges)
                count = 0
                if len(edges) > 1:
                    count = 1
                    for x in range(1, len(edges)):
                        if edges[x-1] + 1 != edges[x]:
                            count += 1
                else:
                    count = len(edges)

                sides_count += count

            total += (sides_count * area)

    print(total)


if __name__ == '__main__':
    main()

