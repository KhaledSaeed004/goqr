package render

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/KhaledSaeed004/goqr/matrix"
)

type PNGRenderer struct{}

func (r PNGRenderer) Render(grid [][]matrix.Module, opts Options) (any, error) {
	opts.Validate()

	size := len(grid)
	scale := opts.Scale
	imgSize := size * scale

	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))
	rect := image.Rectangle{}

	fg := color.RGBA{R: opts.Foreground.R, G: opts.Foreground.G, B: opts.Foreground.B, A: opts.Foreground.A}
	bg := color.RGBA{R: opts.Background.R, G: opts.Background.G, B: opts.Background.B, A: opts.Background.A}

	fgUniform := &image.Uniform{C: fg}
	bgUniform := &image.Uniform{C: bg}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {

			startX := x * scale
			startY := y * scale

			rect.Min.X = startX
			rect.Min.Y = startY
			rect.Max.X = startX + scale
			rect.Max.Y = startY + scale

			src := bgUniform
			if grid[y][x].Value {
				src = fgUniform
			}

			draw.Draw(img, rect, src, image.Point{}, draw.Src)
		}
	}

	return img, nil
}

func RenderPNG(grid [][]matrix.Module, opts Options) (*image.RGBA, error) {
	r := PNGRenderer{}
	img, err := r.Render(grid, opts)
	if err != nil {
		return nil, err
	}
	return img.(*image.RGBA), nil
}
