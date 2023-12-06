package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	seeds_keyword = "seeds"
	map_keyword   = "map"
)

type valRange struct {
	start int
	end   int
}

type crossedRange struct {
	valRange
	overlap bool
}

func (cr crossedRange) String() string {
	return fmt.Sprintf("%d-%d, overlap: %t", cr.start, cr.end, cr.overlap)
}

type crossedRanges []crossedRange

func (crs crossedRanges) String() string {
	var s string
	for _, cr := range crs {
		s += fmt.Sprintf("%s\n", cr)
	}
	return s
}

// returns crossed ranges when overlap exists
// if no overlap, then returns the original range as a crossedRange
// e.g.,
// no overlap: [{1, 5}, {6, 9}] -> [{{1, 5}, false}]
// overlap:    [{1, 5}, {4, 8}, {7, 9}] -> [{{1, 3}, false}, {{4, 8}, true}, {{9, 9}, false}]
func (vr valRange) crossWithAnother(other valRange) crossedRanges {
	start1 := vr.start
	end1 := vr.end
	start2 := other.start
	end2 := other.end

	// non-overlapping ranges -> return the original *unchanged* range
	if end1 < start2 || end2 < start1 {
		return crossedRanges{
			{
				valRange: vr,
				overlap:  false,
			},
		}
	}

	result := make(crossedRanges, 0)

	// non-overlapping range at the beginning
	if start1 < start2 {
		result = append(result,
			crossedRange{
				valRange: valRange{
					start: start1,
					end:   min(start2, end1) - 1,
				},
				overlap: false,
			},
		)
	}

	// overlapping range
	result = append(result,
		crossedRange{
			valRange: valRange{
				start: max(start1, start2),
				end:   min(end1, end2),
			},
			overlap: true,
		},
	)

	// non-overlapping range at the end
	if end2 < end1 {
		result = append(result,
			crossedRange{
				valRange: valRange{
					start: end2 + 1,
					end:   end1,
				},
				overlap: false,
			},
		)
	}

	// cleanup invalid ranges e.g., [7, 6]
	var finalResult crossedRanges

	for _, r := range result {
		if r.start <= r.end {
			finalResult = append(finalResult, r)
		}
	}

	return finalResult
}

