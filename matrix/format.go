package matrix

import "github.com/KhaledSaeed004/goqr/ecc"

func reserveFormatAreas(grid [][]Module) {
	// Row (y = 8, x = 0..8 except x=6)
	for x := 0; x <= 8; x++ {
		if x == 6 {
			continue
		}
		grid[withOffset(8)][withOffset(x)].Reserved = true
	}

	// Column (x = 8, y = 0..8 except y=6)
	for y := 0; y <= 8; y++ {
		if y == 6 {
			continue
		}
		grid[withOffset(y)][withOffset(8)].Reserved = true
	}

	size := len(grid) - 2*QUIET_ZONE

	// Right side (horizontal mirror)
	for x := size - 8; x < size; x++ {
		grid[withOffset(8)][withOffset(x)].Reserved = true
	}

	// Bottom side (vertical mirror)
	for y := size - 8; y < size; y++ {
		grid[withOffset(y)][withOffset(8)].Reserved = true
	}
}

func getFormatBits(mask int, level ecc.ECLevel) int {
	// Level
	levelBits := ecc.ECLevelBits(level)

	// 5 bits: EC level + mask
	data := (levelBits << 3) | mask

	// BCH encode (15,5)
	g := 0b10100110111
	bits := data << 10

	for i := 14; i >= 10; i-- {
		if ((bits >> i) & 1) == 1 {
			bits ^= g << (i - 10)
		}
	}

	format := ((data << 10) | bits) ^ 0b101010000010010

	return format
}

func formatBitsToArray(format int) []bool {
	result := make([]bool, 15)

	for i := 0; i < 15; i++ {
		result[14-i] = ((format >> i) & 1) == 1
	}

	return result
}

func WriteFormatBits(grid [][]Module, mask int, level ecc.ECLevel) {
	size := len(grid) - 2*QUIET_ZONE

	format := getFormatBits(mask, level)
	bits := formatBitsToArray(format)

	// --- top-left ---
	for i := 0; i <= 5; i++ {
		setFunctionModule(grid, withOffset(8), withOffset(i), bits[i], true)
	}
	setFunctionModule(grid, withOffset(8), withOffset(7), bits[6], true)
	setFunctionModule(grid, withOffset(8), withOffset(8), bits[7], true)
	setFunctionModule(grid, withOffset(7), withOffset(8), bits[8], true)

	for i := 9; i < 15; i++ {
		setFunctionModule(grid, withOffset(14-i), withOffset(8), bits[i], true)
	}

	// --- top-right ---
	for i := 0; i < 8; i++ {
		setFunctionModule(grid, withOffset(8), withOffset(size-1-i), bits[i], true)
	}

	// --- bottom-left ---
	for i := 0; i < 7; i++ {
		setFunctionModule(grid, withOffset(size-1-i), withOffset(8), bits[i], true)
	}
}
