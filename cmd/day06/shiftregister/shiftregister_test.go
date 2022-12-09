package shiftregister

import (
	"reflect"
	"strconv"
	"testing"
)

func TestShiftIn(t *testing.T) {
	register := New(2)

	register.ShiftIn('a')
	want := []rune{'a'}
	if !reflect.DeepEqual(register.data, want) {
		t.Errorf("want: %v, got: %v", want, register.data)
	}

	register.ShiftIn('b')
	want = []rune{'a', 'b'}
	if !reflect.DeepEqual(register.data, want) {
		t.Errorf("want: %v, got: %v", want, register.data)
	}

	register.ShiftIn('c')
	want = []rune{'b', 'c'}
	if !reflect.DeepEqual(register.data, want) {
		t.Errorf("want: %v, got: %v", want, register.data)
	}
}

func TestValuesUnique(t *testing.T) {
	cases := []struct {
		desc      string
		inputVals string
		inputLen  int
		want      bool
	}{
		{"SingleCharInLen1IsTrue", "a", 1, true},
		{"SingleCharInLen2IsTrue", "a", 2, true},
		{"DifferentCharsInLen2IsTrue", "ab", 2, true},
		{"SameCharsInLen2IsTrue", "aa", 2, false},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reg := New(tc.inputLen)
			for _, ch := range tc.inputVals {
				reg.ShiftIn(ch)
			}

			got := reg.ValuesUnique()
			if got != tc.want {
				t.Errorf("want: %s, got: %s", strconv.FormatBool(tc.want), strconv.FormatBool(got))
			}
		})
	}
}

func TestFullAndValuesUnique(t *testing.T) {
	cases := []struct {
		desc      string
		inputVals string
		inputLen  int
		want      bool
	}{
		{"SingleCharInLen1IsTrue", "a", 1, true},
		{"SingleCharInLen2IsFalse", "a", 2, false},
		{"DifferentCharsInLen2IsTrue", "ab", 2, true},
		{"SameCharsInLen2IsTrue", "aa", 2, false},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reg := New(tc.inputLen)
			for _, ch := range tc.inputVals {
				reg.ShiftIn(ch)
			}

			got := reg.FullAndValuesUnique()
			if got != tc.want {
				t.Errorf("want: %s, got: %s", strconv.FormatBool(tc.want), strconv.FormatBool(got))
			}
		})
	}
}
