# Planets
> 👉 Planets were introduced in `0.24`

A **planet** is a reusable RocketLang library: a git repository of `.rl` files
that a project pulls in and imports. Planets live in `.planets/` and are
recorded in `planets.yml`.

## Getting started

```
$ rocket-lang planet init
created planets.yml
added .planets/ to .gitignore

$ rocket-lang planet get flipez/rocket-lang-core@main
installed core main (7441c48) to .planets/core
recorded in planets.yml

import it with:  import "core/<module>"
```

```js
import "core/list" as list
import "core/stats" as stats

scores = [42, 17, 99, 63]

puts(list.Filter(scores, def(n) return n > 40 end))  // [42, 99, 63]
puts(stats.Mean(scores))                             // 55.25
```

[flipez/rocket-lang-core](https://github.com/flipez/rocket-lang-core) is a
small planet of generic helpers, and is used as the example throughout this
page.

## Sources

A source can be written four ways:

```
flipez/rocket-lang-core           github.com is assumed
codeberg.org/flipez/utils         any host, when the first segment has a dot
https://example.com/utils.git     an explicit git URL
../sibling-project                a local path, for a monorepo or a private planet
```

Add `@<version>` to pin one:

```
rocket-lang planet get flipez/rocket-lang-core@v1.2.0
```

## Names

A planet is imported under an **alias**, which is the key in `planets.yml` and
the directory name under `.planets/`. It is derived from the source, dropping a
redundant `rocket-lang-` prefix, so `flipez/rocket-lang-core` becomes `core`.

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
extension. A planet with `list.rl` and `stats.rl` offers:

```js
import "core/list" as list
import "core/stats" as stats
```

A planet directory is not importable on its own — `import "core"` does not
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
  core:
    source: flipez/rocket-lang-core
    version: main
    commit: 7441c4811ff048ed4e4ee1a1a371ca837f4f52ef
```

The commit is what makes an install reproducible. Once recorded, `planet
install` checks out **that commit** rather than re-resolving the reference, so a
tag that is force-moved or a branch that advances cannot change what a checkout
gets. Only an explicit `planet get <source>@<version>` moves it.

### Branches and commits

A version does not have to be a tag. Any reference git understands works,
which is useful for a planet that has not tagged a release yet:

```
rocket-lang planet get flipez/rocket-lang-core@main
rocket-lang planet get flipez/rocket-lang-core@6daa0cc
```

A bare `planet get` still requires tags, because there is no sensible default
otherwise:

```
$ rocket-lang planet get flipez/rocket-lang-core
planet get: https://github.com/flipez/rocket-lang-core publishes no version
tags; pass an explicit @version
```

Tracking a branch is not the same as following it: the commit is recorded at
install time and pinned from then on, so everyone who checks the project out
gets the same code.

Running `get` again on an installed planet reports what is there and changes
nothing, so a routine `get` can never move a dependency by surprise:

```
$ rocket-lang planet get flipez/rocket-lang-core
core is already installed at main
pass an explicit version to change it, for example flipez/rocket-lang-core@v1.2.3
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
up to date  core main
```

`planet list` reports when what is installed disagrees with the manifest.

## What is on disk

```
planets.yml
.planets/
  core/
    .planet          <- what is installed here
    list.rl
    stats.rl
    strings.rl
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
