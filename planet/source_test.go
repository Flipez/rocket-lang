package planet

import (
	"path/filepath"
	"testing"
)

func TestParseSourceURL(t *testing.T) {
	tests := map[string]string{
		// A bare owner/name defaults to GitHub.
		"flipez/rocket-lang-utils": "https://github.com/flipez/rocket-lang-utils",
		"flipez/helpers":           "https://github.com/flipez/helpers",
		// A dot in the first segment marks it as a host.
		"codeberg.org/flipez/utils": "https://codeberg.org/flipez/utils",
		"git.example.com/team/a/b":  "https://git.example.com/team/a/b",
		// Explicit URLs and scp-style remotes pass through untouched.
		"https://example.com/u/thing.git": "https://example.com/u/thing.git",
		"git@github.com:flipez/helpers.git": "git@github.com:flipez/helpers.git",
		// A trailing slash is tolerated.
		"flipez/helpers/": "https://github.com/flipez/helpers",
		// Absolute local paths clone directly, which is what makes a planet in a
		// monorepo or on a shared drive usable.
		"/srv/planets/utils": "/srv/planets/utils",
	}

	for raw, want := range tests {
		got, err := ParseSource(raw)
		if err != nil {
			t.Errorf("ParseSource(%q) errored: %s", raw, err)
			continue
		}
		if got.URL != want {
			t.Errorf("ParseSource(%q).URL = %q, want %q", raw, got.URL, want)
		}
		if got.Raw != raw {
			t.Errorf("ParseSource(%q).Raw = %q, want it preserved verbatim", raw, got.Raw)
		}
	}
}

func TestParseSourceResolvesRelativePaths(t *testing.T) {
	got, err := ParseSource("./sibling")
	if err != nil {
		t.Fatal(err)
	}

	if !filepath.IsAbs(got.URL) {
		t.Errorf("ParseSource(\"./sibling\").URL = %q, want an absolute path", got.URL)
	}
}

func TestParseSourceRejects(t *testing.T) {
	for _, raw := range []string{"", "   ", "onlyone", "a/b/c"} {
		if got, err := ParseSource(raw); err == nil {
			t.Errorf("ParseSource(%q) succeeded with %q, want an error", raw, got.URL)
		}
	}
}

func TestDefaultAlias(t *testing.T) {
	tests := map[string]string{
		"flipez/rocket-lang-utils":          "utils",
		"flipez/rocketlang-json":            "json",
		"flipez/rl-http":                    "http",
		"flipez/helpers":                    "helpers",
		"codeberg.org/flipez/utils":         "utils",
		"https://example.com/u/thing.git":   "thing",
		"git@github.com:flipez/helpers.git": "helpers",
		"/srv/planets/utils":                "utils",
	}

	for raw, want := range tests {
		got, err := DefaultAlias(raw)
		if err != nil {
			t.Errorf("DefaultAlias(%q) errored: %s", raw, err)
			continue
		}
		if got != want {
			t.Errorf("DefaultAlias(%q) = %q, want %q", raw, got, want)
		}
	}
}

func TestDefaultAliasRefusesUnusableNames(t *testing.T) {
	// The first is the important one: my-lib cannot be referenced after import,
	// because my-lib.Foo parses as subtraction. Failing here, with a pointer to
	// --as, is better than installing something unusable.
	for _, raw := range []string{"flipez/my-lib", "flipez/2fast", "/srv/planets/001"} {
		if got, err := DefaultAlias(raw); err == nil {
			t.Errorf("DefaultAlias(%q) = %q, want an error", raw, got)
		}
	}
}

func TestIsValidAlias(t *testing.T) {
	for _, alias := range []string{"utils", "http2", "_private", "aB9", "a"} {
		if !IsValidAlias(alias) {
			t.Errorf("IsValidAlias(%q) = false, want true", alias)
		}
	}
	for _, alias := range []string{"", "my-lib", "1st", "has space", "dot.ted", "sla/sh"} {
		if IsValidAlias(alias) {
			t.Errorf("IsValidAlias(%q) = true, want false", alias)
		}
	}
}
