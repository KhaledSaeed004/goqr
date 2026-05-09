package matrix

func ApplyBestMask(grid [][]Module, version int) ([][]Module, int) {
	size := qrSize(version)
	bestMask := 0
	minPenalty := int(^uint(0) >> 1)

	var bestGrid [][]Module

	for m := 0; m < 8; m++ {
		testGrid := cloneGrid(grid)

		applyMask(testGrid, size, m)

		score := calculatePenalty(testGrid, size)

		if score < minPenalty {
			minPenalty = score
			bestMask = m
			bestGrid = testGrid
		}
	}

	return bestGrid, bestMask
}
