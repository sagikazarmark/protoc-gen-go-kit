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
              goreleaser
              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
            ];

            shellHook = ''
              go version
              golangci-lint --version
              gotestsum --version
              protoc --version
            '';
          };

          ci = devShells.default;
        };
      }
    );
}
