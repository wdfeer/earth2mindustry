{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.xclip # for accessing the clipboard on X11 systems
    pkgs.python3
    pkgs.python3Packages.pillow
    pkgs.python3Packages.numpy
  ];

  shellHook = ''
    echo "Welcome to the Python development shell!"
  '';
}
