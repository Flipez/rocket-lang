//go:build !wasm

package planet

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// inProject runs Command with the working directory set to a fresh project, and
// returns the exit code with everything written to stdout and stderr.
func inProject(t *testing.T, dir string, args ...string) (int, string) {
	t.Helper()

	t.Chdir(dir)

	var out strings.Builder
	code := Command(args, &out, &out)

	return code, out.String()
}

func TestCommandInitCreatesTheManifest(t *testing.T) {
	dir := t.TempDir()

	code, out := inProject(t, dir, "init")
	if code != 0 {
		t.Fatalf("init exited %d: %s", code, out)
	}

	if _, err := os.Stat(filepath.Join(dir, ManifestName)); err != nil {
		t.Errorf("init did not create %s", ManifestName)
	}

	// A second init must refuse rather than overwrite a manifest with entries.
	code, out = inProject(t, dir, "init")
	if code == 0 {
		t.Errorf("a second init succeeded: %s", out)
	}
}

func TestCommandInitAmendsAnExistingGitignore(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, ".gitignore"), []byte("build/\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if code, out := inProject(t, dir, "init"); code != 0 {
		t.Fatalf("init exited %d: %s", code, out)
	}

	content, err := os.ReadFile(filepath.Join(dir, ".gitignore"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), DirName+"/") {
		t.Errorf(".gitignore = %q, want it to mention %s/", content, DirName)
	}
	if !strings.Contains(string(content), "build/") {
		t.Error("init clobbered the existing .gitignore content")
	}

	// Running init elsewhere must not duplicate the entry when re-run; check
	// idempotency of the amend by calling the helper directly.
	if err := ignorePlanetsDir(dir, &strings.Builder{}); err != nil {
		t.Fatal(err)
	}
	content, _ = os.ReadFile(filepath.Join(dir, ".gitignore"))
	if strings.Count(string(content), DirName+"/") != 1 {
		t.Errorf(".gitignore has the entry %d times, want 1", strings.Count(string(content), DirName+"/"))
	}
}

func TestCommandGetInstallsAndRecords(t *testing.T) {
	repo := newRepo(t, "v1.0.0", "v2.0.0")
	dir := t.TempDir()

	if code, out := inProject(t, dir, "init"); code != 0 {
		t.Fatalf("init exited %d: %s", code, out)
	}

	code, out := inProject(t, dir, "get", repo, "--as", "utils")
	if code != 0 {
		t.Fatalf("get exited %d: %s", code, out)
	}

	// No version was given, so the highest tag should have been chosen.
	if !strings.Contains(out, "v2.0.0") {
		t.Errorf("get did not resolve to v2.0.0: %s", out)
	}

	manifest, err := LoadFrom(dir)
	if err != nil {
		t.Fatal(err)
	}
	entry := manifest.Planets["utils"]
	if entry.Version != "v2.0.0" || entry.Commit == "" {
		t.Errorf("manifest entry = %+v, want v2.0.0 with a commit", entry)
	}

	if _, err := os.Stat(filepath.Join(dir, DirName, "utils", "mod.rl")); err != nil {
		t.Errorf("planet content missing: %s", err)
	}
}

func TestCommandGetIsIdempotentWithoutAVersion(t *testing.T) {
	repo := newRepo(t, "v1.0.0", "v2.0.0")
	dir := t.TempDir()

	inProject(t, dir, "init")
	inProject(t, dir, "get", repo, "--as", "utils")

	// Pin to the older tag, then a bare get must report rather than upgrade.
	manifest, _ := LoadFrom(dir)
	manifest.Planets["utils"] = Entry{Source: repo, Version: "v1.0.0"}
	if err := manifest.Save(); err != nil {
		t.Fatal(err)
	}
	if _, err := Install(manifest, "utils", &strings.Builder{}); err != nil {
		t.Fatal(err)
	}
	if err := manifest.Save(); err != nil {
		t.Fatal(err)
	}

	code, out := inProject(t, dir, "get", repo, "--as", "utils")
	if code != 0 {
		t.Fatalf("get exited %d: %s", code, out)
	}
	if !strings.Contains(out, "already installed at v1.0.0") {
		t.Errorf("a bare get did not report the installed version: %s", out)
	}

	reloaded, _ := LoadFrom(dir)
	if reloaded.Planets["utils"].Version != "v1.0.0" {
		t.Errorf("a bare get moved the pinned version to %q", reloaded.Planets["utils"].Version)
	}
}

func TestCommandGetRefusesAShadowingAlias(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	dir := t.TempDir()

	inProject(t, dir, "init")
	if err := os.WriteFile(filepath.Join(dir, "utils.rl"), []byte("export X = 1\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out := inProject(t, dir, "get", repo, "--as", "utils")
	if code == 0 {
		t.Fatalf("get succeeded despite a local utils.rl: %s", out)
	}
	if !strings.Contains(out, "shadow") {
		t.Errorf("the error did not explain the shadowing: %s", out)
	}
}

func TestCommandGetRefusesAnUnusableAlias(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	dir := t.TempDir()

	inProject(t, dir, "init")

	code, out := inProject(t, dir, "get", repo, "--as", "my-lib")
	if code == 0 {
		t.Fatalf("get accepted an alias that cannot be imported: %s", out)
	}
}

func TestCommandListAndRemove(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	dir := t.TempDir()

	inProject(t, dir, "init")
	inProject(t, dir, "get", repo, "--as", "utils")

	code, out := inProject(t, dir, "list")
	if code != 0 {
		t.Fatalf("list exited %d: %s", code, out)
	}
	if !strings.Contains(out, "utils") || !strings.Contains(out, "installed") {
		t.Errorf("list output = %q", out)
	}

	if code, out = inProject(t, dir, "remove", "utils"); code != 0 {
		t.Fatalf("remove exited %d: %s", code, out)
	}

	manifest, _ := LoadFrom(dir)
	if len(manifest.Planets) != 0 {
		t.Error("remove left the manifest entry behind")
	}

	code, out = inProject(t, dir, "list")
	if code != 0 || !strings.Contains(out, "no planets") {
		t.Errorf("list after remove: code %d, output %q", code, out)
	}
}

func TestCommandExitCodes(t *testing.T) {
	dir := t.TempDir()

	tests := []struct {
		args []string
		want int
	}{
		{nil, 2},                // no verb prints usage
		{[]string{"nope"}, 2},   // unknown verb
		{[]string{"help"}, 0},   // help is a success
		{[]string{"list"}, 1},   // outside a project
		{[]string{"get"}, 1},    // missing source
		{[]string{"remove"}, 1}, // missing name
	}

	for _, tt := range tests {
		if code, out := inProject(t, dir, tt.args...); code != tt.want {
			t.Errorf("planet %v exited %d, want %d: %s", tt.args, code, tt.want, out)
		}
	}
}
