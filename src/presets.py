from dataclasses import dataclass
from enum import Enum
from mindustry import *
from tiles import Tiles

@dataclass
class Preset():
    land: tuple = Tiles.grass.value
    coast_land: tuple = Tiles.sand.value
    coast_water: tuple = Tiles.sand_water.value
    shallow_water: tuple = Tiles.shallow_water.value
    deep_water: tuple = Tiles.deep_water.value

class Presets(Enum):
    @staticmethod
    def find_preset_by_name(name: str) -> Preset:
        for preset in Presets:
            if preset.name == name:
                return preset.value
        return None
    
    green = Preset()
    cold = Preset(land=Tiles.snow.value)
    arctic = Preset(land=Tiles.snow.value, coast_land=Tiles.ice.value, coast_water=Tiles.ice.value)
    desert = Preset(land=Tiles.sand.value, coast_land=Tiles.dark_sand.value, coast_water=Tiles.dark_sand_water.value)