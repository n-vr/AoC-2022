package main

import "testing"

var input []string = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestOneContainsOther(t *testing.T) {
	cases := []struct {
		desc   string
		inputA *Range
		inputB *Range
		want   bool
	}{
		{"2-83-7Contain", NewRange("2-8"), NewRange("3-7"), true},
		{"6-64-6Contain", NewRange("6-6"), NewRange("4-6"), true},
		{"2-46-8NoContain", NewRange("2-4"), NewRange("6-8"), false},
		{"2-64-8NoContain", NewRange("2-6"), NewRange("4-8"), false},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := oneContainsOther(tc.inputA, tc.inputB)
			if tc.want != got {
				t.Errorf("Want: %v, Got: %v", tc.want, got)
			}
		})
	}
}

func TestOverlap(t *testing.T) {
	cases := []struct {
		desc   string
		inputA *Range
		inputB *Range
		want   bool
	}{
		{"1-33-5Overlap", NewRange("1-3"), NewRange("3-5"), true},
		{"1-22-5Overlap", NewRange("1-2"), NewRange("2-5"), true},
		{"1-43-5Overlap", NewRange("1-4"), NewRange("3-5"), true},
		{"1-23-5NoOverlap", NewRange("1-2"), NewRange("3-5"), false},
		{"1-34-5NoOverlap", NewRange("1-3"), NewRange("4-5"), false},
		{"1-45-9NoOverlap", NewRange("1-4"), NewRange("5-9"), false},
		{"3-51-3Overlap", NewRange("3-5"), NewRange("1-3"), true},
		{"2-51-2Overlap", NewRange("2-5"), NewRange("1-2"), true},
		{"3-51-4Overlap", NewRange("3-5"), NewRange("1-4"), true},
		{"3-51-2NoOverlap", NewRange("3-5"), NewRange("1-2"), false},
		{"4-51-3NoOverlap", NewRange("4-5"), NewRange("1-3"), false},
		{"5-91-4NoOverlap", NewRange("5-9"), NewRange("1-4"), false},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := overlap(tc.inputA, tc.inputB)
			if tc.want != got {
				t.Errorf("Want: %v, Got: %v", tc.want, got)
			}
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	want := "2"

	got, err := solvePartOne(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestSolvePartTwo(t *testing.T) {
	want := "4"

	got, err := solvePartTwo(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
