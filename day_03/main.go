package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

var symbols = []rune{'*', '#', '$', '+'}

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

	log.Printf("SUM: %d", sum)
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
	fmt.Printf("\n---------------------------------------------\n")
	for i, v := range line {
		if unicode.IsDigit(v) || v == '.' {
			continue
		}

		fmt.Printf("symbol: %s\n", string(v))

		if lineIdx > 0 {
			fmt.Printf("line: %s\n", lines[lineIdx-1])
			numbers = append(numbers, getNumbers(lines[lineIdx-1], i)...)
		}

		fmt.Printf("line: %s\n", lines[lineIdx])
		numbers = append(numbers, getNumbers(line, i)...)

		if lineIdx < len(lines)-1 {
			fmt.Printf("line: %s\n", lines[lineIdx+1])
			numbers = append(numbers, getNumbers(lines[lineIdx+1], i)...)
		}

		log.Printf("getSymbolAdjacentNumbers - numbers: %+v, line: %d\n\n", numbers, lineIdx)
	}
	fmt.Printf("---------------------------------------------\n\n\n\n")

	return numbers
}

func getNumbers(line string, symbolIdx int) []int {
	// log.Printf("getNumbers - symbolIdx: %d, line: %s\n", symbolIdx, line)
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

	// log.Printf("getNumbers - numbers: %+v\n", numbers)
	return numbers
}

func isNrAdjacentToSymbol(nrStartIdx, nrEndIdx, symbolIdx int) bool {
	symbolAdjacentToNrStartIdx := math.Abs(float64(nrStartIdx)-float64(symbolIdx)) <= 1.0
	symbolAdjacentToNrEndIdx := math.Abs(float64(nrEndIdx)-float64(symbolIdx)) <= 1.0
	symbolBetweenStartAndEndIdx := symbolIdx > nrStartIdx && symbolIdx < nrEndIdx
	return symbolAdjacentToNrStartIdx || symbolAdjacentToNrEndIdx || symbolBetweenStartAndEndIdx
}
