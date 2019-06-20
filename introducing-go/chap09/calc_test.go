package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	retval := Add(1, 2)
	expect := 3

	if retval != expect {
		t.Error("Expected =", expect, "Got :", retval)
	}
}

func TestSub(t *testing.T) {
	retval := Sub(5, 3)
	expect := 2

	if retval != expect {
		t.Error("Error!")
	}
}
