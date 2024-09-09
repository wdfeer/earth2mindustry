import os
from PIL import ImageGrab
from time import sleep
from mindustry import *

# Ensure the images directory exists
if not os.path.exists('images'):
    os.makedirs('images')

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

    clip_img.save(f'images/in_{counter}.png')
    print(f"Clipboard image saved as in_{counter}.png")

    game_map = convert_to_mindustry_map(clip_img)
    game_map.save(f'images/out_{counter}.png')
    print(f"Mindustry map image saved as out_{counter}.png")

    counter += 1

if __name__ == "__main__":
    while True:
        process_image()
        sleep(0.5)
