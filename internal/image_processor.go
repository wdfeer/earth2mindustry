package internal

import (
	"earth2mindustry/internal/px"
	"earth2mindustry/internal/tiles"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
)

func Try_decode_image(reader io.Reader) (image.Image, error) {
	img, _, err := image.Decode(reader)
	return img, err
}

func Convert_image(image image.Image, pixelmap map[px.Pixel]px.Pixel, output_path string) {
	println("Mapping each pixel to mindustry tiles...")
	pixelmappedImg := map_to_closest(image, pixelmap)

	println("Blending shallow water...")
	outImg := blend_shallow_water(pixelmappedImg, 6)

	outputFile, err := os.Create(output_path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, outImg)
	if err != nil {
		panic(err)
	}
	println("Image successfully written to " + output_path)
}

func map_to_closest(img image.Image, pixelmap map[px.Pixel]px.Pixel) image.Image {
	bounds := img.Bounds()
	newImg := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := get_pixel(img, x, y)

			closest := closest_pixel(original, pixelmap)
			mapped := pixelmap[closest]

			newImg.Set(x, y, color.RGBA{R: mapped.R, G: mapped.G, B: mapped.B, A: 255})
		}
	}

	return newImg
}

func closest_pixel(p px.Pixel, pixelmap map[px.Pixel]px.Pixel) px.Pixel {
	var minDist float64 = math.MaxFloat64
	var closest px.Pixel

	for k := range pixelmap {
		dist := colorDistance(p, k)
		if dist < minDist {
			minDist = dist
			closest = k
		}
	}

	return closest
}

func colorDistance(a, b px.Pixel) float64 {
	dr := float64(a.R) - float64(b.R)
	dg := float64(a.G) - float64(b.G)
	db := float64(a.B) - float64(b.B)
	return math.Sqrt(dr*dr + dg*dg + db*db)
}

func blend_shallow_water(img image.Image, radius int) image.Image {
	bounds := img.Bounds()
	newImg := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := get_pixel(img, x, y)
			var new px.Pixel

			if original.Equal(tiles.Deep_water) && !only_deep_water_nearby(img, x, y, radius) {
				new = tiles.Shallow_water
			} else {
				new = original
			}

			newImg.Set(x, y, color.RGBA{R: new.R, G: new.G, B: new.B, A: 255})
		}
	}

	return newImg
}

func only_deep_water_nearby(img image.Image, posX int, posY int, radius int) bool {
	bounds := img.Bounds()
	for x := max(posX-radius, bounds.Min.X); x <= posX+radius && x < bounds.Max.X; x++ {
		for y := max(posY-radius, bounds.Min.Y); y <= posY+radius && y < bounds.Max.Y; y++ {
			if !get_pixel(img, x, y).Equal(tiles.Deep_water) {
				return false
			}
		}
	}
	return true
}

func get_pixel(img image.Image, x int, y int) px.Pixel {
	r, g, b, _ := img.At(x, y).RGBA()
	return px.Pixel{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
	}
}
