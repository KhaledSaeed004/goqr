package goqr

import (
	"testing"

	"github.com/KhaledSaeed004/goqr/ecc"
	"github.com/KhaledSaeed004/goqr/encode"
	"github.com/KhaledSaeed004/goqr/matrix"
	"github.com/KhaledSaeed004/goqr/render"
)

func BenchmarkGenerate(b *testing.B) {
	opts := DefaultOptions()
	input := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Generate(input, opts)
	}
}

func BenchmarkEncode(b *testing.B) {
	input := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	for i := 0; i < b.N; i++ {
		_, _ = encode.EncodeString(input)
	}
}

func BenchmarkECC(b *testing.B) {
	data := make([]byte, 100)

	for i := 0; i < b.N; i++ {
		_ = ecc.Encode(data, 20)
	}
}

func BenchmarkMatrix(b *testing.B) {
	gridBits := make([]bool, 800)

	for i := 0; i < b.N; i++ {
		grid := matrix.BuildMatrix(gridBits, 6, 4)
		matrix.ApplyBestMask(grid, 6)
	}
}

func BenchmarkPNG(b *testing.B) {
	qrCode, _ := Generate("spaghetti doodles", DefaultOptions())
	opts := render.DefaultOptions()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		render.RenderPNG(qrCode.Grid, opts)
	}
}

func BenchmarkSVG(b *testing.B) {
	qrCode, _ := Generate("spaghetti doodles", DefaultOptions())
	opts := render.DefaultOptions()

	for i := 0; i < b.N; i++ {
		render.RenderSVG(qrCode.Grid, opts)
	}
}
