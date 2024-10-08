{
  description = "MGA: Modern Go Application tool";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    devenv.url = "github:cachix/devenv";
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        inputs.devenv.flakeModule
      ];

      systems = [ "x86_64-linux" "x86_64-darwin" "aarch64-darwin" ];

      perSystem = { config, self', inputs', pkgs, system, ... }: rec {
        devenv.shells = {
          default = {
            languages = {
              go = {
                enable = true;
                package = pkgs.go_1_23;
              };
            };

            packages = with pkgs; [
              gnumake
              go-task

              gotestsum
              golangci-lint

              goreleaser

              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
            ];

            scripts = {
              versions.exec = ''
                go version
                golangci-lint version
              '';
            };

            enterShell = ''
              versions
            '';

            # https://github.com/cachix/devenv/issues/528#issuecomment-1556108767
            containers = pkgs.lib.mkForce { };
          };

          ci = {
            imports = [ devenv.shells.default ];

            packages = with pkgs; [
              goreleaser
              syft
              cosign
            ];
          };
        };
      };
    };
}
