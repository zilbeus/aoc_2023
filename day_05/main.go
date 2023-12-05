package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type categoryMappings struct {
	seedToSoil         []categoryMappingRange
	soilToFertilizer   []categoryMappingRange
	fertilizerToWater  []categoryMappingRange
	waterToLight       []categoryMappingRange
	lightToTemp        []categoryMappingRange
	tempToHumidity     []categoryMappingRange
	humidityToLocation []categoryMappingRange
}

type categoryMappingRange struct {
	sourceDestinationRangeStart int
	sourceDestinationRangeEnd   int
	targetDestinationRangeStart int
	targetDestinationRangeEnd   int
}

func find(needle int, haystack []categoryMappingRange) int {
	for _, v := range haystack {
		if v.sourceDestinationRangeStart <= needle && needle <= v.sourceDestinationRangeEnd {
			idx := needle - v.sourceDestinationRangeStart
			return v.targetDestinationRangeStart + idx
		}
	}
	return needle
}

func main() {
	location := FindLocation("input.txt")
	fmt.Printf("LOWEST LOCATION NR: %d\n", location)
}

func FindLocation(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	defer file.Close()

	lines := readLines(file)
	seeds := findSeeds(lines)
	mappings := &categoryMappings{}
	mappings.seedToSoil = findCategoryMapping("seed-to-soil", lines)
	mappings.soilToFertilizer = findCategoryMapping("soil-to-fertilizer", lines)
	mappings.fertilizerToWater = findCategoryMapping("fertilizer-to-water", lines)
	mappings.waterToLight = findCategoryMapping("water-to-light", lines)
	mappings.lightToTemp = findCategoryMapping("light-to-temperature", lines)
	mappings.tempToHumidity = findCategoryMapping("temperature-to-humidity", lines)
	mappings.humidityToLocation = findCategoryMapping("humidity-to-location", lines)
	return findLowestLocationNr(seeds, mappings)
}

func readLines(f *os.File) []string {
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func findSeeds(input []string) []int {
	seeds := []int{}
	seedsLineIdx := 0
	seedsLineFound := false
	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], "seeds:") {
			seedsLineIdx = i
			seedsLineFound = true
		}
	}

	if !seedsLineFound {
		log.Fatal("expected to find input line starting with 'seeds:'")
	}

	parts := strings.Split(input[seedsLineIdx], ":")
	if len(parts) != 2 {
		log.Fatal("expected ':' separated seeds input line")
	}

	seedNrs := strings.Fields(parts[1])
	for _, v := range seedNrs {
		nr, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("expected numeric input for a seed value, got '%s'", v)
		}
		seeds = append(seeds, nr)
	}

	return seeds
}

func findCategoryMapping(mapping string, input []string) []categoryMappingRange {
	mappingLineIdx := 0
	mappingLineFound := false
	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], fmt.Sprintf("%s map:", mapping)) {
			mappingLineIdx = i
			mappingLineFound = true
		}
	}

	if !mappingLineFound {
		log.Fatalf("expected to find input line starting with '%s'", mapping)
	}

	ranges := []categoryMappingRange{}
	for i := mappingLineIdx + 1; i < len(input); i++ {
		if len(input[i]) == 0 {
			break
		}

		values := strings.Fields(input[i])
		if len(values) != 3 {
			log.Fatalf("expected 3 values for a category mapping line, got '%s'", input[i])
		}

		numericValues := []int{}
		for _, v := range values {
			nr, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("expected numeric input, got '%s'", v)
			}
			numericValues = append(numericValues, nr)
		}

		destinationRangeStart := numericValues[0]
		sourceRangeStart := numericValues[1]
		rangeLength := numericValues[2]

		ranges = append(ranges, categoryMappingRange{
			sourceDestinationRangeStart: sourceRangeStart,
			sourceDestinationRangeEnd:   sourceRangeStart + rangeLength - 1,
			targetDestinationRangeStart: destinationRangeStart,
			targetDestinationRangeEnd:   destinationRangeStart + rangeLength - 1,
		})
	}

	return ranges
}

func findLowestLocationNr(seeds []int, mappings *categoryMappings) int {
	locationMin := -1
	for _, s := range seeds {
		soil := find(s, mappings.seedToSoil)
		fertilizer := find(soil, mappings.soilToFertilizer)
		water := find(fertilizer, mappings.fertilizerToWater)
		light := find(water, mappings.waterToLight)
		temp := find(light, mappings.lightToTemp)
		humidity := find(temp, mappings.tempToHumidity)
		location := find(humidity, mappings.humidityToLocation)

		fmt.Printf("SEED: %d, LOCATION: %d\n", s, location)
		if locationMin < 0 || location < locationMin {
			locationMin = location
		}
		// fmt.Printf("MAPPINGS: %+v\n", mappings)

	}
	return locationMin
}
