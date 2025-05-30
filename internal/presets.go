package internal

import (
	"earth2mindustry/internal/px"
	"earth2mindustry/internal/tiles"
	"os"

	"github.com/samber/mo"
)

type Preset struct {
	name     string
	Pixelmap map[px.Pixel]px.Pixel
}

func Get_preset_from_args() Preset {
	return Default // using default, no other presets yet

	var result mo.Option[Preset]
	name := os.Args[1]
	if Presets[name].name != "" {
		result = mo.Some(Presets[name])
	} else {
		result = mo.None[Preset]()
	}

	preset, present := result.Get()
	if present {
		println("No valid preset specified. Using the default preset.")
		return Default
	} else {
		return preset
	}
}

var Presets map[string]Preset = map[string]Preset{
	Default.name: Default,
}

var Default Preset = Preset{
	name: "Default",
	Pixelmap: map[px.Pixel]px.Pixel{
		{R: 132, G: 215, B: 235}: tiles.Deep_water,
		{R: 207, G: 246, B: 224}: tiles.Grass,
		{R: 154, G: 229, B: 194}: tiles.Grass,
		{R: 245, G: 240, B: 230}: tiles.Sand,
		{R: 247, G: 247, B: 247}: tiles.Snow,
		{R: 167, G: 187, B: 214}: tiles.Stone,
	},
}
