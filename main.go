package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	file, _ := os.Open("input.png")
	defer file.Close()
	img, _, _ := image.Decode(file)

	bounds := img.Bounds()
	outImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			origColor := img.At(x, y)
			r, g, b, a := origColor.RGBA()
			// r, g, b, a are in 16-bit (0-65535), so compare with 65535
			if r == 0xffff && g == 0xffff && b == 0xffff {
				outImg.Set(x, y, color.RGBA{0, 0, 0, uint8(a >> 8)})
			} else {
				outImg.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
			}
		}
	}

	outFile, _ := os.Create("output.png")
	defer outFile.Close()
	png.Encode(outFile, outImg)
}
