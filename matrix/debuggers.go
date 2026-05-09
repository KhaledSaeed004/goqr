package matrix

import "fmt"

type MatrixDebugInfo struct {
	TotalModules    int
	ReservedModules int
	WritableModules int
	DataBits        int
	UnusedWritable  int
}

func countReserved(grid [][]Module) int {
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x].Reserved {
				count++
			}
		}
	}

	return count
}

func countWritable(grid [][]Module) int {
	total := len(grid) * len(grid)

	return total - countReserved(grid)
}

func DebugMatrix(grid [][]Module, finalBits []bool) {

	size := len(grid) - 2*QUIET_ZONE

	total := size * size

	reserved := 0

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {

			y := withOffset(r)
			x := withOffset(c)

			if grid[y][x].Reserved {
				reserved++
			}
		}
	}

	writable := total - reserved

	fmt.Println("========== MATRIX DEBUG ==========")
	fmt.Println("QR Size:", size)
	fmt.Println("Total modules:", total)
	fmt.Println("Reserved modules:", reserved)
	fmt.Println("Writable modules:", writable)
	fmt.Println("Final bits:", len(finalBits))
	fmt.Println("Difference:", writable-len(finalBits))
	fmt.Println("==================================")
}
