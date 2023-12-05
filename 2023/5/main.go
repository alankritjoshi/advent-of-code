package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	seeds_keyword = "seeds"
	map_keyword   = "map"
)

type valueRange struct {
	Start       int
	End         int
	Destination int
}

func (vr valueRange) getDestinationValue(sourceValue int) (int, bool) {
	if sourceValue < vr.Start || sourceValue > vr.End {
		return 0, false
	}

	return vr.Destination + (sourceValue - vr.Start), true
}

type valueRanges []valueRange

func (vrs valueRanges) getDestinationValue(sourceValue int) int {
	for _, vr := range vrs {
		if val, ok := vr.getDestinationValue(sourceValue); ok {
			return val
		}
	}

	return sourceValue
}

type Entity map[string]valueRanges // location -> ValueMap, seed -> ValueMap

type EntityMap map[string]Entity // seed -> Entity, soil -> Entity

func main() {
	input := flag.String("i", "", "Input File Name")

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*input)
	if err != nil {
		log.Fatal("unable to open file: %w", err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var (
		seeds     []int
		entityMap = make(EntityMap)

		lastKey   = ""
		lastValue = ""
	)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("unable to read line: %w", err)
		}

		line = strings.TrimSuffix(line, "\n")

		if strings.Contains(line, seeds_keyword) {
			seeds = getNumFields(line)
		} else if strings.Contains(line, map_keyword) {
			mapName := getFields(line)[0]
			mapNameSplit := strings.Split(mapName, "-")

			mapKey := mapNameSplit[0]
			mapValue := mapNameSplit[2]

			var (
				entity Entity
				ok     bool
			)

			if entity, ok = entityMap[mapKey]; !ok {
				entity = make(Entity)
			}

			entity[mapValue] = make(valueRanges, 0)

			entityMap[mapKey] = entity

			lastKey = mapKey
			lastValue = mapValue
		} else if line == "" {
			continue
		} else {
			nums := getNumFields(line)

			destStart := nums[0]
			sourceStart := nums[1]
			length := nums[2]

			entityMap[lastKey][lastValue] = append(
				entityMap[lastKey][lastValue],
				valueRange{
					Start:       sourceStart,
					End:         sourceStart + length - 1,
					Destination: destStart,
				},
			)
		}
	}

	minLocation := math.MaxInt32

	for _, seed := range seeds {
		current := seed

		entityName := "seed"

		for {
			m := entityMap[entityName]

			var values valueRanges

			for entityName, values = range m {
				current = values.getDestinationValue(current)
				break
			}

			if entityName == "location" {
				break
			}
		}

		minLocation = int(math.Min(float64(minLocation), float64(current)))
	}

	fmt.Println(minLocation)
}

func getFields(line string) []string {
	return strings.Fields(line)
}

func getNumFields(line string) []int {
	var nums []int

	for _, seed := range getFields(line) {
		num, err := strconv.Atoi(seed)
		if err != nil {
			continue
		}

		nums = append(nums, num)
	}

	return nums
}
