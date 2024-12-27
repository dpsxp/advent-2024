package day1

import (
	"advent/lib"
	"strconv"
	"strings"
)

type dataLines struct {
	Left  []int
	Right []int
}

func sortList(data []int) []int {
	sortedData := []int{}

	for _, val := range data {
		sortedData = append(sortedData, val)

		for sidx, currentVal := range sortedData {
			beforeSlice := sortedData[:sidx]
			afterSlice := sortedData[sidx:]
			lastIndex := len(afterSlice) - 1

			if val < currentVal {
				afterData := append(
					[]int{val},
					afterSlice[:lastIndex]...,
				)

				sortedData = append(beforeSlice, afterData...)
				break
			}
		}
	}

	return sortedData
}

func readLines(fileName string) dataLines {
	dataLines := &dataLines{}

	fileLines := lib.GetLinesFromFile(fileName)

	for _, line := range fileLines {
		dataItems := strings.Split(line, "   ")

		left := dataItems[0]
		right := dataItems[1]

		item, err := strconv.Atoi(left)

		if err != nil {
			panic(err)
		}

		itemRight, err := strconv.Atoi(right)

		if err != nil {
			panic(err)
		}

		dataLines.Left = append(dataLines.Left, item)
		dataLines.Right = append(dataLines.Right, itemRight)
	}

	return *dataLines
}
