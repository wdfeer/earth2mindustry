package internal

import (
	"fmt"
	"io"
	"os"

	"github.com/skanehira/clipboard-image/v2"
)

func Write_clipboard_image_to(path string) {
	reader, err := clipboard.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read image from clipboard: %v\n", err)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file %s: %v\n", path, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write image to file: %v\n", err)
		return
	}

	fmt.Printf("Image successfully written to %s\n", path)
}
