#!/bin/bash

FETCH_COOKIE=false
GET_INPUT=false
TEMPLATE=""
TXT=""
YEAR=""
DAY=""

while getopts ":ci:t:y:d:" opt; do
	case $opt in
	c)
		FETCH_COOKIE=true
		;;
	i)
		TXT=$OPTARG
		;;
	t)
		TEMPLATE="$OPTARG"
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

	gum confirm --default=false "Get Cookie from Browser?" && open https://adventofcode.com

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
else
	echo "No Cookie Provided!"
	exit 1
fi

if [ -n "$TXT" ]; then
	if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
		echo "Year and Day must be specified for creating using Input flag!"
		exit 1
	fi

	DIRECTORY="$YEAR/$DAY"

	CANCEL=0
	if [ -e "$DIRECTORY/$TXT" ]; then
		echo "WARNING: $TXT already exists!"

		gum confirm --default=false "Overwrite $DIRECTORY/$TXT anyway?"
		CANCEL=$?
	else
		if [ ! -d "$DIRECTORY" ]; then
			mkdir -p "$DIRECTORY"
		fi
	fi

	if [ $CANCEL -eq 0 ]; then
		http --check-status --ignore-stdin -o $DIRECTORY/${TXT} https://adventofcode.com/$YEAR/day/$DAY/input "Cookie:session=$AOC_COOKIE"
		# If the input is not available, delete the file
		if [ $? -ne 0 ]; then
			rm -rf $DIRECTORY/$TXT
			# If $DIRECTORY is empty, delete it
			if [ ! "$(ls -A $DIRECTORY)" ]; then
				rm -rf $DIRECTORY
			fi

			echo "Input download cancelled/failed!"
			exit 1
		else
			echo "Input downloaded to $DIRECTORY/$TXT"
			# If non-sample.txt file is created, ask to create sample.txt
			if [ $TXT != "sample.txt" ]; then
				if [ -e "$DIRECTORY/sample.txt" ]; then
					WARNING="WARNING: sample.txt already exists! BUT..."
					gum confirm --default=false "WARNING: sample.txt already exists! Overwrite and create new one?"
				else
					gum confirm --default=false "Create sample file?"
				fi

				if [ $? -eq 0 ]; then
					head -n 5 $DIRECTORY/$TXT >$DIRECTORY/sample.txt
					echo "Sample created in $DIRECTORY/sample.txt"
				fi
			fi
		fi
	fi
fi

if [ -n "$TEMPLATE" ]; then
	if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
		echo "Year and Day must be specified for creating using Template flag!"
		exit 1
	fi

	DIRECTORY="$YEAR/$DAY"

	CANCEL=0
	if [ -e "$DIRECTORY/main.$TEMPLATE" ]; then
		echo "WARNING: main.$TEMPLATE already exists!"

		gum confirm --default=false "Overwrite $DIRECTORY/main.$TEMPLATE anyway?"
		CANCEL=$?
	else
		if [ ! -d "$DIRECTORY" ]; then
			mkdir -p "$DIRECTORY"
		fi
	fi

	if [ $CANCEL -eq 0 ]; then
		cp "./templates/main.$TEMPLATE" "./$DIRECTORY/main.$TEMPLATE"
		# If the copy failed
		if [ $? -ne 0 ]; then
			# If $DIRECTORY is empty, delete it
			if [ ! "$(ls -A $DIRECTORY)" ]; then
				rm -rf $DIRECTORY
			fi

			echo "Template copy cancelled/failed!"
			exit 1
		else
			echo "Template copied to $DIRECTORY/main.$TEMPLATE"
		fi
	fi
fi
