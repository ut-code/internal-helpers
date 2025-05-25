{pkgs ? import <nixpkgs> {}}:
pkgs.buildGoModule {
  name = "create-members-page";
  version = "0.1.0";
  src = ./.;
  goPackagePath = "github.com/ut-code/internal-helper/create-members-page";
  vendorHash = "sha256-1BRIrH874SjR4vDQ+6jkE+bkv884SmHrkiZqDc+1Vss=";
  nativeBuildInputs = [
    pkgs.imagemagick
  ];
}
