from img_utils import *
from PIL import Image
import numpy as np
import time

# Colors of Google Maps with Silver style preset (https://mapstyle.withgoogle.com)
input_water = (201, 201, 201)
input_land = (245, 245, 245)

# Mindustry map-image colors
deep_water = (61, 73, 128)
shallow_water = (71, 84, 143)
sand_water = (116, 104, 122)
sand = (167, 137, 111)
grass = (84, 132, 73)
snow = (179, 183, 187)
ice = (162, 162, 194)

def remap_to(array):
    print("Mapping image to deep water and grass...")
    colormap = {
        input_water: deep_water,
        input_land: preset.land
    }
    return remap_colors(array, colormap)

def blend_shallow_water(array, radius = 10):
    print(f"Blending shallow water {radius} blocks around the coastline...")
    return blend(array, preset.deep_water, preset.shallow_water, preset.land, radius)

def blend_sand(array, radius = 6):
    print(f"Blending sand {radius} blocks around the coastline...")
    return blend(array, preset.land, preset.coast_land, preset.shallow_water, radius)

def blend_sand_water(array, radius = 3):
    print(f"Blending sand water {radius} blocks around the coastline...")
    return blend(array, preset.shallow_water, preset.coast_water, preset.coast_land, radius)

preset = None
def set_preset(new_preset):
    global preset
    preset = new_preset

def convert_to_mindustry_map(img):    
    print("Starting conversion...")
    start_time = time.time()
    
    array = remove_alpha(np.array(img))

    current_land = preset.land
    current_coast_water = preset.coast_water

    array = remap_to(array)
    array = blend_shallow_water(array)
    array = blend_sand(array)
    array = blend_sand_water(array)
    
    img = Image.fromarray(array.astype('uint8'))
    
    elapsed = time.time() - start_time
    print(f"Done! ({elapsed:.2f} s)")

    return img