# RocketLang v1.0 — Steps 0 & 1: Retire Exercises, Library Rename

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Retire the exercise apparatus while keeping its 39 programs as golden-file
regression tests, then rename the entire standard library to the v1.0 naming rules —
without changing any syntax.

**Architecture:** This is a pure library-and-names change. The lexer, parser, AST and
evaluator are not touched. Every method is registered into `objectMethods[TYPE]` maps in
`object/*.go` `init()` functions, and into `*Functions` maps in `stdlib/*.go`. A rename
is: change the registration key, change the matching key in `docs/literals/*.yml` or
`docs/builtins/*.yml`, change the Go unit tests, change the `.rl` corpus. A guard test
added in Task 1 makes the docs half impossible to forget.

**Tech Stack:** Go 1.24 (CI uses 1.25), `gopkg.in/yaml.v3`, `stretchr/testify`,
Docusaurus 3 for the site.

**Spec:** `docs/superpowers/specs/2026-08-31-language-redesign-design.md`

## Global Constraints

- **Do not run `go test ./exercises -update`.** It regenerates expectations *from* the
  solutions, which would let a rewrite silently redefine correctness. The expectations
  are the oracle. (Moot after Task 3, which deletes that harness — but fatal before it.)
- **No syntax changes in this plan.** No `#` comments, no `..` ranges, no `->` lambdas,
  no parenless calls. Those are steps 2 and 5–8. Code written here still uses `//`
  comments, `def(x) … end`, and `.method()` with parentheses.
- **`upto`/`downto` are NOT deleted here.** Their replacement is `(1..5).each(f)` and
  ranges do not exist until step 5.
- **Every task ends with a green test run and a commit.** No task may leave `main` broken.
- **`go test ./object` can hang.** One observed run parked on a goroutine at
  `object/http.go:65` (an HTTP listener that never shuts down) and failed after 600s.
  Always pass `-timeout 120s`. Scoped runs (`-run TestName`) are unaffected. Pre-existing;
  out of scope here.
- Method names are `snake_case`; types, modules and constants are `CapCase`.
- **A rename touches YAML keys AND values.** `docs/literals/*.yml` and
  `docs/builtins/*.yml` carry old method names in three places: the `methods:`/
  `functions:` **keys**, the worked examples in `input:`/`output:`, and the English in
  `description:`. 107 lines across 8 files are affected. The Task 1 guard test only
  checks keys, so stale examples pass it silently. Every rename task must therefore end
  with `grep -rn '<old-name>' docs/literals docs/builtins` returning nothing —
  `description:` prose usually needs a real edit rather than a substitution.
- **Never run `git add -A`.** Stage explicit paths. `DESIGN_HTTP_CLIENT.md` is untracked
  scratch belonging to the repo owner and must never be committed.
- A `!` method exists **only** as the counterpart of a same-named non-mutating method.
  `stringPair`/`arrayPair`/`hashCallbackPair` generate both from one base name, so
  renaming the base renames the pair.

---

## File Structure

**Created:**
- `docs/generate_test.go` — guard test: every registered method and stdlib function has
  a documentation entry. Package `main` (same as `generate.go`).
- `tests/lang/*.rl` + `tests/lang/*.expected` — 39 salvaged exercise programs.

**Modified:**
- `main_test.go:20` — widen the fixture glob to include `tests/lang/`.
- `object/object.go` — universal methods (`to_s`, `to_i`, `to_f`, `wat`).
- `object/string.go`, `array.go`, `hash.go`, `integer.go`, `float.go`, `error.go`,
  `matrix.go` — registration keys.
- `stdlib/std.go`, `puts.go`, `os.go`, `math.go` — global and module functions.
- `docs/literals/*.yml`, `docs/builtins/*.yml` — documentation keys must track renames.
- `docs/docs/**/*.md` (hand-written pages) — prose and examples using old names.
- All `.rl` files under `tests/`, `fixtures/`, `examples/`.
- `object/*_test.go`, `stdlib/os_test.go` — Go unit tests.

**Deleted:**
- `exercises/` entirely, `wasm/exercises.json`.

---

### Task 1: Guard test — every method must be documented

`docs/generate.go` silently skips undocumented methods (`if v, ok := docData.Methods[name]; ok`),
so a rename that misses the YAML produces a docs page with an empty description and **no
error**. Across ~60 renames that is certain to happen. This test makes it loud. Verified:
zero methods are undocumented today, so it passes on a clean tree.

**Files:**
- Create: `docs/generate_test.go`

**Interfaces:**
- Consumes: `object.ListObjectMethods()`, `stdlib.Modules` (both already exported).
- Produces: `TestEveryMethodIsDocumented`, run by every later task.

- [ ] **Step 1: Write the test**

Note the paths: `generate.go` is run from the repo root (`go run docs/generate.go`) so it
says `docs/literals/…`, but `go test ./docs` runs with `docs/` as the working directory,
so the test says `literals/…`.

```go
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/stdlib"
)

// docKeys reads the method or function names documented in a YAML file. It
// returns nil when the file is missing, which the caller reports as every
// name being undocumented.
func docKeys(t *testing.T, path, section string) map[string]bool {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("%s: %s", path, err)
		return nil
	}

	var doc struct {
		Methods   map[string]any `yaml:"methods"`
		Functions map[string]any `yaml:"functions"`
	}
	if err := yaml.Unmarshal(content, &doc); err != nil {
		t.Fatalf("%s: %s", path, err)
	}

	source := doc.Methods
	if section == "functions" {
		source = doc.Functions
	}

	keys := make(map[string]bool, len(source))
	for name := range source {
		keys[name] = true
	}

	return keys
}

// reportUndocumented fails with every missing name at once, sorted, rather
// than stopping at the first -- during a rename the whole list is what you
// want to see.
func reportUndocumented(t *testing.T, path string, have map[string]bool, want []string) {
	t.Helper()

	var missing []string
	for _, name := range want {
		if !have[name] {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)

	if len(missing) > 0 {
		t.Errorf("%s is missing %d entries: %s", path, len(missing), strings.Join(missing, " "))
	}
}

// TestEveryMethodIsDocumented guards the rename. generate.go skips a method it
// finds no documentation for, without complaining, so a renamed method quietly
// loses its description on the website instead of failing the build.
func TestEveryMethodIsDocumented(t *testing.T) {
	for objectType, methods := range object.ListObjectMethods() {
		name := "object"
		if objectType != "*" {
			name = strings.ToLower(string(objectType))
		}

		path := fmt.Sprintf("literals/%s.yml", name)

		names := make([]string, 0, len(methods))
		for method := range methods {
			names = append(names, method)
		}

		reportUndocumented(t, path, docKeys(t, path, "methods"), names)
	}
}

// TestEveryBuiltinFunctionIsDocumented is the same guard for the stdlib modules.
func TestEveryBuiltinFunctionIsDocumented(t *testing.T) {
	for _, module := range stdlib.Modules {
		path := fmt.Sprintf("builtins/%s.yml", strings.ToLower(module.Name))

		names := make([]string, 0, len(module.Functions))
		for function := range module.Functions {
			names = append(names, function)
		}

		reportUndocumented(t, path, docKeys(t, path, "functions"), names)
	}
}
```

