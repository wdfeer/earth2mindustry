import osmnx as ox
from PIL import Image
import sys

def fetch_and_render_map(lon1, lat1, lon2, lat2):
    # Fetch the land and water data
    gdf = ox.features_from_bbox(bbox=(lat1, lon1, lat2, lon2), tags={'natural': True, 'landuse': True})

    # Initialize a grid to hold the data
    resolution = 1000
    img_width = resolution
    img_height = resolution

    # Create a blank image with white background
    image = Image.new('RGB', (img_width, img_height), 'white')
    pixels = image.load()

    # Define scaling factors to fit the bounding box into the image dimensions
    lat_min, lon_min, lat_max, lon_max = lat1, lon1, lat2, lon2
    lat_range = lat_max - lat_min
    lon_range = lon_max - lon_min

    def lat_lon_to_pixel(lat, lon):
        x = int((lon - lon_min) / lon_range * (img_width - 1))
        y = int((1 - (lat - lat_min) / lat_range) * (img_height - 1))
        return x, y

    # Render land and water areas
    for _, row in gdf.iterrows():
        geom = row['geometry']
        if geom.is_empty:
            continue
        
        if geom.geom_type == 'Polygon' or geom.geom_type == 'MultiPolygon':
            if 'natural' in row and row['natural'] == 'water':
                color = (0, 0, 255)  # Blue for water
            else:
                color = (0, 255, 0)  # Green for land
            
            if geom.geom_type == 'Polygon':
                for x, y in geom.exterior.coords:
                    px, py = lat_lon_to_pixel(y, x)
                    pixels[px, py] = color
            else:
                for poly in geom:
                    for x, y in poly.exterior.coords:
                        px, py = lat_lon_to_pixel(y, x)
                        pixels[px, py] = color

    # Save the image
    image.save('map_image.png')
    print("Map image saved as 'map_image.png'")

if __name__ == '__main__':
    if len(sys.argv) != 5:
        print("Usage: python src/main.py <lon1> <lat1> <lon2> <lat2>")
        sys.exit(1)

    lon1, lat1, lon2, lat2 = map(float, sys.argv[1:])
    fetch_and_render_map(lon1, lat1, lon2, lat2)
