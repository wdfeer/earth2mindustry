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

func TryDecodeImage(reader io.Reader) (image.Image, error) {
	img, _, err := image.Decode(reader)
	return img, err
}

func ConvertImage(image image.Image, pixelmap map[px.Pixel]px.Pixel, outputPath string) {
	println("Mapping each pixel to mindustry tiles...")
	pixelmappedImg := mapToClosest(image, pixelmap)

	println("Blending shallow water...")
	outImg := blendShallowWater(pixelmappedImg, 6)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, outImg)
	if err != nil {
		panic(err)
	}
	println("Image successfully written to " + outputPath)
}

func mapToClosest(img image.Image, pixelmap map[px.Pixel]px.Pixel) image.Image {
	bounds := img.Bounds()
	newImg := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := getPixel(img, x, y)

			closest := closestPixel(original, pixelmap)
			mapped := pixelmap[closest]

			newImg.Set(x, y, color.RGBA{R: mapped.R, G: mapped.G, B: mapped.B, A: 255})
		}
	}

	return newImg
}

func closestPixel(p px.Pixel, pixelmap map[px.Pixel]px.Pixel) px.Pixel {
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

func blendShallowWater(img image.Image, radius int) image.Image {
	bounds := img.Bounds()
	newImg := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := getPixel(img, x, y)
			var new px.Pixel

			if original.Equal(tiles.DeepWater) && !onlyDeepWaterNearby(img, x, y, radius) {
				new = tiles.ShallowWater
			} else {
				new = original
			}

			newImg.Set(x, y, color.RGBA{R: new.R, G: new.G, B: new.B, A: 255})
		}
	}

	return newImg
}

func onlyDeepWaterNearby(img image.Image, posX int, posY int, radius int) bool {
	bounds := img.Bounds()
	radiusSquared := radius * radius

	for x := max(posX-radius, bounds.Min.X); x <= min(posX+radius, bounds.Max.X-1); x++ {
		for y := max(posY-radius, bounds.Min.Y); y <= min(posY+radius, bounds.Max.Y-1); y++ {
			dx := x - posX
			dy := y - posY
			if dx*dx+dy*dy > radiusSquared {
				continue // outside the circle
			}

			if !getPixel(img, x, y).Equal(tiles.DeepWater) {
				return false
			}
		}
	}
	return true
}

func getPixel(img image.Image, x int, y int) px.Pixel {
	r, g, b, _ := img.At(x, y).RGBA()
	return px.Pixel{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
	}
}
