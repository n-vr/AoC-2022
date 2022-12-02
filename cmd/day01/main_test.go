package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	input := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}

	want := "24000"

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
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}

	want := "45000"

	got, err := solvePartTwo(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
