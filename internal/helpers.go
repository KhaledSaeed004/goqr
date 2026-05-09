package internal

func (spec QRSpec) TotalDataCodewords() int {
	total := 0
	for _, g := range spec.Groups {
		total += g.NumBlocks * g.DataCodewords
	}
	return total
}

func (spec QRSpec) TotalBlocks() int {
	total := 0
	for _, g := range spec.Groups {
		total += g.NumBlocks
	}
	return total
}
