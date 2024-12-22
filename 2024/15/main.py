import argparse
from enum import Enum

class Direction(Enum):
    UP = (-1, 0)
    DOWN = (1, 0)
    LEFT = (0, -1)
    RIGHT = (0, 1)

    @classmethod
    def from_symbol(cls, sym: str) -> 'Direction':
        if sym == '^':
            return cls.UP
        elif sym == 'v':
            return cls.DOWN
        elif sym == '<':
            return cls.LEFT
        elif sym == '>':
            return cls.RIGHT
        
        raise ValueError("Invalid symbol: ", sym)

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    grid: list[list[str]] = []

    grid_processed: bool = False

    directions: list[Direction] = []

    robot: tuple[int, int] = (-1, -1)

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            line = line.strip()

            if not line:
                grid_processed = True
                continue

            if not grid_processed:
                grid.append(list(line))
            else:
                directions.extend([Direction.from_symbol(ch) for ch in line])

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == '@':
                robot = (i, j)

    def move(x: int, y: int, direction: Direction) -> bool:
        xx, yy = direction.value

        new_x, new_y = x + xx, y + yy

        if not (0 <= new_x < len(grid) and 0 <= new_y < len(grid[0])):
                return False

        if grid[new_x][new_y] == '#':
            return False

        if grid[new_x][new_y] == '.' or move(new_x, new_y, direction):
            grid[new_x][new_y], grid[x][y] = grid[x][y], grid[new_x][new_y]
            return True

        return False

    for direction in directions:
        if move(*robot, direction):
            robot = (robot[0] + direction.value[0], robot[1] + direction.value[1])

    total = 0
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == 'O':
                total += (100 * i + j)

    print(total)


if __name__ == '__main__':
    main()

