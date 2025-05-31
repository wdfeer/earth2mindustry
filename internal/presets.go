package internal

import (
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
			{R: 132, G: 215, B: 235}: tiles.DeepWater,
			{R: 207, G: 246, B: 224}: tiles.Grass,
			{R: 154, G: 229, B: 194}: tiles.Grass,
			{R: 245, G: 240, B: 230}: tiles.Sand,
			{R: 247, G: 247, B: 247}: tiles.Stone,
			{R: 25, G: 25, B: 25}:    tiles.DarkSand,
			{R: 167, G: 187, B: 214}: tiles.DarkSand,
		}

	Presets["snow"] =
		map[px.Pixel]px.Pixel{
			{R: 132, G: 215, B: 235}: tiles.DeepWater,
			{R: 207, G: 246, B: 224}: tiles.Snow,
			{R: 154, G: 229, B: 194}: tiles.Grass,
			{R: 245, G: 240, B: 230}: tiles.Sand,
			{R: 247, G: 247, B: 247}: tiles.Snow,
			{R: 167, G: 187, B: 214}: tiles.DarkSand,
		}

	Presets["desert"] =
		map[px.Pixel]px.Pixel{
			{R: 132, G: 215, B: 235}: tiles.DeepWater,
			{R: 207, G: 246, B: 224}: tiles.Snow,
			{R: 154, G: 229, B: 194}: tiles.Grass,
			{R: 245, G: 240, B: 230}: tiles.Sand,
			{R: 247, G: 247, B: 247}: tiles.Sand,
			{R: 167, G: 187, B: 214}: tiles.DarkSand,
		}

	keys := make([]string, 0, len(Presets))
	for k := range Presets {
		keys = append(keys, k)
	}

	presetKeys = strings.Join(keys, ", ")
}
