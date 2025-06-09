{pkgs ? import <nixpkgs> {}}:
pkgs.buildGoModule {
  name = "create-members-page";
  version = "0.1.0";
  src = ./.;
  goPackagePath = "github.com/ut-code/internal-helper/create-members-page";
  vendorHash = "sha256-27V+oK6SAynnw+TM7bbCi8jN/nWk8hGW25DBTAy3WBY=";
  nativeBuildInputs = [
    pkgs.imagemagick
  ];
}
