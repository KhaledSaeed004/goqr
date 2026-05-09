package goqr

import (
	"fmt"

	"github.com/KhaledSaeed004/goqr/ecc"
	"github.com/KhaledSaeed004/goqr/encode"
	"github.com/KhaledSaeed004/goqr/internal"
	"github.com/KhaledSaeed004/goqr/matrix"
)

type QRCode struct {
	Grid    [][]matrix.Module
	Version int
	Level   ecc.ECLevel
}

func Generate(text string, opts Options) (*QRCode, error) {
	defaults := DefaultOptions()

	if opts.QuietZone == 0 {
		opts.QuietZone = defaults.QuietZone
	}

	if opts.QuietZone < 0 {
		return nil, fmt.Errorf("quiet zone cannot be negative")
	}

	if opts.Level == 0 {
		opts.Level = defaults.Level
	}

	version := opts.Version
	if version == 0 {
		var err error
		version, err = matrix.DetectQRVersion(text, opts.Level)
		if err != nil {
			return nil, err
		}
	}

	ecc.Init()

	encodedBits, err := encode.EncodeString(text)
	if err != nil {
		return nil, err
	}
	dataECBits, err := encode.EncodeDataEC(encodedBits, version, opts.Level)
	if err != nil {
		return nil, err
	}

	for i := 0; i < internal.RemainderBits[version]; i++ {
		dataECBits = append(dataECBits, false)
	}

	qrGrid := matrix.BuildMatrix(dataECBits, version, opts.QuietZone)

	qrGrid, qrMask := matrix.ApplyBestMask(qrGrid, version)

	matrix.WriteFormatBits(qrGrid, qrMask, opts.Level)
	matrix.WriteVersionBits(qrGrid, version)

	return &QRCode{
		Grid:    qrGrid,
		Version: version,
		Level:   opts.Level,
	}, nil
}
