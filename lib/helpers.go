package lib

import (
	"io"
	"os"
	"path/filepath"
)

func GetLinesFromFile(fileName string) []string {
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
