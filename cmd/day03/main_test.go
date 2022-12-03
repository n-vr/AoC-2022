package main

import "testing"

var input []string = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestFindTypes(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		want  ItemTypeBitmap
	}{
		{"LowercaseA", "a", ItemTypeBitmap(1 << 0)},
		{"LowercaseZ", "z", ItemTypeBitmap(1 << 25)},
		{"UppercaseA", "A", ItemTypeBitmap(1 << 26)},
		{"UppercaseZ", "Z", ItemTypeBitmap(1 << 51)},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := findTypes(tc.input)
			if tc.want != got {
				t.Errorf("Want: %d, Got: %d", tc.want, got)
			}
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	want := "157"

	got, err := solvePartOne(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestSolvePartTwo(t *testing.T) {
	want := "70"

	got, err := solvePartTwo(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
