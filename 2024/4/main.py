import argparse

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    grid = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            grid.append([ch for ch in line.split()[0]])

    def find_xmas_counts(i: int, j: int, curr: int, direction: tuple[int, int]) -> int:
        if grid[i][j] != 'XMAS'[curr]:
            return 0

        if curr == 3: # S in XMAS
            return 1

        ii, jj = i + direction[0], j + direction[1]

        if not (0 <= ii < len(grid) and 0 <= jj < len(grid[0])):
            return 0

        return find_xmas_counts(ii, jj, curr + 1, direction)

    total = 0

    directions = [
        (1, 0), (-1, 0), # vertical
        (0, 1), (0, -1), # horizontal
        (1, 1), (-1, -1), # diagonal 1
        (1, -1), (-1, 1), # diagonal 2
    ]
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == 'X':
                for direction in directions:
                    total += find_xmas_counts(i, j, 0, direction)

    print(total)

if __name__ == '__main__':
    main()

