.PHONY: cookie input template run air z setup setupz

# set default to current year and day 1
year ?= $(shell date +'%Y')
day ?= 1
lang ?= go
txt ?= input.txt

# get the cookie from adventofcode.com
cookie:
	@ ./setup.sh -c

# get the input
input:
	@ ./setup.sh -i $(txt) -y $(year) -d $(day)

# create main file from template
template:
	@ ./setup.sh -t $(lang) -y $(year) -d $(day)

run:
ifeq ($(lang),go)
	@ go run $(year)/$(day)/main.go -i $(year)/$(day)/$(txt)
else ifeq ($(lang),py)
	@ poetry run python $(year)/$(day)/main.py -i $(year)/$(day)/$(txt)
else
  @ echo "Unsupported language: $(lang)"
endif

hot:
ifeq ($(lang),go)
	@ cd $(year)/$(day) && go run github.com/cosmtrek/air --build.args_bin="-i,$(txt)" || true
	@ cd ..
else ifeq ($(lang),py)
	@ poetry run python hotreload.py -s "$(year)/$(day)/main.py -i $(year)/$(day)/$(txt)"
else
  @ echo "Unsupported language: $(lang)"
endif

# edit using zellij layout (IFF you have zellij installed)
# return exit code 1 if zellij is not installed
z:
	@ zellij -V >/dev/null || exit 1
	@ YEAR=$(year) DAY=$(day) zellij action new-tab -l .zellij/$(lang).kdl

setup: cookie input template

setupz: cookie input template z
