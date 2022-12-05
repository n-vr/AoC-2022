package main

import (
	"errors"
	"log"
	"strconv"
	"strings"

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

	for _, line := range input {
		split := strings.Split(line, ",")
		if len(split) < 2 {
			return "", errors.New("input invalid")
		}

		a, b := NewRange(split[0]), NewRange(split[1])

		if oneContainsOther(a, b) {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

func solvePartTwo(input []string) (string, error) {
	var sum int

	for _, line := range input {
		split := strings.Split(line, ",")
		if len(split) < 2 {
			return "", errors.New("input invalid")
		}

		a, b := NewRange(split[0]), NewRange(split[1])

		if overlap(a, b) {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

// Range implements a range between two numbers.
type Range struct {
	From int
	To   int
}

// NewRange creates a new range from a string line "1-3".
// Will panic if rng is invalid.
func NewRange(rng string) *Range {
	split := strings.Split(rng, "-")
	if len(split) < 2 {
		panic("rng string invalid")
	}

	from, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	to, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	return &Range{
		From: from,
		To:   to,
	}
}

// oneContainsOther return true if
// a fully contains b or b fully contains a.
func oneContainsOther(a, b *Range) bool {
	if a.From <= b.From && a.To >= b.To {
		// A fully contains B.
		return true
	}
	if a.From >= b.From && a.To <= b.To {
		// B fully contains A.
		return true
	}
	return false
}

func overlap(a, b *Range) bool {
	if (a.To >= b.From && b.From <= a.To && a.From <= b.From) || (b.To >= a.From && a.From <= b.To && b.From <= a.From) {
		return true
	}
	return false
}
