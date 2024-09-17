import os
import re
import sys
import webbrowser
from PIL import ImageGrab
from time import sleep
from mindustry import convert_to_mindustry_map, set_preset
from preset import Presets

def get_highest_counter():
    regex = re.compile(r'_(\d+)\.png$')

    highest_number = -1

    for filename in os.listdir(img_dir):
        match = regex.search(filename)
        if match:
            file_number = int(match.group(1))
            if file_number > highest_number:
                highest_number = file_number

    return highest_number

def process_image(img, preset):
    counter = get_highest_counter() + 1
    
    img.save(f'images/in_{counter}.png')
    print(f"Clipboard image saved as in_{counter}.png")

    set_preset(preset)
    game_map = convert_to_mindustry_map(img)
    game_map.save(f'images/out_{counter}.png')
    print(f"Mindustry map image saved as out_{counter}.png")


img_dir = "images"
map_link = "https://mapstyle.withgoogle.com"

if __name__ == "__main__":
    if len(sys.argv) <= 1:
        print("Preset argument not specified. Defaulting to 'green'.")
        preset = Presets.green
    else:
        preset = Presets.find_preset_by_name(sys.argv[1])
        if preset == None:
            print("Invalid preset. Defaulting to 'green'.")
            preset = Presets.green
    
    clipboard_image = ImageGrab.grabclipboard()
    if clipboard_image is None:
        print(f"Opening {map_link}")
        webbrowser.open(map_link)
        print("Waiting for a clipboard image...")
    while clipboard_image == None:
        sleep(1.5)
        clipboard_image = ImageGrab.grabclipboard()
    
    if not os.path.exists(img_dir):
        os.makedirs(img_dir)
    
    process_image(clipboard_image, preset)