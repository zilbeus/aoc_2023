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

const (
	HighCard = iota
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
