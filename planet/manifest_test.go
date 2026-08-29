package planet

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindRootWalksUp(t *testing.T) {
	root := t.TempDir()
	deep := filepath.Join(root, "a", "b", "c")

	if err := os.MkdirAll(deep, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ManifestName), []byte("planets: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, ok := FindRoot(deep)
	if !ok {
		t.Fatalf("FindRoot(%q) found nothing, want %q", deep, root)
	}

	// The temp dir may itself sit under a symlink, so compare resolved paths.
	wantResolved, _ := filepath.EvalSymlinks(root)
	gotResolved, _ := filepath.EvalSymlinks(got)
	if gotResolved != wantResolved {
		t.Errorf("FindRoot(%q) = %q, want %q", deep, gotResolved, wantResolved)
	}
}

func TestFindRootReportsAbsence(t *testing.T) {
	// A temp dir with no manifest anywhere above it inside the sandbox. The
	// walk stops at the filesystem root rather than looping.
	if _, ok := FindRoot(t.TempDir()); ok {
		t.Error("FindRoot found a manifest where none was written")
	}
}

func TestLoadAndSaveRoundTrip(t *testing.T) {
	root := t.TempDir()

	manifest := New(root)
	manifest.RocketLang = ">=0.24"
	manifest.Planets["utils"] = Entry{
		Source:  "github.com/flipez/rocket-lang-utils",
		Version: "v1.2.0",
		Commit:  "a3f21c9",
	}

	if err := manifest.Save(); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadFrom(root)
	if err != nil {
		t.Fatal(err)
	}

	if loaded.RocketLang != ">=0.24" {
		t.Errorf("rocket-lang = %q, want %q", loaded.RocketLang, ">=0.24")
	}

	entry, ok := loaded.Planets["utils"]
	if !ok {
		t.Fatal("utils missing after round trip")
	}
	if entry.Source != "github.com/flipez/rocket-lang-utils" || entry.Version != "v1.2.0" || entry.Commit != "a3f21c9" {
		t.Errorf("entry round-tripped wrong: %+v", entry)
	}
}

func TestLoadTreatsMissingPlanetsAsEmpty(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, ManifestName), []byte("rocket-lang: \">=0.24\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadFrom(root)
	if err != nil {
		t.Fatal(err)
	}

	// A nil map would panic on assignment in planet get.
	if loaded.Planets == nil {
		t.Fatal("Planets is nil, want an empty map")
	}
	if len(loaded.Planets) != 0 {
		t.Errorf("Planets has %d entries, want 0", len(loaded.Planets))
	}
}

func TestAliasesAreSorted(t *testing.T) {
	manifest := New(t.TempDir())
	for _, alias := range []string{"zeta", "alpha", "mid"} {
		manifest.Planets[alias] = Entry{Source: "x"}
	}

	got := manifest.Aliases()
	want := []string{"alpha", "mid", "zeta"}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Aliases() = %v, want %v", got, want)
		}
	}
}

func TestPathHelpers(t *testing.T) {
	manifest := New("/proj")

	if got := manifest.PlanetsDir(); got != filepath.Join("/proj", ".planets") {
		t.Errorf("PlanetsDir() = %q", got)
	}
	if got := manifest.Dir("utils"); got != filepath.Join("/proj", ".planets", "utils") {
		t.Errorf("Dir() = %q", got)
	}
}

func TestLoadReportsNoProject(t *testing.T) {
	if _, err := Load(t.TempDir()); err == nil {
		t.Error("Load succeeded with no manifest present")
	}
}
