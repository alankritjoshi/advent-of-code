import argparse
import heapq

def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    a, b = [], []

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            first, second = line.split()
            heapq.heappush(a, int(first))
            heapq.heappush(b, int(second))

    total = 0
    for _ in range(len(a)):
        total += abs(heapq.heappop(a) - heapq.heappop(b))

    print(total)

if __name__ == '__main__':
    main()

