package encode

import (
	"errors"

	"github.com/KhaledSaeed004/goqr/ecc"
	"github.com/KhaledSaeed004/goqr/internal"
	"github.com/KhaledSaeed004/goqr/utils"
)

func EncodeString(msg string) ([]bool, error) {
	buffer := &BitBuffer{}

	segments := segmentMessage(msg)

	for _, seg := range segments {
		// Mode indicator
		buffer.AppendBits(modeIndicator(seg.Mode), 4)

		// Character count
		countBits := charCountBits(seg.Mode)
		buffer.AppendBits(len(seg.Data), countBits)

		// Data
		switch seg.Mode {
		case Numeric:
			encodeNumericInto(buffer, seg.Data)
		case Alphanumeric:
			encodeAlphanumericInto(buffer, seg.Data)
		case Byte:
			encodeByteInto(buffer, seg.Data)
		default:
			return nil, errors.New("unsupported mode")
		}
	}

	return buffer.Bits, nil
}

func EncodeDataEC(encodedBits []bool, version int, level ecc.ECLevel) ([]bool, error) {
	spec := internal.QRTable[version][level]

	bits := utils.FinalizeBits(encodedBits, spec.TotalDataCodewords()*8)

	dataBytes := utils.BitsToBytes(bits)

	var dataECBytes []byte

	blocks := utils.BuildBlocks(dataBytes, spec)

	if len(blocks) == 0 {
		return nil, errors.New("no blocks generated (capacity mismatch)")
	}

	rsBlocks := ecc.EncodeBlocks(blocks, spec.ECCodewordsPerBlock)

	interData := ecc.InterleaveData(rsBlocks)
	interEC := ecc.InterleaveEC(rsBlocks)

	dataECBytes = append(interData, interEC...)
	dataECBits := utils.BytesToBits(dataECBytes)

	return dataECBits, nil
}
