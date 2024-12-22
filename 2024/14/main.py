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

    for iteration in range(1000000):
        for x in range(len(grid)):
            for y in range(len(grid[0])):
                while grid[x][y] and grid[x][y][0][0] < iteration:
                    _, vel_x, vel_y = grid[x][y].popleft()
                    new_x, new_y = move(len(grid), x, vel_x), move(len(grid[0]), y, vel_y)
                    grid[new_x][new_y].append((iteration, vel_x, vel_y))


        # Print the grid when 10 or more consecutive robots exist in a row and break
        p: list[list[int]] = []
        show = False
        for i in range(len(grid)):
            l: list[int] = []
            prev: bool = False
            m: int = -1
            consec: int = 0
            for j in range(len(grid[0])):
                l.append(len(grid[i][j]))
                if prev:
                    if len(grid[i][j]) > 0:
                        consec = consec + 1
                    else:
                        m = max(m, consec)
                        consec = 0
                        prev = False
                else:
                    if len(grid[i][j]) > 0:
                        consec = 1
                        prev = True

            if m >= 10:
                show = True

            p.append(l)

        if show:
            print(iteration+1)
            for l in p:
                print("".join(str(num) if num else "." for num in l))
            break


if __name__ == '__main__':
    main()

