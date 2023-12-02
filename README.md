# Advent of Code

## Requirements

### Talking with [AoC](https://adventofcode.com)

- bash
- [httpie](https://github.com/httpie/cli)
- [Gum](https://github.com/charmbracelet/gum)

### [OPTIONAL] Edit & Run with Zellij

- [Zellij](https://github.com/zellij-org/zellij)

## Usage

> NOTE: *year* and *day* are optional. If not set, the *current* *year* and *DAY*=*1* are used.

### Setup

Sets up cookie, input, and template.

```sh
make setup day=2
```

### Setup using Zellij

Sets up cookie, input, and template. And then opens corresponding setup in Zellij.

```sh
make setupz day=2
```

### Set Cookie

```sh
make cookie
```

### Get Input

Download the input to $year/$day/input.txt.

```sh
make input year=2023 day=1
```

### Set Template

Copies the specified `TEMPLATE` (`go`, `py`, etc.) to $year/$day/main.`<TEMPLATE>`.

```sh
make template template=go day=1
```

### Edit & Run using Zellij

If you use [Zellij](https://github.com/zellij-org/zellij), you can edit and run the code with the [layout](.zellij/layout.kdl).

```sh
make z year=2023 day=1
```

### Test

```sh
make test year=2023 day=1
```
