package matrix

func drawTimingPatterns(grid [][]Module, size int) {
	for i := range size {
		val := i%2 == 0

		// Horizontal line (row = 6)
		y := withOffset(6)
		x := withOffset(i)

		if grid[y][x].Reserved {
			continue
		}
		setModule(grid, y, x, val, true)

		// vertical
		y = withOffset(i)
		x = withOffset(6)

		if grid[y][x].Reserved {
			continue
		}
		setModule(grid, y, x, val, true)
	}
}
