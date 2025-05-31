package internal

import (
	"earth2mindustry/internal/mapcolors"
	"earth2mindustry/internal/px"
	"earth2mindustry/internal/tiles"
	"os"
	"strings"
)

type Preset = map[px.Pixel]px.Pixel

func GetPresetFromArgs() Preset {
	if len(os.Args) < 2 {
		println("No valid preset specified. Options:", presetKeys)
		return nil
	}

	name := os.Args[1]

	if Presets[name] == nil {
		println("No valid preset specified. Options:", presetKeys)
	} else {
		println("Selected", name, "preset.")
	}

	return Presets[name]
}

var presetKeys string
var Presets map[string]Preset = map[string]Preset{}

func InitializePresets() {
	Presets["default"] = map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.Grass,
		mapcolors.Grass:  tiles.Grass,
		mapcolors.Road:   tiles.Stone,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.Stone,
	}
	Presets["snow"] = map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.Grass,
		mapcolors.Grass:  tiles.Snow,
		mapcolors.Road:   tiles.Stone,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.Snow,
	}
	Presets["desert"] = map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.Grass,
		mapcolors.Grass:  tiles.Grass,
		mapcolors.Road:   tiles.DarkSand,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.Sand,
	}
	Presets["wasteland"] = map[px.Pixel]px.Pixel{
		mapcolors.Border: tiles.DarkSand,
		mapcolors.Forest: tiles.DarkSand,
		mapcolors.Grass:  tiles.Sand,
		mapcolors.Road:   tiles.Stone,
		mapcolors.Sand:   tiles.Sand,
		mapcolors.Water:  tiles.DeepWater,
		mapcolors.White:  tiles.DarkSand,
	}

	keys := make([]string, 0, len(Presets))
	for k := range Presets {
		keys = append(keys, k)
	}

	presetKeys = strings.Join(keys, ", ")
}
