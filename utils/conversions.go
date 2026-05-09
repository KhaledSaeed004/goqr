package utils

func BitsToBytes(bits []bool) []byte {
	if len(bits)%8 != 0 {
		panic("bits not byte-aligned")
	}

	bytes := make([]byte, len(bits)/8)

	for i := 0; i < len(bits); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			if bits[i+j] {
				b |= 1 << (7 - j)
			}
		}
		bytes[i/8] = b
	}

	return bytes
}

func BytesToBits(bytes []byte) []bool {
	bits := make([]bool, 0, len(bytes)*8)

	for _, b := range bytes {
		for i := 7; i >= 0; i-- {
			isBitSet := ((b >> i) & 1) == 1
			bits = append(bits, isBitSet)
		}
	}

	return bits
}
