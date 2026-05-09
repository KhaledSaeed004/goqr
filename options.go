package goqr

import (
	"github.com/KhaledSaeed004/goqr/ecc"
	"github.com/KhaledSaeed004/goqr/matrix"
)

type Options struct {
	Version   int
	Level     ecc.ECLevel
	QuietZone int
}

func DefaultOptions() Options {
	return Options{
		Version:   0,
		Level:     ecc.L,
		QuietZone: matrix.QUIET_ZONE,
	}
}
