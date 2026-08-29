//go:build !wasm

package planet

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// newRepo builds a real git repository with the given tags, so the git layer is
// exercised against git itself rather than a mock, without touching a network.
func newRepo(t *testing.T, tags ...string) string {
	t.Helper()

	return newRepoOnBranch(t, "main", tags...)
}

// newRepoOnBranch builds a real git repository whose default branch is named
// branch, so the default-branch lookup is tested against something other than
// main.
func newRepoOnBranch(t *testing.T, branch string, tags ...string) string {
	t.Helper()

	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not available")
	}

	dir := t.TempDir()

	run := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %s", args, out)
		}
	}

	run("init", "--quiet", "--initial-branch", branch)
	run("config", "user.email", "test@example.com")
	run("config", "user.name", "test")

	for i, tag := range tags {
		name := filepath.Join(dir, "mod.rl")
		if err := os.WriteFile(name, []byte("export V = "+string(rune('0'+i))+"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		run("add", "mod.rl")
		run("commit", "--quiet", "-m", tag)
		run("tag", "-a", tag, "-m", tag)
	}

	return dir
}

func TestTagsFiltersAndOrders(t *testing.T) {
	// Deliberately out of order, with tags that are not versions at all.
	repo := newRepo(t, "v1.0.0", "v1.10.0", "v1.2.0", "nightly", "v2.0.0", "v1.2.3-rc1")

	tags, err := Tags(repo)
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"v1.0.0", "v1.2.0", "v1.10.0", "v2.0.0"}
	if len(tags) != len(want) {
		t.Fatalf("Tags() = %v, want %v", tags, want)
	}
	for i := range want {
		if tags[i] != want[i] {
			t.Fatalf("Tags() = %v, want %v", tags, want)
		}
	}
}

func TestLatestTagIsHighestNotNewest(t *testing.T) {
	// v1.10.0 is committed after v2.0.0, so anything ordering by date or by
	// string would pick the wrong one.
	repo := newRepo(t, "v2.0.0", "v1.10.0")

	latest, err := LatestTag(repo)
	if err != nil {
		t.Fatal(err)
	}

	if latest != "v2.0.0" {
		t.Errorf("LatestTag() = %q, want v2.0.0", latest)
	}
}

// TestLatestTagWithoutVersionTags: no tags is not an error. A planet that has
// not tagged a release is still usable through its default branch, so the empty
// answer lets ResolveVersion fall back rather than failing the install.
func TestLatestTagWithoutVersionTags(t *testing.T) {
	repo := newRepo(t, "nightly")

	tag, err := LatestTag(repo)
	if err != nil {
		t.Fatalf("LatestTag errored on a repo with no version tags: %s", err)
	}
	if tag != "" {
		t.Errorf("LatestTag() = %q, want an empty string", tag)
	}
}

// TestDefaultBranchIsReadFromTheRemote covers not hardcoding "main": the branch
// name is whatever the remote says its HEAD points at.
func TestDefaultBranchIsReadFromTheRemote(t *testing.T) {
	repo := newRepoOnBranch(t, "trunk", "v1.0.0")

	branch, err := DefaultBranch(repo)
	if err != nil {
		t.Fatal(err)
	}
	if branch != "trunk" {
		t.Errorf("DefaultBranch() = %q, want trunk", branch)
	}
}

func TestResolveVersionPrefersATag(t *testing.T) {
	repo := newRepo(t, "v1.0.0", "v2.0.0")

	version, isTag, err := ResolveVersion(repo)
	if err != nil {
		t.Fatal(err)
	}
	if version != "v2.0.0" || !isTag {
		t.Errorf("ResolveVersion() = %q, isTag %v; want v2.0.0, true", version, isTag)
	}
}

func TestResolveVersionFallsBackToTheDefaultBranch(t *testing.T) {
	// Tagged, but not with anything that parses as a version.
	repo := newRepoOnBranch(t, "trunk", "nightly")

	version, isTag, err := ResolveVersion(repo)
	if err != nil {
		t.Fatal(err)
	}
	if version != "trunk" || isTag {
		t.Errorf("ResolveVersion() = %q, isTag %v; want trunk, false", version, isTag)
	}
}

func TestCheckoutReturnsCommitAndDropsHistory(t *testing.T) {
	repo := newRepo(t, "v1.0.0", "v2.0.0")
	dest := filepath.Join(t.TempDir(), "utils")

	commit, err := Checkout(repo, "v1.0.0", dest)
	if err != nil {
		t.Fatal(err)
	}

	if len(commit) != 40 {
		t.Errorf("commit = %q, want a 40-character sha", commit)
	}

	if _, err := os.Stat(filepath.Join(dest, "mod.rl")); err != nil {
		t.Errorf("checkout is missing mod.rl: %s", err)
	}

	// History is not kept; the stamp file records provenance instead.
	if _, err := os.Stat(filepath.Join(dest, ".git")); !os.IsNotExist(err) {
		t.Error(".git survived the checkout")
	}

	// v1.0.0 wrote V = 0 and v2.0.0 wrote V = 1, so the content proves the
	// requested tag was checked out rather than the default branch.
	content, err := os.ReadFile(filepath.Join(dest, "mod.rl"))
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "export V = 0\n" {
		t.Errorf("checked out %q, want the v1.0.0 content", content)
	}
}

func TestCheckoutRejectsUnknownTag(t *testing.T) {
	repo := newRepo(t, "v1.0.0")

	if _, err := Checkout(repo, "v9.9.9", filepath.Join(t.TempDir(), "x")); err == nil {
		t.Error("Checkout succeeded for a tag that does not exist")
	}
}

func TestParseVersion(t *testing.T) {
	ok := map[string]version{
		"v1.2.3": {1, 2, 3},
		"1.2.3":  {1, 2, 3},
		"v0.0.0": {0, 0, 0},
	}
	for tag, want := range ok {
		got, valid := parseVersion(tag)
		if !valid || got != want {
			t.Errorf("parseVersion(%q) = %+v, %v", tag, got, valid)
		}
	}

	// Prereleases are rejected so a bare `planet get` never selects one.
	for _, tag := range []string{"", "v1", "v1.2", "v1.2.3.4", "v1.2.x", "nightly", "v1.2.3-rc1", "v-1.2.3"} {
		if _, valid := parseVersion(tag); valid {
			t.Errorf("parseVersion(%q) accepted it", tag)
		}
	}
}
