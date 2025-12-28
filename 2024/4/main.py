import argparse
import itertools


def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i", "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    grid = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            grid.append(list(line.split()[0]))

    def find_xmas_counts(
        i: int,
        j: int,
        curr: int,
        direction: tuple[int, int],
        answer: list[tuple[int, int]],
    ) -> list[tuple[int, int]]:
        if grid[i][j] != "MAS"[curr]:
            return []

        if curr == 2:  # S in XMAS
            return answer

        ii, jj = i + direction[0], j + direction[1]

        if not (0 <= ii < len(grid) and 0 <= jj < len(grid[0])):
            return []

        return find_xmas_counts(ii, jj, curr + 1, direction, answer + [(ii, jj)])

    answers = []

    # only find diagonals
    directions = [
        (1, 1),
        (-1, -1),  # diagonal 1
        (1, -1),
        (-1, 1),  # diagonal 2
    ]

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            # instead of XMAS, find MAS and retrieve their coordinates
            if grid[i][j] == "M":
                for direction in directions:
                    answer = find_xmas_counts(i, j, 0, direction, [(i, j)])
                    if answer:
                        answers.append(answer)

    # find all pair combinations of MAS found
    mas_pairs = 0
    for mas_a, mas_b in itertools.combinations(answers, 2):
        if mas_a[1] == mas_b[1]:  # middle needs to be A to get a pair which makes an X
            mas_pairs += 1

    print(mas_pairs)


if __name__ == "__main__":
    main()
