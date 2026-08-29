package utilities

import (
	"os"
	"path/filepath"
	"testing"
)

// TestFindModuleFrom_RelativeBranches exercises the five meaningful branches
// of FindModuleFrom: "./"-prefixed existing, "../"-prefixed existing,
// "./"-prefixed missing, the importerDir=="" guard, and delegation to
// FindModule for plain names.
func TestFindModuleFrom_RelativeBranches(t *testing.T) {
	t.Run("dot-slash path that exists relative to importerDir", func(t *testing.T) {
		importerDir := t.TempDir()
		modulePath := filepath.Join(importerDir, "sibling.rl")

		if err := os.WriteFile(modulePath, []byte(""), 0o644); err != nil {
			t.Fatalf("failed to write module file: %v", err)
		}

		got := FindModuleFrom("./sibling", importerDir)

		want := canonicalize(modulePath)
		if got != want {
			t.Fatalf("FindModuleFrom(./sibling) = %q, want %q", got, want)
		}
		if got == "" {
			t.Fatalf("FindModuleFrom(./sibling) returned empty string, expected a resolved path")
		}
	})

	t.Run("dot-dot-slash path that exists relative to importerDir", func(t *testing.T) {
		root := t.TempDir()
		importerDir := filepath.Join(root, "sub")
		if err := os.Mkdir(importerDir, 0o755); err != nil {
			t.Fatalf("failed to create importer dir: %v", err)
		}

		modulePath := filepath.Join(root, "parent.rl")
		if err := os.WriteFile(modulePath, []byte(""), 0o644); err != nil {
			t.Fatalf("failed to write module file: %v", err)
		}

		got := FindModuleFrom("../parent", importerDir)

		want := canonicalize(modulePath)
		if got != want {
			t.Fatalf("FindModuleFrom(../parent) = %q, want %q", got, want)
		}
	})

	t.Run("dot-slash path that does not exist", func(t *testing.T) {
		importerDir := t.TempDir()

		got := FindModuleFrom("./nonexistent", importerDir)

		if got != "" {
			t.Fatalf("FindModuleFrom(./nonexistent) = %q, want empty string", got)
		}
	})

	t.Run("relative path with empty importerDir short-circuits", func(t *testing.T) {
		// Create a file in the current working directory of the test
		// process under the exact relative name being looked up, to prove
		// that the importerDir == "" guard short-circuits before any
		// filesystem work rather than coincidentally missing.
		cwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("failed to get cwd: %v", err)
		}

		findableName := "would_be_findable_from_cwd"
		findablePath := filepath.Join(cwd, findableName+".rl")

		if err := os.WriteFile(findablePath, []byte(""), 0o644); err != nil {
			t.Fatalf("failed to write findable file: %v", err)
		}
		t.Cleanup(func() {
			os.Remove(findablePath)
		})

		got := FindModuleFrom("./"+findableName, "")

		if got != "" {
			t.Fatalf("FindModuleFrom(%q, \"\") = %q, want empty string (guard should short-circuit)", "./"+findableName, got)
		}
	})

	t.Run("plain non-relative name delegates to FindModule", func(t *testing.T) {
		searchDir := t.TempDir()
		if err := AddPath(searchDir); err != nil {
			t.Fatalf("AddPath failed: %v", err)
		}

		modulePath := filepath.Join(searchDir, "delegate_target.rl")
		if err := os.WriteFile(modulePath, []byte(""), 0o644); err != nil {
			t.Fatalf("failed to write module file: %v", err)
		}

		got := FindModuleFrom("delegate_target", "/some/importer/dir")

		want := FindModule("delegate_target")
		if got == "" {
			t.Fatalf("FindModuleFrom(delegate_target) returned empty string, expected successful delegation")
		}
		if got != want {
			t.Fatalf("FindModuleFrom(delegate_target) = %q, want %q (result of FindModule)", got, want)
		}
	})
}

// TestFindModule_CanonicalizesSymlinks proves Finding 1's fix: reaching the
// same file through a symlinked route and a real route must produce
// byte-identical strings, so the module cache cannot end up with two
// entries for a single file.
func TestFindModule_CanonicalizesSymlinks(t *testing.T) {
	root := t.TempDir()

	realDir := filepath.Join(root, "real")
	if err := os.Mkdir(realDir, 0o755); err != nil {
		t.Fatalf("failed to create real dir: %v", err)
	}

	modulePath := filepath.Join(realDir, "shared.rl")
	if err := os.WriteFile(modulePath, []byte(""), 0o644); err != nil {
		t.Fatalf("failed to write module file: %v", err)
	}

	linkDir := filepath.Join(root, "link")
	if err := os.Symlink(realDir, linkDir); err != nil {
		t.Skipf("symlink creation unavailable, skipping: %v", err)
	}

	// Resolve the same target file through both spellings:
	// 1. via a "../" import from a sibling directory that reaches the
	//    symlinked directory.
	sibling := filepath.Join(root, "sibling")
	if err := os.Mkdir(sibling, 0o755); err != nil {
		t.Fatalf("failed to create sibling dir: %v", err)
	}
	viaSymlink := FindModuleFrom("../link/shared", sibling)

	// 2. via a "../" import from a sibling directory that reaches the real
	//    directory directly.
	viaReal := FindModuleFrom("../real/shared", sibling)

	if viaSymlink == "" || viaReal == "" {
		t.Fatalf("expected both resolutions to succeed, got viaSymlink=%q viaReal=%q", viaSymlink, viaReal)
	}

	if viaSymlink != viaReal {
		t.Fatalf("resolutions via symlinked and real routes differ: viaSymlink=%q viaReal=%q, want identical strings", viaSymlink, viaReal)
	}
}

