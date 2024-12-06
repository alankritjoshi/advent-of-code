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

    def fix_topological_order(nodes: list[int]) -> list[int]:
        seen = set()
        for i in range(len(nodes)):
            for neighbor in graph[nodes[i]]:
                if neighbor in seen:
                    to_swap = nodes.index(neighbor) # find the neighbor that shouldn't have been seen first
                    nodes[i], nodes[to_swap] = nodes[to_swap], nodes[i] # and swap
                    nodes = fix_topological_order(nodes) # hacky but ran this program again from beginning lol
                    return nodes
            seen.add(nodes[i])
        return nodes

    middles = 0
    for update in updates:
        if not check_topological_order(update):
            update = fix_topological_order(update)
            middles  += update[len(update) // 2]

    print(middles)



if __name__ == '__main__':
    main()

