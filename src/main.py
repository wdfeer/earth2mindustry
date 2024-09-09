import os
import numpy as np
from PIL import Image, ImageGrab
from time import sleep

# Ensure the images directory exists
if not os.path.exists('images'):
    os.makedirs('images')

# Helper function to map colors
def map_colors(img_array):
    # Define the color mappings
    colormap = {
        (156, 192, 249): (0x47, 0x54, 0x8f, 255), # 9cc0f9 -> 47548f
        (229, 227, 223): (0x54, 0x84, 0x49, 255)  # e5e3df -> 548449
    }
    
    # Function to find the closest color in the colormap
    def closest_color(pixel):
        pixel = tuple(pixel[:3])  # Use only RGB values, ignoring alpha if present
        distances = [np.sqrt(np.sum((np.array(pixel) - np.array(key))**2)) for key in colormap.keys()]
        closest = list(colormap.keys())[distances.index(min(distances))]
        return colormap[closest]

    height, width, _ = img_array.shape
    new_array = np.empty(img_array.shape, dtype=tuple)

    for y in range(height):
        for x in range(width):
            new_array[y, x] = closest_color(img_array[y, x])
    
    return new_array

# Initialize counter for filenames
counter = 1
last_image = None

def process_image():
    global counter
    global last_image
    # Grab image from clipboard
    img = ImageGrab.grabclipboard()
    
    if img is None or img == last_image:
        print("No new image found in clipboard.")
        return

    # Generate file names
    input_filename = f'images/in_{counter}.png'
    output_filename = f'images/out_{counter}.png'

    # Save the image from the clipboard
    img.save(input_filename)

    # Convert PIL image to numpy array
    img_array = np.array(img)

    # Process the image
    processed_array = map_colors(img_array)
    
    # Convert numpy array back to PIL Image
    processed_img = Image.fromarray(processed_array.astype('uint8'))

    # Save the processed image
    processed_img.save(output_filename)

    print(f"Image saved as {input_filename} and processed as {output_filename}")

    counter += 1
    last_image = img

if __name__ == "__main__":
    while True:
        process_image()
        sleep(0.5)
