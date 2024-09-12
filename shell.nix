{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.python3
    pkgs.python3Packages.pillow
    pkgs.python3Packages.numpy
    pkgs.python3Packages.scipy
  ];
}
