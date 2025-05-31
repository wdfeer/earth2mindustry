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
	Presets["default"] =
		map[px.Pixel]px.Pixel{
			mapcolors.Water:  tiles.DeepWater,
			mapcolors.Grass:  tiles.Grass,
			mapcolors.Forest: tiles.Grass,
			mapcolors.Sand:   tiles.Sand,
			mapcolors.White:  tiles.Stone,
			mapcolors.Border: tiles.DarkSand,
			mapcolors.Road:   tiles.DarkSand,
		}

	Presets["snow"] =
		map[px.Pixel]px.Pixel{
			mapcolors.Water:  tiles.DeepWater,
			mapcolors.Grass:  tiles.Snow,
			mapcolors.Forest: tiles.Grass,
			mapcolors.Sand:   tiles.Sand,
			mapcolors.White:  tiles.Snow,
			mapcolors.Road:   tiles.DarkSand,
			mapcolors.Border: tiles.DarkSand,
		}

	Presets["desert"] =
		map[px.Pixel]px.Pixel{
			mapcolors.Water:  tiles.DeepWater,
			mapcolors.Grass:  tiles.Grass,
			mapcolors.Forest: tiles.Grass,
			mapcolors.Sand:   tiles.Sand,
			mapcolors.White:  tiles.Sand,
			mapcolors.Road:   tiles.DarkSand,
			mapcolors.Border: tiles.DarkSand,
		}

	keys := make([]string, 0, len(Presets))
	for k := range Presets {
		keys = append(keys, k)
	}

	presetKeys = strings.Join(keys, ", ")
}
