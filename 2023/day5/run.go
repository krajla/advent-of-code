package day5

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type seedRange struct {
	start  int
	ranger int
}

func Run() {
	reader, err := os.Open("./day5/input")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	seedsLine := scanner.Text()
	seedsLine = seedsLine[7:]

	seeds := make([]int, 0)
	for _, seedStr := range strings.Split(seedsLine, " ") {
		num := util.QuickAtoi(seedStr)
		seeds = append(seeds, num)
	}

	seedMappings := make([]int, len(seeds))
	copy(seedMappings, seeds)

	seedRangeMappings := make([]seedRange, 0, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seedRangeMappings = append(seedRangeMappings, seedRange{
			start:  seeds[i],
			ranger: seeds[i+1],
		})
	}
	seedRangeMappingsAux := seedRangeMappings

	var indexVisited map[int]struct{}
	var rangeIndexVisited map[int]struct{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, ":") {
			indexVisited = make(map[int]struct{})
			seedRangeMappings = seedRangeMappingsAux

			seedRangeMappings = inserOutOfRanges(seedRangeMappings, seedRangeMappingsAux, rangeIndexVisited)
			rangeIndexVisited = make(map[int]struct{})
			seedRangeMappingsAux = make([]seedRange, 0, len(seedRangeMappings))
			continue
		}

		parts := strings.Split(line, " ")
		destStart := util.QuickAtoi(parts[0])
		sourceStart := util.QuickAtoi(parts[1])
		ranger := util.QuickAtoi(parts[2])

		// First Part
		for i, mapping := range seedMappings {
			if mapping < sourceStart || mapping > sourceStart+ranger {
				continue
			}
			if _, ok := indexVisited[i]; !ok {
				indexVisited[i] = struct{}{}
				seedMappings[i] = destStart + (mapping - sourceStart)
			}
		}

		// Second Part
		for i := 0; i < len(seedRangeMappings); i++ {
			rangeMapping := seedRangeMappings[i]
			if _, ok := rangeIndexVisited[i]; ok {
				continue
			}
			if rangeMapping.start+rangeMapping.ranger <= sourceStart || rangeMapping.start >= sourceStart+ranger {
				continue
			}
			rangeIndexVisited[i] = struct{}{}

			var outRangeEntry seedRange
			if rangeMapping.start <= sourceStart {
				outStart := rangeMapping.start
				outRange := sourceStart - rangeMapping.start
				mappedEntry := seedRange{
					start:  destStart,
					ranger: rangeMapping.ranger - outRange,
				}
				seedRangeMappingsAux = mergeInsert(seedRangeMappingsAux, mappedEntry)

				outRangeEntry = seedRange{
					start:  outStart,
					ranger: outRange,
				}
			} else {
				outStart := sourceStart + ranger + 1
				outRange := rangeMapping.start + rangeMapping.ranger - outStart
				mappedEntry := seedRange{
					start:  destStart + rangeMapping.start - sourceStart,
					ranger: min(rangeMapping.ranger-outRange, rangeMapping.ranger),
				}
				seedRangeMappingsAux = mergeInsert(seedRangeMappingsAux, mappedEntry)

				outRangeEntry = seedRange{
					start:  outStart,
					ranger: outRange,
				}
			}

			if outRangeEntry.start >= 0 && outRangeEntry.ranger > 0 {
				seedRangeMappings = append(seedRangeMappings, outRangeEntry)
			}
		}
	}
	seedRangeMappings = inserOutOfRanges(seedRangeMappings, seedRangeMappingsAux, rangeIndexVisited)

	slices.Sort(seedMappings)
	slices.SortFunc(seedRangeMappings, func(a, b seedRange) int { return a.start - b.start })
	fmt.Printf("Lowest height to plant seeds: %d\n", seedMappings[0])
	fmt.Printf("Lowest height to plant seed ranges: %d\n", seedRangeMappings[0].start)
}

func inserOutOfRanges(seedRangeMappings, seedRangeMappingsAux []seedRange, rangeIndexVisited map[int]struct{}) []seedRange {
	for i := 0; i < len(seedRangeMappings); i++ {
		rangeMapping := seedRangeMappings[i]
		if _, ok := rangeIndexVisited[i]; ok {
			continue
		}

		seedRangeMappingsAux = mergeInsert(seedRangeMappingsAux, rangeMapping)
	}

	return seedRangeMappingsAux
}

func mergeInsert(array []seedRange, entry seedRange) []seedRange {
	// Could keep sorted and binary search, won't.
	for i, currRange := range array {
		if currRange.start+currRange.ranger <= entry.start || currRange.start >= entry.start+entry.ranger {
			continue
		}

		newEntry := seedRange{
			start:  min(currRange.start, entry.start),
			ranger: max(currRange.ranger, entry.ranger),
		}
		newSlice := slices.Delete(array, i, i+1)
		return mergeInsert(newSlice, newEntry)

	}

	return append(array, entry)
}
