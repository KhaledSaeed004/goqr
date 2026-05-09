package matrix

func drawDarkModule(grid [][]Module, version int) {
	y := withOffset(4*version + 9)
	x := withOffset(8)

	setModule(grid, y, x, true, true)
}
