//go:build !wasm

package planet

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const usage = `Usage: rocket-lang planet <command> [arguments]

Commands:
  init                        create %s in the current directory
  get <source>[@<version>]    fetch a planet, install it and record it
  install                     install every planet the manifest lists
  list                        show what this project uses
  remove <alias>              delete a planet from disk and the manifest

A source is owner/name (GitHub by default), host/owner/name, a git URL or a
local path:

  rocket-lang planet get flipez/rocket-lang-utils
  rocket-lang planet get flipez/rocket-lang-utils@v1.2.0
  rocket-lang planet get codeberg.org/flipez/utils --as helpers
`

// Command runs the planet subcommand. It returns a process exit code rather
// than exiting, so it stays testable.
func Command(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintf(stderr, usage, ManifestName)
		return 2
	}

	verb, rest := args[0], args[1:]

	var err error
	switch verb {
	case "init":
		err = cmdInit(stdout)
	case "get":
		err = cmdGet(rest, stdout)
	case "install":
		err = cmdInstall(stdout)
	case "list":
		err = cmdList(stdout)
	case "remove":
		err = cmdRemove(rest, stdout)
	case "help", "--help", "-h":
		fmt.Fprintf(stdout, usage, ManifestName)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown planet command %q\n\n", verb)
		fmt.Fprintf(stderr, usage, ManifestName)
		return 2
	}

	if err != nil {
		fmt.Fprintf(stderr, "planet %s: %s\n", verb, err)
		return 1
	}

	return 0
}

func cmdInit(out io.Writer) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := filepath.Join(cwd, ManifestName)
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("%s already exists", ManifestName)
	}

	manifest := New(cwd)
	if err := manifest.Save(); err != nil {
		return err
	}

	fmt.Fprintf(out, "created %s\n", ManifestName)

	if err := ignorePlanetsDir(cwd, out); err != nil {
		return err
	}

	return nil
}

// ignorePlanetsDir adds the planet directory to an existing .gitignore.
// Consumers are expected to ignore it and run planet install; an author who
// vendors dependencies inside their own planet overrides that deliberately.
func ignorePlanetsDir(root string, out io.Writer) error {
	path := filepath.Join(root, ".gitignore")

	content, err := os.ReadFile(path)
	if err != nil {
		// No .gitignore to amend. Creating one is more presumptuous than
		// leaving it alone, so say what to do instead.
		fmt.Fprintf(out, "note: add %s/ to .gitignore, or commit it to vendor your planets\n", DirName)
		return nil
	}

	for _, line := range strings.Split(string(content), "\n") {
		if strings.TrimSpace(strings.TrimSuffix(line, "/")) == DirName {
			return nil
		}
	}

	amended := string(content)
	if !strings.HasSuffix(amended, "\n") {
		amended += "\n"
	}
	amended += DirName + "/\n"

	if err := os.WriteFile(path, []byte(amended), 0o644); err != nil {
		return err
	}

	fmt.Fprintf(out, "added %s/ to .gitignore\n", DirName)

	return nil
}

