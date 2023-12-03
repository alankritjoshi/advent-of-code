import argparse

from dataclasses import dataclass

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    total = 0

    with open(input_file_name, "r") as f:

        while True:
            line = f.readline()

            if not line:
                break

            seen = ""
            two_digits = TwoDigits()

            for char in line:
                if char.isdigit():
                    seen = ""
                    two_digits.update_digit(int(char))
                else:
                    seen += char
                    two_digits.maybe_update_word(seen)

            total += int(two_digits)

    print(total)


DIGIT_TO_WORD = {
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine"
}

@dataclass
class TwoDigits:
    first: int = -1
    second: int = -1

    def update_digit(self, digit: int) -> None:
        if self.first == -1:
            self.first = digit
        self.second = digit

    def maybe_update_word(self, seen: str) -> None:
        for digit, word in DIGIT_TO_WORD.items():
            if seen.endswith(word):
                self.update_digit(digit)
                break

    def __repr__(self) -> str:
        return f"{self.first}{self.second}"

    def __int__(self) -> int:
        return int(self.__repr__())


if __name__ == '__main__':
    main()

