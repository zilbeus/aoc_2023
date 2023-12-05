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
	seedToSoil         map[int]int
	soilToFertilizer   map[int]int
	fertilizerToWater  map[int]int
	waterToLight       map[int]int
	lightToTemp        map[int]int
	tempToHumidity     map[int]int
	humidityToLocation map[int]int
}

func main() {
	file, err := os.Open("input.txt")
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
	location := findLowestLocationNr(seeds, mappings)

	fmt.Printf("LOWEST LOCATION NR: %d\n", location)
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

func findCategoryMapping(mapping string, input []string) map[int]int {
	categoryMap := map[int]int{}
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

		for i := 0; i < rangeLength; i++ {
			categoryMap[sourceRangeStart+i] = destinationRangeStart + i
		}
	}

	return categoryMap
}

func createMapToSeed(categoryMapping, seedToCategoryMapping map[int]int) map[int]int {
	seedToNewCategoryMapping := map[int]int{}
	for newCategoryKey, newCategoryVal := range categoryMapping {
		for prevCategoryKey, prevCategoryVal := range seedToCategoryMapping {
			if prevCategoryVal == newCategoryKey {
				seedToNewCategoryMapping[prevCategoryKey] = newCategoryVal
				break
			}
		}
	}
	return seedToNewCategoryMapping
}

func findLowestLocationNr(seeds []int, mappings *categoryMappings) int {
	locationMin := -1
	for _, s := range seeds {
		soil, exists := mappings.seedToSoil[s]
		if !exists {
			mappings.seedToSoil[s] = s
			soil = s
		}
		fertilizer, exists := mappings.soilToFertilizer[soil]
		if !exists {
			if s == 13 {
				fmt.Printf("SOIL: %d\n", soil)
			}
			mappings.soilToFertilizer[soil] = soil
			fertilizer = soil
		}
		water, exists := mappings.fertilizerToWater[fertilizer]
		if !exists {
			if s == 13 {
				fmt.Printf("FERTILIZER: %d\n", fertilizer)
			}
			mappings.fertilizerToWater[fertilizer] = fertilizer
			water = fertilizer
		}
		light, exists := mappings.waterToLight[water]
		if !exists {
			if s == 13 {
				fmt.Printf("WATER: %d\n", water)
			}
			mappings.waterToLight[water] = water
			light = water
		}
		temp, exists := mappings.lightToTemp[light]
		if !exists {
			if s == 13 {
				fmt.Printf("LIGHT: %d\n", light)
			}
			mappings.lightToTemp[light] = light
			temp = light
		}
		humidity, exists := mappings.tempToHumidity[temp]
		if !exists {
			if s == 13 {
				fmt.Printf("TEMP: %d\n", temp)
			}
			mappings.tempToHumidity[temp] = temp
			humidity = temp
		}
		location, exists := mappings.humidityToLocation[humidity]
		if !exists {
			if s == 13 {
				fmt.Printf("HUMIDITY: %d\n", humidity)
			}
			mappings.humidityToLocation[humidity] = humidity
			location = humidity
		}

		if locationMin < 0 || location < locationMin {
			fmt.Printf("SEED: %d, LOCATION: %d\n", s, location)
			locationMin = location
		}
		// fmt.Printf("MAPPINGS: %+v\n", mappings)

	}
	return locationMin
}
