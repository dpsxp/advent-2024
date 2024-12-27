package day2

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("./data/test.txt")
	expected := 2

	if result != expected {
		t.Fatalf("Expecting %v got %v instead", expected, result)
	}
}

func TestSolutionPart1(t *testing.T) {
	result := Part1("./data/input.txt")
	expected := 686

	if result != expected {
		t.Fatalf("Result %v different from expected %v", result, expected)
	}
}
