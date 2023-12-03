.PHONY: cookie input template test z setup setupz

# set default to current year and day 1
year ?= $(shell date +'%Y')
day ?= 1
template ?= go

# get the cookie from adventofcode.com
cookie:
	@ ./run.sh -c

# get the input
input:
	@ ./run.sh -i -y $(year) -d $(day)

# create main file from template
template:
	@ ./run.sh -t $(template) -y $(year) -d $(day)

test:
	@ go run $(year)/$(day)/main.go -i $(year)/$(day)/input.txt

# edit using zellij layout (IFF you have zellij installed)
# return exit code 1 if zellij is not installed
z:
	@ zellij -V >/dev/null || exit 1
	@ YEAR=$(year) DAY=$(day) zellij action new-tab -l .zellij/$(template).kdl

setup: cookie input template

setupz: cookie input template z
