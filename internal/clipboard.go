package internal

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/skanehira/clipboard-image/v2"
)

func AwaitClipboardImage() image.Image {
	reader, err := readClipboard()
	if err != nil {
		panic(err)
	}

	img, err := TryDecodeImage(reader)
	if err != nil {
		println("Awaiting an image in the clipboard...")
	}

	for err != nil {
		time.Sleep(200 * time.Millisecond)

		reader, err = readClipboard()
		if err != nil {
			panic(err)
		}

		img, err = TryDecodeImage(reader)
	}

	return img
}

func readClipboard() (io.Reader, error) {
	// Wayland detection (Linux only)
	if runtime.GOOS == "linux" && os.Getenv("WAYLAND_DISPLAY") != "" {
		return readWaylandClipboard()
	}

	// Fallback to X11 / macOS / Windows via library
	return clipboard.Read()
}

func readWaylandClipboard() (io.Reader, error) {
	// Check that wl-paste exists. If not, this is a real provider failure.
	if _, err := exec.LookPath("wl-paste"); err != nil {
		return nil, fmt.Errorf("wl-paste not found: %w", err)
	}

	cmd := exec.Command("wl-paste", "--no-newline", "--type", "image")
	out, err := cmd.Output()

	// IMPORTANT:
	// wl-paste exits with code 1 when clipboard does not contain this MIME type.
	// That is NOT a provider failure — it just means "no image yet".
	if err != nil {
		// If process started but exited non-zero → treat as non-image clipboard.
		if _, ok := err.(*exec.ExitError); ok {
			return bytes.NewReader(nil), nil
		}
		// Real failure: cannot start process
		return nil, fmt.Errorf("wl-paste failed: %w", err)
	}

	// Clipboard had image data
	return bytes.NewReader(out), nil
}
