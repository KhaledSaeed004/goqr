package matrix

type Module struct {
	Value    bool
	Reserved bool
}

func setModule(grid [][]Module, y, x int, val bool, reserved bool) {
	if grid[y][x].Reserved {
		return
	}
	grid[y][x].Value = val
	grid[y][x].Reserved = reserved
}

func setFunctionModule(grid [][]Module, y, x int, val bool, reserved bool) {
	grid[y][x].Value = val
	grid[y][x].Reserved = reserved
}
