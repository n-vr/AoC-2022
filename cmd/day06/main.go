package main

import (
	"log"
	"strconv"

	aoc "github.com/n-vr/AoC-2022"
	"github.com/n-vr/AoC-2022/cmd/day06/shiftregister"
)

type ItemTypeBitmap uint64

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
	return strconv.Itoa(findConsecutiveUniqueCharsPos(input[0], 4)), nil
}

func solvePartTwo(input []string) (string, error) {
	return strconv.Itoa(findConsecutiveUniqueCharsPos(input[0], 14)), nil
}

// Find position of the first occurance where a consecutive number (length) of chars are unique
// in a string of random characters.
// Returns the position when found, or -1 if not found.
func findConsecutiveUniqueCharsPos(input string, length int) int {
	register := shiftregister.New(length)

	for i, ch := range input {
		register.ShiftIn(ch)
		if register.FullAndValuesUnique() {
			return i + 1
		}
	}

	return -1
}
