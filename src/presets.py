from dataclasses import dataclass
from enum import Enum
from mindustry import *
from tiles import Tiles

@dataclass
class Preset():
    name: str
    land: tuple = Tiles.grass
    coast_land: tuple = Tiles.sand
    coast_water: tuple = Tiles.sand_water
    shallow_water: tuple = Tiles.shallow_water
    deep_water: tuple = Tiles.deep_water

class Presets(Enum):
    @staticmethod
    def find_preset_by_name(name: str) -> Preset:
        for preset in Presets:
            if preset.value.name == name:
                return preset.value
        return None
    
    green = Preset(name="green")
    cold = Preset(name="cold", land=Tiles.snow)
    arctic = Preset(name="arctic", land=Tiles.snow, coast_land=Tiles.ice, coast_water=Tiles.ice)