- [ ] **Step 2: Run it — it must PASS on the unmodified tree**

Run: `go test ./docs -v -run 'Documented'`
Expected: `PASS` for both tests. If it fails, the tree is not clean — stop and
investigate before renaming anything, because the safety net is the point.

- [ ] **Step 3: Prove the guard actually catches a rename**

Temporarily rename one registration key and confirm the test fails:

```bash
sed -i '' 's/"nil?": ObjectMethod{/"is_nil?": ObjectMethod{/' object/object.go
go test ./docs -run 'TestEveryMethodIsDocumented'
```

Expected: FAIL, reporting `literals/object.yml is missing 1 entries: is_nil?`

- [ ] **Step 4: Revert the probe**

```bash
git checkout object/object.go
go test ./docs -run 'Documented'
```

Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add docs/generate_test.go
git commit -m "test(docs): fail the build when a method has no documentation

generate.go skips a method it finds no docs entry for and says nothing, so a
renamed method quietly loses its description on the website. The v1.0 rename
moves about sixty names; without this the first missed YAML key would only be
noticed by a reader."
```

---

### Task 2: Salvage the exercise solutions as golden-file tests

The 39 solutions are 245 lines covering combinations the Go unit tests do not — control
flow inside functions, hashes inside loops, errors across imports. They become fixtures
for `TestRocketlangCode`.

**Critical semantic difference:** `exercises/expected/*.txt` were generated with the
program's trailing `nil` **stripped**; `tests/*.expected` contain the **raw** output of
`runProgram`, which prints the final value. Compare `tests/if.expected` (ends with a
`nil` line) against any exercise expectation. So expectations must be **regenerated**
under `tests/` semantics, not copied.

Verified: no solution contains `import` or `IO.open`, so relocating them breaks no paths.

**Files:**
- Modify: `main_test.go:20`
- Create: `tests/lang/*.rl`, `tests/lang/*.expected`

**Interfaces:**
- Consumes: `runProgram(string, string) int` from `main.go:114`.
- Produces: 39 fixture pairs under `tests/lang/`, picked up by `TestRocketlangCode`.

- [ ] **Step 1: Confirm the oracle holds before touching anything**

Run: `go test ./exercises`
Expected: `ok`. If this fails, stop — the programs are not currently correct and must
not be promoted to fixtures.

- [ ] **Step 2: Widen the fixture glob**

`main_test.go:20` currently reads:

```go
	matches, err := fs.Glob(os.DirFS(testDir), "*.rl")
```

Replace with a walk, so fixtures may be grouped in subdirectories:

```go
	var matches []string
	err := fs.WalkDir(os.DirFS(testDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".rl") {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk test dir 'tests/': %s", err)
	}
```

`io/fs` and `strings` are already imported in `main_test.go`. The `err` declaration
changes from `:=` to being declared by the walk — check the surrounding lines compile.

- [ ] **Step 3: Run the existing fixtures to prove the walk still finds them**

Run: `go test -run TestRocketlangCode . -timeout 120s`
Expected: PASS, same 14 fixtures as before.

- [ ] **Step 4: Copy the solutions and generate expectations under tests/ semantics**

```bash
mkdir -p tests/lang
go build -o /tmp/rl .

for src in exercises/solutions/*/*.rl; do
  id="${src#exercises/solutions/}"        # 01_basics/01_puts.rl
  flat="${id%.rl}"                        # 01_basics/01_puts
  dest="tests/lang/$(echo "$flat" | tr '/' '_').rl"
  cp "$src" "$dest"
  /tmp/rl "$dest" > "${dest%.rl}.expected" 2>&1
done

