import argparse
import re
from collections import deque

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    pattern = r"^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$"

    HEIGHT = 103
    WIDTH = 101

    grid: list[list[deque[tuple[int, int, int]]]] = [[deque([]) for _ in range(WIDTH)] for _ in range(HEIGHT)]

    def move(bound: int, position: int, velocity: int) -> int:
        return (position + velocity) % bound

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            line = line.strip()

            nums: list[int] = [int(num) for num in re.findall(pattern, line)[0]]

            x = nums[1]
            y = nums[0]
            velocity: tuple[int, int] = (nums[3], nums[2])

            grid[x][y].append((-1, *velocity))

    for iteration in range(100):
        for x in range(len(grid)):
            for y in range(len(grid[0])):
                while grid[x][y] and grid[x][y][0][0] < iteration:
                    _, vel_x, vel_y = grid[x][y].popleft()
                    new_x, new_y = move(len(grid), x, vel_x), move(len(grid[0]), y, vel_y)
                    grid[new_x][new_y].append((iteration, vel_x, vel_y))


    rows, cols = len(grid), len(grid[0])
    central_row, central_col = rows // 2, cols // 2

    quadrants = [
        (0, central_row, 0, central_col), # Top left
        (0, central_row, central_col + 1, cols), # Top right
        (central_row + 1, rows, 0, central_col), # Bottom left
        (central_row + 1, rows, central_col + 1, cols), # Bottom right
    ]

    total = 1
    for row_start, row_end, col_start, col_end in quadrants:
        quad_total = 0
        for x in range(row_start, row_end):
            for y in range(col_start, col_end):
                quad_total += len(grid[x][y])
        total *= quad_total

    print(total)


if __name__ == '__main__':
    main()

