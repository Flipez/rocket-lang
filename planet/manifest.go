package planet

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

// ManifestName is the file whose presence marks a project root.
const ManifestName = "planets.yml"

// DirName holds installed planets, relative to the project root.
const DirName = ".planets"

// StampName records what was installed in a planet's directory.
const StampName = ".planet"

// Entry is one planet as the manifest records it. Version is the tag that was
// asked for; Commit is what that tag resolved to, which is what makes an
// install reproducible when a tag is later moved.
type Entry struct {
	Source  string `yaml:"source"`
	Version string `yaml:"version"`
	Commit  string `yaml:"commit"`
}

// Manifest is planets.yml. Keys of Planets are aliases chosen by the consumer,
// and double as directory names under .planets.
type Manifest struct {
	RocketLang string           `yaml:"rocket-lang,omitempty"`
	Planets    map[string]Entry `yaml:"planets"`

	// root is the directory the manifest was read from, not serialised.
	root string
}

// Root returns the project root: the directory holding the manifest.
func (m *Manifest) Root() string { return m.root }

// PlanetsDir returns the directory installed planets live in.
func (m *Manifest) PlanetsDir() string { return filepath.Join(m.root, DirName) }

// Dir returns where a given alias is installed.
func (m *Manifest) Dir(alias string) string { return filepath.Join(m.PlanetsDir(), alias) }

// Aliases returns the aliases in sorted order, so output and iteration are
// deterministic rather than following Go's map ordering.
func (m *Manifest) Aliases() []string {
	aliases := make([]string, 0, len(m.Planets))
	for alias := range m.Planets {
		aliases = append(aliases, alias)
	}
	sort.Strings(aliases)

	return aliases
}

// FindRoot walks up from dir looking for a manifest, and returns the directory
// containing it. It reports whether one was found rather than erroring, since
// "not in a project" is an ordinary state.
func FindRoot(dir string) (string, bool) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return "", false
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ManifestName)); err == nil {
			return dir, true
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}

// Load reads the manifest for the project containing dir.
func Load(dir string) (*Manifest, error) {
	root, ok := FindRoot(dir)
	if !ok {
		return nil, fmt.Errorf("no %s found in %s or any parent directory", ManifestName, dir)
	}

	return LoadFrom(root)
}

// LoadFrom reads the manifest in exactly root, without searching upwards.
func LoadFrom(root string) (*Manifest, error) {
	path := filepath.Join(root, ManifestName)

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	manifest := &Manifest{root: root}
	if err := yaml.Unmarshal(content, manifest); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	if manifest.Planets == nil {
		manifest.Planets = map[string]Entry{}
	}

	return manifest, nil
}

// Save writes the manifest back. yaml.v3 does not preserve comments or key
// order, so the file is tool-managed by design.
func (m *Manifest) Save() error {
	content, err := yaml.Marshal(m)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(m.root, ManifestName), content, 0o644)
}

// New returns an empty manifest rooted at root, without writing it.
func New(root string) *Manifest {
	return &Manifest{root: root, Planets: map[string]Entry{}}
}
