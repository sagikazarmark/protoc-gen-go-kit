{
  description = "Go Kit Protoc Compiler";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      rec
      {
        devShells = {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              go_1_20
              gnumake
              golangci-lint
              gotestsum
              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
            ];

            shellHook = ''
              make versions
            '';
          };

          ci = devShells.default.overrideAttrs (final: prev: {
            buildInputs = prev.buildInputs ++ (with pkgs; [
              goreleaser
              syft
              cosign
            ]);

            shellHook = ''
              ${prev.shellHook}
              goreleaser --version
              syft --version
              cosign version
            '';
          });
        };
      }
    );
}
