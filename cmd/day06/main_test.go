package main

import "testing"

var input []string = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
}

func TestFindConsecutiveUniqueCharsPos(t *testing.T) {
	cases := []struct {
		desc     string
		inputStr string
		inputLen int
		want     int
	}{
		{"", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"", "bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"", "nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := findConsecutiveUniqueCharsPos(tc.inputStr, tc.inputLen)
			if tc.want != got {
				t.Errorf("Want: %d, Got: %d", tc.want, got)
			}
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	want := "7"

	got, err := solvePartOne(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestSolvePartTwo(t *testing.T) {
	want := "19"

	got, err := solvePartTwo(input)
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
