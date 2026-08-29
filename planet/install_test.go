//go:build !wasm

package planet

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func projectWith(t *testing.T, repo, alias, version string) *Manifest {
	t.Helper()

	m := New(t.TempDir())
	m.Planets[alias] = Entry{Source: repo, Version: version}
	if err := m.Save(); err != nil {
		t.Fatal(err)
	}

	return m
}

func TestInstallWritesTheStamp(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	m := projectWith(t, repo, "utils", "v1.0.0")

	commit, err := Install(m, "utils", io.Discard)
	if err != nil {
		t.Fatal(err)
	}

	stamp, ok := ReadStamp(m.Dir("utils"))
	if !ok {
		t.Fatal("no stamp written")
	}
	if stamp.Version != "v1.0.0" || stamp.Commit != commit || stamp.Source != repo {
		t.Errorf("stamp = %+v, want source %q version v1.0.0 commit %q", stamp, repo, commit)
	}
}

func TestInstallAllSkipsWhatAlreadyMatches(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	m := projectWith(t, repo, "utils", "v1.0.0")

	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	// Leave a marker inside the installed directory. A skipped install must not
	// disturb it; a re-clone would wipe it out.
	marker := filepath.Join(m.Dir("utils"), "marker")
	if err := os.WriteFile(marker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(marker); err != nil {
		t.Error("a second InstallAll re-cloned an up-to-date planet")
	}
}

func TestInstallAllReinstallsOnVersionChange(t *testing.T) {
	repo := newRepo(t, "v1.0.0", "v2.0.0")
	m := projectWith(t, repo, "utils", "v1.0.0")

	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	entry := m.Planets["utils"]
	entry.Version = "v2.0.0"
	entry.Commit = ""
	m.Planets["utils"] = entry

	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	stamp, _ := ReadStamp(m.Dir("utils"))
	if stamp.Version != "v2.0.0" {
		t.Errorf("stamp version = %q, want v2.0.0", stamp.Version)
	}

	// v2.0.0 was the second commit, so its content differs from v1.0.0's.
	content, _ := os.ReadFile(filepath.Join(m.Dir("utils"), "mod.rl"))
	if string(content) != "export V = 1\n" {
		t.Errorf("content = %q, want the v2.0.0 content", content)
	}
}

func TestInstallAllRecordsTheResolvedCommit(t *testing.T) {
	repo := newRepo(t, "v1.0.0", "v2.0.0")

	// A manifest entry with no version, as a hand-written one might be.
	m := projectWith(t, repo, "utils", "")

	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	reloaded, err := LoadFrom(m.Root())
	if err != nil {
		t.Fatal(err)
	}

	entry := reloaded.Planets["utils"]
	if entry.Commit == "" {
		t.Error("the resolved commit was not written back to the manifest")
	}
	if entry.Version != "v2.0.0" {
		t.Errorf("version = %q, want the highest tag v2.0.0", entry.Version)
	}
}

func TestInstallLeavesTheOldCopyOnFailure(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	m := projectWith(t, repo, "utils", "v1.0.0")

	if _, err := Install(m, "utils", io.Discard); err != nil {
		t.Fatal(err)
	}

	// Ask for a tag that does not exist. The working copy must survive.
	entry := m.Planets["utils"]
	entry.Version = "v9.9.9"
	m.Planets["utils"] = entry

	if _, err := Install(m, "utils", io.Discard); err == nil {
		t.Fatal("Install succeeded for a missing tag")
	}

	if _, err := os.Stat(filepath.Join(m.Dir("utils"), "mod.rl")); err != nil {
		t.Error("a failed install destroyed the previous copy")
	}
	if _, err := os.Stat(m.Dir("utils") + ".incoming"); !os.IsNotExist(err) {
		t.Error("a failed install left its staging directory behind")
	}
}

func TestRemoveDeletesFromDiskAndManifest(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	m := projectWith(t, repo, "utils", "v1.0.0")

	if _, err := Install(m, "utils", io.Discard); err != nil {
		t.Fatal(err)
	}
	if err := Remove(m, "utils", io.Discard); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(m.Dir("utils")); !os.IsNotExist(err) {
		t.Error("the directory survived remove")
	}

	reloaded, _ := LoadFrom(m.Root())
	if _, ok := reloaded.Planets["utils"]; ok {
		t.Error("the manifest entry survived remove")
	}
}

func TestLocalModuleConflict(t *testing.T) {
	m := New(t.TempDir())

	if _, conflict := LocalModuleConflict(m, "utils"); conflict {
		t.Error("reported a conflict with nothing on disk")
	}

	if err := os.WriteFile(filepath.Join(m.Root(), "utils.rl"), []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, conflict := LocalModuleConflict(m, "utils"); !conflict {
		t.Error("missed a conflict with a local utils.rl")
	}

	if _, conflict := LocalModuleConflict(m, "other"); conflict {
		t.Error("reported a conflict for an unrelated alias")
	}
}

// moveBranch advances main in a repo built by newRepo, so drift can be tested.
func moveBranch(t *testing.T, repo, content string) {
	t.Helper()

	if err := os.WriteFile(filepath.Join(repo, "mod.rl"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	for _, args := range [][]string{{"add", "mod.rl"}, {"commit", "--quiet", "-m", "moved"}} {
		cmd := exec.Command("git", args...)
		cmd.Dir = repo
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %s", args, out)
		}
	}
}

// TestInstallPinsToTheRecordedCommit is the property that makes the commit
// field load-bearing rather than decorative. A branch advances and a tag can be
// force-moved, so re-resolving the ref would hand back different code than the
// manifest describes.
func TestInstallPinsToTheRecordedCommit(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	m := projectWith(t, repo, "dep", "main")

	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	pinned := m.Planets["dep"].Commit
	if pinned == "" {
		t.Fatal("no commit was recorded")
	}

	moveBranch(t, repo, "export V = 99\n")

	// A fresh checkout, as a CI job would have.
	if err := os.RemoveAll(m.PlanetsDir()); err != nil {
		t.Fatal(err)
	}
	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	if got := m.Planets["dep"].Commit; got != pinned {
		t.Errorf("the recorded commit changed from %s to %s", pinned, got)
	}

	stamp, _ := ReadStamp(m.Dir("dep"))
	if stamp.Commit != pinned {
		t.Errorf("installed %s, want the pinned %s", stamp.Commit, pinned)
	}

	content, _ := os.ReadFile(filepath.Join(m.Dir("dep"), "mod.rl"))
	if string(content) == "export V = 99\n" {
		t.Error("install followed the branch instead of the recorded commit")
	}
}

// TestInstallFollowsTheRefWhenNoCommitIsRecorded is the counterpart: without a
// commit there is nothing to pin to, so the ref decides.
func TestInstallFollowsTheRefWhenNoCommitIsRecorded(t *testing.T) {
	repo := newRepo(t, "v1.0.0")
	m := projectWith(t, repo, "dep", "main")

	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	moveBranch(t, repo, "export V = 99\n")

	// Clearing the commit is what planet get does for an explicit version.
	entry := m.Planets["dep"]
	entry.Commit = ""
	m.Planets["dep"] = entry

	if err := os.RemoveAll(m.PlanetsDir()); err != nil {
		t.Fatal(err)
	}
	if err := InstallAll(m, io.Discard); err != nil {
		t.Fatal(err)
	}

	content, _ := os.ReadFile(filepath.Join(m.Dir("dep"), "mod.rl"))
	if string(content) != "export V = 99\n" {
		t.Errorf("content = %q, want the advanced branch content", content)
	}
}

// TestCheckoutAcceptsABranch covers using a branch name as a version, which
// git's clone accepts as readily as a tag.
func TestCheckoutAcceptsABranch(t *testing.T) {
	repo := newRepo(t, "v1.0.0")

	if _, err := Checkout(repo, "main", filepath.Join(t.TempDir(), "x")); err != nil {
		t.Errorf("Checkout of a branch failed: %s", err)
	}
}
