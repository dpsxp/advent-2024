package day2

import (
	"advent/lib"
	"strconv"
	"strings"
)

func IsLineSafe(line []int) (bool, int) {
	diffThreshold := 3
	changeType := ""

	for idx, val := range line {
		if idx == 0 {
			continue
		}

		previous := line[idx-1]

		if idx == 1 {
			isIncrease := val > previous

			if isIncrease {
				changeType = "increase"
			} else {
				changeType = "decrease"
			}
		}

		change := max(val, previous) - min(val, previous)

		if val < previous && changeType == "increase" {
			return false, idx
		}

		if val > previous && changeType == "decrease" {
			return false, idx
		}

		if change == 0 || change > diffThreshold {
			return false, idx
		}
	}

	return true, 0
}

func ReadData(fileName string) [][]int {
	lines := lib.GetLinesFromFile(fileName)

	data := [][]int{}

	for _, val := range lines {
		numbers := strings.Split(val, " ")

		result := []int{}

		for _, num := range numbers {
			num, err := strconv.Atoi(num)

			if err != nil {
				panic(err)
			}

			result = append(result, num)
		}

		data = append(data, result)
	}

	return data
}
