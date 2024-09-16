import os
import re
import webbrowser
from PIL import ImageGrab
from time import sleep
from mindustry import convert_to_mindustry_map

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

def process_image(img):
    counter = get_highest_counter() + 1
    
    img.save(f'images/in_{counter}.png')
    print(f"Clipboard image saved as in_{counter}.png")

    game_map = convert_to_mindustry_map(img)
    game_map.save(f'images/out_{counter}.png')
    print(f"Mindustry map image saved as out_{counter}.png")


img_dir = "images"
map_link = "https://mapstyle.withgoogle.com"

if __name__ == "__main__":
    webbrowser.open(map_link)
    
    if not os.path.exists(img_dir):
        os.makedirs(img_dir)
    
    clipboard_image = ImageGrab.grabclipboard()
    if clipboard_image is None:
        print("Waiting for a clipboard image...")
    while clipboard_image == None:
        sleep(1.5)
        clipboard_image = ImageGrab.grabclipboard()
        
    process_image(clipboard_image)