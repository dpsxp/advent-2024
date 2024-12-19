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
