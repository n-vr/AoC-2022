package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	want := "15"

	got, err := solvePartOne(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	want := "12"

	got, err := solvePartTwo(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
