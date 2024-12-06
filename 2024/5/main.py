import argparse
from collections import defaultdict
import math

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    update_ready = False

    graph: dict[int, list[int]] = defaultdict(list)
    updates: list[list[int]] = [] 

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            if line == "\n":
                update_ready = True
                continue

            line = line.strip()

            if not update_ready:
                a, b = line.split("|")
                graph[int(a)].append(int(b))
            else:
                updates.append([int(num) for num in line.split(",")])

    def check_topological_order(nodes: list[int]) -> bool:
        seen = set()
        for node in nodes:
            for neighbor in graph[node]:
                if neighbor in seen:
                    return False
            seen.add(node)
        return True

    middles = 0
    for update in updates:
        if check_topological_order(update):
            middles  += update[len(update) // 2]

    print(middles)



if __name__ == '__main__':
    main()

