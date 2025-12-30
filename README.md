# [My](https://github.com/alankritjoshi) Advent of Code

## Requirements

### To setup [AoC](https://adventofcode.com)

```sh
brew install bash httpie gum # zellij (optional)
```

> NOTE: use package manager of your choice.

### Languages

Install [mise](https://mise.jdx.dev), then run:

```sh
make
```

## Usage

### TL;DR

```sh
make setup year=2025 day=2 lang=rb # If Zellij, make setupz year=2025 day=2 lang=rb
make run year=2025 day=2 lang=rb   # If Zellij, make z year=2025 day=2 lang=rb
```

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
