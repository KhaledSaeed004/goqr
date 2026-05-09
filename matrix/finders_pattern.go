package matrix

func isFinderOverlap(x, y, size int) bool {
	// top-left
	if x == 6 && y == 6 {
		return true
	}

	// top-right
	if x == size-7 && y == 6 {
		return true
	}

	// bottom-left
	if x == 6 && y == size-7 {
		return true
	}

	return false
}

func finderAnchors(size int) [][2]int {
	return [][2]int{
		{0, 0},
		{0, size - FINDER_SIZE},
		{size - FINDER_SIZE, 0},
	}
}

func drawFinder(grid [][]Module, startRow, startCol int) {
	for r := range FINDER_SIZE {
		for c := range FINDER_SIZE {
			row := startRow + r
			col := startCol + c
			isBlack := r == 0 || r == FINDER_SIZE-1 || c == 0 || c == FINDER_SIZE-1 || (r >= 2 && r <= 4 && c >= 2 && c <= 4)

			setModule(grid, row, col, isBlack, true)
		}
	}
}

func drawSeparator(grid [][]Module, startRow, startCol int) {
	min := -SEPARATOR_SIZE
	max := FINDER_SIZE + SEPARATOR_SIZE

	for r := min; r < max; r++ {
		for c := min; c < max; c++ {

			if r >= 0 && r < FINDER_SIZE && c >= 0 && c < FINDER_SIZE {
				continue
			}

			row := startRow + r
			col := startCol + c

			if row < 0 || col < 0 || row >= len(grid) || col >= len(grid) {
				continue
			}

			setModule(grid, row, col, false, true)
		}
	}
}

func drawFinders(grid [][]Module, size int) {
	for _, anchor := range finderAnchors(size) {
		startRow := withOffset(anchor[0])
		startCol := withOffset(anchor[1])

		drawFinder(grid, startRow, startCol)
		drawSeparator(grid, startRow, startCol)
	}
}
