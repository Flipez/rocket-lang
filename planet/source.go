package planet

import (
	"fmt"
	"path/filepath"
	"strings"
)

// defaultHost is prepended to a bare owner/name source.
const defaultHost = "github.com"

// Source is a resolved planet location.
type Source struct {
	// Raw is what the user typed, and what the manifest records.
	Raw string
	// URL is what git is given.
	URL string
}

// ParseSource turns what the user typed into something git can clone.
//
//	flipez/rocket-lang-utils   -> https://github.com/flipez/rocket-lang-utils
//	codeberg.org/flipez/utils  -> https://codeberg.org/flipez/utils
//	https://example.com/u.git  -> as written
//	git@github.com:u/n.git     -> as written
//	/srv/planets/utils         -> as written, a local clone
//	../sibling-project         -> resolved against the working directory
//
// Deriving an alias is deliberately separate: a source can be perfectly
// cloneable while its last path segment makes no sense as an identifier.
func ParseSource(raw string) (Source, error) {
	raw = strings.TrimSpace(raw)
	trimmed := strings.TrimSuffix(raw, "/")

	if trimmed == "" {
		return Source{}, fmt.Errorf("empty planet source")
	}

	switch {
	case strings.Contains(trimmed, "://"), strings.HasPrefix(trimmed, "git@"):
		return Source{Raw: raw, URL: trimmed}, nil

	case isLocalPath(trimmed):
		absolute, err := filepath.Abs(trimmed)
		if err != nil {
			return Source{}, err
		}

		return Source{Raw: raw, URL: absolute}, nil

	case looksHostQualified(trimmed):
		return Source{Raw: raw, URL: "https://" + trimmed}, nil
	}

	parts := strings.Split(trimmed, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return Source{}, fmt.Errorf("cannot read %q as a planet source: expected owner/name, host/owner/name, a URL or a local path", raw)
	}

	return Source{Raw: raw, URL: "https://" + defaultHost + "/" + trimmed}, nil
}

// isLocalPath reports whether raw names a directory on this machine rather than
// a remote. An explicitly relative or absolute path is the signal; a bare
// owner/name is never treated as local.
func isLocalPath(raw string) bool {
	return filepath.IsAbs(raw) ||
		strings.HasPrefix(raw, "./") ||
		strings.HasPrefix(raw, "../") ||
		raw == "." ||
		raw == ".."
}

// looksHostQualified reports whether the first segment names a host. A dot in
// the first segment is the signal, which distinguishes codeberg.org/owner/name
// from owner/name without maintaining a list of known hosts.
func looksHostQualified(raw string) bool {
	parts := strings.SplitN(raw, "/", 2)

	return len(parts) == 2 && strings.Contains(parts[0], ".")
}

// DefaultAlias derives a binding name from a source: the last path segment with
// a .git suffix and any redundant rocket-lang prefix removed. It fails rather
// than returning something unusable, because an alias that is not a legal
// identifier produces an import that can never be referenced -- "my-lib.Foo"
// parses as subtraction.
func DefaultAlias(raw string) (string, error) {
	cleaned := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSpace(raw), "/"), ".git")
	cleaned = strings.ReplaceAll(cleaned, "\\", "/")

	// git@host:owner/name puts the interesting part after the colon.
	if _, after, found := strings.Cut(cleaned, ":"); found && !strings.Contains(after, "//") {
		cleaned = after
	}

	segments := strings.Split(cleaned, "/")
	alias := segments[len(segments)-1]

	for _, prefix := range []string{"rocket-lang-", "rocketlang-", "rl-"} {
		if len(alias) > len(prefix) && strings.HasPrefix(alias, prefix) {
			alias = strings.TrimPrefix(alias, prefix)
			break
		}
	}

	if !IsValidAlias(alias) {
		return "", fmt.Errorf("cannot derive a name from %q: %q would not be usable in an import, pass one with --as", raw, alias)
	}

	return alias, nil
}

// IsValidAlias reports whether alias can be written in an import and then
// referenced. A hyphen is the common failure.
func IsValidAlias(alias string) bool {
	if alias == "" {
		return false
	}

	for i, r := range alias {
		isLetter := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
		isDigit := r >= '0' && r <= '9'

		if !isLetter && !(isDigit && i > 0) {
			return false
		}
	}

	return true
}
