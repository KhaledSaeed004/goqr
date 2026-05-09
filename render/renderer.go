package render

import (
	"github.com/KhaledSaeed004/goqr/matrix"
)

const (
	DEFAULT_SCALE = 10
)

var (
	DEFAULT_FOREGROUND = Color{0, 0, 0, 255}
	DEFAULT_BACKGROUND = Color{255, 255, 255, 255}
)

type Renderer interface {
	Render(grid [][]matrix.Module, opts Options) (any, error)
}

type Options struct {
	Scale int

	Foreground *Color
	Background *Color

	// ASCII specific
	DarkChar  string
	LightChar string
}

type Color struct {
	R, G, B, A uint8
}

func (opts *Options) Validate() {
	if opts.IsUnset() {
		*opts = DefaultOptions()
	}
	if opts.Scale <= 0 {
		opts.Scale = DEFAULT_SCALE
	}
	if opts.Foreground == nil {
		FG := DEFAULT_FOREGROUND
		opts.Foreground = &FG
	}
	if opts.Background == nil {
		BG := DEFAULT_BACKGROUND
		opts.Background = &BG
	}
}

func (opts Options) IsUnset() bool {
	return opts == Options{}
}

func DefaultOptions() Options {
	FG := DEFAULT_FOREGROUND
	BG := DEFAULT_BACKGROUND

	return Options{
		Scale:      DEFAULT_SCALE,
		Foreground: &FG,
		Background: &BG,
	}
}
