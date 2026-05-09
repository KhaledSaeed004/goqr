package matrix

import (
	"errors"

	"github.com/KhaledSaeed004/goqr/ecc"
	"github.com/KhaledSaeed004/goqr/encode"
	"github.com/KhaledSaeed004/goqr/internal"
)

func DetectQRVersion(msg string, level ecc.ECLevel) (int, error) {
	for version := 1; version <= 40; version++ {
		spec := internal.QRTable[version][level]

		requiredBits := encode.EstimateBits(msg)
		capacity := spec.TotalDataCodewords() * 8

		if requiredBits <= capacity {
			return version, nil
		}
	}

	return 0, errors.New("message too long for any QR version at the specified error correction level")
}
