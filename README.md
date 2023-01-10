<h1 align="center">PolyGen</h1>

<p align="center">
    CLI for generating Polybase Collections API code
</p>

## Install PolyGen

```bash
go install github.com/durudex/polygen/cmd/polygen@latest
```

## Usage

1) To start working with PolyGen, you need to create a configuration according
to which the code will be generated.

```yml
collection:
  - "Collection"

language:
  go:
    package: "generated"
    directory: "generated"
```

> You can find the complete configuration in the 
> [.polygen.example.yml](.polygen.example.yml) file.

2) Now you need to run PolyGen with the `--config` or `-c` flag and specify the
path to your configuration file.

```bash
polygen --config .polygen.yml
```

## ⚠️ License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
