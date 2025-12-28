import argparse
import re


def solve_linear_equations(eq1: list[int], eq2: list[int]) -> tuple[int, int] | None:
    assert len(eq1) == 3
    assert len(eq2) == 3

    a1, b1, c1 = eq1
    a2, b2, c2 = eq2

    determinant = a1 * b2 - a2 * b1

    if determinant == 0:
        return None

    x = (c1 * b2 - c2 * b1) / determinant
    y = (a1 * c2 - a2 * c1) / determinant

    if not (x >= 0 and y >= 0 and x.is_integer() and y.is_integer()):
        return None

    return (int(x), int(y))

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    total = 0

    curr_eq1: list[int] = []
    curr_eq2: list[int] = []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            line = line.strip()

            if not line:
                curr_eq1 = []
                curr_eq2 = []

            numbers = re.findall(r'\+(\d+)', line)
            if numbers:
                a, b = [int(num) for num in numbers]
                curr_eq1.append(a)
                curr_eq2.append(b)

            numbers = re.findall(r'\=(\d+)', line)
            if numbers:
                a, b = [int(num) for num in numbers]
                curr_eq1.append(a + 10**13)
                curr_eq2.append(b + 10**13)

            if len(curr_eq1) == 3:
                assert len(curr_eq2) == 3

                ans = solve_linear_equations(curr_eq1, curr_eq2)
                if ans is not None:
                    x, y = ans
                    total += (3*x + y)


    print(total)

if __name__ == '__main__':
    main()

