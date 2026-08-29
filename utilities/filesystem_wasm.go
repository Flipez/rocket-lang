//go:build wasm

package utilities

import "syscall/js"

// PlaygroundRoot is where the playground's files live. A browser has no
// working directory, so imports need somewhere to resolve against.
const PlaygroundRoot = "/play"

// The playground has no filesystem, so it hands its files over as
// globalThis.rocketFiles: an object of absolute path to source text. Reading
// them here means nothing else has to know where its sources came from --
// resolving an import, reading a module and reading the entry file all go
// through the same FileSystem either way.
//
// A fresh WebAssembly instance runs each program, so taking a copy at init is
// enough; there is no later change to observe.
func init() {
	files := js.Global().Get("rocketFiles")
	if !files.Truthy() {
		return
	}

	names := js.Global().Get("Object").Call("keys", files)

	loaded := make(map[string][]byte, names.Length())
	for i := 0; i < names.Length(); i++ {
		name := names.Index(i).String()
		loaded[name] = []byte(files.Get(name).String())
	}

	SetFileSystem(MapFileSystem{Files: loaded})

	// initSearchPaths is a no-op under wasm, so this is the only entry in the
	// search path and a bare `import "util"` resolves to /play/util.rl. A
	// relative `import "./util"` resolves against the importing file's
	// directory, which is this same root.
	SearchPaths = append(SearchPaths, PlaygroundRoot)
}
