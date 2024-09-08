INPUT=$1
OUTPUT=$2

magick -size 2x1 xc:"#e5e3df" xc:"#9cc0f9" +append colormap.png
magick $INPUT -fuzz 20% -remap colormap.png $OUTPUT
magick $OUTPUT -fill "#548449" -opaque "#e5e3df" -fill "#47548f" -opaque "#9cc0f9" $OUTPUT

rm colormap.png
