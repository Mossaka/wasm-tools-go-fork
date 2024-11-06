// Code generated by wit-bindgen-go. DO NOT EDIT.

package streams

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:io@0.2.0".

//go:wasmimport wasi:io/streams@0.2.0 [resource-drop]input-stream
//go:noescape
func wasmimport_InputStreamResourceDrop(self0 uint32)

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.blocking-read
//go:noescape
func wasmimport_InputStreamBlockingRead(self0 uint32, len0 uint64, result *cm.Result[cm.List[uint8], cm.List[uint8], StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.blocking-skip
//go:noescape
func wasmimport_InputStreamBlockingSkip(self0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.read
//go:noescape
func wasmimport_InputStreamRead(self0 uint32, len0 uint64, result *cm.Result[cm.List[uint8], cm.List[uint8], StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.skip
//go:noescape
func wasmimport_InputStreamSkip(self0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.subscribe
//go:noescape
func wasmimport_InputStreamSubscribe(self0 uint32) (result0 uint32)

//go:wasmimport wasi:io/streams@0.2.0 [resource-drop]output-stream
//go:noescape
func wasmimport_OutputStreamResourceDrop(self0 uint32)

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-flush
//go:noescape
func wasmimport_OutputStreamBlockingFlush(self0 uint32, result *cm.Result[StreamError, struct{}, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-splice
//go:noescape
func wasmimport_OutputStreamBlockingSplice(self0 uint32, src0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-write-and-flush
//go:noescape
func wasmimport_OutputStreamBlockingWriteAndFlush(self0 uint32, contents0 *uint8, contents1 uint32, result *cm.Result[StreamError, struct{}, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-write-zeroes-and-flush
//go:noescape
func wasmimport_OutputStreamBlockingWriteZeroesAndFlush(self0 uint32, len0 uint64, result *cm.Result[StreamError, struct{}, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.check-write
//go:noescape
func wasmimport_OutputStreamCheckWrite(self0 uint32, result *cm.Result[uint64, uint64, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.flush
//go:noescape
func wasmimport_OutputStreamFlush(self0 uint32, result *cm.Result[StreamError, struct{}, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.splice
//go:noescape
func wasmimport_OutputStreamSplice(self0 uint32, src0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.subscribe
//go:noescape
func wasmimport_OutputStreamSubscribe(self0 uint32) (result0 uint32)

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.write
//go:noescape
func wasmimport_OutputStreamWrite(self0 uint32, contents0 *uint8, contents1 uint32, result *cm.Result[StreamError, struct{}, StreamError])

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.write-zeroes
//go:noescape
func wasmimport_OutputStreamWriteZeroes(self0 uint32, len0 uint64, result *cm.Result[StreamError, struct{}, StreamError])
