package main

import (
	"errors"
	"log"
	"sort"
	"strconv"

	aoc "github.com/n-vr/AoC-2022"
)

func main() {
	log.SetFlags(0)

	challengeOne := aoc.NewChallenge(1, 1, solvePartOne)

	err := challengeOne.Solve()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(challengeOne)

	challengeTwo := aoc.NewChallenge(1, 2, solvePartTwo)

	err = challengeTwo.Solve()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("-----\n", challengeTwo)
}

func solvePartOne(input []string) (string, error) {
	elvesCalories, err := countAndSortCalories(input)
	if err != nil {
		return "error", err
	}

	mostCalories := elvesCalories[len(elvesCalories)-1]

	return strconv.Itoa(mostCalories), nil
}

func solvePartTwo(input []string) (string, error) {
	elvesCalories, err := countAndSortCalories(input)
	if err != nil {
		return "error", err
	}

	if len(elvesCalories) < 3 {
		return "error", errors.New("too few elves")
	}

	sumOfThree := elvesCalories[len(elvesCalories)-1] + elvesCalories[len(elvesCalories)-2] + elvesCalories[len(elvesCalories)-3]

	return strconv.Itoa(sumOfThree), nil
}

func countAndSortCalories(input []string) ([]int, error) {
	elvesCalories := make([]int, 0)

	var currentSum int

	for _, line := range input {
		if line == "" {
			// Empty line means this elve's calories ended.
			elvesCalories = append(elvesCalories, currentSum)
			currentSum = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return []int{}, err
		}
		currentSum += calories
	}
	// Always make sure to append the last sum.
	elvesCalories = append(elvesCalories, currentSum)

	sort.Slice(elvesCalories, func(i, j int) bool {
		return elvesCalories[i] < elvesCalories[j]
	})

	return elvesCalories, nil
}
