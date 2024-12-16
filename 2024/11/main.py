import argparse
from dataclasses import dataclass
from typing import Optional

@dataclass
class Node:
    val: int
    next: Optional['Node'] = None

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

            stones: list[int] = [int(ch) for ch in line.strip().split()]

            head: Node | None = None

            curr = None
            for stone in stones:
                node = Node(val=stone, next=None)
                if head is None:
                    head = node
                    curr = node
                else:
                    assert curr is not None
                    curr.next = node
                    curr = node

            for _ in range(25):
                curr = head
                while curr:
                    val = curr.val
                    str_val = str(val)
                    if val == 0:
                        curr.val = 1
                        curr = curr.next
                    elif len(str_val) % 2 == 0:
                        first, second = str_val[:len(str_val) // 2], str_val[len(str_val) // 2:]
                        curr.val = int(first)
                        old_next = curr.next
                        curr.next = Node(val=int(second), next=old_next)
                        curr = old_next
                    else:
                        curr.val *= 2024
                        curr = curr.next

            total = 0
            curr = head
            while curr:
                total += 1
                curr = curr.next

            print(total)


if __name__ == '__main__':
    main()

