{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.xclip
    pkgs.python3
    pkgs.python3Packages.pillow
    pkgs.python3Packages.numpy
  ];

  shellHook = ''
    echo "Welcome to the Python development shell!"
  '';
}
