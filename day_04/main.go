package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	defer file.Close()

	cardValues := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cardValues = append(cardValues, getCardValue(line))
	}

	sum := 0
	for _, v := range cardValues {
		sum += v
	}

	fmt.Printf("SUM OF CARD VALUES: %d\n", sum)
}

func getCardValue(input string) int {
	winningNrs := getWinningNumbers(input)
	givenNrs := getGivenNumbers(input)
	matchingNrs := findMatchingNumbers(winningNrs, givenNrs)
	return int(math.Pow(2.0, float64(len(matchingNrs)-1)))
}

func getWinningNumbers(input string) []int {
	numbers := []int{}
	parts := strings.Split(input, "|")
	if len(parts) != 2 {
		log.Fatalf("expected '|' separated input, got %s", input)
	}

	parts = strings.Split(parts[0], ":")
	if len(parts) != 2 {
		log.Fatalf("expected ':' separated input, got %s", parts[0])
	}

	inputNrs := strings.Fields(parts[1])
	for _, v := range inputNrs {
		nr, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("getWinningNumbers: expected numeric input, got %s", v)
		}
		numbers = append(numbers, nr)
	}

	return numbers
}

func getGivenNumbers(input string) []int {
	numbers := []int{}
	parts := strings.Split(input, "|")
	if len(parts) != 2 {
		log.Fatalf("expected '|' separated input, got %s", input)
	}

	inputNrs := strings.Fields(parts[1])
	fmt.Printf("inputNrs: %s\n", inputNrs)
	for _, v := range inputNrs {
		nr, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("getGivenNumbers: expected numeric input, got %s", v)
		}
		numbers = append(numbers, nr)
	}

	return numbers
}

func findMatchingNumbers(winningNrs, givenNrs []int) []int {
	matchingNrs := []int{}
	for _, v := range givenNrs {
		if slices.Contains(winningNrs, v) {
			matchingNrs = append(matchingNrs, v)
		}
	}
	fmt.Printf("matchingNrs: %+v\n", matchingNrs)
	return matchingNrs
}
