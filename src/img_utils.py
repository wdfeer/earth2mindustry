import numpy as np

def remove_alpha(array):
    """
    Remove the alpha channel from each pixel in the array.
    
    Parameters:
        array (np.ndarray): A 3D NumPy array representing the image with shape (height, width, 4).
                            The last dimension contains RGBA values.
    
    Returns:
        np.ndarray: A 3D NumPy array with the alpha channel removed (RGB values only).
    """
    # Ensure the input array has an alpha channel to remove (i.e., shape (height, width, 4))
    if array.shape[-1] == 4:
        # Remove the alpha channel by slicing the array to keep only the first three channels (RGB)
        return array[:, :, :3]
    else:
        # Return the array as is if it doesn't have an alpha channel
        return array

def closest_color(pixel, colormap):
    distances = [np.sqrt(np.sum((pixel - np.array(key))**2)) for key in colormap.keys()]
    closest = list(colormap.keys())[distances.index(min(distances))]
    return colormap[closest]

def remap_colors(array, colormap):
    height, width, _ = array.shape
    new_array = np.empty(array.shape, dtype=tuple)

    for y in range(height):
        for x in range(width):
            new_array[y, x] = closest_color(array[y, x], colormap)
    
    return new_array

def blend(array, target_color, new_color, required_color, radius):
    def is_within_radius(p1, p2, r):
        return np.sqrt(np.sum((np.array(p1) - np.array(p2))**2)) <= r

    height, width, _ = array.shape
    new_array = np.copy(array)
    
    target_color = tuple(target_color)
    required_color = tuple(required_color)
    new_color = tuple(new_color)

    for y in range(height):
        for x in range(width):
            if tuple(array[y, x]) == target_color:
                # Check the surrounding pixels within the radius
                found_required_color = False
                for dy in range(-radius, radius + 1):
                    for dx in range(-radius, radius + 1):
                        ny, nx = y + dy, x + dx
                        if 0 <= ny < height and 0 <= nx < width:
                            if is_within_radius((x, y), (nx, ny), radius) and tuple(array[ny, nx]) == required_color:
                                found_required_color = True
                                break
                    if found_required_color:
                        break
                
                if found_required_color:
                    new_array[y, x] = new_color

    return new_array