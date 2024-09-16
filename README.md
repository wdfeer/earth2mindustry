# Earth to Mindustry Converter

## Usage

1. **Install dependencies:**
   - If you have [Nix](https://nixos.org/download) installed, run:
     ```bash
     nix-shell
     ```
   - Otherwise install dependencies manually based on the `shell.nix` file.

2. **Run the script:**
   - Execute:
     ```bash
     python src/main.py
     ```
   - It opens the [Google Map Styler](https://mapstyle.withgoogle.com).

3. **Capture the screenshot:**
   - Select the **Silver** style, and disable **roads**, **landmarks**, and **labels**.
   - Take a screenshot of the desired map region and copy it to the clipboard.

4. **Import to Mindustry:**
   - Locate `out_1.png` in the `images` folder.
   - In Mindustry, create a new map and use the **Import** button to add your map.