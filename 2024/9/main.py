import argparse
from collections import deque
from dataclasses import dataclass


@dataclass
class Segment:
    index: int
    vals: deque[tuple[int, int]]
    free: int = 0


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

            block: list[Segment] = []
            spaces: deque[Segment] = deque([])

            def printer():
                for segment in block:  # noqa
                    print("(", end=" ")
                    for file_number, count in segment.vals:
                        print(" ".join(str(file_number) for _ in range(count)), end=" ")
                    print(" ".join("*" for _ in range(segment.free)), end=" ")
                    print(")", end=" ")
                print()

            file_number = 0

            for index, ch in enumerate(line.strip("\n")):
                segment: Segment | None = None
                if index % 2 != 0:
                    segment = Segment(index=index, vals=deque([]), free=int(ch))
                    spaces.append(segment)
                else:
                    segment = Segment(index=index, vals=deque([(file_number, int(ch))]))
                    file_number += 1

                block.append(segment)

            file_index = len(block) - 1

            while spaces and file_index >= 0:
                segment = block[file_index]

                empty_spaces: set[int] = set()

                for space_index, space in enumerate(spaces):
                    if not segment.vals:
                        break

                    if space.free < segment.vals[-1][1]:
                        continue

                    if space.index > segment.index:
                        break

                    file_number, size = segment.vals.pop()

                    space.free -= size
                    if not space.free:
                        empty_spaces.add(space_index)

                    segment.free += size

                    block[space.index].vals.append((file_number, size))

                spaces = deque([space for index, space in enumerate(spaces) if index not in empty_spaces])
                file_index -= 1

            total = 0
            index = 0
            for segment in block:
                for file_number, size in segment.vals:
                    # print(file_number, end=" ")
                    for _ in range(size):
                        total += file_number * index
                        index += 1

                for _ in range(segment.free):
                    index += 1

            print(total)


if __name__ == "__main__":
    main()
