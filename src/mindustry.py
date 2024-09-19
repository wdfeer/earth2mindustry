from img_utils import *
from PIL import Image
import numpy as np
import time

# Colors of Google Maps with Silver style preset (https://mapstyle.withgoogle.com)
input_water = (201, 201, 201)
input_land = (245, 245, 245)

def create_land_and_water(array):
    print("Mapping image to deep water and land...")
    colormap = {
        input_water: preset.deep_water,
        input_land: preset.land
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

    array = create_land_and_water(array)
    array = blend_shallow_water(array)
    array = blend_coast_land(array)
    array = blend_coast_water(array)
    
    img = Image.fromarray(array.astype('uint8'))
    
    elapsed = time.time() - start_time
    print(f"Done! ({elapsed:.2f} s)")

    return img