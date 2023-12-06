package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input_test.txt"
	file := getFile(fileName)
	defer file.Close()
	lines := getLines(file)
	result := calculateErrorMarginMultiplication(lines)
	fmt.Printf("RESULT: %d\n", result)

	result = calculateErrorMarginMultiplicationForSingleRace(lines)
	fmt.Printf("RESULT: %d\n", result)
}

func calculateErrorMarginMultiplication(lines []string) int {
	raceTimesRow := lines[0]
	raceDistancesRow := lines[1]

	parts := strings.Split(raceTimesRow, ":")
	raceTimes := getNumbersFromInput(parts[1])

	parts = strings.Split(raceDistancesRow, ":")
	raceDistances := getNumbersFromInput(parts[1])

	result := 1
	for i := 0; i < len(raceTimes); i++ {
		time := raceTimes[i]
		recordDistance := raceDistances[i]
		x1 := int(math.Floor((float64(time) + math.Sqrt(float64(time*time-4*(recordDistance+1)))) / 2))
		x2 := int(math.Ceil((float64(time) - math.Sqrt(float64(time*time-4*(recordDistance+1)))) / 2))
		fmt.Printf("TIME: %d, RECORD DISTANCE: %d\n", time, recordDistance)
		fmt.Printf("X1: %d, X2: %d\n", x1, x2)
		result *= int(math.Abs(float64(x1-x2))) + 1
	}

	return result
}

func calculateErrorMarginMultiplicationForSingleRace(lines []string) int {
	raceTimesRow := lines[0]
	raceDistancesRow := lines[1]

	parts := strings.Split(raceTimesRow, ":")
	raceTime := getAsSingleNumberFromInput(parts[1])

	parts = strings.Split(raceDistancesRow, ":")
	raceDistance := getAsSingleNumberFromInput(parts[1])

	x1 := int(math.Floor((float64(raceTime) + math.Sqrt(float64(raceTime*raceTime-4*(raceDistance+1)))) / 2))
	x2 := int(math.Ceil((float64(raceTime) - math.Sqrt(float64(raceTime*raceTime-4*(raceDistance+1)))) / 2))
	fmt.Printf("raceTime: %d, RECORD DISTANCE: %d\n", raceTime, raceDistance)
	fmt.Printf("X1: %d, X2: %d\n", x1, x2)
	return int(math.Abs(float64(x1-x2))) + 1
}

func getLines(file *os.File) []string {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getNumbersFromInput(input string) []int {
	fields := strings.Fields(input)
	numbers := []int{}
	for _, v := range fields {
		nr, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("expected numeric input, got '%s'", v)
		}
		numbers = append(numbers, nr)
	}
	return numbers
}

func getAsSingleNumberFromInput(input string) int {
	fields := strings.Fields(input)
	inputNumber := strings.Join(fields, "")
	nr, err := strconv.Atoi(inputNumber)
	if err != nil {
		log.Fatalf("expected numeric input, got '%s'", inputNumber)
	}
	return nr
}

func getFile(fileName string) *os.File {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	return file
}
