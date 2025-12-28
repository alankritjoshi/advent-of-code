import argparse
from collections import Counter


def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i", "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            stones: Counter[int] = Counter([int(ch) for ch in line.strip().split()])

            for _ in range(75):
                current: Counter[int] = Counter()

                for stone, count in stones.items():
                    if stone == 0:
                        current[1] += count
                    elif len(str(stone)) % 2 == 0:
                        str_val = str(stone)
                        first, second = str_val[: len(str_val) // 2], str_val[len(str_val) // 2 :]
                        current[int(first)] += count
                        current[int(second)] += count
                    else:
                        current[stone * 2024] += count

                stones = current

            total = 0
            for count in stones.values():
                total += count

            print(total)


if __name__ == "__main__":
    main()
