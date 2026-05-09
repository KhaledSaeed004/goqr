package render

import (
	"strings"

	"github.com/KhaledSaeed004/goqr/matrix"
)

type ASCIIRenderer struct{}

func DefaultASCIIOptions(dark bool) Options {
	opts := Options{}

	if dark {
		opts.DarkChar = "██"
		opts.LightChar = "  "
	} else {
		opts.DarkChar = "  "
		opts.LightChar = "██"
	}

	return opts
}

func (r ASCIIRenderer) Render(grid [][]matrix.Module, opts Options) (any, error) {
	if opts.DarkChar == "" || opts.LightChar == "" {
		opts = DefaultASCIIOptions(false)
	}

	fgChar := opts.DarkChar
	bgChar := opts.LightChar

	var sb strings.Builder

	size := len(grid)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if grid[y][x].Value {
				sb.WriteString(fgChar)
			} else {
				sb.WriteString(bgChar)
			}
		}
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func RenderASCII(grid [][]matrix.Module, opts Options) (string, error) {
	r := ASCIIRenderer{}
	out, err := r.Render(grid, opts)
	if err != nil {
		return "", err
	}
	return out.(string), nil
}
