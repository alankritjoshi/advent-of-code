.PHONY: cookie input template run air z setup setupz

# set default to current year and day 1
year ?= $(shell date +'%Y')
day ?= 1
lang ?= go
txt ?= input.txt

# get the cookie from adventofcode.com
cookie:
	@ ./run.sh -c

# get the input
input:
	@ ./run.sh -i -y $(year) -d $(day)

# create main file from template
template:
	@ ./run.sh -t $(lang) -y $(year) -d $(day)

run:
	@ go run $(year)/$(day)/main.go -i $(year)/$(day)/$(txt)

air:
	@ cd $(year)/$(day) && go run github.com/cosmtrek/air --build.args_bin="-i,$(txt)" || true
	@ cd ..

# edit using zellij layout (IFF you have zellij installed)
# return exit code 1 if zellij is not installed
z:
	@ zellij -V >/dev/null || exit 1
	@ YEAR=$(year) DAY=$(day) zellij action new-tab -l .zellij/$(lang).kdl

setup: cookie input template

setupz: cookie input template z
