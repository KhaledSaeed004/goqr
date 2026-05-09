package matrix

const (
	FINDER_SIZE    = 7
	SEPARATOR_SIZE = 1
	QUIET_ZONE     = 4
)

func newGrid(version int, quiteZone int) [][]Module {
	size := qrSize(version)
	fullSize := size + quiteZone*2

	grid := make([][]Module, fullSize)
	for i := range grid {
		grid[i] = make([]Module, fullSize)
	}

	return grid
}

func cloneGrid(grid [][]Module) [][]Module {
	clone := make([][]Module, len(grid))
	for i := range grid {
		clone[i] = make([]Module, len(grid[i]))
		copy(clone[i], grid[i])
	}

	return clone
}

func BuildMatrix(bits []bool, version int, quietZone int) [][]Module {
	size := qrSize(version)

	grid := newGrid(version, quietZone)

	count := 0
	for y := range grid {
		for x := range grid[y] {
			if !grid[y][x].Reserved {
				count++
			}
		}
	}

	drawFinders(grid, size)
	drawTimingPatterns(grid, size)
	drawAlignmentPatterns(grid, version)
	reserveFormatAreas(grid)
	reserveVersionAreas(grid, version)
	drawDarkModule(grid, version)

	placeData(grid, bits)

	return grid
}
