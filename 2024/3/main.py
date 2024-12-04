import argparse
import re

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    pattern = r"mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)"

    total = 0
    enabled = True

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            expressions = re.findall(pattern, line)

            for expr in expressions:
                int_pattern = r"\d+"
                ints = re.findall(int_pattern, expr)
                if ints and enabled:
                    total += int(ints[0]) * int(ints[1])
                else:
                    if expr == "do()":
                        enabled = True
                    elif expr == "don't()":
                        enabled = False

    print(total)

if __name__ == '__main__':
    main()

