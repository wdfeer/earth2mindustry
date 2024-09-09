from img_utils import remap_colors
import numpy as np
from PIL import Image

def create_water_and_grass(img_array):
    colormap = {
        (0x9c, 0xc0, 0xf9, 255): (0x47, 0x54, 0x8f, 255), # water
        (0xe5, 0xe3, 0xdf, 255): (0x54, 0x84, 0x49, 255) # grass
    }
    return remap_colors(img_array, colormap)

def convert_to_mindustry_map(img):
    print("Starting conversion...")
    
    old_array = np.array(img)

    new_array = create_water_and_grass(old_array)

    return Image.fromarray(new_array.astype('uint8'))