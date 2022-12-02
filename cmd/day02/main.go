package main

import (
	"log"
	"strconv"

	aoc "github.com/n-vr/AoC-2022"
	"github.com/n-vr/AoC-2022/cmd/day02/rps"
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
	var sum int

	for _, line := range input {
		opponent := rps.NewShape(rune(line[0]))
		me := rps.NewShape(rune(line[2]))

		sum += rps.RoundScore(*opponent, *me)
	}

	return strconv.Itoa(sum), nil
}

func solvePartTwo(input []string) (string, error) {
	var sum int

	for _, line := range input {
		opponent := rps.NewShape(rune(line[0]))
		me := rps.ChooseShapeForOutcome(*opponent, rps.Outcome(rune(line[2])))

		sum += rps.RoundScore(*opponent, me)
	}

	return strconv.Itoa(sum), nil
}
