import argparse
from types import GeneratorType

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

            grid.append([ch for ch in line.strip()])

    def search(curr_x: int, curr_y: int, char: str, perimeter: int, area: int, visited: set[tuple[int, int]]) -> tuple[int, int]:
        grid[curr_x][curr_y] = "*"
        visited.add((curr_x, curr_y))

        for i, j in ((1, 0), (-1, 0), (0, 1), (0, -1)):
            neigh_x, neigh_y = curr_x + i, curr_y + j
            if not (0 <= neigh_x < len(grid) and 0 <= neigh_y < len(grid[0])):
                perimeter += 1
            elif (neigh_x, neigh_y) in visited:
                    continue
            elif grid[neigh_x][neigh_y] == char:
                curr_perimeter, curr_area = search(neigh_x, neigh_y, char, 0, 0, visited)
                perimeter += curr_perimeter
                area += curr_area
            else:
                perimeter += 1

        grid[curr_x][curr_y] = None

        return perimeter , area + 1

    total = 0
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            symbol = grid[i][j]
            if symbol is None:
                continue

            perimeter, area = search(i, j, symbol, 0, 0, set())

            total += (perimeter * area)

    print(total)


if __name__ == '__main__':
    main()

