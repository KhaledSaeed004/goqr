package matrix

func qrSize(version int) int {
	return 21 + (version-1)*4
}

func withOffset(p int) int {
	return p + QUIET_ZONE
}

func abs[T ~int | ~int32 | ~int64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
