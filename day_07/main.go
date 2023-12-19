package main

import (
	"bufio"
	"log"
	"os"
)

var cardStrengths = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAkind
	FiveOfAKind
)

func main() {
	fileName := "input_test.txt"
	file := getFile(fileName)
	defer file.Close()
	lines := getLines(file)
}

func getFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	return file
}

func getLines(file *os.File) []string {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getHandType(hand string) HandType {
	nrOfCardsInHand := map[rune]int{}
	for _, c := range hand {
		val, ok := nrOfCardsInHand[c]
		if !ok {
			val = 0
		}
		nrOfCardsInHand[c] = val + 1
	}

	var handType HandType = -1
	max := 0
	for _, v := range nrOfCardsInHand {
		if (v == 3 && max == 2) || (v == 2 && max == 3) {
			handType = FullHouse
		}

		if v == 2 && max == 2 {
			handType = TwoPair
		}

		if v == 4 {
			handType = FourOfAkind
		}

		if v == 5 {
			handType = FiveOfAKind
		}

		if v == 2 && max < v {
			handType = OnePair
		}

		if v == 3 && max < v-1 {
			handType = ThreeOfAKind
		}

		if max < v {
			max = v
		}
	}

	if max < 2 {
		handType = HighCard
	}

	return handType
}
