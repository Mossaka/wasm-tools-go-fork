package wit

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// LoadJSON loads a [WIT] JSON file from path.
// If path is "" or "-", it reads from os.Stdin.
//
// [WIT]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
func LoadJSON(path string) (*Resolve, error) {
	r := reader(path)
	if r != nil {
		return DecodeJSON(r)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return DecodeJSON(f)
}

// LoadWIT loads [WIT] data from path by processing it through [wasm-tools].
// This will fail if wasm-tools is not in $PATH.
// If path is "" or "-", it reads from os.Stdin.
//
// [WIT]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
// [wasm-tools]: https://crates.io/crates/wasm-tools
func LoadWIT(path string) (*Resolve, error) {
	r := reader(path)
	return loadWIT(path, r)
}

// ParseWIT parses [WIT] data from a buffer by processing it through [wasm-tools].
// This will fail if wasm-tools is not in $PATH.
//
// [WIT]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
// [wasm-tools]: https://crates.io/crates/wasm-tools
func ParseWIT(buffer []byte) (*Resolve, error) {
	r := bytes.NewReader(buffer)
	return loadWIT("", r)
}

// loadWIT loads WIT data from path or reader by processing it through wasm-tools.
// It accepts either a path or an io.Reader as input, but not both.
// If the path is not "" and "-", it will be used as the input file.
// Otherwise, the reader will be used as the input.
func loadWIT(path string, reader io.Reader) (*Resolve, error) {
	if (path != "" && path != "-") && reader != nil {
		return nil, errors.New("cannot set both path and reader; provide only one")
	}

	wasmTools, err := exec.LookPath("wasm-tools")
	if err != nil {
		return nil, err
	}

	var stdout, stderr bytes.Buffer

	cmdArgs := []string{"component", "wit", "-j", "--all-features"}
	if path != "" && path != "-" {
		cmdArgs = append(cmdArgs, path)
	}

	cmd := exec.Command(wasmTools, cmdArgs...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = reader

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, stderr.String())
		return nil, err
	}

	return DecodeJSON(&stdout)
}

func reader(path string) io.ReadCloser {
	if path == "" || path == "-" {
		return os.Stdin
	}
	return nil
}
