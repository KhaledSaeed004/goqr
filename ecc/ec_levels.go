package ecc

type ECLevel int

const (
	L ECLevel = iota
	M
	Q
	H
)

func ParseECLevel(level string) ECLevel {
	switch level {
	case "L":
		return L
	case "M":
		return M
	case "Q":
		return Q
	case "H":
		return H
	default:
		return L
	}
}

func ECLevelBits(level ECLevel) int {
	switch level {
	case L:
		return 0b01
	case M:
		return 0b00
	case Q:
		return 0b11
	case H:
		return 0b10
	default:
		return 0b01
	}
}
