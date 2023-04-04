# [Go Kit](https://github.com/go-kit/kit/) Protoc Compiler

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/protoc-gen-go-kit/ci.yaml?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/protoc-gen-go-kit?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/protoc-gen-go-kit)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/protoc-gen-go-kit)
[![built with nix](https://img.shields.io/badge/builtwith-nix-7d81f7?style=flat-square)](https://builtwithnix.org)

**Protobuf compiler plugin for [Go Kit](https://github.com/go-kit/kit/).**


**IMPORTANT: This project is work in progress.**


## Installation

Download prebuilt binaries from the [releases](https://github.com/sagikazarmark/protoc-gen-go-kit/releases) page or install it from source:

```shell
go install github.com/sagikazarmark/protoc-gen-go-kit@latest
```


## Development

When all coding and testing is done, please run the test suite:

```shell
make check
```

For the best developer experience, install [Nix](https://builtwithnix.org/) and [direnv](https://direnv.net/).

Alternatively, install Go manually or using a package manager. Install the rest of the dependencies by running `make deps`.


## License

The project is licensed under the [MIT License](LICENSE).
