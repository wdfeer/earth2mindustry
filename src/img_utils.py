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
	"""
	Find the closest color in the colormap to the given pixel using Euclidean distance.

	Parameters:
		pixel (array-like): The RGB values of the pixel.
		colormap (dict): A dictionary where keys are color tuples and values are mapped colors.

	Returns:
		tuple: The color from the colormap closest to the pixel.
	"""
	# Calculate the Euclidean distances between the pixel and each color in the colormap
	distances = [np.sqrt(np.sum((pixel - np.array(key))**2)) for key in colormap.keys()]
	closest = list(colormap.keys())[distances.index(min(distances))]
	return colormap[closest]

def remap_colors(array, colormap):
	"""
	Remap the colors of an image array using a colormap.

	Parameters:
		array (np.ndarray): A 3D NumPy array representing the image.
		colormap (dict): A dictionary where keys are color tuples and values are the new mapped colors.

	Returns:
		np.ndarray: A new image array with remapped colors.
	"""
	height, width, _ = array.shape
	new_array = np.empty(array.shape, dtype=tuple)

	# Iterate over each pixel and remap it using the closest color from the colormap
	for y in range(height):
		for x in range(width):
			new_array[y, x] = closest_color(array[y, x], colormap)
	
	return new_array

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
	def is_within_radius(p1, p2, r):
		# Check if two points are within a given radius
		return np.sqrt(np.sum((np.array(p1) - np.array(p2))**2)) <= r

	height, width, _ = array.shape
	new_array = np.copy(array)
	
	# Convert the color lists to tuples for comparison
	target_color = tuple(target_color)
	required_color = tuple(required_color)
	new_color = tuple(new_color)

	def process_pixel(x, y):
		if tuple(array[y, x]) == target_color:
			# Check the surrounding pixels within the radius
			found_required_color = False
			for dy in range(-radius, radius + 1):
				for dx in range(-radius, radius + 1):
					ny, nx = y + dy, x + dx
					# Ensure the surrounding pixel is within bounds and within the radius
					if 0 <= ny < height and 0 <= nx < width:
						if tuple(array[ny, nx]) == required_color: # and is_within_radius((x, y), (nx, ny), radius):
							found_required_color = True
							break
				if found_required_color:
					break
			
			# Replace the target pixel with the new color if the required color is found
			if found_required_color:
				new_array[y, x] = new_color

	# Iterate over each pixel in the image
	for y in range(height):
		for x in range(width):
			process_pixel(x, y)

	return new_array