// TestSearchPathList covers the three defects the old initSearchPaths had:
// it substituted the working directory instead of adding to it, it split on a
// hardcoded ":", and it treated a bad entry as fatal.
func TestSearchPathList(t *testing.T) {
	sep := string(os.PathListSeparator)
	cwd := "/work"

	tests := []struct {
		name     string
		env      string
		cwd      string
		expected []string
	}{
		{
			name:     "unset falls back to the working directory",
			env:      "",
			cwd:      cwd,
			expected: []string{cwd},
		},
		{
			// The old behaviour dropped cwd entirely here, so setting the
			// variable silently broke every working-directory-relative import.
			name:     "entries come first and the working directory is kept",
			env:      "/opt/rl",
			cwd:      cwd,
			expected: []string{"/opt/rl", cwd},
		},
		{
			name:     "multiple entries keep their order",
			env:      "/a" + sep + "/b",
			cwd:      cwd,
			expected: []string{"/a", "/b", cwd},
		},
		{
			// Split on a hardcoded ":" turned an empty leading entry into "",
			// which AddPath then resolved to the working directory.
			name:     "empty entries are dropped",
			env:      sep + "/a" + sep + sep + "/b" + sep,
			cwd:      cwd,
			expected: []string{"/a", "/b", cwd},
		},
		{
			name:     "no working directory yields only the configured entries",
			env:      "/a",
			cwd:      "",
			expected: []string{"/a"},
		},
		{
			name:     "nothing configured and no working directory yields nothing",
			env:      "",
			cwd:      "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		got := searchPathList(tt.env, tt.cwd)

		if len(got) != len(tt.expected) {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, got)
			continue
		}

		for i := range got {
			if got[i] != tt.expected[i] {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, got)
				break
			}
		}
	}
}

// TestInitSearchPathsThenAddPathAppends covers the ordering the planet
// directory depends on. FindModule builds the search path lazily, so a caller
// that appends without forcing initialisation first would land at the front of
// the list instead of the back — and an installed planet would then shadow the
// project's own modules.
func TestInitSearchPathsThenAddPathAppends(t *testing.T) {
	before := len(SearchPaths)

	InitSearchPaths()

	built := len(SearchPaths)
	if built == 0 {
		t.Fatal("InitSearchPaths produced no search paths")
	}

	dir := t.TempDir()
	if err := AddPath(dir); err != nil {
		t.Fatal(err)
	}

	if len(SearchPaths) != built+1 {
		t.Fatalf("AddPath added %d entries, want 1", len(SearchPaths)-built)
	}

	resolved, _ := filepath.EvalSymlinks(dir)
	last, _ := filepath.EvalSymlinks(SearchPaths[len(SearchPaths)-1])
	if last != resolved {
		t.Errorf("AddPath put %q at position %d, want it last", resolved, before)
	}

	// InitSearchPaths is idempotent, so a second call must not rebuild and
	// push the appended entry out of last place.
	InitSearchPaths()
	last, _ = filepath.EvalSymlinks(SearchPaths[len(SearchPaths)-1])
	if last != resolved {
		t.Errorf("a second InitSearchPaths call disturbed the appended entry")
	}
}

func TestMapFileSystem(t *testing.T) {
	previous := SetFileSystem(MapFileSystem{Files: map[string][]byte{
		"/play/main.rl": []byte(`import "util"`),
	}})
	defer SetFileSystem(previous)

	if !Exists("/play/main.rl") {
		t.Error("Exists should find a file the map has")
	}
	if Exists("/play/missing.rl") {
		t.Error("Exists should not find a file the map lacks")
	}

	content, err := ReadFile("/play/main.rl")
	if err != nil {
		t.Fatalf("ReadFile: %s", err)
	}
	if string(content) != `import "util"` {
		t.Errorf("ReadFile returned %q", content)
	}

	if _, err := ReadFile("/play/missing.rl"); err == nil {
		t.Error("ReadFile should fail for a file the map lacks")
	}
}

// TestSetFileSystemRestores checks that the previous filesystem comes back, so
// a test installing one cannot leak it into the next.
func TestSetFileSystemRestores(t *testing.T) {
	previous := SetFileSystem(MapFileSystem{Files: map[string][]byte{}})
	restored := SetFileSystem(previous)

	if _, ok := restored.(MapFileSystem); !ok {
		t.Errorf("SetFileSystem should return what it replaced, got %T", restored)
	}
	if !Exists("utilities.go") {
		t.Error("the real filesystem should be back, and utilities.go should exist")
	}
}
