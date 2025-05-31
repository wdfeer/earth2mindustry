package main

import (
	"earth2mindustry/internal"
)

const inputPath = "images/in.png"
const outputPath = "images/out.png"

func main() {
	internal.InitializePresets()
	preset := internal.GetPresetFromArgs()
	if preset.Valid {
		image := internal.AwaitClipboardImage()
		internal.ConvertImage(image, preset, outputPath)
	}
}
