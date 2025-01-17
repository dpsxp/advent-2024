package day1

import (
	"testing"
)

func TestPart2(t *testing.T) {
	result := Part2("./data/test.txt")
	expected := 31

	if result != expected {
		t.Fatalf("Result %v different from expected %v", result, expected)
	}
}

func TestSolutionPart2(t *testing.T) {
	result := Part2("./data/input.txt")
	expected := 21024792

	if result != expected {
		t.Fatalf("Result %v different from expected %v", result, expected)
	}
}
