package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 1 + 2
	actual := Addition(1, 2)
	if Addition(1, 2) != 3 {
		t.Errorf("wrong test, expected : %d, actual : %d", expected, actual)
	}
	if Addition(-1, -2) != -3 {
		t.Error("Wrong")
	}
}

func TestSub(t *testing.T) {
	if Subtraction(1, 2) != -1 {
		t.Error("123")
	}
}

func TestDiv(t *testing.T) {
	if Division(2, 2) != 1 {
		t.Error("123")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(1, 2) != 2 {
		t.Error("123")
	}
}
