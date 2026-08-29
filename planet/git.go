//go:build !wasm

package planet

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// requireGit reports a usable git binary, or explains what is missing. This is
// the only external process the interpreter depends on.
func requireGit() error {
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("planet needs the `git` command, which was not found on PATH")
	}

	return nil
}

func git(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir

	var stderr strings.Builder
	cmd.Stderr = &stderr

	out, err := cmd.Output()
	if err != nil {
		message := strings.TrimSpace(stderr.String())
		if message == "" {
			message = err.Error()
		}

		return "", fmt.Errorf("git %s: %s", strings.Join(args, " "), message)
	}

	return strings.TrimSpace(string(out)), nil
}

// Tags returns the semver-looking tags a remote publishes, highest last. Tags
// that do not parse as a version are ignored rather than rejected, since
// repositories carry all sorts of tags.
func Tags(url string) ([]string, error) {
	if err := requireGit(); err != nil {
		return nil, err
	}

	out, err := git("", "ls-remote", "--tags", url)
	if err != nil {
		return nil, err
	}

	var tags []string
	seen := map[string]bool{}

	for _, line := range strings.Split(out, "\n") {
		_, ref, found := strings.Cut(line, "\t")
		if !found {
			continue
		}

		// Annotated tags also appear as refs/tags/x^{}; both name one version.
		name := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(ref), "refs/tags/"), "^{}")
		if name == "" || seen[name] {
			continue
		}

		if _, ok := parseVersion(name); !ok {
			continue
		}

		seen[name] = true
		tags = append(tags, name)
	}

	sort.Slice(tags, func(i, j int) bool {
		a, _ := parseVersion(tags[i])
		b, _ := parseVersion(tags[j])

		return a.less(b)
	})

	return tags, nil
}

// LatestTag returns the highest semver tag a remote publishes, or an empty
// string when it publishes none. That is not an error: a planet that has not
// tagged a release yet is still usable through its default branch.
func LatestTag(url string) (string, error) {
	tags, err := Tags(url)
	if err != nil {
		return "", err
	}

	if len(tags) == 0 {
		return "", nil
	}

	return tags[len(tags)-1], nil
}

// DefaultBranch asks the remote which branch its HEAD points at. The answer is
// read from the remote rather than assumed to be "main", since plenty of
// repositories use master, trunk or something else entirely.
func DefaultBranch(url string) (string, error) {
	if err := requireGit(); err != nil {
		return "", err
	}

	out, err := git("", "ls-remote", "--symref", url, "HEAD")
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(out, "\n") {
		if !strings.HasPrefix(line, "ref:") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		if branch := strings.TrimPrefix(fields[1], "refs/heads/"); branch != "" {
			return branch, nil
		}
	}

	return "", fmt.Errorf("%s does not report a default branch", url)
}

// ResolveVersion picks what to install when no version was asked for: the
// highest semver tag if the planet publishes any, otherwise its default branch.
// isTag says which, so a caller can point out that it settled for a branch.
func ResolveVersion(url string) (version string, isTag bool, err error) {
	tag, err := LatestTag(url)
	if err != nil {
		return "", false, err
	}
	if tag != "" {
		return tag, true, nil
	}

	branch, err := DefaultBranch(url)
	if err != nil {
		return "", false, err
	}

	return branch, false, nil
}

// Checkout clones url into dest and checks out ref, which may be a tag, a
// branch or a commit. It returns the resolved commit.
//
// The clone is deliberately not shallow. A full clone is what lets an arbitrary
// recorded commit be checked out, which is how an install is pinned against a
// moved tag or a branch that has advanced. Fetching a single commit shallowly
// is possible, but only when the server sets uploadpack.allowReachableSHA1InWant
// -- GitHub does, a self-hosted Gitea may not -- so it would need a full-clone
// fallback anyway.
//
// The cost is small for the libraries planets tend to be: cloning a two-commit
// repository takes the same time either way, because network latency dominates.
// Shallow only pays on a repository with real history, which is worth revisiting
// if that becomes common.
func Checkout(url, ref, dest string) (string, error) {
	if err := requireGit(); err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return "", err
	}

	if _, err := git("", "clone", "--quiet", url, dest); err != nil {
		return "", err
	}

	if _, err := git(dest, "checkout", "--quiet", ref); err != nil {
		return "", err
	}

	commit, err := git(dest, "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}

	// Provenance lives in the stamp file, so the history is not kept.
	if err := os.RemoveAll(filepath.Join(dest, ".git")); err != nil {
		return "", err
	}

	return commit, nil
}

type version struct{ major, minor, patch int }

func (a version) less(b version) bool {
	if a.major != b.major {
		return a.major < b.major
	}
	if a.minor != b.minor {
		return a.minor < b.minor
	}

	return a.patch < b.patch
}

// parseVersion reads a v-prefixed or bare major.minor.patch tag. Anything with
// a prerelease or build suffix is rejected, so `planet get` without a version
// never selects a release candidate by accident.
func parseVersion(tag string) (version, bool) {
	parts := strings.Split(strings.TrimPrefix(tag, "v"), ".")
	if len(parts) != 3 {
		return version{}, false
	}

	var parsed version
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil || n < 0 {
			return version{}, false
		}

		switch i {
		case 0:
			parsed.major = n
		case 1:
			parsed.minor = n
		case 2:
			parsed.patch = n
		}
	}

	return parsed, true
}