func cmdGet(args []string, out io.Writer) error {
	sources, alias, err := parseGetArgs(args)
	if err != nil {
		return err
	}

	if len(sources) != 1 {
		return fmt.Errorf("expected one source, for example flipez/rocket-lang-core")
	}

	raw, version, _ := strings.Cut(sources[0], "@")

	source, err := ParseSource(raw)
	if err != nil {
		return err
	}

	if alias == "" {
		if alias, err = DefaultAlias(raw); err != nil {
			return err
		}
	}
	if !IsValidAlias(alias) {
		return fmt.Errorf("%q would not be usable in an import; choose a name made of letters, digits and underscores", alias)
	}

	manifest, err := currentManifest()
	if err != nil {
		return err
	}

	if conflict, found := LocalModuleConflict(manifest, alias); found {
		return fmt.Errorf("this project already has %s, which would shadow a planet named %q; install it under a different name with --as", conflict, alias)
	}

	// A bare get on an installed planet reports and stops, so a routine get can
	// never silently move a dependency.
	if existing, ok := manifest.Planets[alias]; ok && version == "" {
		if stamp, stamped := ReadStamp(manifest.Dir(alias)); stamped {
			fmt.Fprintf(out, "%s is already installed at %s\n", alias, stamp.Version)
			fmt.Fprintf(out, "pass an explicit version to change it, for example %s@%s\n", raw, "v1.2.3")
			return nil
		}

		version = existing.Version
	}

	if version == "" {
		isTag := false
		if version, isTag, err = ResolveVersion(source.URL); err != nil {
			return err
		}

		if isTag {
			fmt.Fprintf(out, "resolved %s to %s\n", raw, version)
		} else {
			// Worth saying out loud: a branch is not a release, and it will
			// move. The commit is pinned, so this install stays reproducible.
			fmt.Fprintf(out, "resolved %s to the %s branch; it publishes no version tags\n", raw, version)
		}
	}

	manifest.Planets[alias] = Entry{Source: raw, Version: version}

	commit, err := Install(manifest, alias, out)
	if err != nil {
		delete(manifest.Planets, alias)
		return err
	}

	manifest.Planets[alias] = Entry{Source: raw, Version: version, Commit: commit}

	if err := manifest.Save(); err != nil {
		return err
	}

	fmt.Fprintf(out, "recorded in %s\n", ManifestName)
	fmt.Fprintf(out, "\nimport it with:  import \"%s/<module>\"\n", alias)

	return nil
}

// parseGetArgs reads `get` arguments. It is hand-rolled rather than using
// flag.FlagSet because the standard library stops parsing flags at the first
// positional argument, which would make `get <source> --as name` fail -- and
// that reads more naturally than putting the flag first.
func parseGetArgs(args []string) (sources []string, alias string, err error) {
	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "--as", arg == "-as":
			if i+1 >= len(args) {
				return nil, "", fmt.Errorf("--as needs a name")
			}
			alias = args[i+1]
			i++

		case strings.HasPrefix(arg, "--as="):
			alias = strings.TrimPrefix(arg, "--as=")

		case strings.HasPrefix(arg, "-"):
			return nil, "", fmt.Errorf("unknown option %q", arg)

		default:
			sources = append(sources, arg)
		}
	}

	return sources, alias, nil
}

func cmdInstall(out io.Writer) error {
	manifest, err := currentManifest()
	if err != nil {
		return err
	}

	return InstallAll(manifest, out)
}

func cmdList(out io.Writer) error {
	manifest, err := currentManifest()
	if err != nil {
		return err
	}

	aliases := manifest.Aliases()
	if len(aliases) == 0 {
		fmt.Fprintf(out, "no planets in %s\n", ManifestName)
		return nil
	}

	width := 0
	for _, alias := range aliases {
		if len(alias) > width {
			width = len(alias)
		}
	}

	for _, alias := range aliases {
		entry := manifest.Planets[alias]

		state := "not installed"
		if stamp, ok := ReadStamp(manifest.Dir(alias)); ok {
			switch {
			case stamp.Matches(entry):
				state = "installed"
			default:
				state = fmt.Sprintf("installed %s, manifest says %s", stamp.Version, entry.Version)
			}
		}

		fmt.Fprintf(out, "%-*s  %-10s  %s  (%s)\n", width, alias, entry.Version, entry.Source, state)
	}

	return nil
}

func cmdRemove(args []string, out io.Writer) error {
	if len(args) != 1 {
		return fmt.Errorf("expected one planet name")
	}

	manifest, err := currentManifest()
	if err != nil {
		return err
	}

	return Remove(manifest, args[0], out)
}

func currentManifest() (*Manifest, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	manifest, err := Load(cwd)
	if err != nil {
		return nil, fmt.Errorf("%w\nrun `rocket-lang planet init` to start one", err)
	}

	return manifest, nil
}
