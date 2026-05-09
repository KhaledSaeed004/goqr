package ecc

type RSBlock struct {
	data []byte
	ec   []byte
}

func rsEncode(data []byte, ecBytes int) []byte {
	gen := generatePoly(ecBytes)

	buffer := make([]int, len(data)+ecBytes)

	for i := 0; i < len(data); i++ {
		buffer[i] = int(data[i])
	}

	for i := 0; i < len(data); i++ {
		factor := buffer[i]
		if factor == 0 {
			continue
		}

		for j := 0; j < len(gen); j++ {
			buffer[i+j] ^= gfMul(gen[j], factor)
		}
	}

	// remainder
	ec := make([]byte, ecBytes)
	for i := 0; i < ecBytes; i++ {
		ec[i] = byte(buffer[len(data)+i])
	}

	return ec
}

func EncodeBlocks(blocks [][]byte, ecBytes int) []RSBlock {
	resultBlocks := make([]RSBlock, len(blocks))

	for i, block := range blocks {
		ec := rsEncode(block, ecBytes)
		resultBlocks[i] = RSBlock{
			data: block,
			ec:   ec,
		}
	}

	return resultBlocks
}

func InterleaveData(blocks []RSBlock) []byte {
	var resultData []byte

	maxLen := 0
	for _, b := range blocks {
		if len(b.data) > maxLen {
			maxLen = len(b.data)
		}
	}

	for i := 0; i < maxLen; i++ {
		for _, b := range blocks {
			if i < len(b.data) {
				resultData = append(resultData, b.data[i])
			}
		}
	}

	return resultData
}

func InterleaveEC(blocks []RSBlock) []byte {
	if len(blocks) == 0 {
		return nil
	}

	var resultEC []byte
	ecLen := len(blocks[0].ec)

	for i := 0; i < ecLen; i++ {
		for _, b := range blocks {
			resultEC = append(resultEC, b.ec[i])
		}
	}

	return resultEC
}

func Encode(data []byte, ecBytes int) []byte {
	return rsEncode(data, ecBytes)
}
