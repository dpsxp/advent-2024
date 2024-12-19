package day1

import (
	"io"
	"os"
	"path/filepath"
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

	fileLines := getLinesFromFile(fileName)

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

func getLinesFromFile(fileName string) []string {
	path, err := filepath.Abs(fileName)

	if err != nil {
		panic(err)
	}

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	MAX_BYTES := 1

	bytes := make([]byte, MAX_BYTES)

	dataBuffer := ""

	lines := []string{}

	for {
		n, err := file.Read(bytes)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if n < MAX_BYTES {
			break
		}

		char := string(bytes)

		if char != "\n" {
			dataBuffer += char
			continue
		}

		lines = append(lines, dataBuffer)

		dataBuffer = ""
	}

	return lines
}
