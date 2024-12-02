import argparse
from collections import Counter

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    a, b = [],[]

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            first, second = line.split()
            a.append(int(first))
            b.append(int(second))

    b_counter = Counter(b)

    total = 0
    for num in a:
        if num in b_counter:
            total += num * b_counter[num]

    print(total)

if __name__ == '__main__':
    main()

