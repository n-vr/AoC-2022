package main

import (
	"log"
	"strconv"

	aoc "github.com/n-vr/AoC-2022"
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
	var sum int

	// Find common item between three elves.
	for _, line := range input {

		left := line[:len(line)/2]
		right := line[len(line)/2:]

		leftItems := findTypes(left)
		rightItems := findTypes(right)

		// Find the duplicates.
		duplicate := leftItems & rightItems

		// Calculate the score from duplicates.
		for i := 0; i < 52; i++ {
			if duplicate&(1<<i) != 0 {
				// Remember the bitmap is shifted with priority-1.
				sum += i + 1
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func solvePartTwo(input []string) (string, error) {
	var sum int

	// Range over groups of three elves.
	for i := 0; i < len(input); i += 3 {
		elveOneItems := findTypes(input[i])
		elveTwoItems := findTypes(input[i+1])
		elveThreeItems := findTypes(input[i+2])

		commonItem := elveOneItems & elveThreeItems & elveTwoItems

		// Get score for the set bit.
		for i := 0; i < 52; i++ {
			if commonItem&(1<<i) != 0 {
				// Remember the bitmap is shifted with priority-1.
				sum += i + 1
			}
		}
	}

	return strconv.Itoa(sum), nil
}

// findTypes finds all types in a sequence
// and returns a bitmap of the found types, shifted with their priority-1.
func findTypes(sequence string) ItemTypeBitmap {
	var items ItemTypeBitmap

	for _, ch := range sequence {
		if ch >= 'a' && ch <= 'z' {
			items |= 1 << (ItemTypeBitmap(ch) - 'a')
		}

		if ch >= 'A' && ch <= 'Z' {
			const minShiftForUppercase = 26
			items |= 1 << (minShiftForUppercase + ItemTypeBitmap(ch) - 'A')
		}
	}

	return items
}
