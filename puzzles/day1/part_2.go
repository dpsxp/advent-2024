package day1

func Part2(fileName string) int {
	type cacheEntry struct {
		Left  int
		Right int
	}

	var cache = map[int]*cacheEntry{}
	lines := readLines(fileName)

	listLeft := lines.Left
	listRight := lines.Right

	for idx, line := range listLeft {
		entry, ok := cache[line]

		if ok {
			entry.Left = entry.Left + 1
		} else {
			cache[line] = &cacheEntry{
				Left:  1,
				Right: 0,
			}
		}

		rightValue := listRight[idx]

		rightEntry, ok := cache[rightValue]

		if ok {
			rightEntry.Right = rightEntry.Right + 1
		} else {
			cache[rightValue] = &cacheEntry{
				Left:  0,
				Right: 1,
			}
		}
	}

	total := 0

	for key, entry := range cache {
		total += (key * entry.Right) * entry.Left
	}

	return total
}
