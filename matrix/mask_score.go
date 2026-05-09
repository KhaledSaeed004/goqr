package matrix

// Consecutive modules rule
func penaltyRule1(grid [][]Module, size int) int {
	penalty := 0

	// Rows
	for r := 0; r < size; r++ {
		count := 1
		prev := grid[withOffset(r)][withOffset(0)].Value

		for c := 1; c < size; c++ {
			curr := grid[withOffset(r)][withOffset(c)].Value
			if curr == prev {
				count++
			} else {
				if count >= 5 {
					penalty += 3 + (count - 5)
				}
				count = 1
				prev = curr
			}
		}
		if count >= 5 {
			penalty += 3 + (count - 5)
		}
	}

	// Columns
	for c := 0; c < size; c++ {
		count := 1
		prev := grid[withOffset(0)][withOffset(c)].Value

		for r := 1; r < size; r++ {
			curr := grid[withOffset(r)][withOffset(c)].Value
			if curr == prev {
				count++
			} else {
				if count >= 5 {
					penalty += 3 + (count - 5)
				}
				count = 1
				prev = curr
			}
		}
		if count >= 5 {
			penalty += 3 + (count - 5)
		}
	}

	return penalty
}

// 2x2 blocks rule
func penaltyRule2(grid [][]Module, size int) int {
	penalty := 0

	for r := 0; r < size-1; r++ {
		for c := 0; c < size-1; c++ {
			v := grid[withOffset(r)][withOffset(c)].Value

			if grid[withOffset(r)][withOffset(c+1)].Value == v &&
				grid[withOffset(r+1)][withOffset(c)].Value == v &&
				grid[withOffset(r+1)][withOffset(c+1)].Value == v {

				penalty += 3
			}
		}
	}

	return penalty
}

// Finder-like patterns rule
func penaltyRule3(grid [][]Module, size int) int {
	penalty := 0

	var pattern = []bool{true, false, true, true, true, false, true, false, false, false, false}

	// Rows
	for r := 0; r < size; r++ {
		for c := 0; c <= size-11; c++ {
			match := true
			for i := 0; i < 11; i++ {
				if grid[withOffset(r)][withOffset(c+i)].Value != pattern[i] {
					match = false
					break
				}
			}
			if match {
				penalty += 40
			}
		}
	}

	// Columns
	for c := 0; c < size; c++ {
		for r := 0; r <= size-11; r++ {
			match := true
			for i := 0; i < 11; i++ {
				if grid[withOffset(r+i)][withOffset(c)].Value != pattern[i] {
					match = false
					break
				}
			}
			if match {
				penalty += 40
			}
		}
	}

	return penalty
}

// Balance of black/white rule
func penaltyRule4(grid [][]Module, size int) int {
	total := size * size
	black := 0

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if grid[withOffset(r)][withOffset(c)].Value {
				black++
			}
		}
	}

	percent := (black * 100) / total
	diff := abs(percent-50) / 5

	return diff * 10
}

func calculatePenalty(grid [][]Module, size int) int {
	return penaltyRule1(grid, size) +
		penaltyRule2(grid, size) +
		penaltyRule3(grid, size) +
		penaltyRule4(grid, size)
}
