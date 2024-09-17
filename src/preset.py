from dataclasses import dataclass
from enum import Enum
from mindustry import *

@dataclass
class Preset():
    name: str
    land: tuple = grass
    coast_land: tuple = sand
    coast_water: tuple = sand_water
    shallow_water: tuple = shallow_water
    deep_water: tuple = deep_water

class Presets(Enum):
    @staticmethod
    def find_preset_by_name(name: str) -> Preset:
        for preset in Presets:
            if preset.value.name == name:
                return preset.value
        return None
    
    green = Preset(name="green")
    cold = Preset(name="cold", land=snow)
    arctic = Preset(name="arctic", land=snow, coast_land=ice, coast_water=ice)