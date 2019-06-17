package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappyNumber(t *testing.T) {
	tests := map[string]struct {
		input  int
		answer bool
	}{
		"Happy Number -7-":      {7, true},
		"Not Happy Number -4-":  {4, false},
		"Base Case -1-":         {1, true},
		"Before Algorithm -11-": {11, true},
		"Long Number -5432-":    {5423, true},
	}
	for name, test := range tests {
		got := HappyNumber(test.input)
		t.Run(name, func(t *testing.T) {
			if test.answer != got {
				t.Fatalf("expected %#v, got %#v", test.answer, got)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input  int
		answer []int
	}{
		"Standard Number 24": {24, []int{2, 4}},
		"Double Number 11":   {11, []int{1, 1}},
		"Single number 1":    {1, []int{1}},
	}

	for name, test := range tests {
		got := Parse(test.input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.answer, got, "Results are not equal")
		})

	}
}
