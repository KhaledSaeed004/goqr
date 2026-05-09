package matrix

func mask(mask int, r, c int) bool {
	switch mask {
	case 0:
		return (r+c)%2 == 0
	case 1:
		return r%2 == 0
	case 2:
		return c%3 == 0
	case 3:
		return (r+c)%3 == 0
	case 4:
		return (r/2+c/3)%2 == 0
	case 5:
		return ((r*c)%2 + (r*c)%3) == 0
	case 6:
		return (((r*c)%2 + (r*c)%3) % 2) == 0
	case 7:
		return (((r+c)%2 + (r*c)%3) % 2) == 0
	default:
		panic("invalid mask")
	}
}

func applyMask(grid [][]Module, size int, maskPattern int) {
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			y := withOffset(r)
			x := withOffset(c)

			if grid[y][x].Reserved {
				continue
			}

			if mask(maskPattern, r, c) {
				grid[y][x].Value = !grid[y][x].Value
			}
		}
	}
}
