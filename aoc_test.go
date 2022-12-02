package aoc2022

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	challenge := Challenge{
		Day:       0,
		Part:      0,
		InputFile: "./testdata/test_input.txt",
		SolverFunc: func(input []string) (string, error) {
			var sum int
			for _, line := range input {
				n, err := strconv.Atoi(line)
				if err != nil {
					return "", err
				}
				sum += n
			}
			return strconv.Itoa(sum), nil
		},
	}

	want := "15"

	err := challenge.Solve()
	if err != nil {
		t.Error(err)
	}

	got := challenge.Answer

	if want != got {
		t.Errorf("Want: %v, Got: %v\n", want, got)
	}
}

func TestString(t *testing.T) {
	challenge := Challenge{
		Day:       30,
		Part:      45,
		InputFile: "testfile",
		Answer:    "42",
	}

	want := `Challenge: 	day 30 part 45
InputFile: 	testfile
Answer: 	42`

	got := challenge.String()

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestGetInputFromFile(t *testing.T) {
	want := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	got, err := getInputFromFile("./testdata/test_input.txt")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v, Got: %v\n", want, got)
	}
}
