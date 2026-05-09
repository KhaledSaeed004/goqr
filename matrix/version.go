package matrix

func reserveVersionAreas(grid [][]Module, version int) {
	if version < 7 {
		return
	}

	size := len(grid) - 2*QUIET_ZONE

	// Top-right area
	for r := 0; r < 6; r++ {
		for c := 0; c < 3; c++ {

			y := withOffset(r)
			x := withOffset(size - 11 + c)

			grid[y][x].Reserved = true
		}
	}

	// Bottom-left area
	for r := 0; r < 3; r++ {
		for c := 0; c < 6; c++ {

			y := withOffset(size - 11 + r)
			x := withOffset(c)

			grid[y][x].Reserved = true
		}
	}
}

func getVersionBits(version int) int {
	v := version << 12

	// generator polynomial: 0x1F25
	g := 0x1F25

	for i := 17; i >= 12; i-- {
		if ((v >> i) & 1) == 1 {
			v ^= g << (i - 12)
		}
	}

	return (version << 12) | (v & 0xFFF)
}

func WriteVersionBits(grid [][]Module, version int) {
	if version < 7 {
		return
	}

	bits := getVersionBits(version)
	size := len(grid) - 2*QUIET_ZONE

	// 18 bits in total, written in two places
	for i := 0; i < 18; i++ {
		bit := ((bits >> i) & 1) == 1

		// bottom-left
		row := withOffset(size - 11 + (i % 3))
		col := withOffset(i / 3)
		setFunctionModule(grid, row, col, bit, true)

		// top-right
		row = withOffset(i / 3)
		col = withOffset(size - 11 + (i % 3))
		setFunctionModule(grid, row, col, bit, true)
	}
}
