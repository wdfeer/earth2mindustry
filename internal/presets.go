package internal

import (
	"earth2mindustry/internal/mapcolors"
	"earth2mindustry/internal/px"
	"earth2mindustry/internal/tiles"
	"os"
	"strings"
)

type Preset struct {
	pixelmap     map[px.Pixel]px.Pixel
	shallowWater px.Pixel
}

func (self Preset) IsValid() bool {
	return self.pixelmap != nil
}

func GetPresetFromArgs() Preset {
	if len(os.Args) < 2 {
		println("No valid preset specified. Options:", presetKeys)
		return Preset{}
	}

	name := os.Args[1]

	if Presets[name].IsValid() {
		println("Selected", name, "preset.")
	} else {
		println("No valid preset specified. Options:", presetKeys)
	}

	return Presets[name]
}

var presetKeys string
var Presets map[string]Preset = map[string]Preset{}

func InitializePresets() {
	Presets["default"] = Preset{pixelmap: map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.Grass,
		mapcolors.Grass:  tiles.Grass,
		mapcolors.Road:   tiles.Stone,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.Stone,
	}, shallowWater: tiles.ShallowWater}
	Presets["snow"] = Preset{pixelmap: map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.Grass,
		mapcolors.Grass:  tiles.Snow,
		mapcolors.Road:   tiles.Stone,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.Snow,
	}, shallowWater: tiles.ShallowWater}
	Presets["arctic"] = Preset{pixelmap: map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.Grass,
		mapcolors.Grass:  tiles.Snow,
		mapcolors.Road:   tiles.Ice,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.Snow,
	}, shallowWater: tiles.Ice}
	Presets["desert"] = Preset{pixelmap: map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.Grass,
		mapcolors.Grass:  tiles.Grass,
		mapcolors.Road:   tiles.DarkSand,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.Sand,
	}, shallowWater: tiles.ShallowWater}
	Presets["wasteland"] = Preset{pixelmap: map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.DarkSand,
		mapcolors.Grass:  tiles.Sand,
		mapcolors.Road:   tiles.Stone,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.DarkSand,
	}, shallowWater: tiles.ShallowWater}

	keys := make([]string, 0, len(Presets))
	for k := range Presets {
		keys = append(keys, k)
	}

	presetKeys = strings.Join(keys, ", ")
}
