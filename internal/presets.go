package internal

import (
	"github.com/samber/mo"
	"os"
)

type Preset struct {
	name     string
	Pixelmap map[Pixel]Pixel
}

func Get_preset_from_args() Preset {
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
	Pixelmap: map[Pixel]Pixel{
		{1, 2, 3}: {1, 2, 3}, // TODO: get actual pixel values
	},
}
