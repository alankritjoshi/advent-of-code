import argparse

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    tests: list[tuple[int, list[int]]] = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            verify_str, vals_str = line.strip().split(":")
            verify = int(verify_str)
            vals = [int(val) for val in vals_str.strip().split()]
            tests.append((verify, vals))

    def calc(index: int, so_far: int | None, verify: int, vals: list[int]) -> bool:
        if index == len(vals):
            return so_far is not None and so_far == verify

        multi_so_far = so_far if so_far is not None else 1
        if multi_so_far * vals[index] <= verify and calc(index+1, multi_so_far * vals[index], verify, vals):
            return True

        plus_so_far = so_far if so_far is not None else 0
        if plus_so_far + vals[index] <= verify and calc(index+1, plus_so_far+ vals[index], verify, vals):
            return True

        return False

    total = 0
    for verify, vals in tests:
        if calc(0, None, verify, vals):
            total += verify

    print(total)

if __name__ == '__main__':
    main()

