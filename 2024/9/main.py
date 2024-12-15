import argparse
from dataclasses import dataclass
from collections import deque

@dataclass
class Segment:
    index: int
    vals: deque[tuple[int, int]]
    free: int = 0

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

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
                for segment in block:
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

            while spaces and file_index > spaces[0].index:
                segment = block[file_index]

                while True:
                    if not spaces:
                        break

                    if not segment.vals:
                        break

                    file_number, size = segment.vals.pop()
                    space = spaces.popleft()

                    filled = 0
                    if space.free <= size:
                        filled = space.free
                        size -= space.free
                        space.free = 0
                        if size != 0:
                            segment.vals.append((file_number, size))
                    else:
                        filled = size
                        space.free -= filled
                        spaces.appendleft(space)
                        size = 0

                    block[space.index].vals.append((file_number, filled))

                file_index -= 1

            total = 0
            index = 0
            for segment in block:
                for file_number, size in segment.vals:
                    # print(file_number, end=" ")
                    for _ in range(size):
                        total += (file_number * index)
                        index += 1

            print(total)

if __name__ == '__main__':
    main()

