import numpy as np

def closest_color(pixel, colormap):
    pixel = tuple(pixel)
    distances = [np.sqrt(np.sum((np.array(pixel) - np.array(key))**2)) for key in colormap.keys()]
    closest = list(colormap.keys())[distances.index(min(distances))]
    return colormap[closest]

def remap_colors(img_array, colormap):
    height, width, _ = img_array.shape
    new_array = np.empty(img_array.shape, dtype=tuple)

    for y in range(height):
        for x in range(width):
            new_array[y, x] = closest_color(img_array[y, x], colormap)
    
    return new_array
