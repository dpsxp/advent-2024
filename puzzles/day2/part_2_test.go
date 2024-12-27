package day2

import "testing"

func TestPart2(t *testing.T) {
	result := Part2("./data/test.txt")
	expected := 4

	if result != expected {
		t.Fatalf("Expecting %v got %v instead", expected, result)
	}
}

func TestSolutionPart2(t *testing.T) {
	result := Part2("./data/input.txt")
	expected := 717

	if result != expected {
		t.Fatalf("Result %v different from expected %v", result, expected)
	}
}
