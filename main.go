package main

import (
	"earth2mindustry/internal"
)

const input_path = "images/in.png"
const output_path = "images/out.png"

func main() {
	internal.Initialize_presets()
	preset := internal.Get_preset_from_args()
	if preset != nil {
		image := internal.Await_clipboard_image()
		internal.Convert_image(image, preset, output_path)
	}
}
