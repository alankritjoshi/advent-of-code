# [My](https://github.com/alankritjoshi) Advent of Code

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

### [OPTIONAL] [Zellij](https://github.com/zellij-org/zellij)

```sh
brew install zellij
```

## Usage

> NOTE: *year*, *day*, *lang* & *txt* are optional. If not set, the *year*=<*current*>, *day*=*1*, *lang*=*go* and *txt*=*input.txt* are used.
> e.g., `make run year=2023 day=2 lang=py txt=sample.txt`

### Setup

Setup cookie, input, and template.

```sh
make setup
```

### Setup using Zellij

[Setup](#setup) and then open corresponding setup in Zellij.

```sh
make setupz
```

### Set Cookie

```sh
make cookie
```

### Download Input

Download the input to *year*/*day*/*txt*.

```sh
make input
```

### Copy Template

Copy the template for *lang* (`go` or `py`) to *year*/*day*/main.*lang*.

```sh
make template
```

### Run

```sh
make run
```

### Run with Hot Reload

```sh
make hot
```

### Edit & Run using Zellij

If using Zellij, edit and run the code with one of the [layouts](.zellij/) corresponding to the *lang*.

```sh
make z
```

