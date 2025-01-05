from img_utils import *
from PIL import Image
import numpy as np
import time
from presets import Presets
from tiles import Tiles

# Standard Google Maps colors
standard_input_water = (132, 215, 235)
standard_input_grass = (207, 246, 224)
standard_input_forest = (154, 229, 194)
standard_input_sand = (245, 240, 230)
standard_input_white = (247, 247, 247) # either city or snow
standard_input_road = (167, 187, 214)
# Colors of Google Maps with Silver style, for custom presets
silver_input_water = (201, 201, 201)
silver_input_land = (245, 245, 245)

# for Silver => Custom
def create_land_and_water(array):
    print("Mapping image to deep water and land...")
    colormap = {
        silver_input_water: preset.deep_water,
        silver_input_land: preset.land
    }
    return remap_colors(array, colormap)
def blend_shallow_water(array, radius = 10):
    print(f"Blending shallow water {radius} blocks around the coastline...")
    return blend(array, preset.deep_water, preset.shallow_water, preset.land, radius)
def blend_coast_land(array, radius = 6):
    print(f"Blending coast {radius} blocks around the coastline...")
    return blend(array, preset.land, preset.coast_land, preset.shallow_water, radius)
def blend_coast_water(array, radius = 3):
    print(f"Blending coast water {radius} blocks around the coastline...")
    return blend(array, preset.shallow_water, preset.coast_water, preset.coast_land, radius)

preset = None
def set_preset(new_preset):
    global preset
    preset = new_preset

def convert_to_mindustry_map(img):    
    print("Starting conversion...")
    start_time = time.time()
    
    array = remove_alpha(np.array(img))

    if (preset == Presets.standard.value):
        # Standard Google Maps style
        colormap = {
            standard_input_water: Tiles.shallow_water.value,
            standard_input_grass: Tiles.grass.value,
            standard_input_forest: Tiles.grass.value,
            standard_input_white: Tiles.snow.value,
            standard_input_road: Tiles.stone.value,
            standard_input_sand: Tiles.sand.value
        }
        print("Converting the standard Google Maps image to mindustry tiles...")
        array = remap_colors(array, colormap)
    else:
        # Silver Google Maps style => custom mindustry style
        array = create_land_and_water(array)
        array = blend_shallow_water(array)
        array = blend_coast_land(array)
        array = blend_coast_water(array)
    
    img = Image.fromarray(array.astype('uint8'))
    
    elapsed = time.time() - start_time
    print(f"Done! ({elapsed:.2f} s)")

    return img