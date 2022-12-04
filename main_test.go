package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if add(3, 5) != 8 {
		t.Errorf("add() = %v, want %v", add(3, 5), 8)
	}
}

func TestSub(t *testing.T) {
	if sub(4, 1) != 3 {
		t.Errorf("sub() = %v, want %v", add(4, 1), 3)
	}
}
