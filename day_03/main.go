package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	defer file.Close()

	numbers := []int{}
	lines := readLines(file)
	for i := range lines {
		numbers = append(numbers, getSymbolAdjacentNumbers(lines, i)...)
	}

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	log.Printf("SUM: %d\n", sum)

	gearRatios := []int{}
	for i := range lines {
		gearRatios = append(gearRatios, getGearRatios(lines, i)...)
	}

	sum = 0
	for _, v := range gearRatios {
		sum += v
	}
	log.Printf("GEAR RATIO SUM: %d\n", sum)
}

func readLines(f *os.File) []string {
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getSymbolAdjacentNumbers(lines []string, lineIdx int) []int {
	numbers := []int{}
	line := lines[lineIdx]
	for i, v := range line {
		if unicode.IsDigit(v) || v == '.' {
			continue
		}

		if lineIdx > 0 {
			numbers = append(numbers, getNumbers(lines[lineIdx-1], i)...)
		}

		numbers = append(numbers, getNumbers(line, i)...)

		if lineIdx < len(lines)-1 {
			numbers = append(numbers, getNumbers(lines[lineIdx+1], i)...)
		}
	}

	return numbers
}

func getNumbers(line string, symbolIdx int) []int {
	numbers := []int{}
	var currentNrStartIdx int
	var currentNr string

	for i, v := range line {
		if unicode.IsDigit(v) {
			if len(currentNr) == 0 {
				currentNrStartIdx = i
			}
			currentNr += string(v)
		} else if len(currentNr) != 0 {
			if isNrAdjacentToSymbol(currentNrStartIdx, i-1, symbolIdx) {
				nr, err := strconv.Atoi(currentNr)
				if err != nil {
					log.Fatalf("getNumbers - should be able to convert to int, currentNr: %s", currentNr)
				}
				numbers = append(numbers, nr)
			}
			currentNr = ""
		}
	}

	if len(currentNr) != 0 {
		if isNrAdjacentToSymbol(currentNrStartIdx, len(line)-1, symbolIdx) {
			nr, err := strconv.Atoi(currentNr)
			if err != nil {
				log.Fatalf("getNumbers - should be able to convert to int, currentNr: %s", currentNr)
			}
			numbers = append(numbers, nr)
		}
	}

	return numbers
}

func isNrAdjacentToSymbol(nrStartIdx, nrEndIdx, symbolIdx int) bool {
	symbolAdjacentToNrStartIdx := math.Abs(float64(nrStartIdx)-float64(symbolIdx)) <= 1.0
	symbolAdjacentToNrEndIdx := math.Abs(float64(nrEndIdx)-float64(symbolIdx)) <= 1.0
	symbolBetweenStartAndEndIdx := symbolIdx > nrStartIdx && symbolIdx < nrEndIdx
	return symbolAdjacentToNrStartIdx || symbolAdjacentToNrEndIdx || symbolBetweenStartAndEndIdx
}

func getGearRatios(lines []string, lineIdx int) []int {
	ratios := []int{}
	line := lines[lineIdx]

	for i, v := range line {
		if v != '*' {
			continue
		}

		partNumbers := []int{}

		if lineIdx > 0 {
			partNumbers = append(partNumbers, getNumbers(lines[lineIdx-1], i)...)
		}

		partNumbers = append(partNumbers, getNumbers(line, i)...)

		if lineIdx < len(lines)-1 {
			partNumbers = append(partNumbers, getNumbers(lines[lineIdx+1], i)...)
		}

		if len(partNumbers) == 2 {
			ratios = append(ratios, partNumbers[0]*partNumbers[1])
		}
	}

	return ratios
}
