package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := Part1("./data/test.txt")
	expected := 11

	if result != expected {
		t.Fatalf("Result %v different from expected %v", result, expected)
	}
}
