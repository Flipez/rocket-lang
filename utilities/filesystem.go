package utilities

import "os"

// FileSystem is where the interpreter looks for module sources. The default
// reads the real filesystem, which is what every caller wants; the playground
// installs one backed by the editor's tabs, because a browser has no
// filesystem for an import to resolve against.
//
// Two methods is the whole surface: resolving an import stats candidate paths
// and reads the one that exists. Everything else on the import path -- joining,
// absolutising, canonicalising -- is string work that needs no filesystem.
type FileSystem interface {
	// Exists reports whether a file can be read at path.
	Exists(path string) bool
	// ReadFile returns the contents of path.
	ReadFile(path string) ([]byte, error)
}

// osFileSystem is the real thing, and the default.
type osFileSystem struct{}

func (osFileSystem) Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func (osFileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

var fileSystem FileSystem = osFileSystem{}

// SetFileSystem replaces where sources are read from, and returns the previous
// one so a caller can put it back -- which is what a test wants.
func SetFileSystem(replacement FileSystem) FileSystem {
	previous := fileSystem
	fileSystem = replacement

	return previous
}

// Exists reports whether a file exists, through the installed FileSystem.
func Exists(path string) bool {
	return fileSystem.Exists(path)
}

// ReadFile returns a file's contents, through the installed FileSystem.
func ReadFile(path string) ([]byte, error) {
	return fileSystem.ReadFile(path)
}

// MapFileSystem serves a fixed set of files from memory. The playground builds
// one from the editor's tabs, and a test can use it to exercise imports without
// putting fixtures on disk.
type MapFileSystem struct {
	Files map[string][]byte
}

func (m MapFileSystem) Exists(path string) bool {
	_, ok := m.Files[path]

	return ok
}

func (m MapFileSystem) ReadFile(path string) ([]byte, error) {
	content, ok := m.Files[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	return content, nil
}
