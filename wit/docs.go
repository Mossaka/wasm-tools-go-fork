// Package wit contains a Go representation of the WIT (WebAssembly Interface Type) specification.
//
// # WIT
//
// WIT ([WebAssembly Interface Type]) is an interface definition language with rich types, functions, and methods,
// used to define the interface of a [Component].
//
// Note: this package depends on the [wasm-tools] exectuable to parse WIT into an intermediary JSON representation.
//
// # Structure
//
// A fully-decoded [Resolve] struct is a cyclical data structure containing [World], [Interface], and [TypeDef]
// values that reference each other. Tools to decode polymorphic, cyclical JSON are implemented in the [codec] package.
// Types defined in this package largely map 1:1 to their equivalent types in the [wit-parser] crate ([source]),
// accommodating differences between the Go and Rust type systems.
//
// Types that are represented as Rust enums are implemented via sealed Go interfaces, implemented by other types
// in this package. An example is [WorldItem], which is the set of types that a [World] may
// import or export, currently [InterfaceRef], [TypeDef], and [Function].
//
// # JSON
//
// [DecodeJSON] decodes fully-resolved WIT descriptions in JSON format generated by [wasm-tools] into a [Resolve]:
//
//	f, err := os.Open("wasi-clocks.wit.json")
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	res, err := wit.DecodeJSON(f)
//	if err != nil {
//		return err
//	}
//
//	// Do something with res
//
// [wit-parser]: https://docs.rs/wit-parser/latest/wit_parser/
// [source]: https://github.com/bytecodealliance/wasm-tools/tree/main/crates/wit-parser
// [wasm-tools]: https://crates.io/crates/wasm-tools
// [WebAssembly Interface Type]: https://component-model.bytecodealliance.org/design/wit.html
// [Component]: https://component-model.bytecodealliance.org/introduction.html
package wit

// Docs represent WIT documentation text extracted from comments.
type Docs struct {
	Contents string // may be empty
}
