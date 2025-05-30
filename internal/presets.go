package internal

type Preset struct {
	name     string
	Pixelmap map[Pixel]Pixel
}

func Get_preset_from_args() Preset {
	// TODO: read CLI args and return a preset
	return Default
}

var Presets map[string]Preset = map[string]Preset{
	Default.name: Default,
}

var Default Preset = Preset{
	name: "Default",
	Pixelmap: map[Pixel]Pixel{
		{1, 2, 3}: {1, 2, 3},
	},
}
