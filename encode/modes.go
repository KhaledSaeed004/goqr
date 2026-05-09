package encode

import "github.com/KhaledSaeed004/goqr/internal"

type QRMode int

const (
	Numeric QRMode = iota
	Alphanumeric
	Byte
)

func modeIndicator(mode QRMode) int {
	switch mode {
	case Numeric:
		return 0b0001
	case Alphanumeric:
		return 0b0010
	case Byte:
		return 0b0100
	default:
		panic("unsupported mode")
	}
}

func encodeNumericInto(buffer *BitBuffer, msg string) {
	length := len(msg)

	for i := 0; i < length; {
		remaining := length - i

		if remaining >= 3 {
			num := int(msg[i]-'0')*100 +
				int(msg[i+1]-'0')*10 +
				int(msg[i+2]-'0')

			buffer.AppendBits(num, 10)
			i += 3

		} else if remaining == 2 {
			num := int(msg[i]-'0')*10 +
				int(msg[i+1]-'0')

			buffer.AppendBits(num, 7)
			i += 2

		} else {
			num := int(msg[i] - '0')

			buffer.AppendBits(num, 4)
			i += 1
		}
	}
}

func encodeAlphanumericInto(buffer *BitBuffer, msg string) {
	length := len(msg)

	for i := 0; i < length; {
		remaining := length - i

		if remaining >= 2 {
			val1 := internal.AlphanumericCharset[rune(msg[i])]
			val2 := internal.AlphanumericCharset[rune(msg[i+1])]

			value := val1*45 + val2
			buffer.AppendBits(value, 11)

			i += 2
		} else {
			val := internal.AlphanumericCharset[rune(msg[i])]
			buffer.AppendBits(val, 6)

			i += 1
		}
	}
}

func encodeByteInto(buffer *BitBuffer, data string) {
	for i := 0; i < len(data); i++ {
		buffer.AppendBits(int(data[i]), 8)
	}
}
