{pkgs ? import <nixpkgs> {}}:
pkgs.buildGoModule {
  name = "disallow-large-dir";
  version = "0.1.0";
  src = ./.;
  goPackagePath = "github.com/ut-code/internal-helper/disallow-large-dir";
  vendorHash = "sha256-7d8+WhF1LZi0AjAFiM3/h8w8ttlLrTcolHwQLVquokc=";
}
