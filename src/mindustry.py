from img_utils import *
from PIL import Image
import numpy as np
import time

# Colors of Google Maps with Silver style preset (https://mapstyle.withgoogle.com)
input_water = (201, 201, 201)
input_grass = (245, 245, 245)

# Mindustry map-image colors
deep_water = (61, 73, 128)
water = (71, 84, 143)
grass = (84, 132, 73)
sand = (167, 137, 111)

def remap_to_water_and_grass(array):
    print("Mapping image to deep water and grass...")
    colormap = {
        input_water: deep_water,
        input_grass: grass
    }
    return remap_colors(array, colormap)

def blend_shallow_water(array, radius = 8):
    print(f"Blending shallow water {radius} blocks around the coastline...")
    return blend(array, deep_water, water, grass, radius)

def blend_sand(array, radius = 4):
    print(f"Blending sand {radius} blocks around the coastline...")
    return blend(array, grass, sand, water, radius)

def convert_to_mindustry_map(img):
    print("Starting conversion...")
    start_time = time.time()
    
    array = remove_alpha(np.array(img))

    array = remap_to_water_and_grass(array)
    array = blend_shallow_water(array)
    array = blend_sand(array)
    
    img = Image.fromarray(array.astype('uint8'))
    
    elapsed = time.time() - start_time
    print(f"Done! ({elapsed:.2f} s)")

    return img