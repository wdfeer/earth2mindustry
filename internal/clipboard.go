package internal

import (
	"image"
	"time"

	"github.com/skanehira/clipboard-image/v2"
)

func AwaitClipboardImage() image.Image {
	reader, err := clipboard.Read()
	if err != nil {
		panic(err)
	}

	image, err := TryDecodeImage(reader)

	if err != nil {
		println("Awaiting an image in the clipboard...")
	}
	for err != nil {
		time.Sleep(200 * time.Millisecond)

		reader, err = clipboard.Read()
		if err != nil {
			panic(err)
		}
		image, err = TryDecodeImage(reader)
	}

	return image
}
