# Planets
> 👉 Planets were introduced in `0.25`

A **planet** is a reusable RocketLang library: a git repository of `.rl` files
that a project pulls in and imports. Planets live in `.planets/` and are
recorded in `planets.yml`.

## Getting started

```
$ rocket-lang planet init
created planets.yml
added .planets/ to .gitignore

$ rocket-lang planet get flipez/rocket-lang-utils
resolved flipez/rocket-lang-utils to v1.1.0
installed utils v1.1.0 (364d720) to .planets/utils
recorded in planets.yml

import it with:  import "utils/<module>"
```

```js
import "utils/strings" as strings

puts(strings.Snake("hello world"))  // hello_world
```

## Sources

A source can be written four ways:

```
flipez/rocket-lang-utils          github.com is assumed
codeberg.org/flipez/utils         any host, when the first segment has a dot
https://example.com/utils.git     an explicit git URL
../sibling-project                a local path, for a monorepo or a private planet
```

Add `@<version>` to pin one:

```
rocket-lang planet get flipez/rocket-lang-utils@v1.2.0
```

## Names

A planet is imported under an **alias**, which is the key in `planets.yml` and
the directory name under `.planets/`. It is derived from the source, dropping a
redundant `rocket-lang-` prefix, so `flipez/rocket-lang-utils` becomes `utils`.

Choose your own with `--as`:

```
rocket-lang planet get codeberg.org/flipez/utils --as helpers
```

An alias has to be usable in an import. `my-lib` is not, because `my-lib.Foo`
parses as subtraction, so `--as` is required for a source whose name would
produce one.

## Importing

```js
import "<alias>/<module>"
```

The module is a file inside the planet, written **without** the `.rl`
extension. A planet with `strings.rl` and `math.rl` offers:

```js
import "utils/strings" as strings
import "utils/math" as math
```

A planet directory is not importable on its own — `import "utils"` does not
resolve. Imports name a file, as they do everywhere else in the language.

Your own modules always win a name clash: the working directory is searched
before `.planets/`. `planet get` refuses an alias that a local module would
shadow, so the conflict surfaces at install time rather than as a mystery at
run time.

## Versions

Versions are git tags read as semver, and always exact — there are no ranges.
`planet get` without a version picks the highest `major.minor.patch` tag,
ignoring prereleases, and records both the tag and the commit it resolved to:

```yaml
planets:
  utils:
    source: flipez/rocket-lang-utils
    version: v1.1.0
    commit: 364d7201a12f07ce31a46671b8499535b403894f
```

The commit is what makes an install reproducible, because a tag can be moved
after the fact.

Running `get` again on an installed planet reports what is there and changes
nothing, so a routine `get` can never move a dependency by surprise:

```
$ rocket-lang planet get flipez/rocket-lang-utils
utils is already installed at v1.1.0
pass an explicit version to change it, for example flipez/rocket-lang-utils@v1.2.3
```

## Commands

```
rocket-lang planet init                       create planets.yml
rocket-lang planet get <source>[@<version>]   fetch, install and record
rocket-lang planet install                    install everything the manifest lists
rocket-lang planet list                       show what this project uses
rocket-lang planet remove <alias>             delete from disk and the manifest
```

`planet install` is what a checkout or a CI job runs. It skips anything already
correct, so repeated runs are cheap:

```
$ rocket-lang planet install
up to date  helpers v1.1.0
up to date  utils v1.1.0
```

`planet list` reports when what is installed disagrees with the manifest.

## What is on disk

```
planets.yml
.planets/
  utils/
    .planet          <- what is installed here
    strings.rl
    math.rl
main.rl
```

`.planet` records the source, version and commit of that directory. It carries
no timestamp, so it is a pure function of what was installed.

`.planets/` is gitignored by default, and `planet init` adds it to an existing
`.gitignore`. Commit it instead if you want the tree self-contained — see
publishing below.

## Publishing a planet

A planet is an ordinary git repository of `.rl` files with semver tags. Nothing
else is required: no manifest, no registration.

```
strings.rl
math.rl
```

Export what callers should see, and keep the rest private:

```js
// strings.rl
export def Snake(s)
  return s.replace(" ", "_")
end
```

Tag a release:

```
git tag -a v1.0.0 -m v1.0.0 && git push --tags
```

**If your planet depends on another planet**, install it and commit your own
`.planets/`, then import it relatively:

```js
import "./.planets/other/mod" as other
```

`planet get` clones and does not install recursively, so a dependency that is
not committed will be missing for anyone who fetches your planet.

## Requirements and limits

- `planet` needs the `git` command on `PATH`.
- Planets are unavailable in the browser playground: there is no filesystem to
  install into.
- Version ranges, a lockfile and automatic transitive dependencies are not
  implemented.
