import argparse


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

            print(line)

if __name__ == '__main__':
    main()

