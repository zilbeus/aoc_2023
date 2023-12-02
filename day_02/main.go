package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id       int
	cubeSets []CubeSet
}

type CubeSet struct {
	cubes []Cubes
}

type Cubes struct {
	amount int
	color  string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't find file.")
	}
	defer file.Close()

	games := []Game{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := Game{}
		line := scanner.Text()
		game.id = getGameId(line)
		game.cubeSets = getCubeSets(line)
		games = append(games, game)
	}

	cubesInBag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  15,
	}

	possibleGames := findPossibleGames(games, cubesInBag)
	sum := 0
	for _, game := range possibleGames {
		sum += game.id
	}

	fmt.Printf("SUM FOR POSSIBLE GAME IDS: %d\n", sum)

	cubeSetPowers := findPowerOfCubeSetsForEachGame(games)
	sum = 0
	for _, v := range cubeSetPowers {
		sum += v
	}

	fmt.Printf("SUM OF POWER SETS: %d\n", sum)
}

func getGameId(input string) int {
	result := strings.Split(input, ":")
	if len(result) < 2 {
		log.Fatal("Invalid input:", input)
	}

	result = strings.Split(result[0], " ")
	if len(result) < 2 {
		log.Fatal("Invalid game ID:", result[0])
	}

	id, err := strconv.Atoi(result[1])

	if err != nil {
		log.Fatal("Game ID should be numeric:", result[1])
	}

	return id
}

func getCubeSets(input string) []CubeSet {
	result := strings.Split(input, ":")
	if len(result) < 2 {
		log.Fatal("Invalid input:", input)
	}

	sets := strings.Split(result[1], ";")
	if len(sets) < 1 {
		log.Fatal("invalid nr of sets in input (3 sets required): ", sets)
	}

	cubeSets := []CubeSet{}
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		if len(cubes) < 1 {
			log.Fatal("Invalid comma-separated set: ", set)
		}

		cubeSet := CubeSet{}
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			nrAndColor := strings.Split(cube, " ")
			if len(nrAndColor) != 2 {
				log.Fatal("Invalid cube (expected amount and color): ", cube)
			}

			quantity, err := strconv.Atoi(nrAndColor[0])
			if err != nil {
				log.Fatal("Cube quantity should be numeric: ", cube)
			}

			cubeSet.cubes = append(cubeSet.cubes, Cubes{amount: quantity, color: nrAndColor[1]})
		}
		cubeSets = append(cubeSets, cubeSet)
	}

	return cubeSets
}

func findPossibleGames(games []Game, cubesInBag map[string]int) []Game {
	possibleGames := []Game{}
	for _, game := range games {
		if isGamePossible(game.cubeSets, cubesInBag) {
			possibleGames = append(possibleGames, game)
		}
	}

	return possibleGames
}

func isGamePossible(cubeSets []CubeSet, cubesInBag map[string]int) bool {
	for _, cubeSet := range cubeSets {
		for _, cubes := range cubeSet.cubes {
			nrOfCubesInBag := cubesInBag[cubes.color]
			if nrOfCubesInBag < cubes.amount {
				return false
			}
		}

	}

	return true
}

func findPowerOfCubeSetsForEachGame(games []Game) map[int]int {
	powers := map[int]int{}
	for _, game := range games {
		powers[game.id] = findPowerOfGameCubes(game)
	}
	return powers
}

func findPowerOfGameCubes(game Game) int {
	maxQtys := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, cubeSets := range game.cubeSets {
		for _, cubes := range cubeSets.cubes {
			currentMaxQty := maxQtys[cubes.color]
			if currentMaxQty < cubes.amount {
				maxQtys[cubes.color] = cubes.amount
			}
		}
	}

	power := 1
	for _, v := range maxQtys {
		power *= v
	}

	return power
}
