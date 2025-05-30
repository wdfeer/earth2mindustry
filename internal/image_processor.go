package internal

import (
	"earth2mindustry/internal/px"
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
	outImg := map_to_closest(image, pixelmap)

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

// map_to_closest maps each pixel of the image to the closest key in the pixelmap
func map_to_closest(img image.Image, pixelmap map[px.Pixel]px.Pixel) image.Image {
	bounds := img.Bounds()
	newImg := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			original := px.Pixel{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
			}

			// Find closest pixel in pixelmap keys
			closest := closest_pixel(original, pixelmap)
			mapped := pixelmap[closest]

			newImg.Set(x, y, color.RGBA{R: mapped.R, G: mapped.G, B: mapped.B, A: 255})
		}
	}

	return newImg
}

// closest_pixel finds the key in pixelmap closest to the given pixel
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

// colorDistance calculates Euclidean distance between two pixels
func colorDistance(a, b px.Pixel) float64 {
	dr := float64(a.R) - float64(b.R)
	dg := float64(a.G) - float64(b.G)
	db := float64(a.B) - float64(b.B)
	return math.Sqrt(dr*dr + dg*dg + db*db)
}