ls tests/lang/*.rl | wc -l    # expect 39
```

The expectation is generated by running the current interpreter, which is sound *only*
because Step 1 proved these programs already produce their documented output and no
behaviour has changed yet. Do not repeat this trick after a rename lands.

- [ ] **Step 5: Verify the new fixtures pass**

Run: `go test -run TestRocketlangCode . -timeout 120s -v 2>&1 | tail -20`
Expected: PASS. A failure here means a program is not deterministic (time, randomness) —
delete that fixture rather than pinning a flaky expectation, and note which.

- [ ] **Step 6: Commit**

```bash
git add main_test.go tests/lang
git commit -m "test: keep the exercise programs as golden-file fixtures

The 39 exercise solutions cover combinations the unit tests do not -- control
flow inside functions, hashes inside loops, errors across imports -- and the
v1.0 rewrite needs an end-to-end oracle. The exercise apparatus around them is
going away; the programs are worth keeping.

Expectations are regenerated rather than copied: the exercise runner strips the
program's trailing nil and this harness does not."
```

---

### Task 3: Delete the exercise apparatus

**Files:**
- Delete: `exercises/`, `wasm/exercises.json`

- [ ] **Step 1: Check what references the bundle before deleting it**

```bash
grep -rn "exercises" --include="*.go" --include="*.js" --include="*.html" --include="*.py" --include="*.sh" . | grep -v "^./docs/superpowers" | grep -v "^./exercises/"
```

Expected hits: `wasm/index.html`, `wasm/term.js` (the playground UI loads the bundle).
Note each one — the playground must still load with the bundle gone.

- [ ] **Step 2: Remove the exercise tree and bundle**

```bash
git rm -r --quiet exercises wasm/exercises.json
```

- [ ] **Step 3: Remove the playground's exercise UI**

Good news, verified: the playground **already degrades gracefully**. `wasm/index.html:398`
wraps the fetch in `try`/`catch` returning `[]`, and line 404 guards the whole UI behind
`if (exercises.length)`. The dropdown at line 60 carries `d-none` and is only revealed
when exercises load. So deleting the bundle leaves a working REPL — this step is dead-code
removal, not a repair.

Remove, in `wasm/index.html`:
- line 60, the `exerciseSelect` element, and the prev/next buttons at lines 64–65
- the `try`/`catch` fetch and the whole `if (exercises.length) { … }` block from line 397
- the `params.delete('exercise')` handling at line 283
- the stale comment at lines 187–189 referring to `exercises/generate.go` and `run.sh`

Check `wasm/term.js` for any `exerciseSelect`/`prevButton`/`nextButton` handlers left
dangling.

- [ ] **Step 3a: Load the playground and confirm it still works**

```bash
GOOS=js GOARCH=wasm go build -tags wasm -o wasm/main.wasm .
cd wasm && python3 server.py
```

Open the page, run `print("hi")` in the REPL, and confirm the browser console shows no
404 and no `ReferenceError`. Stop the server when done.

- [ ] **Step 4: Verify the build and the whole suite**

```bash
go build ./... && go vet ./...
go test -run TestRocketlangCode . -timeout 120s
go test ./docs ./lexer ./token ./parser ./evaluator ./planet ./utilities ./repl -timeout 120s
```

Expected: all PASS. `./exercises` no longer exists, so it is absent from the list.

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "chore: retire the exercise apparatus ahead of the v1.0 rewrite

The exercises teach a syntax that is about to change entirely; translating them
mechanically would produce bad teaching material. They are one day old and are
recoverable from history (c5c0f19, 13095da). Fresh exercises will be written
for v1.0 and given a proper home in the playground.

The 39 programs themselves were kept as fixtures in the previous commit. This
removes the task files, the runner, the expectation generator and the
repo-to-playground bundle coupling."
```

---

### Task 4: Universal methods — `to_s`/`to_i`/`to_f`, and `wat` → `help`

**Files:**
- Modify: `object/object.go` (the `objectMethods["*"]` block, ~line 367)
- Modify: `docs/literals/object.yml`
- Modify: `object/object_test.go`, and every `.rl` file using `.to_s()`/`.to_i()`/`.to_f()`

**Interfaces:**
- Produces: `to_string`, `to_integer`, `to_float` on every value; `help` with `wat` as an
  alias registered to the same `ObjectMethod` value.

- [ ] **Step 1: Find the blast radius**

```bash
grep -rn "to_s()\|to_i()\|to_f()\|\.wat()" --include="*.rl" --include="*.go" --include="*.md" --include="*.yml" . | grep -v "^./docs/superpowers" | wc -l
grep -rln "to_s()\|to_i()\|to_f()\|\.wat()" --include="*.rl" . | sort
```

Record the count; Step 6 checks it reaches zero.

- [ ] **Step 2: Update the Go tests first**

In `object/object_test.go`, rename every RocketLang snippet using `.to_s()`, `.to_i()`,
`.to_f()` and `.wat()` to the new names. Example of the shape:

```go
	tests := []inputTestCase{
		{`42.to_string()`, "42"},
		{`"42".to_integer()`, 42},
		{`"3.5".to_float()`, 3.5},
	}
```

- [ ] **Step 3: Run — expect failure**

Run: `go test ./object -run TestObject -timeout 120s`
Expected: FAIL, `undefined method .to_string()` and similar.

- [ ] **Step 4: Rename the registration keys**

In `object/object.go`, in the `objectMethods["*"]` map: `"to_s"` → `"to_string"`,
`"to_i"` → `"to_integer"`, `"to_f"` → `"to_float"`.

For `wat`, register the same method under both names. Extract the existing
`ObjectMethod` value into a variable so the two keys cannot drift:

```go
	// help prints the receiver's method list. It is printed rather than
	// returned because this listing exists to be read, and Inspect() would
	// escape the newlines onto one line.
	helpMethod := ObjectMethod{
		Layout: MethodLayout{
			ReturnPattern: Args(
				Arg(NIL_OBJ),
			),
		},
		method: func(o Object, _ []Object, _ Environment) Object {
			oms := objectMethods[o.Type()]
			names := make([]string, 0, len(oms))
			for name := range oms {
				names = append(names, name)
			}
			sort.Strings(names)

			fmt.Printf("%s supports the following methods:\n", o.Type())
			for _, name := range names {
				fmt.Printf("\t%s\n", oms[name].Layout.Usage(name))
			}

			return NIL
		},
	}

	objectMethods["*"]["help"] = helpMethod
	// wat is the only alias in the language: a deliberate easter egg, kept
	// because it predates help. It is not a precedent -- every other method
	// has exactly one name.
	objectMethods["*"]["wat"] = helpMethod
```

Place these two assignments after the map literal is assigned to `objectMethods["*"]`,
and delete the `"wat"` entry from inside the literal.

- [ ] **Step 5: Run — expect pass**

Run: `go test ./object -run TestObject -timeout 120s`
Expected: PASS

- [ ] **Step 6: Update the docs YAML and the `.rl` corpus**

In `docs/literals/object.yml`, rename the `methods:` keys `to_s` → `to_string`,
`to_i` → `to_integer`, `to_f` → `to_float`, `wat` → `help`, and add a `wat` entry:

```yaml
  wat:
    description: |
      An alias of `help`, kept as an easter egg. This is the only alias in
      RocketLang; every other method has exactly one name.
    input: '"test".wat()'
    output: ""
```

Then rewrite the corpus:

```bash
find tests fixtures examples docs/docs docs/literals docs/builtins -type f \
  \( -name '*.rl' -o -name '*.md' -o -name '*.yml' \) -print0 |
xargs -0 sed -i '' \
  -e 's/\.to_s()/.to_string()/g' \
  -e 's/\.to_i()/.to_integer()/g' \
  -e 's/\.to_f()/.to_float()/g'

grep -rn "to_s()\|to_i()\|to_f()" tests fixtures examples docs/docs docs/literals docs/builtins | wc -l
```

Expected: `0`. Note `.wat()` is deliberately *not* rewritten — it still works.

`docs/literals/object.yml` carries 20 lines mentioning these names, including
`description:` prose. Read the grep output and fix any prose the substitution mangled.

- [ ] **Step 7: Verify everything**

```bash
go test ./docs -run Documented
go test -run TestRocketlangCode . -timeout 120s
go test ./object -timeout 120s -run 'TestObject|TestString|TestInteger|TestFloat'
go run docs/generate.go && git diff --stat docs/docs
```

Expected: all PASS; `generate.go` produces a diff in `docs/docs/literals/*.md` reflecting
the new names.

- [ ] **Step 8: Commit**

```bash
git add -A
git commit -m "feat(object)!: name conversions after their type

to_s, to_i and to_f are abbreviations you have to learn; to_string,
to_integer and to_float follow one rule -- to_ plus the type name -- so a
conversion can be guessed rather than remembered. That rule already covers
to_json and now covers all of them.

wat becomes help, which is the name someone would try. wat is kept as an
alias, the only one in the language.

BREAKING CHANGE: to_s, to_i and to_f are renamed."
```

---

### Task 5: String renames

**Files:**
- Modify: `object/string.go`, `docs/literals/string.yml`, `object/string_test.go`, corpus

**Interfaces:**
- Produces: `index_of`, `last_index_of`, `codepoints`, `uppercase`/`!`, `lowercase`/`!`,
  `swap_case`/`!`, `trim`/`!`, `trim_start`/`!`, `trim_end`/`!`, `remove_last`/`!`,
  `trim_line_end`/`!`, `contains?`, `starts_with?`, `ends_with?`.

The full mapping, from the spec:

| Old | New |
|---|---|
| `find` | `index_of` |
| `ascii` | `codepoints` |
| `upcase` | `uppercase` |
| `downcase` | `lowercase` |
| `swapcase` | `swap_case` |
| `strip` | `trim` |
| `lstrip` | `trim_start` |
| `rstrip` | `trim_end` |
| `chop` | `remove_last` |
| `chomp` | `trim_line_end` |
| `include?` | `contains?` |
| `start_with?` | `starts_with?` |
| `end_with?` | `ends_with?` |

`upcase` … `chomp` are registered via `stringPair(name, …)`, which also generates
`name + "!"`. Renaming the string literal passed to `stringPair` renames both halves.
`include?`, `start_with?`, `end_with?` are registered via `stringPredicate`.

- [ ] **Step 1: Add the new `last_index_of` test plus the renamed cases**

In `object/string_test.go`:

```go
func TestStringIndexOf(t *testing.T) {
	tests := []inputTestCase{
		{`"hello".index_of("l")`, 2},
		{`"hello".last_index_of("l")`, 3},
		{`"hello".index_of("z")`, -1},
		{`"hello".last_index_of("z")`, -1},
	}

	testInput(t, tests)
}
```

Rename the existing cases for the other twelve methods in the same file.

- [ ] **Step 2: Run — expect failure**

Run: `go test ./object -run TestString -timeout 120s`
Expected: FAIL, `undefined method .index_of()`

- [ ] **Step 3: Rename the keys and add `last_index_of`**

Rename the map keys and `stringPair`/`stringPredicate` name arguments per the table.
Add, beside the renamed `index_of`:

```go
		"last_index_of": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(STRING_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				s := o.(*String)
				arg := args[0].(*String).Value

				return NewInteger(strings.LastIndex(s.Value, arg))
			},
		},
```

- [ ] **Step 4: Run — expect pass**

Run: `go test ./object -run TestString -timeout 120s`
Expected: PASS

- [ ] **Step 5: Update `docs/literals/string.yml` and the corpus**

Rename the same thirteen `methods:` keys, plus their `!` counterparts where documented,
and add a `last_index_of` entry with `description`, `input` and `output`. Then:

```bash
find tests fixtures examples docs/docs docs/literals docs/builtins -type f \
  \( -name '*.rl' -o -name '*.md' -o -name '*.yml' \) -print0 |
xargs -0 sed -i '' \
  -e 's/\.find(/.index_of(/g'      -e 's/\.ascii(/.codepoints(/g' \
  -e 's/\.upcase(/.uppercase(/g'   -e 's/\.upcase!(/.uppercase!(/g' \
  -e 's/\.downcase(/.lowercase(/g' -e 's/\.downcase!(/.lowercase!(/g' \
  -e 's/\.swapcase(/.swap_case(/g' -e 's/\.swapcase!(/.swap_case!(/g' \
  -e 's/\.lstrip(/.trim_start(/g'  -e 's/\.lstrip!(/.trim_start!(/g' \
  -e 's/\.rstrip(/.trim_end(/g'    -e 's/\.rstrip!(/.trim_end!(/g' \
  -e 's/\.strip(/.trim(/g'         -e 's/\.strip!(/.trim!(/g' \
  -e 's/\.chop(/.remove_last(/g'   -e 's/\.chop!(/.remove_last!(/g' \
  -e 's/\.chomp(/.trim_line_end(/g' -e 's/\.chomp!(/.trim_line_end!(/g' \
  -e 's/\.start_with?(/.starts_with?(/g' -e 's/\.end_with?(/.ends_with?(/g'
```

`lstrip`/`rstrip` are rewritten **before** `strip` on purpose — `s/\.strip(/` would
otherwise not match them, but a careless reordering that used `s/strip(/` would corrupt
them. Leave `include?` alone here; it is shared with Array and Hash and is renamed once
in Task 7.

- [ ] **Step 6: Verify**

```bash
go test ./docs -run Documented
go test ./object -run TestString -timeout 120s
go test -run TestRocketlangCode . -timeout 120s
```

Expected: all PASS.

- [ ] **Step 7: Commit**

```bash
git add -A
git commit -m "feat(string)!: spell string methods out

chop and chomp differ by one character and nobody remembers which trims the
newline. lstrip and rstrip say which end in a prefix nobody reads aloud.
upcase is a word only Ruby uses. trim, trim_start, trim_end, remove_last and
trim_line_end say what they do.

find is renamed index_of, matching Array, and gains the last_index_of it
always lacked.

BREAKING CHANGE: find, ascii, upcase, downcase, swapcase, strip, lstrip,
rstrip, chop, chomp, start_with? and end_with? are renamed."
```

---

### Task 6: Array renames

**Files:**
- Modify: `object/array.go`, `docs/literals/array.yml`, `object/array_test.go`, corpus

**Interfaces:**
- Produces: `index_of`, `last_index_of`, `append`, `remove_last`, `prepend`,
  `remove_first`, `remove`, `remove_at`, `unique`/`!`, `chunks`, `to_matrix`, `skip`,
  `skip_last`, `filter`/`!`. Deletes `take`.

| Old | New |
|---|---|
| `index` | `index_of` |
| `rindex` | `last_index_of` |
| `push` | `append` |
| `pop` | `remove_last` |
| `unshift` | `prepend` |
| `shift` | `remove_first` |
| `delete` | `remove` |
| `delete_at` | `remove_at` |
| `uniq` / `uniq!` | `unique` / `unique!` |
| `slices` | `chunks` |
| `to_m` | `to_matrix` |
| `select` / `select!` | `filter` / `filter!` |
| `drop` | `skip` |
| `take` | *deleted* |
| — | `skip_last` (new) |

`take(n)` and `first(n)` were verified to return identical results, which is why `take`
goes rather than being renamed.

- [ ] **Step 1: Write the tests, including the new `skip_last` and the removal of `take`**

```go
func TestArraySkip(t *testing.T) {
	tests := []inputTestCase{
		{`[1,2,3,4,5].skip(2)`, "[3, 4, 5]"},
		{`[1,2,3,4,5].skip_last(2)`, "[1, 2, 3]"},
		{`[1,2,3].skip(0)`, "[1, 2, 3]"},
		{`[1,2,3].skip(99)`, "[]"},
		{`[1,2,3].skip_last(99)`, "[]"},
	}

	testInput(t, tests)
}

// take was identical to first(n). Keeping both meant two names for one
// behaviour, so it is gone; this pins that it stays gone.
func TestArrayTakeIsGone(t *testing.T) {
	evaluated := testEval(`[1,2,3].take(2)`)

	if !object.IsError(evaluated) {
		t.Errorf("take should no longer exist, got %s", evaluated.Inspect())
	}
}
```

Rename the existing cases for the other thirteen methods.

- [ ] **Step 2: Run — expect failure**

Run: `go test ./object -run TestArray -timeout 120s`
Expected: FAIL, `undefined method .skip()`

- [ ] **Step 3: Rename the keys, delete `take`, add `skip_last`**

Apply the table to the registration keys and to the `arrayPair` name arguments. Delete
the `"take"` entry. Add `skip_last` beside the renamed `skip`, mirroring `skip`'s
existing bounds handling:

```go
		"skip_last": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				elements := o.(*Array).Elements
				count := args[0].(*Integer).Value

				if count < 0 {
					return NewErrorFormat("skip_last needs a count of zero or more, got %d", count)
				}

				keep := len(elements) - count
				if keep < 0 {
					keep = 0
				}

				return NewArray(copyElements(elements[:keep]))
			},
		},
```

`copyElements` already exists at `object/array.go:866`.

- [ ] **Step 4: Run — expect pass**

Run: `go test ./object -run TestArray -timeout 120s`
Expected: PASS

- [ ] **Step 5: Update `docs/literals/array.yml` and the corpus**

Rename the keys, delete the `take` entry, add `skip_last`. Then:

```bash
find tests fixtures examples docs/docs docs/literals docs/builtins -type f \
  \( -name '*.rl' -o -name '*.md' -o -name '*.yml' \) -print0 |
xargs -0 sed -i '' \
  -e 's/\.rindex(/.last_index_of(/g' -e 's/\.index(/.index_of(/g' \
  -e 's/\.push(/.append(/g'          -e 's/\.pop(/.remove_last(/g' \
  -e 's/\.unshift(/.prepend(/g'      -e 's/\.shift(/.remove_first(/g' \
  -e 's/\.delete_at(/.remove_at(/g'  -e 's/\.delete(/.remove(/g' \
  -e 's/\.uniq!(/.unique!(/g'        -e 's/\.uniq(/.unique(/g' \
  -e 's/\.slices(/.chunks(/g'        -e 's/\.to_m(/.to_matrix(/g' \
  -e 's/\.select!(/.filter!(/g'      -e 's/\.select(/.filter(/g' \
  -e 's/\.drop(/.skip(/g'            -e 's/\.take(/.first(/g'
```

Order matters twice: `rindex` before `index`, and `delete_at` before `delete`. `take(n)`
is rewritten to `first(n)`, which is what it always did.

- [ ] **Step 6: Verify**

```bash
go test ./docs -run Documented
go test ./object -run 'TestArray|TestHash' -timeout 120s
go test -run TestRocketlangCode . -timeout 120s
```

Expected: all PASS.

- [ ] **Step 7: Commit**

```bash
git add -A
git commit -m "feat(array)!: name array operations after what they do

shift and unshift are the classic pair nobody can order correctly, and push
and pop only make sense once you already picture a stack. append, prepend,
remove_first and remove_last are unambiguous and pair visibly.

take(n) is deleted rather than renamed: it returned exactly what first(n)
returns, so it was a second name for one behaviour.

select becomes filter, which is what every language outside Ruby calls it.

BREAKING CHANGE: index, rindex, push, pop, shift, unshift, delete, delete_at,
uniq, slices, to_m, select and drop are renamed; take is removed in favour of
first."
```

---

### Task 7: Hash renames, and the shared `include?` → `contains?`

`include?` exists on String, Array and Hash. String and Array become `contains?`; Hash
becomes `has_key?`, because it checks keys and `include?` hid that.

**Files:**
- Modify: `object/hash.go`, `object/array.go`, `object/string.go`,
  `docs/literals/hash.yml`, `array.yml`, `string.yml`, tests, corpus

**Interfaces:**
- Produces: `Hash#has_key?`, `Hash#remove`, `Hash#filter`/`!`, `Array#contains?`,
  `String#contains?`.

- [ ] **Step 1: Write the tests**

```go
func TestContains(t *testing.T) {
	tests := []inputTestCase{
		{`"hello".contains?("ell")`, true},
		{`"hello".contains?("z")`, false},
		{`[1,2,3].contains?(2)`, true},
		{`[1,2,3].contains?(9)`, false},
		{`{"a": 1}.has_key?("a")`, true},
		{`{"a": 1}.has_key?("b")`, false},
	}

	testInput(t, tests)
}
```

Rename existing `Hash#delete` and `Hash#select` cases to `remove` and `filter`.

- [ ] **Step 2: Run — expect failure**

Run: `go test ./object -run 'TestContains|TestHash' -timeout 120s`
Expected: FAIL, `undefined method .contains?()`

- [ ] **Step 3: Rename the keys**

- `object/string.go`: the `stringPredicate("include?", …)` name argument → `"contains?"`
- `object/array.go`: `"include?"` → `"contains?"`
- `object/hash.go`: `"include?"` → `"has_key?"`, `"delete"` → `"remove"`,
  `"select"`/`"select!"` → `"filter"`/`"filter!"` (via `hashCallbackPair` where used)

- [ ] **Step 4: Run — expect pass**

Run: `go test ./object -run 'TestContains|TestHash|TestArray|TestString' -timeout 120s`
Expected: PASS

- [ ] **Step 5: Document the `get`/`fetch` rule**

The spec requires the difference to be stated. In `docs/literals/hash.yml`, replace the
`get` and `fetch` descriptions with:

```yaml
  get:
    description: |
      Returns the value for `key`, or `nil` when the key is absent. Pass a
      second argument to get that instead of `nil`. Use `fetch` when a missing
      key is a mistake rather than an expected case.
    input: '{"a": 1}.get("b", 0)'
    output: "0"
  fetch:
    description: |
      Returns the value for `key`, and raises when the key is absent. Use
      `get` when a missing key is an expected case.
    input: '{"a": 1}.fetch("a")'
    output: "1"
```

Verify the claim before writing it — if `fetch` does not currently raise on a missing
key, make it do so as part of this task and add a test:

```bash
/tmp/rl -e 'puts({"a": 1}.fetch("zzz"))'
```

- [ ] **Step 6: Update the remaining YAML and the corpus**

Rename the keys in `hash.yml`, `array.yml`, `string.yml`. Then, because `include?` maps
to two different new names depending on receiver type, rewrite by hand rather than with
a blanket `sed`:

```bash
grep -rn "include?" --include="*.rl" --include="*.md" tests fixtures examples docs/docs
```

Inspect each hit and choose `contains?` (String, Array) or `has_key?` (Hash). Then:

```bash
find tests fixtures examples docs/docs docs/literals docs/builtins -type f \
  \( -name '*.rl' -o -name '*.md' -o -name '*.yml' \) -print0 |
xargs -0 sed -i '' -e 's/\.select!(/.filter!(/g' -e 's/\.select(/.filter(/g'
```

- [ ] **Step 7: Verify**

```bash
go test ./docs -run Documented
go test ./object -timeout 120s -run 'TestContains|TestHash|TestArray|TestString'
go test -run TestRocketlangCode . -timeout 120s
grep -rn "include?" --include="*.rl" --include="*.md" tests fixtures examples docs/docs | wc -l
```

Expected: tests PASS, final count `0`.

- [ ] **Step 8: Commit**

```bash
git add -A
git commit -m "feat(hash)!: say which half of a hash include? was checking

Hash#include? tested keys, which the name never said -- a reader had to guess
whether {\"a\": 1}.include?(1) was true. It is has_key? now. String and Array
get contains?, which is the ordinary word for the question.

The difference between get and fetch is finally written down: get answers nil
for a missing key, fetch raises.

BREAKING CHANGE: include? is renamed contains? on String and Array and
has_key? on Hash; Hash#delete becomes remove and Hash#select becomes filter."
```

---

### Task 8: Integer and Float renames, and the `Math` consolidation

Rule 4: a number's own operations are methods on the number; `Math` keeps only constants
and randomness.

**Verified duplicates** (both a `Math` function and a number method exist today):
`abs`, `ceil`, `floor`, `round`, `pow` — the `Math` copies are deleted outright.

**Verified as `Math`-only**, so they move to `Integer` and `Float`: `sqrt`, `exp`, `log`,
`log2`, `log10`, `sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `copysign`, `remainder`.

`Math.max`/`Math.min` are deleted — `[a, b].max` already exists.

**Files:**
- Modify: `object/integer.go`, `object/float.go`, `stdlib/math.go`,
  `docs/literals/integer.yml`, `float.yml`, `docs/builtins/math.yml`, tests, corpus

**Interfaces:**
- Produces: `Integer#successor`, `#predecessor`, `#to_character`, and the thirteen
  absorbed math methods on both `Integer` and `Float`. Deletes `Integer#base`,
  `Math.abs/ceil/floor/round/pow/max/min`; renames `Math.rand` → `Math.random`.

| Old | New |
|---|---|
| `succ` | `successor` |
| `pred` | `predecessor` |
| `chr` | `to_character` |
| `base` | *deleted* — `to_base(n)` remains |
| `Math.rand` | `Math.random` |

- [ ] **Step 1: Write the tests**

```go
func TestNumberMathMethods(t *testing.T) {
	tests := []inputTestCase{
		{`16.sqrt()`, 4.0},
		{`16.0.sqrt()`, 4.0},
		{`1.exp()`, 2.718281828459045},
		{`8.log2()`, 3.0},
		{`100.log10()`, 2.0},
		{`0.sin()`, 0.0},
		{`0.cos()`, 1.0},
		{`5.successor()`, 6},
		{`5.predecessor()`, 4},
		{`65.to_character()`, "A"},
	}

	testInput(t, tests)
}

// Math keeps constants and randomness only. Everything that operates on a
// number is a method on the number, so there is one place to look.
func TestMathKeepsOnlyConstantsAndRandom(t *testing.T) {
	for _, gone := range []string{"abs", "ceil", "floor", "round", "pow", "sqrt", "max", "min", "rand"} {
		evaluated := testEval("Math." + gone + "(1)")

		if !object.IsError(evaluated) {
			t.Errorf("Math.%s should be gone, got %s", gone, evaluated.Inspect())
		}
	}
}
```

Check each expected value against the current `Math` equivalent before pinning it —
`/tmp/rl -e 'puts(Math.sqrt(16.0))'` — so the test encodes real behaviour, not a guess.

- [ ] **Step 2: Run — expect failure**

Run: `go test ./object -run 'TestNumber|TestMath' -timeout 120s`
Expected: FAIL, `undefined method .sqrt()`

- [ ] **Step 3: Rename on the number types and absorb the math functions**

In `object/integer.go` and `object/float.go`, rename `succ`/`pred`/`chr` and delete
`base`. Add the thirteen absorbed methods. They all take the receiver as a float and
return a float, so register them from one table rather than writing thirteen blocks:

```go
// numberMath registers a method that treats the receiver as a float and
// returns a float. Written as a table because thirteen near-identical blocks
// is thirteen chances for one of them to call the wrong function -- which is
// how Math.log2 and Math.log10 would differ only by a character.
func numberMath(objectType ObjectType, toFloat func(Object) float64) {
	unary := map[string]func(float64) float64{
		"sqrt":  math.Sqrt,
		"exp":   math.Exp,
		"log":   math.Log,
		"log2":  math.Log2,
		"log10": math.Log10,
		"sin":   math.Sin,
		"cos":   math.Cos,
		"tan":   math.Tan,
		"asin":  math.Asin,
		"acos":  math.Acos,
		"atan":  math.Atan,
	}

	for name, fn := range unary {
		apply := fn
		objectMethods[objectType][name] = ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewFloat(apply(toFloat(o)))
			},
		}
	}

	binary := map[string]func(float64, float64) float64{
		"copysign":  math.Copysign,
		"remainder": math.Remainder,
	}

	for name, fn := range binary {
		apply := fn
		objectMethods[objectType][name] = ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(NUMERIC),
				),
				ReturnPattern: Args(
					Arg(FLOAT_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				other, ok := args[0].(Floatable)
				if !ok {
					return NewErrorFormat("expected a number, got %s", args[0].Type())
				}

				value, ok := other.ToFloatObj().(*Float)
				if !ok {
					return NewErrorFormat("expected a number, got %s", args[0].Type())
				}

				return NewFloat(apply(toFloat(o), value.Value))
			},
		}
	}
}
```

Call it at the end of each `init()`:

```go
	// integer.go
	numberMath(INTEGER_OBJ, func(o Object) float64 { return float64(o.(*Integer).Value) })

	// float.go
	numberMath(FLOAT_OBJ, func(o Object) float64 { return o.(*Float).Value })
```

Put `numberMath` in `object/number.go` (a new file — `object/number_test.go` already
exists with no partner) and import `math` there.

- [ ] **Step 4: Strip `stdlib/math.go` down**

Delete `abs`, `acos`, `asin`, `atan`, `ceil`, `copysign`, `cos`, `exp`, `floor`, `log`,
`log10`, `log2`, `max`, `min`, `pow`, `remainder`, `round`, `sin`, `sqrt`, `tan` from
`mathFunctions`. Rename `rand` → `random`. The properties block is untouched.

- [ ] **Step 5: Run — expect pass**

Run: `go test ./object ./stdlib -run 'TestNumber|TestMath' -timeout 120s`
Expected: PASS

- [ ] **Step 6: Update the YAML and the corpus**

In `docs/builtins/math.yml`, delete the twenty removed function entries and rename
`rand` → `random`. In `docs/literals/integer.yml` and `float.yml`, rename `succ`, `pred`,
`chr`, delete `base`, and add the thirteen absorbed methods to both.

```bash
grep -rn "Math\." --include="*.rl" --include="*.md" tests fixtures examples docs/docs
```

Rewrite each by hand — `Math.sqrt(x)` becomes `x.sqrt()`, so this is a restructure rather
than a substitution. Then:

```bash
find tests fixtures examples docs/docs docs/literals docs/builtins -type f \
  \( -name '*.rl' -o -name '*.md' -o -name '*.yml' \) -print0 |
xargs -0 sed -i '' -e 's/\.succ(/.successor(/g' -e 's/\.pred(/.predecessor(/g' -e 's/\.chr(/.to_character(/g'
```

- [ ] **Step 7: Verify**

```bash
go test ./docs -run Documented
go test ./object ./stdlib -timeout 120s
go test -run TestRocketlangCode . -timeout 120s
```

Expected: all PASS.

- [ ] **Step 8: Commit**

```bash
git add -A
git commit -m "feat(number)!: give numbers their own arithmetic, and shrink Math

Math.abs, .ceil, .floor, .round and .pow each duplicated a method the number
already had, so which one you called was a coin toss. The duplicates are gone
and the remaining Math functions have moved onto Integer and Float, leaving
Math with the two things that are not operations on a receiver: its constants
and randomness.

Math.max and Math.min are gone too -- [a, b].max already existed.

BREAKING CHANGE: most Math functions are now number methods; Math.rand is
Math.random; succ, pred and chr become successor, predecessor and
to_character; Integer#base is removed in favour of to_base."
```

---

### Task 9: `Error#message`, Matrix renames

**Files:**
- Modify: `object/error.go`, `object/matrix.go`, `docs/literals/error.yml`, `matrix.yml`,
  tests, corpus

**Interfaces:**
- Produces: `Error#message`, `Matrix#to_array`. Deletes `Matrix#t` and the duplicate
  `rows`/`cols` properties.

- [ ] **Step 1: Write the tests**

```go
func TestErrorMessage(t *testing.T) {
	evaluated := testEval(`begin
  nil.nope()
rescue e
  e.message()
end`)

	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("expected a string, got %T (%s)", evaluated, evaluated.Inspect())
	}
	if !strings.Contains(str.Value, "undefined method") {
		t.Errorf("unexpected message: %q", str.Value)
	}
}

// t was an abbreviation of transpose, which is short already.
func TestMatrixTIsGone(t *testing.T) {
	evaluated := testEval(`[[1,2],[3,4]].to_matrix().t()`)

	if !object.IsError(evaluated) {
		t.Errorf("Matrix#t should be gone, got %s", evaluated.Inspect())
	}
}
```

- [ ] **Step 2: Run — expect failure**

Run: `go test ./object -run 'TestError|TestMatrix' -timeout 120s`
Expected: FAIL, `undefined method .message()`

- [ ] **Step 3: Rename**

- `object/error.go`: `"msg"` → `"message"`
- `object/matrix.go`: `"to_a"` → `"to_array"`; delete `"t"`; delete the `rows`/`cols`
  **property** entries (around `object/matrix.go:405`), keeping the methods

- [ ] **Step 4: Run — expect pass**

Run: `go test ./object -run 'TestError|TestMatrix' -timeout 120s`
Expected: PASS

- [ ] **Step 5: Update the YAML and corpus**

```bash
find tests fixtures examples docs/docs docs/literals docs/builtins -type f \
  \( -name '*.rl' -o -name '*.md' -o -name '*.yml' \) -print0 |
xargs -0 sed -i '' -e 's/\.msg(/.message(/g' -e 's/\.to_a(/.to_array(/g'
```

Rename the same keys in `docs/literals/error.yml` and `matrix.yml`, and delete the `t`
entry.

- [ ] **Step 6: Verify and commit**

```bash
go test ./docs -run Documented
go test ./object -timeout 120s
go test -run TestRocketlangCode . -timeout 120s
git add -A
git commit -m "feat(error)!: msg becomes message

msg was the only abbreviated method name left in the language. Matrix#to_a
becomes to_array to match every other conversion, and Matrix#t goes -- it
abbreviated transpose, which is not long.

BREAKING CHANGE: Error#msg is Error#message; Matrix#to_a is Matrix#to_array;
Matrix#t is removed."
```

---

### Task 10: `puts` → `print`, and delete `OS.raise`

**Files:**
- Modify: `stdlib/std.go`, `stdlib/puts.go` → `stdlib/print.go`, `stdlib/os.go`,
  `docs/builtins/os.yml`, `stdlib/os_test.go`, every `.rl` file, every `.md` file

**Interfaces:**
- Produces: global `print`. Removes `OS.raise`. Global `raise` is unchanged.

This is the largest corpus change in the plan — `puts` appears in nearly every `.rl`
file — so it is deliberately last, when everything else is already green.

- [ ] **Step 1: Write the tests**

`stdlib/os_test.go` is `package stdlib` (an internal test) and has **no** `testEval`
helper — the evaluator is not importable from here without a cycle. Assert against the
registration maps directly, which is simpler anyway:

```go
// OS.raise duplicated the global raise. Two spellings of one behaviour is
// exactly what v1.0 removes.
func TestOSRaiseIsGone(t *testing.T) {
	if _, exists := Modules["OS"].Functions["raise"]; exists {
		t.Error("OS.raise should be gone; the global raise already does this")
	}
}

// print replaces puts. Checked here rather than through the evaluator for the
// same reason: this package registers the name, so this package can see it.
func TestPrintIsRegisteredAndPutsIsNot(t *testing.T) {
	if _, exists := Functions["print"]; !exists {
		t.Error("print should be registered")
	}
	if _, exists := Functions["puts"]; exists {
		t.Error("puts should be gone")
	}
}
```

Also rename the existing `OS.raise` cases in this file.

- [ ] **Step 2: Run — expect failure**

Run: `go test ./stdlib -run TestOS -timeout 120s`
Expected: FAIL — `OS.raise` still works.

- [ ] **Step 3: Rename `puts` and remove `OS.raise`**

```bash
git mv stdlib/puts.go stdlib/print.go
```

In `stdlib/print.go`, rename `putsFunction` → `printFunction`. In `stdlib/std.go`, change
the registration name and the referenced function:

```go
	RegisterFunction("print", object.MethodLayout{ArgPattern: object.Args(object.Arg(object.ANY))}, printFunction)
```

In `stdlib/os.go`, delete the `raise` entry from `osFunctions`.

- [ ] **Step 4: Run — expect pass**

Run: `go test ./stdlib -timeout 120s`
Expected: PASS

- [ ] **Step 5: Rewrite the corpus**

```bash
find tests fixtures examples docs/docs wasm -type f \
  \( -name '*.rl' -o -name '*.md' -o -name '*.html' -o -name '*.js' \) -print0 |
xargs -0 sed -i '' -e 's/\bputs(/print(/g'

find docs/literals docs/builtins -name '*.yml' -print0 | xargs -0 sed -i '' -e 's/\bputs(/print(/g'

grep -rn "\bputs(" --include="*.rl" --include="*.md" --include="*.yml" \
  tests fixtures examples docs/docs docs/literals docs/builtins | wc -l
```

Expected: `0`. Note this also rewrites the `input:` examples in the YAML, which the
generated docs display.

In `docs/builtins/os.yml`, delete the `raise` entry.

- [ ] **Step 6: Verify the whole tree**

```bash
go build ./... && go vet ./...
go test ./docs -run Documented
go test ./object ./stdlib ./lexer ./token ./parser ./evaluator ./planet ./utilities ./repl -timeout 120s
go test -run TestRocketlangCode . -timeout 120s
```

Expected: all PASS.

- [ ] **Step 7: Commit**

```bash
git add -A
git commit -m "feat(stdlib)!: puts becomes print, and OS.raise is deleted

puts is a name inherited from C by way of Ruby, and it stands for \"put
string\" -- which it has not meant for a long time, since it prints anything.
print is the word.

OS.raise did exactly what the global raise does. A language where the same
behaviour has two spellings makes the reader wonder what the difference is.

BREAKING CHANGE: puts is renamed print; OS.raise is removed, use raise."
```

---

### Task 11: New stdlib functions — `IO.write`, `IO.read_line`, `Time.now`

The spec's globals table requires three additions. `IO.write` is the counterpart to
`print`: `print` adds a newline, `IO.write` does not, and it lives beside `file.write`
where the other precise output tool is. `Time.now` fills an obvious hole — `Time.unix`
gives a timestamp but nothing gives the current time.

**Files:**
- Modify: `stdlib/io.go`, `stdlib/time.go`, `docs/builtins/io.yml`, `docs/builtins/time.yml`
- Create: tests in `stdlib/io_test.go` (new file), `stdlib/time_test.go` (new file)

**Interfaces:**
- Produces: `IO.write(ANY) -> NIL`, `IO.read_line() -> STRING`, `Time.now() -> INTEGER`.

- [ ] **Step 1: Write the tests**

`stdlib` internal tests cannot reach the evaluator, so call the registered functions
directly, the way `stdlib/os_test.go` does.

```go
package stdlib

import (
	"io"
	"os"
	"testing"

	"github.com/flipez/rocket-lang/object"
)

// TestIOWriteAddsNoNewline is the whole point of IO.write existing beside
// print, so it is what gets pinned.
func TestIOWriteAddsNoNewline(t *testing.T) {
	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = write

	Modules["IO"].Functions["write"].Call(
		[]object.Object{object.NewString("ab")},
		*object.NewEnvironment(),
	)

	write.Close()
	os.Stdout = original

	// io.ReadAll, not strings.Builder.ReadFrom -- strings.Builder is a Writer
	// and has no ReadFrom method.
	out, err := io.ReadAll(read)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "ab" {
		t.Errorf("IO.write should emit exactly %q, got %q", "ab", string(out))
	}
}

func TestIOReadLineIsRegistered(t *testing.T) {
	if _, exists := Modules["IO"].Functions["read_line"]; !exists {
		t.Error("IO.read_line should be registered")
	}
}
```

```go
package stdlib

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

// TestTimeNowIsPlausible avoids pinning a clock value: it checks the result is
// an integer in a range that cannot be produced by a stub returning zero.
func TestTimeNowIsPlausible(t *testing.T) {
	result := Modules["Time"].Functions["now"].Call(nil, *object.NewEnvironment())

	seconds, ok := result.(*object.Integer)
	if !ok {
		t.Fatalf("Time.now should return an Integer, got %s", result.Type())
	}

	// 1 January 2020. Any real clock is past it; a stub is not.
	if seconds.Value < 1577836800 {
		t.Errorf("Time.now returned %d, which is not a current unix timestamp", seconds.Value)
	}
}
```

- [ ] **Step 2: Run — expect failure**

Run: `go test ./stdlib -run 'TestIO|TestTimeNow' -timeout 120s`
Expected: FAIL — nil map entry panic or "should be registered".

- [ ] **Step 3: Implement**

In `stdlib/io.go`, following the shape of the existing `open`:

```go
	ioFunctions["write"] = object.NewBuiltinFunction("write",
		object.MethodLayout{
			ArgPattern:    object.Args(object.OverloadArg(object.ANY)),
			ReturnPattern: object.Args(object.Arg(object.NIL_OBJ)),
		},
		func(_ object.Environment, args ...object.Object) object.Object {
			for _, arg := range args {
				// A string writes its value, not its quoted form -- the same
				// choice print makes, for the same reason.
				if str, ok := arg.(*object.String); ok {
					fmt.Print(str.Value)

					continue
				}

				fmt.Print(arg.Inspect())
			}

			return object.NIL
		})

	ioFunctions["read_line"] = object.NewBuiltinFunction("read_line",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.STRING_OBJ, object.NIL_OBJ)),
		},
		func(_ object.Environment, _ ...object.Object) object.Object {
			line, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil && line == "" {
				// End of input is not a failure: it is how a piped program
				// ends. nil says "nothing more", which the caller can test.
				return object.NIL
			}

			return object.NewString(strings.TrimRight(line, "\r\n"))
		})
```

In `stdlib/time.go`:

```go
	timeFunctions["now"] = object.NewBuiltinFunction("now",
		object.MethodLayout{
			ReturnPattern: object.Args(object.Arg(object.INTEGER_OBJ)),
		},
		func(_ object.Environment, _ ...object.Object) object.Object {
			return object.NewInteger(int(time.Now().Unix()))
		})
```

Add `bufio`, `fmt`, `os` and `strings` imports to `io.go` as needed.

- [ ] **Step 4: Run — expect pass**

Run: `go test ./stdlib -timeout 120s`
Expected: PASS

- [ ] **Step 5: Document all three**

Add entries to `docs/builtins/io.yml` under `functions:` for `write` and `read_line`, and
to `docs/builtins/time.yml` for `now`, each with `description`, `input` and `output`. The
guard test from Task 1 fails if any is missing.

- [ ] **Step 6: Verify and commit**

```bash
go test ./docs -run Documented
go test ./stdlib -timeout 120s
git add -A
git commit -m "feat(stdlib): add IO.write, IO.read_line and Time.now

print always ends a line, which is right for almost every use and wrong for
the rest -- a progress dot, a prompt. IO.write is the version that does not,
sitting beside file.write where the other precise output tool already lives.

read_line gives a program a way to take input at all, which it had no way to
do. Time.now fills the obvious hole beside Time.unix."
```

---

### Task 12: Regenerate the reference docs and fix the hand-written pages

**Files:**
- Regenerate: `docs/docs/literals/*.md`, `docs/docs/builtins/*.md`
- Modify: `docs/docs/**/*.md` (hand-written), `README.md`

- [ ] **Step 1: Regenerate**

```bash
go run docs/generate.go
git diff --stat docs/docs
```

Expected: every file under `docs/docs/literals` and `docs/docs/builtins` changes.

- [ ] **Step 2: Find old names surviving in hand-written prose**

```bash
grep -rn "puts(\|to_s()\|to_i()\|to_f()\|\.msg()\|include?\|upcase\|downcase\|lstrip\|rstrip\|chomp\|\.chop\|succ()\|pred()\|\.push(\|\.pop(\|\.shift(\|unshift\|uniq\|\.select(\|Math\.\(sqrt\|abs\|pow\|floor\|ceil\|round\|max\|min\|rand\)" \
  docs/docs README.md | grep -v "docs/docs/literals\|docs/docs/builtins"
