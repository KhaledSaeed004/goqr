package main

import (
	"fmt"
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
		fmt.Println("Error generating QR code:", err)
		return
	}

	img, err := render.RenderPNG(
		qrCode.Grid,
		render.DefaultOptions(),
	)

	if err != nil {
		fmt.Println("Error rendering QR code:", err)
		return
	}

	file, err := os.Create("qrcode.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	png.Encode(file, img)
}
