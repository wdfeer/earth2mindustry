package internal

import (
	"earth2mindustry/internal/px"
	"earth2mindustry/internal/tiles"
	"os"
	"strings"
)

type Preset struct {
	Valid    bool
	name     string
	Pixelmap map[px.Pixel]px.Pixel
}

func Get_preset_from_args() Preset {
	name := os.Args[1]

	if !Presets[name].Valid {
		println("No valid preset specified. Options:", preset_keys)
	} else {
		println("Selected", Presets[name].name, "preset.")
	}

	return Presets[name]
}

var preset_keys string
var Presets map[string]Preset = map[string]Preset{}

func register_preset(name string, pixelmap map[px.Pixel]px.Pixel) {
	Presets[name] = Preset{
		Valid:    true,
		name:     name,
		Pixelmap: pixelmap,
	}
}

func Initialize_presets() {
	register_preset("default",
		map[px.Pixel]px.Pixel{
			{R: 132, G: 215, B: 235}: tiles.Deep_water,
			{R: 207, G: 246, B: 224}: tiles.Grass,
			{R: 154, G: 229, B: 194}: tiles.Grass,
			{R: 245, G: 240, B: 230}: tiles.Sand,
			{R: 247, G: 247, B: 247}: tiles.Snow,
			{R: 167, G: 187, B: 214}: tiles.Stone,
		})

	register_preset("urban",
		map[px.Pixel]px.Pixel{
			{R: 132, G: 215, B: 235}: tiles.Shallow_water,
			{R: 207, G: 246, B: 224}: tiles.Grass,
			{R: 154, G: 229, B: 194}: tiles.Grass,
			{R: 245, G: 240, B: 230}: tiles.Sand,
			{R: 247, G: 247, B: 247}: tiles.Stone,
			{R: 167, G: 187, B: 214}: tiles.Dark_sand,
		})

	keys := make([]string, 0, len(Presets))
	for k := range Presets {
		keys = append(keys, k)
	}

	preset_keys = strings.Join(keys, ", ")
}
