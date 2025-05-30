package internal

type Pixel struct {
	R, G, B int
}

type Preset struct {
	name     string
	pixelmap map[Pixel]Pixel
}

var Presets map[string]Preset = map[string]Preset{
	Default.name: Default,
}

var Default Preset = Preset{
	name: "Default",
	pixelmap: map[Pixel]Pixel{
		{1, 2, 3}: {1, 2, 3},
	},
}
