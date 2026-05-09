package utils

func SplitIntoBlocks(data []byte, blockSizes []int) [][]byte {
	var blocks [][]byte
	offset := 0

	for _, size := range blockSizes {
		block := make([]byte, size)
		copy(block, data[offset:offset+size])
		blocks = append(blocks, block)
		offset += size
	}

	return blocks
}
