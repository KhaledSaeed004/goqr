package render

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/KhaledSaeed004/goqr/matrix"
)

type SVGRenderer struct{}

func (r SVGRenderer) Render(grid [][]matrix.Module, opts Options) (any, error) {
	opts.Validate()

	size := len(grid)
	scale := opts.Scale
	dim := size * scale

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(
		`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d" shape-rendering="crispEdges">`,
		dim, dim,
	))

	sb.WriteString(fmt.Sprintf(
		`<rect width="100%%" height="100%%" fill="rgba(%d,%d,%d,%f)"/>`,
		opts.Background.R,
		opts.Background.G,
		opts.Background.B,
		float64(opts.Background.A)/255,
	))

	sb.WriteString(fmt.Sprintf(
		`<path fill="rgba(%d,%d,%d,%f)" d="`,
		opts.Foreground.R,
		opts.Foreground.G,
		opts.Foreground.B,
		float64(opts.Foreground.A)/255,
	))

	for y := 0; y < size; y++ {
		x := 0
		for x < size {
			if !grid[y][x].Value {
				x++
				continue
			}

			start := x

			for x < size && grid[y][x].Value {
				x++
			}

			runLength := x - start

			startX := start * scale
			startY := y * scale
			width := runLength * scale

			// "M%d %dh%dv%dh-%dz "
			sb.WriteString("M")
			sb.WriteString(strconv.Itoa(startX))
			sb.WriteString(" ")
			sb.WriteString(strconv.Itoa(startY))
			sb.WriteString("h")
			sb.WriteString(strconv.Itoa(width))
			sb.WriteString("v")
			sb.WriteString(strconv.Itoa(scale))
			sb.WriteString("h-")
			sb.WriteString(strconv.Itoa(width))
			sb.WriteString("z ")
		}
	}

	sb.WriteString(`"/>`)
	sb.WriteString(`</svg>`)

	return sb.String(), nil
}

func RenderSVG(grid [][]matrix.Module, opts Options) (string, error) {
	r := SVGRenderer{}
	out, err := r.Render(grid, opts)
	if err != nil {
		return "", err
	}
	return out.(string), nil
}
