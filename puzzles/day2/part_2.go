package day2

func Part2(fileName string) int {
	reportsData := ReadData(fileName)
	totalSafe := 0

	for _, line := range reportsData {
		safe := getSafetyLevel(line, true)

		if safe {
			totalSafe++
		}
	}

	return totalSafe
}

func getSafetyLevel(line []int, firstFailure bool) bool {
	isSafe, badIndex := IsLineSafe(line)

	if isSafe {
		return true
	}

	if !isSafe && !firstFailure {
		return false
	}

	withoutLevels := append([]int{}, line[:badIndex]...)
	withoutLevels = append(withoutLevels, line[badIndex+1:]...)

	previousLevels := append([]int{}, line[:badIndex-1]...)
	previousLevels = append(previousLevels, line[badIndex:]...)

	initLevels := append([]int{}, line[1:]...)

	return getSafetyLevel(withoutLevels, false) ||
		getSafetyLevel(previousLevels, false) ||
		getSafetyLevel(initLevels, false)

}
