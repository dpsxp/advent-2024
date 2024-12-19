package day1

func Part1(filename string) int {
	lines := readLines(filename)

	var total int = 0

	listLeft := sortList(lines.Left)
	listRight := sortList(lines.Right)

	totalLineDistance := 0

	for idx, item := range listLeft {
		otherNumber := listRight[idx]
		distance := item - otherNumber

		if distance < 0 {
			distance = distance * -1
		}

		totalLineDistance += distance
	}

	total += totalLineDistance

	return total
}
