{pkgs ? import <nixpkgs> {}}:
pkgs.buildGoModule {
  name = "disallow-large-dir";
  version = "0.1.0";
  src = ./.;
  goPackagePath = "github.com/ut-code/internal-helper/disallow-large-dir";
  vendorHash = "sha256-nb6HFTfngEMF2n0bZj+Lz/U6rVHd87kvgu07QExPt8g=";
}
