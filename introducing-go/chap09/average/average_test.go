package average

import (
	"testing"
)

type TestTargets struct {
	s        S
	expected float64
}

var testTargets = []TestTargets{
	{S{[]int{1, 2, 3, 4, 5}}, 3.0},
	{S{[]int{1, 1}}, 1.0},
	{S{[]int{-1, 1}}, 0.0},
	{S{[]int{}}, 0.0},
}

func TestAverage(t *testing.T) {
	for _, testTarget := range testTargets {
		calculated := testTarget.s.Average()
		expected := testTarget.expected

		if calculated != expected {
			t.Error(
				"For", testTarget.s,
				"Expected", expected,
				"Got", calculated,
			)
		}
	}
}
