import argparse

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input
    
    grid: list[list[str]] = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            grid.append([ch for ch in line.strip()])

    start: tuple[int, int] = (-1, -1)
    
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == "^":
                start = (i, j)


    DEFAULT_DIRECTION = (-1, 0)

    def traverse(x: int, y: int, direction: tuple[int, int]) -> int:
        seen = set()
        
        while True:
            seen.add((x, y))

            ahead = (x + direction[0], y + direction[1])

            # Check bounds. If out, it means patrol is over
            if not (0 <= ahead[0] < len(grid) and 0 <= ahead[1] < len(grid[0])):
                break

            # Change direction if obstacle ahead
            if grid[ahead[0]][ahead[1]] == "#":
                direction = (direction[1], -direction[0]) # Turn 90 deg to the right

            # Move forward
            x += direction[0]
            y += direction[1]

        return len(seen)


    print(traverse(*start, DEFAULT_DIRECTION))

if __name__ == '__main__':
    main()

