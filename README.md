# goqr

A QR Code generator engine written from scratch in Go.

This project was built as a low-level coding exercise focused on understanding the QR specification and implementing the full generation pipeline manually without relying on external libraries.

---

## Features

- Numeric encoding
- Alphanumeric encoding
- Byte encoding
- Dynamic programming based mode segmentation
- Reed–Solomon error correction
- Error correction levels:
  - L
  - M
  - Q
  - H
- Automatic version detection
- QR mask evaluation and selection
- PNG renderer
- SVG renderer
- ASCII renderer
- Benchmarks and validation tests

---

## Project Status

- QR versions `1–9` are fully tested and scannable.
- Higher versions are implemented structurally and generate successfully, but are considered experimental.

---

## Installation

```bash
go get github.com/KhaledSaeed004/goqr
```

---

## Quick Example

```go
package main

import (
	"image/png"
	"os"

	"github.com/KhaledSaeed004/goqr"
	"github.com/KhaledSaeed004/goqr/ecc"
	"github.com/KhaledSaeed004/goqr/render"
)

func main() {

	qrCode, err := goqr.Generate(
		"https://github.com/KhaledSaeed004/goqr",
		goqr.Options{
			Level: ecc.M,
		},
	)

	if err != nil {
		panic(err)
	}

	img, err := render.RenderPNG(
		qrCode.Grid,
		render.DefaultOptions(),
	)

	if err != nil {
		panic(err)
	}

	file, err := os.Create("qrcode.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}
```

Expected result:

![QR Code](./testdata/qrcode.png "Example QR Code")

---

# API

## Generate

```go
func Generate(text string, opts Options) (*QRCode, error)
```

Generates a QR matrix from the provided input string.

---

## Options

```go
type Options struct {
	Version   int
	Level     ecc.ECLevel
	QuietZone int
}
```

### Fields

| Field | Description |
|---|---|
| `Version` | QR version. `0` enables automatic version detection |
| `Level` | Error correction level |
| `QuietZone` | Padding around the QR matrix |

---

## Return Type

```go
type QRCode struct {
	Grid    [][]matrix.Module
	Version int
	Level   ecc.ECLevel
}
```
---

# Renderers

## PNG

```go
img, err := render.RenderPNG(qrCode.Grid, render.DefaultOptions())
```

---

## SVG

```go
svg, err := render.RenderSVG(qrCode.Grid, render.DefaultOptions())
```

---

## ASCII

```go
ascii, err := render.RenderASCII(
	qrCode.Grid,
	render.DefaultASCIIOptions(false),
)
```

---

## Render Options

```go
type Options struct {
	Scale int

	Foreground *Color
	Background *Color

	// ASCII specific options
	DarkChar  string
	LightChar string
}
```

### Fields

| Field | Description |
|---|---|
| `Scale` | Image scale factor|
| `Foreground` | Module color, Default: Color{0, 0, 0, 255} |
| `Background` | Background color, Default: Color{255, 255, 255, 255} |
| `DarkChar` | ASCII module color |
| `LightChar` | ASCII background color |

---

# Encoding Pipeline

The generator follows the standard QR generation pipeline:

```text
Input
  ↓
Mode Detection / Segmentation
  ↓
Bitstream Encoding
  ↓
Error Correction (Reed–Solomon)
  ↓
Block Interleaving
  ↓
Matrix Construction
  ↓
Mask Selection
  ↓
Format / Version Information
  ↓
Rendering
```

---

# Design Goals

In this project I focused on:
- Implementing the generation pipeline manually
- Modular architecture
- Minimal external dependencies

---

# Things to Improve

- Kanji mode support
- ECI support
- CLI utility

---

# Acknowledgements

References I used for research:
- [Thonky QR Code Tutorial](https://www.thonky.com/qr-code-tutorial/introduction)
- [ISO/IEC 18004 documentation](https://cdi.mecon.gob.ar/bases/docelec/az2182.pdf)

---

# License

MIT License