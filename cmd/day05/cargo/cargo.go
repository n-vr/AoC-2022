// Package cargo implements a stack of crates,
// as described in advent of code day 5.
package cargo

import (
	"fmt"
	"unicode"
)

// Crate is a box of cargo that can be moved.
type Crate rune

// Stack can have multiple crates which can move between stacks.
type Stack []Crate

// ParseStacks parses input into stack representations.
func ParseStacks(input []string) []Stack {
	stacks := make([]Stack, len(input[0])/4+1)

	for _, line := range input {
		for i, ch := range line {
			if unicode.IsLetter(ch) {
				stackNum := i / 4
				stacks[stackNum] = append(Stack{Crate(ch)}, stacks[stackNum]...)
			}
		}
	}

	return stacks
}

// Instruction encapsulates a command of moving an x amount of crates
// from one stack to another.
type Instruction struct {
	Amount int
	From   int
	To     int
}

// Create a new move from a sentence.
func NewInstruction(sentence string) *Instruction {
	move := &Instruction{}
	fmt.Sscanf(sentence, "move %d from %d to %d", &move.Amount, &move.From, &move.To)
	move.From -= 1
	move.To -= 1
	return move
}

// Parse instructions from input and output a slice of [Instruction].
func ParseInstructions(input []string) (instructions []*Instruction) {
	for _, line := range input {
		instructions = append(instructions, NewInstruction(line))
	}
	return
}

// Crane implements a crane that can move crates.
type Crane interface {
	Move(stacks []Stack, instruction *Instruction) []Stack
}

// CrateMover9000 implements a crane.
type CrateMover9000 struct{}

// Move crates between stacks, based on instructions and return the final result.
func (c *CrateMover9000) Move(stacks []Stack, instruction *Instruction) []Stack {
	for i := 0; i < instruction.Amount; i++ {
		stacks[instruction.To] = append(stacks[instruction.To], stacks[instruction.From][len(stacks[instruction.From])-1])
		stacks[instruction.From] = stacks[instruction.From][:len(stacks[instruction.From])-1]
	}

	return stacks
}

// CrateMover9000 implements a crane.
type CrateMover9001 struct{}

// Move crates between stacks, based on instructions and return the final result.
func (c *CrateMover9001) Move(stacks []Stack, instruction *Instruction) []Stack {
	stacks[instruction.To] = append(stacks[instruction.To], stacks[instruction.From][len(stacks[instruction.From])-instruction.Amount:]...)
	stacks[instruction.From] = stacks[instruction.From][:len(stacks[instruction.From])-instruction.Amount]

	return stacks
}

// Execute all instructions using the crane, instructions and stacks.
func DoAllMoves(crane Crane, stacks []Stack, instructions []*Instruction) []Stack {
	newStacks := stacks
	for _, instruction := range instructions {
		newStacks = crane.Move(stacks, instruction)
	}
	return newStacks
}

// Split input into stacks lines that can be parsed into stacks and instructions.
func SplitStacksAndInstructions(input []string) ([]Stack, []*Instruction) {
	var stacksSection []string
	instructionsSection := input

	for i, line := range input {
		if line != "" {
			stacksSection = append(stacksSection, line)
		} else {
			instructionsSection = instructionsSection[i+1:]
			break
		}
	}

	return ParseStacks(stacksSection), ParseInstructions(instructionsSection)
}

func TopCratesLetters(stacks []Stack) (answer string) {
	for _, stack := range stacks {
		answer += string(stack[len(stack)-1])
	}
	return
}
