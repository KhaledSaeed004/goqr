package encode

type BitBuffer struct {
	Bits []bool
}

func (b *BitBuffer) AppendBits(value int, length int) {
	for i := length - 1; i >= 0; i-- {
		b.Bits = append(b.Bits, ((value>>i)&1) == 1)
	}
}
