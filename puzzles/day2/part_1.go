package day2

func Part1(fileName string) int {
	reportsData := ReadData(fileName)
	totalSafe := 0

	for _, line := range reportsData {
		if safe, _ := IsLineSafe(line); safe {
			totalSafe++
		}
	}

	return totalSafe
}
