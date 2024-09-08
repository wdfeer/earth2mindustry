{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.python3
    pkgs.python3Packages.pillow
    pkgs.python3Packages.osmnx
  ];

  shellHook = ''
    echo "Welcome to the Python development shell!"
  '';
}
