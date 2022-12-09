package cargo

import (
	"reflect"
	"testing"
)

func TestParseStacks(t *testing.T) {
	cases := []struct {
		desc  string
		input []string
		want  []Stack
	}{
		{
			"test",
			[]string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
			},
			[]Stack{
				{'Z', 'N'},
				{'M', 'C', 'D'},
				{'P'},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := ParseStacks(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Want: %v, Got: %v", tc.want, got)
			}
		})
	}
}
