import argparse

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    grid: list[list[int | str]] = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            grid.append([int(ch) if ch.isnumeric() else ch for ch in line.strip()])

    def climb(curr_x: int, curr_y: int, visited: set[tuple[int, int]]) -> int:
        curr_val = grid[curr_x][curr_y] 
        assert type(curr_val) is int

        if curr_val == 9:
            return 1

        summits = 0
        for ii, jj in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            neigh_x, neigh_y = curr_x + ii, curr_y + jj
            if (
                not (0 <= neigh_x < len(grid) and 0 <= neigh_y < len(grid[0]))
                or (neigh_x, neigh_y) in visited
                or type(grid[neigh_x][neigh_y]) is not int
                or grid[neigh_x][neigh_y] != curr_val + 1
            ):
                continue
            visited.add((neigh_x, neigh_y))
            summits += climb(neigh_x, neigh_y, visited)

        return summits

    total = 0
    for x in range(len(grid)):
        for y in range(len(grid)):
            if grid[x][y] == 0:
                total += climb(x, y, set([(x, y)]))

    print(total)

if __name__ == '__main__':
    main()

