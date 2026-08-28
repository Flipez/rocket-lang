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
