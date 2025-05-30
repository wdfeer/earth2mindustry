package main

import (
	"earth2mindustry/internal"
)

const input_path = "images/in.png"
const output_path = "images/out.png"

func main() {
	preset := internal.Get_preset_from_args()
	if preset.IsAbsent() {
		println("No valid preset specified. Using the default preset.")
		preset = internal.Default
	}
	internal.Write_clipboard_image_to(input_path)
	internal.Convert_image(input_path, output_path, preset.Pixelmap)
}
