package ecc

var exp [512]int
var log [256]int

func initGalois() {
	x := 1
	for i := 0; i < 255; i++ {
		exp[i] = x
		log[x] = i

		x <<= 1
		if x&0x100 != 0 {
			x ^= 0x11D
		}
	}

	for i := 255; i < 512; i++ {
		exp[i] = exp[i-255]
	}
}

func gfMul(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return exp[log[a]+log[b]]
}

func polyMul(p1, p2 []int) []int {
	result := make([]int, len(p1)+len(p2)-1)

	for i := range p1 {
		for j := range p2 {
			result[i+j] ^= gfMul(p1[i], p2[j])
		}
	}

	return result
}

func generatePoly(ecBytes int) []int {
	g := []int{1}

	for i := 0; i < ecBytes; i++ {
		term := []int{1, exp[i]}
		g = polyMul(g, term)
	}

	return g
}

func Init() {
	initGalois()
}
