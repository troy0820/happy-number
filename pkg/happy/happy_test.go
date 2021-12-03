package happy

import (
	"testing"
)

func TestProcessNumber(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"Thirteen 13:":          {13, 10},
		"Forty-Four 44":         {44, 32},
		"One 1":                 {1, 1},
		"Zero 0":                {0, 0},
		"One Thousand One 1001": {1001, 2},
	}
	for name, test := range tests {
		got := processNumber(test.input)
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if test.want != got {
				t.Fatalf("This is not equal.  Wanted %d got %d", test.want, got)
			}
		})
	}
}

func TestHappyNumber(t *testing.T) {
	tests := map[string]struct {
		input  int
		answer bool
	}{
		"Happy Number -7-":      {7, true},
		"Not Happy Number -4-":  {4, false},
		"Base Case -1-":         {1, true},
		"Before Algorithm -11-": {11, false},
		"Long Number -5432-":    {5423, false},
	}
	for name, test := range tests {
		test := test
		got := HappyNumber(test.input)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if test.answer != got {
				t.Fatalf("expected %#v, got %#v", test.answer, got)
			}
		})
	}
}
