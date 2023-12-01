package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type StringDigit struct {
	Word  string
	Value string
}

var DigitStrings []StringDigit = []StringDigit{
	{
		Word:  "one",
		Value: "1",
	},
	{
		Word:  "two",
		Value: "2",
	},
	{
		Word:  "three",
		Value: "3",
	},
	{
		Word:  "four",
		Value: "4",
	},
	{
		Word:  "five",
		Value: "5",
	},
	{
		Word:  "six",
		Value: "6",
	},
	{
		Word:  "seven",
		Value: "7",
	},
	{
		Word:  "eight",
		Value: "8",
	},
	{
		Word:  "nine",
		Value: "9",
	},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := []rune(text)
		var firstDigitFound bool
		var lastDigitFound bool
		var firstDigit rune
		var lastDigit rune
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(line[i]) {
				firstDigitFound = true
				firstDigit = line[i]
				break
			}
			if isStringDigit(string(line[i:])) {
				firstDigitFound = true
				firstDigit = rune(getStringDigitValue(string(line[i:])))
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(line[i]) {
				lastDigitFound = true
				lastDigit = line[i]
				break
			}
			if isStringDigit(string(line[i:])) {
				lastDigitFound = true
				lastDigit = rune(getStringDigitValue(string(line[i:])))
				break
			}
		}

		if !firstDigitFound || !lastDigitFound {
			continue
		}

		number := string(firstDigit) + string(lastDigit)
		value, err := strconv.Atoi(number)

		if err != nil {
			log.Fatal("Not a number: " + number)
		}

		sum += value
	}

	fmt.Printf("SUM: %d", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isStringDigit(input string) bool {
	for _, el := range DigitStrings {
		if strings.HasPrefix(input, el.Word) {
			return true
		}
	}

	return false
}

func getStringDigitValue(input string) rune {
	for _, el := range DigitStrings {
		if strings.HasPrefix(input, el.Word) {
			return []rune(el.Value)[0]
		}
	}

	log.Fatal("Should have found a value")
	return 0
}
