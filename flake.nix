{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    inherit (nixpkgs) lib;
    systems = ["x86_64-linux" "x86_64-darwin" "64-linux" "aarch64-darwin"];
    eachSystem = fn:
      lib.listToAttrs (
        lib.map (system: {
          name = system;
          value = fn system;
        })
        systems
      );
  in {
    packages = eachSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      hello = pkgs.hello;
      disallow-large-dir = pkgs.callPackage ./disallow-large-dir {};
    });
    devShells = eachSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      default = pkgs.mkShell {
        packages = with pkgs; [
          go
        ];
      };
    });
  };
}
