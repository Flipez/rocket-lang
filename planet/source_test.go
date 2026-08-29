package planet

import "testing"

func TestParseSource(t *testing.T) {
	tests := []struct {
		raw   string
		url   string
		alias string
	}{
		// A bare owner/name defaults to GitHub.
		{"flipez/rocket-lang-utils", "https://github.com/flipez/rocket-lang-utils", "utils"},
		{"flipez/helpers", "https://github.com/flipez/helpers", "helpers"},
		// A dot in the first segment marks it as a host.
		{"codeberg.org/flipez/utils", "https://codeberg.org/flipez/utils", "utils"},
		{"git.example.com/team/tools", "https://git.example.com/team/tools", "tools"},
		// Explicit URLs pass through untouched.
		{"https://example.com/u/thing.git", "https://example.com/u/thing.git", "thing"},
		{"git@github.com:flipez/helpers.git", "git@github.com:flipez/helpers.git", "helpers"},
		// A trailing slash is tolerated.
		{"flipez/helpers/", "https://github.com/flipez/helpers", "helpers"},
		// Redundant rocket-lang prefixes are dropped from the default alias.
		{"flipez/rocketlang-json", "https://github.com/flipez/rocketlang-json", "json"},
		{"flipez/rl-http", "https://github.com/flipez/rl-http", "http"},
	}

	for _, tt := range tests {
		got, err := ParseSource(tt.raw)
		if err != nil {
			t.Errorf("ParseSource(%q) errored: %s", tt.raw, err)
			continue
		}
		if got.URL != tt.url {
			t.Errorf("ParseSource(%q).URL = %q, want %q", tt.raw, got.URL, tt.url)
		}
		if got.Alias != tt.alias {
			t.Errorf("ParseSource(%q).Alias = %q, want %q", tt.raw, got.Alias, tt.alias)
		}
	}
}

func TestParseSourceRejects(t *testing.T) {
	// The last case is the one that matters: an alias with a hyphen left in it
	// would produce an import binding that can never be referenced, because
	// "my-lib.Foo" parses as subtraction.
	for _, raw := range []string{"", "   ", "onlyone", "a/b/c/d/e/f", "flipez/my-lib"} {
		if got, err := ParseSource(raw); err == nil {
			t.Errorf("ParseSource(%q) succeeded with alias %q, want an error", raw, got.Alias)
		}
	}
}

func TestIsValidAlias(t *testing.T) {
	valid := []string{"utils", "http2", "_private", "aB9", "a"}
	invalid := []string{"", "my-lib", "1st", "has space", "dot.ted", "sla/sh"}

	for _, alias := range valid {
		if !IsValidAlias(alias) {
			t.Errorf("IsValidAlias(%q) = false, want true", alias)
		}
	}
	for _, alias := range invalid {
		if IsValidAlias(alias) {
			t.Errorf("IsValidAlias(%q) = true, want false", alias)
		}
	}
}
