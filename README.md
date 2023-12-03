# Advent of Code

## Requirements

> NOTE: install instructions are for macOS using [Homebrew](https://brew.sh). Use your package manager of choice.

### Talking with [AoC](https://adventofcode.com)

```sh
brew install bash httpie gum
```

### Languages

```sh
brew install go python pyenv poetry
```

#### [Go](https://golang.org) 1.21+

```sh
go mod download
```

#### [Python](https://www.python.org) 3.12+

```sh
pyenv install
pip install poetry
poetry install
```

### [OPTIONAL] Edit & Run with Zellij

- [Zellij](https://github.com/zellij-org/zellij)

## Usage

> NOTE: *year*, *day* and *template* are optional. If not set, the *year*=<*current*>, *day*=*1* and *template*=*go* are used.

### Setup

Set up cookie, input, and template.

```sh
make setup day=2 template=py
```

### Setup using Zellij

[Setup](#setup) and then open corresponding setup in Zellij.

```sh
make setupz day=2
```

### Set Cookie

```sh
make cookie
```

### Get Input

Download the input to *year*/*day*/input.txt.

```sh
make input year=2023 day=1
```

### Set Template

Copies the specified *template* (`go`, `py`, etc.) to *year*/*day*/main.*template*.

```sh
make template template=py day=1
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
