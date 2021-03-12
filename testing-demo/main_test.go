package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(3) != 5 {
		t.Error("Excepted x + 2 is equal to 5")
	}
}
func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{5, 7},
		{-1, 1},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Expected {} + 2 = {} but getting {}", test.input, test.expected, output)
		}
	}
}
