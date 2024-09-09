import numpy as np
from scipy.ndimage import binary_dilation

def remove_alpha(array):
    """
    Remove the alpha channel from each pixel in the array.

    Parameters:
        array (np.ndarray): A 3D NumPy array representing the image with shape (height, width, 4).
                            The last dimension contains RGBA values.

    Returns:
        np.ndarray: A 3D NumPy array with the alpha channel removed (RGB values only).
    """
    # Remove the alpha channel if present (shape (height, width, 4))
    return array[:, :, :3] if array.shape[-1] == 4 else array

def closest_color(pixel, colormap):
    """
    Find the closest color in the colormap to the given pixel using Euclidean distance.

    Parameters:
        pixel (array-like): The RGB values of the pixel.
        colormap (dict): A dictionary where keys are color tuples and values are mapped colors.

    Returns:
        tuple: The color from the colormap closest to the pixel.
    """
    # Vectorize color keys and compute distances using broadcasting
    keys = np.array(list(colormap.keys()))
    distances = np.linalg.norm(keys - pixel, axis=1)
    closest_idx = np.argmin(distances)
    return colormap[tuple(keys[closest_idx])]

def remap_colors(array, colormap):
    """
    Remap the colors of an image array using a colormap.

    Parameters:
        array (np.ndarray): A 3D NumPy array representing the image.
        colormap (dict): A dictionary where keys are color tuples and values are the new mapped colors.

    Returns:
        np.ndarray: A new image array with remapped colors.
    """
    # Reshape the array for easier vectorized operations
    reshaped_array = array.reshape(-1, 3)
    remapped_array = np.array([closest_color(pixel, colormap) for pixel in reshaped_array])
    return remapped_array.reshape(array.shape)

def blend(array, target_color, new_color, required_color, radius):
    """
    Blend pixels of the target color into a new color if the required color is found within a given radius.

    Parameters:
        array (np.ndarray): A 3D NumPy array representing the image.
        target_color (list): The RGB values of the target color.
        new_color (list): The RGB values of the new color to blend to.
        required_color (list): The RGB values of the required color for blending.
        radius (int): The radius within which the required color should be found.

    Returns:
        np.ndarray: A new image array with blended colors.
    """
    height, width, _ = array.shape
    target_color = np.array(target_color)
    required_color = np.array(required_color)
    new_color = np.array(new_color)

    # Create masks for target and required colors
    target_mask = np.all(array == target_color, axis=-1)
    required_mask = np.all(array == required_color, axis=-1)

    # Dilate the required color mask by the given radius
    struct_elem = np.ones((2 * radius + 1, 2 * radius + 1), dtype=bool)
    required_dilated = binary_dilation(required_mask, structure=struct_elem)

    # Find where the target color overlaps with the dilated required color mask
    blend_mask = target_mask & required_dilated

    # Apply the new color where blending conditions are met
    new_array = array.copy()
    new_array[blend_mask] = new_color

    return new_array