// MergeValRange merges overlapping ranges
// e.g., [{1, 1}, {1, 5}, {4, 6}, {8, 9}] -> [{1, 6}, {8, 9}]
func MergeValRange(r []valRange) []valRange {
	sort.Slice(r, func(i, j int) bool {
		return r[i].start < r[j].start
	})

	merged := make([]valRange, 0)

	for _, current := range r {
		if len(merged) == 0 {
			merged = append(merged, current)
			continue
		}

		last := merged[len(merged)-1]

		if current.start <= last.end+1 {
			merged[len(merged)-1] = valRange{
				start: last.start,
				end:   max(last.end, current.end),
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

type mapRange struct {
	valRange
	destinationStart int
}

// returns the mapped destination ranges and the unchanged ranges when crossing
// a mapRange (valRange, destinationStart) with a given valRange (start, end)
// e.g., {{5, 10}, 25} with {7, 12} -> [{27, 30}], [{11, 12}]
// because
//
//	{7, 10} is the crossedRange that can be mapped to {27, 30} using destinationStart = 25
//	    i.e., {25 + (7 -5), 25 + (10 - 5)} = {27, 30}
//	{11, 12} is unchanged range
func (mr mapRange) getDestinationRanges(other valRange) ([]valRange, []valRange) {
	mappedDestinations := make([]valRange, 0)
	unchanged := make([]valRange, 0)

	for _, cr := range other.crossWithAnother(mr.valRange) {
		if cr.overlap {
			mappedDestinations = append(mappedDestinations, valRange{
				start: mr.destinationStart + (cr.start - mr.start),
				end:   mr.destinationStart + (cr.end - mr.start),
			})
		} else {
			unchanged = append(unchanged, cr.valRange)
		}
	}

	return MergeValRange(mappedDestinations), MergeValRange(unchanged)
}

type mapRanges []mapRange

// returns the mapped destination ranges and the unchanged ranges when crossing
// the given input valRanges with the mapRanges
//
// first all the input valRanges (currentUnchanged) are crossed with each mapRange
// while crossing with a mappedRange, the crossedRanges that mapped to the destinations are stored in allMappedDestinations
// while crossing with a mappedRange, the unchanged ranges are stored for the next mapRange
//
// e.g., [{{5, 10}, 25}, {{16, 18}, 3}] with [{7, 12}, {15, 17}] -> [{27, 30}], [{11, 12}]
// {{5, 10}, 25}
//
//	is crossed with {7, 12} -> mappedDestinations: [{27, 30}], unchanged: [{11, 12}]
//	is crossed with {15, 17} -> mappedDestinations: [], unchanged: [{11, 12}, {15, 17]}
//
// {{16, 18}, 3}
//
// is crossed with {11, 12} -> mappedDestinations: [], unchanged: [{11, 12}]
// is crossed with {15, 17} -> mappedDestinations: [{3, 4}], unchanged: [{11, 12}, {15, 15}]
//
// final unchanged = {11, 12}, {15, 15}
// final mappedDestinations = {3, 4}, {27, 30}
//
// final [{3, 4}, {11, 12}, {15, 15}, {27, 30}] -> these are merged if they overlap
func (mrs mapRanges) getDestinationValue(others []valRange) []valRange {
	var (
		allMappedDestinations = make([]valRange, 0)
		currentUnchanged      = others
		nextUnchanged         []valRange
	)

	for _, mr := range mrs {
		nextUnchanged = make([]valRange, 0)

		for len(currentUnchanged) != 0 {
			current := currentUnchanged[0]
			currentUnchanged = currentUnchanged[1:]

			mappedDestinations, unchanged := mr.getDestinationRanges(current)

			allMappedDestinations = append(allMappedDestinations, mappedDestinations...)

			nextUnchanged = append(nextUnchanged, unchanged...)
		}

		currentUnchanged = MergeValRange(nextUnchanged)
	}

	final := append(allMappedDestinations, currentUnchanged...)

	final = MergeValRange(final)

	return final
}

type entity map[string]mapRanges

type entityMap map[string]entity

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
		seedRanges []*valRange
		em         = make(entityMap)

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
			nums := getNumFields(line)

			for i, num := range nums {
				if i%2 == 0 {
					seedRanges = append(seedRanges, &valRange{start: num})
				} else {
					r := seedRanges[len(seedRanges)-1]
					r.end = r.start + num - 1
				}
			}
		} else if strings.Contains(line, map_keyword) {
			mapName := getFields(line)[0]
			mapNameSplit := strings.Split(mapName, "-")

			mapKey := mapNameSplit[0]
			mapValue := mapNameSplit[2]

			var (
				e  entity
				ok bool
			)

			if e, ok = em[mapKey]; !ok {
				e = make(entity)
			}

			e[mapValue] = make(mapRanges, 0)

			em[mapKey] = e

			lastKey = mapKey
			lastValue = mapValue
		} else if line == "" {
			continue
		} else {
			nums := getNumFields(line)

			destStart := nums[0]
			sourceStart := nums[1]
			length := nums[2]

			em[lastKey][lastValue] = append(
				em[lastKey][lastValue],
				mapRange{
					valRange: valRange{
						start: sourceStart,
						end:   sourceStart + length - 1,
					},
					destinationStart: destStart,
				},
			)
		}
	}

	minLocation := math.MaxInt32

	for _, seedRange := range seedRanges {
		current := []valRange{
			{
				start: seedRange.start,
				end:   seedRange.end,
			},
		}

		entityName := "seed"

		minCurrent := math.MaxInt32

		for {
			m := em[entityName]

			var values mapRanges

			for entityName, values = range m {
				current = values.getDestinationValue(current)
				break
			}

			if entityName == "location" {
				break
			}
		}

		for _, c := range current {
			minCurrent = min(minCurrent, c.start)
		}

		minLocation = min(minLocation, minCurrent)
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
