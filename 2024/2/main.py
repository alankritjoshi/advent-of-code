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

            numbers = []
            for num_str in line.split():
                num = int(num_str)
                numbers.append(num)

            def is_valid(nums: list[int]) -> bool:
                is_valid_diff = all(abs(nums[i] - nums[i+1]) >= 1 and abs(nums[i] - nums[i+1]) <= 3 for i in range(len(nums)-1))
                is_increasing = all(nums[i] >= nums[i+1] for i in range(len(nums)-1))
                is_decreasing = all(nums[i] <= nums[i+1] for i in range(len(nums)-1))
                return is_valid_diff and (is_increasing or is_decreasing)

            def can_make_valid(nums: list[int]) -> bool:
                for i in range(len(nums)):
                    if is_valid(nums[:i] + nums[i+1:]):
                        return True
                return False

            if is_valid(numbers) or can_make_valid(numbers):
                total += 1
                
    print(total)


if __name__ == '__main__':
    main()

