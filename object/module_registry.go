package object

import (
	"os"
	"path/filepath"
	"strings"
)

// ModuleRegistry caches evaluated modules by resolved absolute path and
// tracks which are mid-evaluation so cycles can be reported instead of
// recursing forever.
type ModuleRegistry struct {
	cache      map[string]*Module
	inProgress map[string]struct{}
	chain      []string
}

func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		cache:      make(map[string]*Module),
		inProgress: make(map[string]struct{}),
	}
}

func (r *ModuleRegistry) Get(path string) (*Module, bool) {
	m, ok := r.cache[path]
	return m, ok
}

func (r *ModuleRegistry) Put(path string, m *Module) {
	r.cache[path] = m
}

func (r *ModuleRegistry) InProgress(path string) bool {
	_, ok := r.inProgress[path]
	return ok
}

func (r *ModuleRegistry) Begin(path string) {
	r.inProgress[path] = struct{}{}
	r.chain = append(r.chain, path)
}

func (r *ModuleRegistry) End(path string) {
	delete(r.inProgress, path)
	if n := len(r.chain); n > 0 && r.chain[n-1] == path {
		r.chain = r.chain[:n-1]
	}
}

// Chain renders the in-progress import stack with path appended, for the
// circular-import error message. Each entry is rendered relative to the
// current working directory to avoid leaking absolute filesystem paths;
// the cache and in-progress maps continue to use absolute paths as keys.
func (r *ModuleRegistry) Chain(path string) string {
	full := append(append([]string{}, r.chain...), path)

	rendered := make([]string, len(full))
	for i, p := range full {
		rendered[i] = relativizePath(p)
	}

	return strings.Join(rendered, " -> ")
}

// relativizePath renders p relative to the current working directory,
// falling back to the unmodified path if that cannot be determined.
func relativizePath(p string) string {
	wd, err := os.Getwd()
	if err != nil {
		return p
	}

	rel, err := filepath.Rel(wd, p)
	if err != nil {
		return p
	}

	return rel
}
