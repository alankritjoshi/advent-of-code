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

            grid.append(list(line.strip()))

    start: tuple[int, int] = (-1, -1)

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == "^":
                start = (i, j)


    DEFAULT_DIRECTION = (-1, 0)

    def has_loop(x: int, y: int, direction: tuple[int, int]) -> int:
        seen: set[tuple[int, int, int, int]] = set()

        while True:
            if (x, y, *direction) in seen:
                return True

            seen.add((x, y, *direction))

            ahead = (x + direction[0], y + direction[1])

            # Check bounds. If out, it means patrol is over
            if not (0 <= ahead[0] < len(grid) and 0 <= ahead[1] < len(grid[0])):
                break

            # Change direction if obstacle ahead
            if grid[ahead[0]][ahead[1]] == "#":
                direction = (direction[1], -direction[0]) # Turn 90 deg to the right
                # This is important, you don't want to walk because your new direction
                #   might be pointing to another obstacle
                # Test with
                # ..#.
                # ...#
                # ..^.
                continue

            # Move forward
            x += direction[0]
            y += direction[1]

        return False

    def traverse(x: int, y: int, direction: tuple[int, int]) -> int:
        possible_loop_obstructions: set[tuple[int, int]] = set()

        while True:
            ahead = (x + direction[0], y + direction[1])

            # Check bounds. If out, it means patrol is over
            if not (0 <= ahead[0] < len(grid) and 0 <= ahead[1] < len(grid[0])):
                break

            # Change direction if obstacle ahead
            if grid[ahead[0]][ahead[1]] == "#":
                direction = (direction[1], -direction[0]) # Turn 90 deg to the right
                # This is important, you don't want to walk because your new direction
                #   might be pointing to another obstacle
                # Test with
                # ..#.
                # ...#
                # ..^.
                continue

            if ahead != start and ahead not in possible_loop_obstructions:
                # Place temporary obstruction
                grid[ahead[0]][ahead[1]] = "#"
                if has_loop(*start, DEFAULT_DIRECTION): # Check if the new obstruction causes a loop
                    possible_loop_obstructions.add(ahead)
                # Remove temporary obstruction
                grid[ahead[0]][ahead[1]] = "."

            # Move forward
            x += direction[0]
            y += direction[1]

        return len(possible_loop_obstructions)

    print(traverse(*start, DEFAULT_DIRECTION))

if __name__ == '__main__':
    main()

