package main

import "testing"

func TestAdd(t *testing.T) {
	if add(2, 3) != 5 {
		t.Errorf("Expected 5, got %d", add(2, 3))
	}
}

func TestAddNegative(t *testing.T) {
	if add(-1, 1) != 0 {
		t.Errorf("Expected 0, got %d", add(-1, 1))
	}
}
