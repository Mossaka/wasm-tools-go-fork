// Package wit contains a Go representation of the WIT ([WebAssembly Interface Type])
// specification as defined in the [WebAssembly Component Model].
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
// [WebAssembly Interface Type]: https://component-model.bytecodealliance.org/design/wit.html
// [WebAssembly Component Model]: https://component-model.bytecodealliance.org/introduction.html
// [wit-parser]: https://docs.rs/wit-parser/latest/wit_parser/
// [source]: https://github.com/bytecodealliance/wasm-tools/tree/main/crates/wit-parser
// [wasm-tools]: https://crates.io/crates/wasm-tools
package wit
