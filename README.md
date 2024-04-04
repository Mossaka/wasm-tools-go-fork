# wasm-tools-go

[![build status](https://img.shields.io/github/actions/workflow/status/ydnar/wasm-tools-go/test.yaml?branch=main)](https://github.com/ydnar/wasm-tools-go/actions)
[![pkg.go.dev](https://img.shields.io/badge/docs-pkg.go.dev-blue.svg)](https://pkg.go.dev/github.com/ydnar/wasm-tools-go)

## [WebAssembly](https://webassembly.org) + [WASI](https://wasi.dev) tools for [Go](https://go.dev)

## About

This repository contains code to generate Go bindings for [Component Model](https://component-model.bytecodealliance.org/) interfaces defined in [WIT](https://component-model.bytecodealliance.org/design/wit.html) (WebAssembly Interface Type) files. The goal of this project is to accelerate adoption of the Component Model and development of [WASI Preview 2](https://bytecodealliance.org/articles/WASI-0.2) in Go.

### Component Model

Package [cm](./cm) contains helper types and functions used by generated packages, such as `option<t>`, `result<ok, err>`, `variant`, `list`, and `resource`. These are intended for use by generated [Component Model](https://github.com/WebAssembly/component-model/blob/main/design/mvp/Explainer.md#type-definitions) bindings, where the caller converts to a Go equivalent. It attempts to map WIT semantics to their equivalent in Go where possible.

#### Note on Memory Safety

Package `cm` and generated bindings from `wit-bindgen-go` may have compatibility issues with the Go garbage collector, as they directly represent `variant` and `result` types as tagged unions where a pointer shape may be occupied by a non-pointer value. The GC may detect and throw an error if it detects a non-pointer value in an area it expects to see a pointer. This is an area of active development.

## `wit-bindgen-go`

### WIT → Go

The `wit-bindgen-go` tool can generate Go bindings for WIT interfaces and worlds. If [`wasm-tools`](https://crates.io/crates/wasm-tools) is installed and in `$PATH`, then `wit-bindgen-go` can load WIT directly.

```sh
wit-bindgen-go generate ../wasi-cli/wit
```

Otherwise, pass the JSON representation of a fully-resolved WIT package:

```sh
wit-bindgen-go generate wasi-cli.wit.json
```

Or pipe via `stdin`:

```sh
wasm-tools component wit -j ../wasi-cli/wit | wit-bindgen-go generate
```

### JSON → WIT

For debugging purposes, `wit-bindgen-go` can also convert a JSON representation back into WIT. This is useful for validating that the intermediate representation faithfully represents the original WIT source.

```sh
wit-bindgen-go wit example.wit.json
```

### WIT → JSON

The [wit](./wit) package can decode a JSON representation of a fully-resolved WIT file. Serializing WIT into JSON requires [wasm-tools](https://crates.io/crates/wasm-tools) v1.0.42 or higher. To convert a WIT file into JSON, run `wasm-tools` with the `-j` argument:

```sh
wasm-tools component wit -j example.wit
```

This will emit JSON on `stdout`, which can be piped to a file or another program.

```sh
wasm-tools component wit -j example.wit > example.wit.json
```

## License

This project is licensed under the Apache 2.0 license with the LLVM exception. See [LICENSE](LICENSE) for more details.

### Contributing

Developing and testing `wit-bindgen-go` requires an up-to-date installation of [Go](https://go.dev), [TinyGo](https://tinygo.org), and [`wasm-tools`](https://crates.io/crates/wasm-tools).

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in this project by you, as defined in the Apache-2.0 license, shall be licensed as above, without any additional terms or conditions.
