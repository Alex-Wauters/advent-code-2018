package main

import "testing"

func TestCookTime(t *testing.T) {
	tests := []struct {
		n string
		i int
	}{{"A", 61}, {"X", 84}, {"Z", 86}}
	for _, test := range tests {
		if c := cookTime(test.n); c != test.i {
			t.Errorf("Expected %v and got %v", test.i, c)
		}
	}
}
