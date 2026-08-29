package planet

import (
	"fmt"
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
	// Alias is the default binding name derived from the source.
	Alias string
}

// ParseSource turns what the user typed into a git URL and a default alias.
//
//	flipez/rocket-lang-utils   -> https://github.com/flipez/rocket-lang-utils
//	codeberg.org/flipez/utils  -> https://codeberg.org/flipez/utils
//	https://example.com/u.git   -> as written
func ParseSource(raw string) (Source, error) {
	raw = strings.TrimSuffix(strings.TrimSpace(raw), "/")

	if raw == "" {
		return Source{}, fmt.Errorf("empty planet source")
	}

	var url string

	switch {
	case strings.Contains(raw, "://"), strings.HasPrefix(raw, "git@"):
		url = raw
	case looksHostQualified(raw):
		url = "https://" + raw
	default:
		parts := strings.Split(raw, "/")
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return Source{}, fmt.Errorf("cannot read %q as a planet source: expected owner/name, host/owner/name or a URL", raw)
		}
		url = "https://" + defaultHost + "/" + raw
	}

	alias, err := aliasFor(raw)
	if err != nil {
		return Source{}, err
	}

	return Source{Raw: raw, URL: url, Alias: alias}, nil
}

// looksHostQualified reports whether the first segment names a host. A dot in
// the first segment is the signal, which is what distinguishes
// codeberg.org/owner/name from owner/name.
func looksHostQualified(raw string) bool {
	parts := strings.SplitN(raw, "/", 2)

	return len(parts) == 2 && strings.Contains(parts[0], ".")
}

// aliasFor derives a default alias from a source: the last path segment, with a
// .git suffix and any rocket-lang prefix removed, since "rocket-lang-utils" as
// an alias is redundant inside a RocketLang project.
func aliasFor(raw string) (string, error) {
	segments := strings.Split(strings.TrimSuffix(raw, ".git"), "/")
	alias := segments[len(segments)-1]

	for _, prefix := range []string{"rocket-lang-", "rocketlang-", "rl-"} {
		if len(alias) > len(prefix) && strings.HasPrefix(alias, prefix) {
			alias = strings.TrimPrefix(alias, prefix)
			break
		}
	}

	if !IsValidAlias(alias) {
		return "", fmt.Errorf("cannot derive an alias from %q: %q is not usable as an identifier, pass one with --as", raw, alias)
	}

	return alias, nil
}

// IsValidAlias reports whether alias can be written in an import and used as a
// binding. A hyphen is the common failure: "str-help.Foo" parses as
// subtraction, so the binding would be unreachable.
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
