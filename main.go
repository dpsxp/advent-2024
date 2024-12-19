package main

import (
	"advent/puzzles/day1"
	"fmt"
)

const SEPARATOR = "------"

func main() {
	fmt.Printf("%v Day 1 %v\n", SEPARATOR, SEPARATOR)
	fmt.Printf("Part 1: %v\n", day1.Part1("./puzzles/day1/data/input.txt"))
	fmt.Printf("Part 2: %v\n", day1.Part2("./puzzles/day1/data/input.txt"))
}
