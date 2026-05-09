package matrix

func traverse(size int, visit func(r, c int)) {
	col := size - 1
	up := true

	for col > 0 {

		// Skip timing column
		if col == 6 {
			col--
		}

		for i := 0; i < size; i++ {
			var row int
			if up {
				row = size - 1 - i
			} else {
				row = i
			}

			// Visit 2 columns (right → left)
			for dx := 0; dx < 2; dx++ {
				c := col - dx
				visit(row, c)
			}
		}

		col -= 2
		up = !up
	}
}

func placeData(grid [][]Module, bits []bool) {
	size := len(grid) - QUIET_ZONE*2

	bitIndex := 0
	visited := 0

	traverse(size, func(r, c int) {

		y := withOffset(r)
		x := withOffset(c)

		if grid[y][x].Reserved {
			return
		}

		visited++

		if bitIndex >= len(bits) {
			return
		}

		setModule(grid, y, x, bits[bitIndex], false)
		bitIndex++
	})

}
