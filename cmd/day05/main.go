package main

import (
	"log"

	aoc "github.com/n-vr/AoC-2022"
	"github.com/n-vr/AoC-2022/cmd/day05/cargo"
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
	stacks, instructions := cargo.SplitStacksAndInstructions(input)

	crane := &cargo.CrateMover9000{}
	stacks = cargo.DoAllMoves(crane, stacks, instructions)

	return cargo.TopCratesLetters(stacks), nil
}

func solvePartTwo(input []string) (string, error) {
	stacks, instructions := cargo.SplitStacksAndInstructions(input)

	crane := &cargo.CrateMover9001{}
	stacks = cargo.DoAllMoves(crane, stacks, instructions)

	return cargo.TopCratesLetters(stacks), nil
}
