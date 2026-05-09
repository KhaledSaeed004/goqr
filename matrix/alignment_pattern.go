package matrix

import "github.com/KhaledSaeed004/goqr/internal"

func alignmentCenters(version int) []int {
	return internal.AlignmentTable[version]
}

func drawAlignment(grid [][]Module, centerRow, centerCol int) {
	for r := -2; r <= 2; r++ {
		for c := -2; c <= 2; c++ {

			row := centerRow + r
			col := centerCol + c

			if row < 0 || col < 0 || row >= len(grid) || col >= len(grid) {
				continue
			}

			isBlack := r == -2 || r == 2 ||
				c == -2 || c == 2 ||
				(r == 0 && c == 0)

			setModule(grid, row, col, isBlack, true)
		}
	}
}

func drawAlignmentPatterns(grid [][]Module, version int) {
	centers := alignmentCenters(version)
	if centers == nil {
		return
	}

	last := len(centers) - 1

	for rowIdx, y := range centers {
		for colIdx, x := range centers {

			centerRow := withOffset(y)
			centerCol := withOffset(x)

			// Skip finder overlaps ONLY
			if (rowIdx == 0 && colIdx == 0) || // top-left
				(rowIdx == 0 && colIdx == last) || // top-right
				(rowIdx == last && colIdx == 0) { // bottom-left
				continue
			}

			drawAlignment(grid, centerRow, centerCol)
		}
	}
}
