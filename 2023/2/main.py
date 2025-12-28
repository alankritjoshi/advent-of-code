import argparse
from dataclasses import dataclass, field


def main() -> None:
    args = argparse.ArgumentParser(description="AoC runner")

    args.add_argument("-i" , "--input", type=str, required=True, help="Input File Name")

    args = args.parse_args()

    input_file_name = args.input

    power = 0

    with open(input_file_name, "r") as f:
        while True:
            line = f.readline()

            if not line:
                break

            line = line.strip()

            split_line = line.split(":")

            game_id = int(split_line[0].split(" ")[1])

            game = Game(game_id)

            for balls_set in split_line[1].strip().split("; "):
                for ball in balls_set.split(", "):
                    count, color = ball.split(" ")
                    game.add_ball(color, int(count))

            power += game.power()

    print(power)

@dataclass
class Ball:
    count: int = 0
    min: int = -1

    def add(self, count: int) -> None:
        self.count += count
        self.min = max(self.min, count)

@dataclass
class Game:
    id: int
    red: Ball = field(default_factory=lambda: Ball())
    green: Ball = field(default_factory=lambda: Ball())
    blue: Ball = field(default_factory=lambda: Ball())

    def add_ball(self, color: str, count: int) -> None:
        if color == "red":
            self.red.add(count)
        elif color == "green":
            self.green.add(count)
        elif color == "blue":
            self.blue.add(count)
        else:
            raise Exception(f"Invalid color: {color}")

    def is_valid(self) -> bool:
        return self.red.count <= 12 and self.green.count <= 13 and self.blue.count > 14

    def power(self) -> int:
        if self.red.count == -1 or self.green.count == -1 or self.blue.count == -1:
            return 0

        red = self.red.min if self.red.min != -1 else 0
        green = self.green.min if self.green.min != -1 else 0
        blue = self.blue.min if self.blue.min != -1 else 0

        return red * green * blue

    def reset(self) -> None:
        self.red = Ball()
        self.green = Ball()
        self.blue = Ball()

    def __repr__(self) -> str:
        return f"Game {self.id}: {self.red} red, {self.green} green, {self.blue} blue, {self.power()} power"


if __name__ == '__main__':
    main()

