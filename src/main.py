import os
from img_utils import *
import numpy as np
from PIL import Image, ImageGrab
from time import sleep

# Ensure the images directory exists
if not os.path.exists('images'):
    os.makedirs('images')

def water_grass(img_array):
    colormap = {
        (0x9c, 0xc0, 0xf9, 255): (0x47, 0x54, 0x8f, 255),
        (0xe5, 0xe3, 0xdf, 255): (0x54, 0x84, 0x49, 255)
    }
    return remap_colors(img_array, colormap)

def transform_image(img):
    old_array = np.array(img)

    new_array = water_grass(old_array)

    return Image.fromarray(new_array.astype('uint8'))

last_image = None
def get_clipboard_image():
    global last_image
    img = ImageGrab.grabclipboard()

    if img is None:
        print("No image in clipboard found.")
        return None
    if img == last_image:
        print("Duplicate image found.")
        return None
    
    last_image = img
    return img

# Initialize counter for filenames
counter = 1
def process_image():
    global counter

    clip_img = get_clipboard_image()
    if clip_img is None:
        return

    # Generate file names
    input_filename = f'images/in_{counter}.png'
    output_filename = f'images/out_{counter}.png'

    clip_img.save(input_filename)
    transform_image(clip_img).save(output_filename)

    print(f"Image saved as {input_filename} and processed as {output_filename}")

    counter += 1

if __name__ == "__main__":
    while True:
        process_image()
        sleep(0.5)
