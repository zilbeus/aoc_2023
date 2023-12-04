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

	nrOfMatches := map[int]int{}
	cardValues := []int{}
	cards := getCards(file)

	for i, c := range cards {
		matchingNrs := findMatchingNumbers(c)
		cardValues = append(cardValues, getCardValue(matchingNrs))
		nrOfMatches[i] = len(matchingNrs)
	}

	sum := 0
	for _, v := range cardValues {
		sum += v
	}

	fmt.Printf("SUM OF CARD VALUES: %d\n", sum)

	nrOfCards := len(cards)
	for i := range nrOfMatches {
		nrOfCards += calculateNrOfCopiesReceived(i, nrOfMatches)
	}

	fmt.Printf("NR OF CARDS WITH COPIES: %d\n", nrOfCards)
}

func calculateNrOfCopiesReceived(cardNr int, nrOfMatches map[int]int) int {
	matches := nrOfMatches[cardNr]
	if matches == 0 {
		return 0
	}

	nrOfCopies := int(math.Min(float64(matches),
		float64(len(nrOfMatches)-cardNr)))

	if nrOfCopies < 1 {
		return 0
	}

	for i := 0; i < matches; i++ {
		nrOfCopies += calculateNrOfCopiesReceived((cardNr + i + 1), nrOfMatches)
	}

	return nrOfCopies
}

func getCards(f *os.File) []string {
	cards := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}
	return cards
}

func getCardValue(matchingNrs []int) int {
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

func findMatchingNumbers(input string) []int {
	winningNrs := getWinningNumbers(input)
	givenNrs := getGivenNumbers(input)
	matchingNrs := []int{}
	for _, v := range givenNrs {
		if slices.Contains(winningNrs, v) {
			matchingNrs = append(matchingNrs, v)
		}
	}
	fmt.Printf("matchingNrs: %+v\n", matchingNrs)
	return matchingNrs
}
