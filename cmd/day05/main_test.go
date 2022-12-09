package main

import "testing"

var input []string = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestSolvePartOne(t *testing.T) {
	want := "CMZ"

	got, err := solvePartOne(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestSolvePartTwo(t *testing.T) {
	want := "MCD"

	got, err := solvePartTwo(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
