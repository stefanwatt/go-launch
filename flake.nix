# TODO: Write README

# TODO: build out this flake
{
  description = "go-launch app launcher";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/release-23.11";
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs-unstable, nixpkgs, home-manager, ... }@inputs:
}

