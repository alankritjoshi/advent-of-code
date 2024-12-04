import argparse

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

            prev = None
            nums = []
            diffs = []
            for num_str in line.split():
                num = int(num_str)
                if prev is not None:
                    diff = num - prev
                    diffs.append(diff)
                prev = num
                nums.append(num)


            is_valid_diff = all(abs(diff) >= 1 and abs(diff) <= 3 for diff in diffs)
            is_increasing = all(nums[i] >= nums[i+1] for i in range(len(nums)-1))
            is_decreasing = all(nums[i] <= nums[i+1] for i in range(len(nums)-1))

            if is_valid_diff and (is_increasing or is_decreasing):
                total += 1
                
    print(total)


if __name__ == '__main__':
    main()

