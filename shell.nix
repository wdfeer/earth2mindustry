{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.python3
    pkgs.python3Packages.pillow
    pkgs.python3Packages.numpy
    pkgs.python3Packages.scipy
  ] ++ (if builtins.getEnv "XDG_SESSION_TYPE" == "x11" then
          [ pkgs.xclip ]
        else if builtins.getEnv "XDG_SESSION_TYPE" == "wayland" then
          [ pkgs.wl-clipboard ]
        else
          []
       );
}