```

Fix each by hand — these are prose and worked examples, not mechanical substitutions.
`docs/docs/language/methods.md` and `docs/docs/specification/builtins.md` are the most
likely to need real rewriting.

- [ ] **Step 3: Build the site, which fails on a broken link**

```bash
cd docs && yarn install --frozen-lockfile && yarn build; cd ..
```

Expected: build succeeds. `onBrokenLinks: 'throw'` and `onBrokenMarkdownLinks: 'throw'`
are set, so a renamed anchor becomes a build failure rather than a dead link.

- [ ] **Step 4: Full verification**

```bash
go build ./... && go vet ./...
go test ./docs -run Documented
go test ./object ./stdlib ./lexer ./token ./parser ./evaluator ./planet ./utilities ./repl -timeout 120s
go test -run TestRocketlangCode . -timeout 120s
git status --porcelain
```

Expected: all PASS, working tree clean after commit.

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "docs: regenerate the reference for the v1.0 names

The literal and builtin pages are generated from the registered method
layouts, so they follow the rename automatically. The hand-written pages and
the README quoted the old names in prose and worked examples and are updated
by hand."
```

---

## Notes for the executor

- **Do not batch the renames.** Each task is one type, so a mistake is bisectable and a
  reviewer can reject one rename without rejecting the rest.
- **`sed` ordering is load-bearing** in Tasks 5, 6 and 8 (`lstrip` before `strip`,
  `rindex` before `index`, `delete_at` before `delete`). Those orderings are written into
  the commands; do not tidy them.
- **The corpus is the oracle.** `tests/lang/*.expected` pins byte-exact output for 39
  programs. If a rename changes output, it changed behaviour — investigate rather than
  regenerating the expectation.
- **`go test ./object` may hang** on the HTTP listener at `object/http.go:65`; always
  pass `-timeout 120s`. This is pre-existing and out of scope.
- **`docs/superpowers/` is excluded** from every `grep`/`sed` in this plan. Docusaurus
  does not scan it (its docs path is `docs/docs`), and the spec quotes old names on
  purpose.
