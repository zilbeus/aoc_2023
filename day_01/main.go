package main

import (
	"bufio"
	"fmt"
	"log"
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

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
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
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(line[i]) {
				lastDigitFound = true
				lastDigit = line[i]
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
