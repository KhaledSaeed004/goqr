package utils

func FinalizeBits(bits []bool, capacity int) []bool {
	remaining := min(capacity-len(bits), 4)
	for range remaining {
		bits = append(bits, false)
	}

	for len(bits)%8 != 0 {
		bits = append(bits, false)
	}

	padBytes := []int{0b11101100, 0b00010001}

	for i := 0; len(bits) < capacity; i++ {
		byteVal := padBytes[i%2]
		for j := 7; j >= 0; j-- {
			bits = append(bits, ((byteVal>>j)&1) == 1)
		}
	}

	return bits
}
