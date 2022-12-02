// Package aoc2022 implements common abstractions for solving Advent of Code 2022 challenges.
package aoc2022

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Solver can solve a challenge.
type Solver interface {
	Solve() error
}

// Challenge represents a challenge and implements the [Solver] interface.
type Challenge struct {
	Day        int
	Part       int
	InputFile  string
	SolverFunc SolverFunc
	Answer     string
}

// SolverFunc implements a function that can solve a challenge.
type SolverFunc = func(input []string) (output string, err error)

// NewChallenge creates a new challenge.
func NewChallenge(day, part int, solverFunc SolverFunc) *Challenge {
	if len(os.Args) < 2 {
		log.Fatalf("inputFile is not defined. please run the same command, but append the filename: <command> <filename>")
	}

	return &Challenge{
		Day:        day,
		Part:       part,
		InputFile:  os.Args[1],
		SolverFunc: solverFunc,
		Answer:     "unsolved",
	}
}

// Solve the challenge using the input.
func (c *Challenge) Solve() (err error) {
	input, err := getInputFromFile(c.InputFile)
	if err != nil {
		return err
	}

	c.Answer, err = c.SolverFunc(input)
	if err != nil {
		return err
	}

	return nil
}

func (c *Challenge) String() string {
	return fmt.Sprintf(`Challenge: 	day %d part %d
InputFile: 	%s
Answer: 	%s`, c.Day, c.Part, c.InputFile, c.Answer)
}

// getInputFromFile reads an entire file and returns it as a slice of strings
// which represent the lines in the file.
func getInputFromFile(fileName string) ([]string, error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(contents), "\n"), nil
}
