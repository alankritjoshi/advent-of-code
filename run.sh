#!/bin/bash

FETCH_COOKIE=false
YEAR=""
DAY=""

while getopts ":cy:d:" opt; do
	case $opt in
	c)
		FETCH_COOKIE=true
		;;
	y)
		YEAR="$OPTARG"
		;;
	d)
		DAY="$OPTARG"
		;;
	\?)
		echo "Invalid option: -$OPTARG" >&2
		exit 1
		;;
	esac
done

# Fetch cookie if requested
if [ "$FETCH_COOKIE" = true ]; then
	# Load cookie if it exists
	if [ -f .env ]; then
		source .env

		echo "WARNING: Cookie is already set: $AOC_COOKIE"
	fi

	gum confirm "Get Cookie from Browser?" && open https://adventofcode.com

	AOC_COOKIE=$(gum input --placeholder "Input Session Cookie..." --value "$AOC_COOKIE")

	if [ -z "$AOC_COOKIE" ]; then
		echo "No Cookie Provided!"
		exit 1
	else
		echo "AOC_COOKIE=$AOC_COOKIE" >.env
	fi
fi

if [ -f .env ]; then
	source .env

	echo "Cookie is: $AOC_COOKIE"
else
	echo "No Cookie Provided!"
	exit 1
fi

if [ -n "$YEAR" ] && [ -n "$DAY" ]; then
	DIRECTORY="$YEAR/$DAY"
	if [ ! -d "$DIRECTORY" ]; then
		mkdir -p "$DIRECTORY"
	fi

	http --check-status --ignore-stdin -o $DIRECTORY/input.txt https://adventofcode.com/$YEAR/day/$DAY/input "Cookie:session=$AOC_COOKIE"

	# If the input is not available, delete the file
	if [ $? -ne 0 ]; then
		rm -rf $DIRECTORY/input.txt
		# If $DIRECTORY is empty, delete it
		if [ ! "$(ls -A $DIRECTORY)" ]; then
			rm -rf $DIRECTORY
		fi
		exit 1
	else
		echo "Input downloaded to $DIRECTORY/input.txt"
	fi
fi
