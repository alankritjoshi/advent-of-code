import argparse
from enum import Enum, auto


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

    class Op(Enum):
        PLUS = auto()
        MULTI = auto()

    def calc(index: int, so_far: int | None, last_op: Op | None, verify: int, vals: list[int]) -> bool:
        if index == len(vals):
            return so_far is not None and so_far == verify

        multi_so_far = so_far if so_far is not None else 1
        if multi_so_far * vals[index] <= verify and calc(index+1, multi_so_far * vals[index], Op.MULTI, verify, vals):
            return True

        plus_so_far = so_far if so_far is not None else 0
        if plus_so_far + vals[index] <= verify and calc(index+1, plus_so_far + vals[index], Op.PLUS, verify, vals):
            return True

        if index > 0:
            assert last_op is not None
            assert so_far is not None
            curr = int(str(so_far) + str(vals[index]))
            if last_op is Op.MULTI:
                return calc(index+1, curr, last_op, verify, vals)
            if last_op is Op.PLUS:
                return calc(index+1, curr, last_op, verify, vals)

        return False

    total = 0
    for verify, vals in tests:
        if calc(0, None, None, verify, vals):
            total += verify

    print(total)

if __name__ == '__main__':
    main()

