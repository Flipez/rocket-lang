package planet

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Stamp records what an installed directory actually contains. It carries no
// timestamp on purpose: the content is then a pure function of what was
// installed, so it diffs cleanly if .planets is committed and can be compared
// against the manifest without false positives.
type Stamp struct {
	Source  string `yaml:"source"`
	Version string `yaml:"version"`
	Commit  string `yaml:"commit"`
}

// WriteStamp records provenance inside an installed planet.
func WriteStamp(dir string, stamp Stamp) error {
	content, err := yaml.Marshal(stamp)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, StampName), content, 0o644)
}

// ReadStamp reads the stamp from an installed planet. A missing stamp is
// reported as not-ok rather than an error: an unstamped directory is something
// to reinstall, not something to fail on.
func ReadStamp(dir string) (Stamp, bool) {
	content, err := os.ReadFile(filepath.Join(dir, StampName))
	if err != nil {
		return Stamp{}, false
	}

	var stamp Stamp
	if err := yaml.Unmarshal(content, &stamp); err != nil {
		return Stamp{}, false
	}

	return stamp, true
}

// Matches reports whether an installed directory already holds what the
// manifest asks for, which is what makes planet install idempotent.
func (s Stamp) Matches(entry Entry) bool {
	if s.Source != entry.Source || s.Version != entry.Version {
		return false
	}

	// A manifest with no recorded commit still matches on source and version;
	// this covers a hand-written entry that has not been installed yet.
	return entry.Commit == "" || s.Commit == entry.Commit
}
