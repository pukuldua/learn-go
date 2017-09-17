package fib

import (
	"testing"
)

type fibTest struct {
	n        int
	expected int
}

func TestFib(t *testing.T) {

	var fibTest = []fibTest {
		{1, 1},
		{2, 1}, 
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tt := range fibTest {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
