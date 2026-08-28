package object

import "strings"

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
	if len(r.chain) > 0 {
		r.chain = r.chain[:len(r.chain)-1]
	}
}

// Chain renders the in-progress import stack with path appended, for the
// circular-import error message.
func (r *ModuleRegistry) Chain(path string) string {
	return strings.Join(append(append([]string{}, r.chain...), path), " -> ")
}
