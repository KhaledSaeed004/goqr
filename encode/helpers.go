package encode

import "github.com/KhaledSaeed004/goqr/internal"

func charCountBits(mode QRMode) int {
	switch mode {
	case Numeric:
		return 10
	case Alphanumeric:
		return 9
	case Byte:
		return 8
	default:
		panic("unsupported mode")
	}
}

func detectBestMode(msg string) QRMode {
	if isNumeric(msg) {
		return Numeric
	}
	if isAlphanumeric(msg) {
		return Alphanumeric
	}
	return Byte
}

func isNumeric(msg string) bool {
	for _, c := range msg {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(msg) > 0
}

func isAlphanumeric(msg string) bool {
	for _, c := range msg {
		if _, ok := internal.AlphanumericCharset[c]; !ok {
			return false
		}
	}
	return len(msg) > 0
}

func canEncode(mode QRMode, c byte) bool {
	switch mode {
	case Numeric:
		return c >= '0' && c <= '9'
	case Alphanumeric:
		_, ok := internal.AlphanumericCharset[rune(c)]
		return ok
	case Byte:
		return true
	default:
		return false
	}
}

func incrementalCost(mode QRMode, runLength int) int {
	switch mode {

	case Numeric:
		if runLength%3 == 1 {
			return 4
		} else if runLength%3 == 2 {
			return 3 // completing 2-digit group (7 total)
		} else {
			return 3 // completing 3-digit group (10 total)
		}

	case Alphanumeric:
		if runLength%2 == 1 {
			return 6
		} else {
			return 5 // completing pair (11 total)
		}

	case Byte:
		return 8
	}

	return 0
}

func EstimateBits(msg string) int {
	total := 0

	segments := segmentMessage(msg)

	for _, seg := range segments {
		countBits := charCountBits(seg.Mode)
		total += 4 + countBits // mode + length

		switch seg.Mode {
		case Numeric:
			length := len(seg.Data)
			full := length / 3
			rem := length % 3

			total += full * 10
			switch rem {
			case 2:
				total += 7
			case 1:
				total += 4
			}

		case Alphanumeric:
			length := len(seg.Data)
			full := length / 2
			rem := length % 2

			total += full * 11
			if rem == 1 {
				total += 6
			}

		case Byte:
			total += len(seg.Data) * 8

		default:
			panic("unsupported mode")
		}
	}

	return total
}
