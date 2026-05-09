package utils

import "github.com/KhaledSaeed004/goqr/internal"

func BuildBlocks(data []byte, spec internal.QRSpec) [][]byte {
	var blocks [][]byte

	offset := 0

	for _, group := range spec.Groups {
		for i := 0; i < group.NumBlocks; i++ {
			size := group.DataCodewords

			block := make([]byte, size)
			copy(block, data[offset:offset+size])

			blocks = append(blocks, block)
			offset += size
		}
	}

	return blocks
}
