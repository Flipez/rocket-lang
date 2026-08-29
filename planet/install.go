//go:build !wasm

package planet

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Install materialises one alias from the manifest, replacing whatever is
// there. It returns the resolved commit.
func Install(m *Manifest, alias string, out io.Writer) (string, error) {
	entry, ok := m.Planets[alias]
	if !ok {
		return "", fmt.Errorf("no planet named %q in %s", alias, ManifestName)
	}

	source, err := ParseSource(entry.Source)
	if err != nil {
		return "", err
	}

	version := entry.Version
	if version == "" {
		if version, err = LatestTag(source.URL); err != nil {
			return "", err
		}
	}

	dir := m.Dir(alias)

	// Clone beside the target and swap, so a failed fetch cannot leave a
	// half-written planet where a working one used to be.
	staging := dir + ".incoming"
	if err := os.RemoveAll(staging); err != nil {
		return "", err
	}

	commit, err := Checkout(source.URL, version, staging)
	if err != nil {
		os.RemoveAll(staging)
		return "", err
	}

	if err := WriteStamp(staging, Stamp{Source: entry.Source, Version: version, Commit: commit}); err != nil {
		os.RemoveAll(staging)
		return "", err
	}

	if err := os.RemoveAll(dir); err != nil {
		os.RemoveAll(staging)
		return "", err
	}
	if err := os.Rename(staging, dir); err != nil {
		return "", err
	}

	fmt.Fprintf(out, "installed %s %s (%s) to %s\n", alias, version, shortCommit(commit), relativeTo(m.Root(), dir))

	return commit, nil
}

// InstallAll materialises every planet the manifest lists, skipping any whose
// stamp already matches. That is what makes repeated runs cheap and what lets a
// vendored planet directory be left alone.
func InstallAll(m *Manifest, out io.Writer) error {
	aliases := m.Aliases()
	if len(aliases) == 0 {
		fmt.Fprintf(out, "no planets in %s\n", ManifestName)
		return nil
	}

	for _, alias := range aliases {
		entry := m.Planets[alias]

		if stamp, ok := ReadStamp(m.Dir(alias)); ok && stamp.Matches(entry) {
			fmt.Fprintf(out, "up to date  %s %s\n", alias, stamp.Version)
			continue
		}

		commit, err := Install(m, alias, out)
		if err != nil {
			return fmt.Errorf("%s: %w", alias, err)
		}

		// Record the commit a bare manifest entry resolved to, so the next
		// install is reproducible.
		if entry.Commit != commit {
			entry.Commit = commit
			if entry.Version == "" {
				if stamp, ok := ReadStamp(m.Dir(alias)); ok {
					entry.Version = stamp.Version
				}
			}
			m.Planets[alias] = entry
			if err := m.Save(); err != nil {
				return err
			}
		}
	}

	return nil
}

// Remove deletes an alias from disk and from the manifest.
func Remove(m *Manifest, alias string, out io.Writer) error {
	if _, ok := m.Planets[alias]; !ok {
		return fmt.Errorf("no planet named %q in %s", alias, ManifestName)
	}

	if err := os.RemoveAll(m.Dir(alias)); err != nil {
		return err
	}

	delete(m.Planets, alias)
	if err := m.Save(); err != nil {
		return err
	}

	fmt.Fprintf(out, "removed %s\n", alias)

	return nil
}

// LocalModuleConflict reports a module in the project root that would shadow an
// alias. The search path puts the working directory first, so such a planet
// would be installed and then silently never used -- better to refuse at
// install time, where it is actionable.
func LocalModuleConflict(m *Manifest, alias string) (string, bool) {
	candidates := []string{
		filepath.Join(m.Root(), alias+".rl"),
		filepath.Join(m.Root(), alias),
	}

	for _, candidate := range candidates {
		info, err := os.Stat(candidate)
		if err != nil {
			continue
		}
		if strings.HasSuffix(candidate, ".rl") || info.IsDir() {
			return relativeTo(m.Root(), candidate), true
		}
	}

	return "", false
}

func shortCommit(commit string) string {
	if len(commit) > 7 {
		return commit[:7]
	}

	return commit
}

func relativeTo(root, path string) string {
	if rel, err := filepath.Rel(root, path); err == nil {
		return rel
	}

	return path
}
