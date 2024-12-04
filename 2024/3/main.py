import argparse
import re

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    pattern = r"mul\(\d{1,3},\d{1,3}\)"

    total = 0

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            expressions = re.findall(pattern, line)

            for expr in expressions:
                int_pattern = r"\d+"
                ints = re.findall(int_pattern, expr)
                total += int(ints[0]) * int(ints[1])

    print(total)

if __name__ == '__main__':
    main()

