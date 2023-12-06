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
	fileName := "input.txt"
	result := calculateErrorMarginMultiplication(fileName)
	fmt.Printf("RESULT: %d", result)
}

func calculateErrorMarginMultiplication(fileName string) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	raceTimesRow := lines[0]
	raceDistancesRow := lines[1]

	parts := strings.Split(raceTimesRow, ":")
	timeFields := strings.Fields(parts[1])
	raceTimes := []int{}
	for _, v := range timeFields {
		nr, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("expected numeric input, got '%s'", v)
		}
		raceTimes = append(raceTimes, nr)
	}

	parts = strings.Split(raceDistancesRow, ":")
	distanceFields := strings.Fields(parts[1])
	raceDistances := []int{}
	for _, v := range distanceFields {
		nr, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("expected numeric input, got '%s'", v)
		}
		raceDistances = append(raceDistances, nr)
	}

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
