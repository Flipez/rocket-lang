# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

RocketLang is a tree-walking interpreter written in Go for a Ruby-flavoured
descendant of MonkeyLang. There is no bytecode VM; the parser builds an AST and
the evaluator walks it.

## Commands

```bash
go build -o rocket-lang .        # build the interpreter
go run . path/to/program.rl      # run a program
go run . -e 'puts("hi")'         # run a snippet
go run .                         # REPL

go test ./...                    # all tests
go test ./object -run TestString # one test in one package
./coverage.sh                    # what CI runs (per-package cover profiles -> coverage.txt)
```

Language-level tests live in `tests/*.rl` with a sibling `*.expected` holding
the exact stdout. `go test -run TestRocketlangCode .` runs all of them; adding a
pair of files is all it takes to add a case. Note the interpreter prints the
program's final value, so most fixtures end in a bare `nil`.

### Docs

`docs/` is a Docusaurus site with `onBrokenLinks: 'throw'`.

```bash
go run docs/generate.go          # from the repo root; regenerates the reference pages
cd docs && yarn && yarn build    # link check + site build
```

`docs/docs/literals/*.md` and `docs/docs/builtins/*.md` are **generated** — never
edit them. Signatures come from the registered `MethodLayout`s; prose, examples
and expected output come from `docs/literals/*.yml` and `docs/builtins/*.yml`
rendered through `docs/templates/`. Everything else under `docs/docs/` is
hand-written.

### WebAssembly

```bash
GOOS=js GOARCH=wasm go build -tags wasm -o wasm/main.wasm .
```

The `wasm` tag is required, not optional: `repl`, `planet` and `utilities` each
have `//go:build wasm` / `//go:build !wasm` file pairs that redeclare each other
if the tag and GOOS disagree.

Releases are cut by tagging (`git tag -a v0.24.0 -m v0.24.0`); see `RELEASE.md`
for the full checklist, including the docs versioning step.

## Architecture

`lexer` → `parser` (Pratt, precedence table at the top of `parser/parser.go`) →
`ast` → `evaluator` → `object`. Each of `ast/`, `parser/` and `evaluator/` has
roughly one file per node kind (`if.go`, `foreach.go`, `import.go`, …), so a new
piece of syntax means touching the same-named file in each.

Tokens carry `File`, `LineNumber` and `LinePosition`; runtime errors are
formatted `file:line:col: ... Error: ...` from them. Keep that when adding
errors.

### The object package is the centre of gravity

`object.Object` is `Type() / Inspect() / InvokeMethod()`. Every value type
registers its methods in its own file's `init()` into the package-level
`objectMethods[TYPE]` map; `objectMethods["*"]` holds the methods every value
has (`to_string`, `to_integer`, `is_a?`, `methods`, `help`, …) and is the fallback in
`objectMethodLookup`.

Argument checking is declarative, via `MethodLayout`: `Arg(...)`,
`OptArg(...)` (may be omitted) and `OverloadArg(...)` (variadic). Types may be
concrete (`STRING_OBJ`) or **type groups** (`ANY`, `HASHABLE`, `STRINGABLE`,
`INTEGERABLE`, `COMPARABLE`, `NUMERIC`, `CALLABLE`) resolved by asking the
object, usually through a Go interface. Prefer a group over enumerating types —
that is what stops a newly added type from being silently rejected. `is_a?`,
`type_groups` and the generated docs all read the same layout, so the signature
cannot drift from the implementation.

A new value type must be added to `knownObjectTypes` in `object/object.go`, or
`is_a?("YOURTYPE")` reports it as an unknown name.

`object` cannot import `evaluator` (the dependency runs the other way), so
`evaluator/register.go`'s `init()` injects `Eval` and `applyFunction` back into
`object`. That is how an object method can call a RocketLang callback
(`Array#each` and friends).

### Builtins, modules and imports

`stdlib/std.go`'s `init()` registers global functions (`puts`, `raise`) and
builtin modules (`Math`, `HTTP`, `JSON`, `IO`, `OS`, `Time`), each defined in
its own `stdlib/<name>.go` as a `map[string]*object.BuiltinFunction` plus a
properties map.

RocketLang-level modules are `.rl` files: `export` populates a module's
attribute hash, `import "x" as y only(a, b)` binds it. `utilities.SearchPaths`
resolves a bare name (`ROCKETLANGPATH` entries, then the working directory, then
the project's `planets/` directory — see `configureSearchPath` in `main.go`),
while `./`-relative paths resolve against the importing file. An
`object.ModuleRegistry` on the environment caches each resolved file so repeated
imports share one `*object.Module`, and detects circular imports.

`planet/` is the package manager reached as `rocket-lang planet <cmd>`
(subcommands are dispatched in `main.go` before flag parsing). It fetches git
sources into `planets/` and records them in a manifest.

## Conventions

- Comments here explain *why* a thing is the way it is, usually naming the bug
  or the wrong behaviour that motivated it. Match that when adding code;
  restating what the line does is out of place.
- The language uses `end`-terminated blocks, not braces, and conditions need no
  parentheses (`if x < 3` … `elif` … `else` … `end`). Curly braces are hash
  literals only.
- Exit codes matter and are tested: an uncaught error, a parse error or an
  unreadable file must exit non-zero (`main_test.go`).
- Evaluator/object tests drive whole snippets of RocketLang through
  `testEval`/`testInput` table cases rather than constructing AST nodes.
- Commits are conventional and lowercase: `fix(object): …`, `feat(planet): …`,
  `test(stdlib): …`.
