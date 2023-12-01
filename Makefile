.PHONY: cookie input z

# set default to current year and day 1
year ?= $(shell date +'%Y')
day ?= 1

# get the cookie from adventofcode.com
cookie:
	@ ./run.sh -c

# get the input
input:
	@ ./run.sh -y $(year) -d $(day)

# edit using zellij layout (IFF you have zellij installed)
# return exit code 1 if zellij is not installed
z:
	@ zellij -V || exit 1
	@ YEAR=$(year) DAY=$(day) zellij action new-tab -l .zellij/layout.kdl

test:
	@ go run $(year)/$(day)/main.go -i $(year)/$(day)/input.txt

