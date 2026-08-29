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

// LatestTag returns the highest semver tag a remote publishes.
func LatestTag(url string) (string, error) {
	tags, err := Tags(url)
	if err != nil {
		return "", err
	}

	if len(tags) == 0 {
		return "", fmt.Errorf("%s publishes no version tags; pass an explicit @version", url)
	}

	return tags[len(tags)-1], nil
}

// Checkout clones url into dest and checks out ref, which may be a tag, a
// branch or a commit. It returns the resolved commit.
//
// The clone is not shallow, and deliberately so on two counts. Shallow
// fetching is unsupported over git's dumb HTTP protocol, which is what lets a
// plain static file server act as a planet mirror; and a full clone is what
// makes checking out a recorded commit possible, which is how an install is
// pinned against a moved tag or a branch that has advanced.
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
