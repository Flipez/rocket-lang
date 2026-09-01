# RocketLang v1.0 — syntax and naming redesign

**Status:** approved design, not yet implemented
**Date:** 2026-08-31

## Why

RocketLang grew out of MonkeyLang and drifted toward Ruby feature by feature. Each
addition was reasonable on its own, but nothing ever revisited the whole, so the
language now contradicts itself in ways a newcomer meets on the first day:

- Two unrelated range syntaxes. `0 -> 5 ^ 2` in a loop header, `s[1:-2]` in a slice.
- `String#find` and `Array#index` do the same job under different names.
- `String#count("ab")` counts occurrences; `Array#count` counts elements.
- `raise` exists as a global *and* as `OS.raise`. `Math.pow` exists *and* `Integer#pow`.
- `Hash#get` and `Hash#fetch` are two lookups with no stated rule for choosing.
- `?` is both a predicate suffix and the ternary operator; `!` is both a mutation
  suffix and logical not.
- Method names mix registers: `msg`, `succ`, `chomp`, `lstrip`, `unshift`, `to_m`
  alongside `capitalize`, `transpose`, `bit_length`.
- No string interpolation, so every message is a `+` chain or a printf template.
- `Array#take(n)` and `Array#first(n)` are behaviourally identical. `Math.abs`,
  `.ceil`, `.floor`, `.round` and `.pow` each duplicate a number method.
- Scoping is arbitrary: `while` and `for` create a child environment, `if` and `begin`
  do not. Nobody decided that; it is where `NewEnclosedEnvironment` happened to be
  called. It is undocumented, so the first time anyone learns it is when a variable
  is missing.

The goal is a language that is easy to use and consistent with itself, with as few
special characters as the job allows. Where those two pull apart, this document
records which won and why.

## Non-goals

- Not becoming Ruby. Familiar shapes are kept where they are genuinely good, not
  because Ruby has them.
- Not a semantic redesign. The object model, evaluator, module system and planet
  package manager keep working as they do. This is surface: syntax and names.
- Not backward compatible. v0.23 code will not parse. This is a deliberate one-time
  break, tagged v1.0.0.

## The language

### Four rules that generate the rest

1. **A dot calls.** Parentheses only ever hold arguments. `list.size` calls;
   `list.size()` is an error; a bare `print` is a value.
2. **Conversions are `to_` + the type name.** `to_string`, `to_integer`, `to_float`,
   `to_array`, `to_hash`, `to_matrix`, `to_json`. No exceptions, so nothing has to be
   memorised.
3. **`?` means "returns a boolean". `!` means "mutates the receiver and returns it".**
   A `!` method exists *only* as the counterpart of a same-named non-mutating one.
4. **A global exists only if a hello-world-sized program needs it.** That is `print`
   and `raise`. Everything else is namespaced.

### A program

```rocketlang
# comments start with a hash

import "stats" as Stats only mean

names = ["ada", "grace", "alan"]

for name in names
  print("hello, #{name.capitalize}")
end

long  = names.filter(n -> n.size > 3)
loud  = names.map(n -> n.uppercase)
total = (1..100).sum

names.each do name
  if name.starts_with?("a")
    next
  end
  print(name)
end

for i in 0..<names.size
  print("#{i}: #{names[i]}")
end

try
  n = "abc".to_integer
rescue e
  print(e.message)
end

sorted = names.sort     # a new array
names.sort!             # mutates, returns the receiver

print(names[0..1])      # first two
print(names[-1..])      # last one
print(names.size)
```

### Calls

A dot always calls. Parentheses appear only when arguments are passed.

| Written | Meaning |
|---|---|
| `print` | look up the name, yield the bound value; nothing is called |
| `print(x)` | look up the name, then call the result with `x` |
| `list.size` | send `size` to `list` |
| `list.contains?(3)` | send `contains?` to `list` with an argument |
| `list.size()` | **error** — empty parentheses are never written |

Method names are not values: `upcase` alone is not a name in scope. To pass behaviour,
use a lambda (`x -> x.uppercase`) or a variable holding a function (`numbers.each(print)`).

This is two triggers for a call — the dot, and the parentheses — rather than one. That
was chosen knowingly: the single-rule alternative (`x.empty?()` everywhere) is more
uniform but noisier on every line, and ease of use won.

### Comments

`#` to end of line, replacing `//`. Inside a double-quoted string `#{` begins an
interpolation; the lexer already knows whether it is inside a string, so there is no
ambiguity.

### Strings

Double-quoted strings process escapes and interpolate. Single-quoted strings are
completely raw. This extends a distinction the lexer already makes
(`lexer.go:250` vs `lexer.go:297`).

```rocketlang
name = "ada"
print("hello, #{name}")            # hello, ada
print("#{2 + 3} items")            # 5 items
print("#{names.size} names")       # any expression is allowed
print('literal #{name} here')      # literal #{name} here
print("a set: {1, 2}")             # braces are ordinary characters
```

`#{}` was chosen over `{}` because every language in the brace family (Python, C#, F#,
Nim, Rust) gates interpolation behind a prefix or a macro, precisely because a lone `{`
is common in real text. `#{` essentially never occurs by accident, so it can be
always-on without a prefix letter to remember and without a doubling rule to teach.

`String#format` survives for padding and precision, which interpolation does not cover.

### Ranges and slices

One range concept, used in loop headers, in slices, and as a standalone value.

```rocketlang
0..5        # 0 1 2 3 4 5   inclusive
0..<5       # 0 1 2 3 4     exclusive
(0..100).by(2)

evens = 0..100          # a Range is a value
evens.sum
evens.to_array

s = "abcdef"
s[0..2]     # "abc"
s[0..<3]    # "abc"
s[1..]      # "bcdef"
s[-2..]     # "ef"
```

`Range` is a new object type. `by(n)` returns a new `Range` rather than mutating, in
keeping with rule 3.

The open-ended forms `a..` and `..b` are **only valid inside an index expression**,
where the collection supplies the missing bound. `x = 1..` as a standalone value is a
parse error; there is no infinite range.

`..<` rather than `...` for the exclusive form: `0..5` and `0...5` differ by a single
dot, which is the classic off-by-one trap, while `..<` reads as "up to, less than" and
cannot be misread.

This retires `->`, `=>`, `^` and colon slicing. Retiring `->` is what frees it for
lambdas.

### Functions and lambdas

```rocketlang
def add(a, b)
  return a + b
end

double = x -> x * 2                 # single expression
sum    = (a, b) -> a + b            # several parameters
anon   = def (x)                    # multi-statement, unchanged
  x * 2
end

numbers.map(x -> x * 2)
numbers.filter(n -> n > 3)
numbers.reduce(0, (sum, x) -> sum + x)

numbers.each do n                   # trailing block
  if n.even?
    next
  end
  print(n)
end
```

The trailing `do params … end` is **pure sugar** for passing a function as the last
argument: `f do x … end` is exactly `f(def (x) … end)`. It is deliberately not a Ruby
block — there is no `yield`, no `block_given?`, and no second concept beside arguments.

### Types are values

`String`, `Integer`, `Float`, `Boolean`, `Nil`, `Array`, `Hash`, `Matrix`, `File`,
`Module`, `HttpServer`, `Error`, `Function`, `Range`, `Type` and the type groups `Any`,
`Numeric`, `Comparable`, `Hashable`, `Stringable`, `Integerable`, `Callable` are
protected global bindings holding `Type` objects.

Two names in the current registry need attention before this can be generated:

- **`HTTP` collides.** `HTTP` is both a stdlib module and an object type today. Once
  types are global bindings, one name cannot be both. The object type is renamed
  **`HttpServer`** — it has `listen` and `handle`, so that is what it is — leaving the
  `HTTP` module name free. This also makes room for the `HttpClient` sketched in
  `DESIGN_HTTP_CLIENT.md`.
- **Internal types are excluded.** `knownObjectTypes` contains `RETURN_VALUE`,
  `BREAK_VALUE` and `NEXT_VALUE`, which are evaluator control-flow objects and must not
  become user-visible globals. The registry gains a user-facing flag; only user-facing
  entries produce a binding. `BuiltinFunction`, `BuiltinModule` and `BuiltinProperty`
  are user-facing (a value can genuinely be one) and do produce bindings.

```rocketlang
x.is_a?(String)
x.is_a?(Numeric)        # groups work identically to concrete types
x.type == String
x.type.name             # "String"
print(42.type)          # Integer

String = 5              # error: cannot rebind a builtin type
```

`.type` returns a `Type`, not a string, so `x.type == "String"` is no longer true.

Two properties fall out. The bindings are generated from `knownObjectTypes` and
`typeGroups` in `object/object.go`, which `docs/generate.go` already reads — so the
names in `is_a?`, in the generated docs, and in the Go-side `Arg(NUMERIC)` checks are
one list with no parallel copy to maintain. And it closes on itself: `42.type` is
`Integer`, `Integer.type` is `Type`, `Type.type` is `Type`.

A type group is named after the method it guarantees: `Stringable` means `to_string`
exists, `Integerable` means `to_integer` exists.

### Operators

Unchanged from v0.23. `!` remains logical not and `? :` remains the ternary, alongside
their roles as method-name suffixes. Replacing `!` with `not` and deleting the ternary
was considered and rejected.

This makes one lexer rule mandatory rather than optional: **an identifier may end in at
most one `?` or `!`.** Without it `x.empty??"a":"b"` is ambiguous, because the lexer
currently lets an identifier absorb every trailing `?`.

### Control flow

`if / elif / else / end`, `while / end`, `break`, `next`, `return` are unchanged.
`foreach` becomes `for`; nothing else `for` could mean here.

### Errors

`begin` becomes **`try`**. Investigation showed `begin` has no purpose beyond providing
a bare block for a `rescue` to attach to: it does not introduce a scope
(`evaluator.go:24` passes the same `env`), and `rescue` is parsed by `parseBlock`, so
*every* block already accepts one — `if … rescue … end`, `while … rescue … end` and
`def … rescue … end` all work today. Since the keyword carries no other meaning, `try`
is strictly clearer.

`rescue` remains available on any block. That is uniform and already true; the redesign
documents it as a feature rather than leaving it an undocumented accident.

```rocketlang
try
  risky()
rescue e
  print(e.message)
end

def parse_config(text)      # still valid on any block
  return JSON.parse(text)
rescue e
  return {}
end
```

### Scoping

Today's behaviour is arbitrary and undocumented: `while` and `for` create a child
environment, `if` and `begin` do not. The rule becomes:

> **Assignments are function-scoped.** A name assigned anywhere in a function body is
> visible for the rest of it.
>
> **Binding forms are scoped to their construct.** Function parameters, lambda
> parameters, the `for` variable and the `rescue` variable behave like parameters, not
> assignments. The `for` variable is a fresh binding on each iteration.

```rocketlang
if found
  result = x
end
print(result)      # works — an assignment

for i in 0..5
  total = total + i
end
print(total)       # works — an assignment
print(i)           # error — a binding form

try
  risky()
rescue e
end
print(e)           # error — a binding form
```

This fixes two live bugs:

- **Closure capture.** `foreach.go:42` creates the child environment once, *before* the
  loop, and reassigns the same binding each iteration, so every closure made in a loop
  shares one variable. Verified against v0.23: three closures over a loop variable all
  return `2`. This is the JavaScript `var` bug, and it is about to become far more
  visible now that `->` lambdas and `map`/`filter` exist. Go 1.22, C# 5 and JavaScript
  `let` all changed their semantics to fix the same problem.
- **Rescue variable leak.** `evalBlock` does `env.Set(block.Rescue.ErrorIdent.Literal, result)`
  on the enclosing environment, so `e` stays bound after the block ends. Under the rule
  above it is a binding form and belongs in a child environment.

### Modules

Unchanged. `import "path" as Alias only A, B`, and `export` prefixing a declaration.
This is already one of the better-designed parts of the language and is deliberately
left alone.

### Program output

A script no longer prints its final value; the REPL still echoes it. Today every file
in `tests/` ends with a bare `nil` purely to suppress that output, which is the clearest
possible signal the behaviour is wrong for scripts.

## Naming

The tables below were taken from the running interpreter (`x.methods` per type, and
the `*Functions` maps in `stdlib/`), not from reading source. Several methods are
registered through helpers — `arrayPair`, `arrayPredicate`, `integerPredicate`,
`integerRounding`, `floatPredicate`, `floatRounding`, `hashCallbackPair`, `stringPair`,
`stringPredicate` — so a grep for literal method names silently misses them. Re-derive
the same way if these tables need checking.

### Rules

1. Conversions are `to_` + the type name.
2. `size` never takes arguments and means "how many are there". `count(…)` always takes
   an argument and means "how many match".
3. Opposites are spelled as opposites. If `append` exists, its partner is `prepend`.
4. A number's own operations are methods on the number. `Math` keeps only constants and
   randomness.
5. Spell words out. Only long-established mathematical abbreviations stay short: `abs`,
   `min`, `max`, `sqrt`, `gcd`, `lcm`, `pow`, `divmod`, `exp`, `log`.
6. Casing: types and modules are `CapCase`, constants are `CapCase`, methods and
   variables are `snake_case`.

### Universal methods

| v0.23 | v1.0 |
|---|---|
| `to_s` | `to_string` |
| `to_i` | `to_integer` |
| `to_f` | `to_float` |
| `to_json` | `to_json` |
| `is_a?("STRING")` | `is_a?(String)` |
| `type` → `"STRING"` | `type` → `String` |
| `type_groups` → strings | `type_groups` → `Type`s |
| `nil?` | `nil?` |
| `methods` | `methods` |
| `wat` | `help`, with `wat` retained as an alias |

`wat`/`help` is **the only alias in the language** — a deliberate easter egg, not a
precedent.

### String

| v0.23 | v1.0 |
|---|---|
| `find` | `index_of` |
| — | `last_index_of` (new, pairs with the above) |
| `ascii` | `codepoints` |
| `upcase` / `upcase!` | `uppercase` / `uppercase!` |
| `downcase` / `downcase!` | `lowercase` / `lowercase!` |
| `swapcase` / `swapcase!` | `swap_case` / `swap_case!` |
| `strip` / `strip!` | `trim` / `trim!` |
| `lstrip` / `lstrip!` | `trim_start` / `trim_start!` |
| `rstrip` / `rstrip!` | `trim_end` / `trim_end!` |
| `chop` / `chop!` | `remove_last` / `remove_last!` |
| `chomp` / `chomp!` | `trim_line_end` / `trim_line_end!` |
| `include?` | `contains?` |
| `start_with?` | `starts_with?` |
| `end_with?` | `ends_with?` |

Unchanged: `count`, `format`, `size`, `split`, `lines`, `reverse`/`!`, `capitalize`/`!`,
`replace`/`!`, `empty?`.

### Array

| v0.23 | v1.0 |
|---|---|
| `index` | `index_of` |
| `rindex` | `last_index_of` |
| `push` | `append` |
| `pop` | `remove_last` |
| `unshift` | `prepend` |
| `shift` | `remove_first` |
| `delete` | `remove` |
| `delete_at` | `remove_at` |
| `include?` | `contains?` |
| `uniq` / `uniq!` | `unique` / `unique!` |
| `slices` | `chunks` |
| `to_m` | `to_matrix` |
| `take(n)` | **deleted** — behaviourally identical to `first(n)`, verified |
| `drop(n)` | `skip(n)` |
| `select` / `select!` | `filter` / `filter!` |
| — | `skip_last(n)` (new, pairs with `skip`) |

Unchanged: `join`, `reverse`/`!`, `size`, `sort`/`!`, `sort_by`/`!`, `sum`, `first`,
`last`, `min`, `max`, `min_by`, `max_by`, `count`, `insert`, `clear`, `concat`,
`empty?`, `reduce`, `each`, `map`/`!`, `reject`/`!`, `compact`/`!`, `flatten`/`!`,
`rotate`/`!`, `all?`, `any?`, `none?`.

### Hash

| v0.23 | v1.0 |
|---|---|
| `include?` | `has_key?` |
| `delete` | `remove` |
| `select` / `select!` | `filter` / `filter!` |

`get` and `fetch` keep their names and gain a stated rule: **`get(key)` returns `nil`
when the key is missing, `get(key, default)` returns the default, and `fetch(key)`
raises.**

Unchanged: `keys`, `values`, `each`, `size`, `empty?`, `clear`, `invert`, `merge`/`!`,
`reject`/`!`, `compact`/`!`, `transform_keys`/`!`, `transform_values`/`!`.

### Integer and Float

| v0.23 | v1.0 |
|---|---|
| `succ` | `successor` (`next` is a keyword) |
| `pred` | `predecessor` |
| `chr` | `to_character` |
| `base` + `to_base(n)` | `to_base(n)` only |
| `upto(n, f)` / `downto(n, f)` | **deleted** — `(a..b).each(f)` says it, once ranges exist |

Methods absorbed from `Math` (rule 4), none of which have a number-method twin today:
`sqrt`, `exp`, `log`, `log2`, `log10`, `sin`, `cos`, `tan`, `asin`, `acos`, `atan`,
`copysign`, `remainder`.

Unchanged: `abs`, `times`, `digits`, `divmod`, `gcd`, `lcm`, `pow`, `bit_length`,
`ceil`, `floor`, `round`, `truncate`, `even?`, `odd?`, `zero?`, `positive?`,
`negative?`, `nan?`, `finite?`, `infinite?`.

`upto`/`downto` are deleted rather than renamed: they exist only because there was no
range value to iterate. `1.upto(5, f)` becomes `(1..5).each(f)`.

**Sequencing:** their deletion belongs to step 5, not step 1. Ranges do not exist until
step 5, so removing them earlier would leave the behaviour unreachable in between.

### Error, File, Matrix

| v0.23 | v1.0 |
|---|---|
| `Error#msg` | `Error#message` |
| `Matrix#to_a` | `Matrix#to_array` |
| `Matrix#t` | **deleted** — abbreviation of `transpose` |

`Matrix#rows` and `#cols` currently exist as both a method and a property; only the
method survives. `File` is unchanged.

### Globals and modules

| v0.23 | v1.0 |
|---|---|
| `puts` | `print` |
| `raise` | `raise` |
| `OS.raise` | **deleted** — not a clean duplicate: `OS.raise(code, message)` let the caller pick the exit code, the global `raise` always exits 1 when uncaught. Reconstructible as `print(msg); OS.exit(n)`, so nothing is lost, but the two were not equivalent |
| `Math.rand` | `Math.random` |
| `Math.abs`, `.ceil`, `.floor`, `.round` | **deleted** — already duplicated on both Integer and Float |
| `Math.pow` | **deleted**, but only after adding `Float#pow` — verified missing, so `Integer` had it and `Float` did not |
| `Math.sqrt`, `.exp`, `.log`, `.log2`, `.log10`, `.sin`, `.cos`, `.tan`, `.asin`, `.acos`, `.atan`, `.copysign`, `.remainder` | **moved** to number methods |
| `Math.max`, `Math.min` | **deleted** — `[a, b].max` |
| — | `IO.write` (new — output without a newline) |
| — | `IO.read_line` (new) |
| — | `Time.now` (new) |

`Math` retains its constants: `Pi`, `E`, `Phi`, `Sqrt2`, `SqrtE`, `SqrtPi`, `SqrtPhi`,
`Ln2`, `Log2E`, `Ln10`, `Log10E`. `IO.open`, `JSON.parse`, `OS.exit`, `Time.format`,
`Time.parse`, `Time.sleep`, `Time.unix`, `HTTP.new` and the `Time` layout constants are
unchanged.

## Implementation notes

Three details that will otherwise surface as mystery bugs:

- **Identifier suffix limit.** An identifier may end in at most one `?` or `!`, so
  `empty??` lexes as `empty?` followed by the ternary `?`.
- **Nested quotes in interpolation.** `"#{h.get("k")}"` requires the string lexer to
  track quote and brace nesting inside `#{…}` rather than stopping at the first `"`.
- **Type bindings are generated, not written.** They come from `knownObjectTypes` and
  `typeGroups`, so a new object type joins `is_a?` automatically. `TestKnownObjectTypesAreComplete`
  already guards the registry. Only entries flagged user-facing produce a binding.
- **Per-iteration binding costs an allocation per loop iteration.** `NewEnclosedEnvironment`
  moves inside the iteration in `foreach.go`. If that shows up in profiling, the fix is
  to allocate a fresh environment only when the body actually creates a closure, but
  correctness comes first and the simple version should ship.

## Migration

### Corpus

1,744 lines of RocketLang across 120 files. Nearly all of it has byte-pinned expected
output (`tests/*.expected`, `exercises/expected/*.txt`), so a rewrite that keeps those
outputs identical is provably semantics-preserving.

**Do not run `go test ./exercises -update` during the migration.** That command
regenerates expectations *from* the solutions, so running it would let a rewrite
silently redefine what is correct. The expectations are the oracle and must stay frozen.

### Step 0 — retire the exercises

`exercises/` was added on 2026-08-30 (`c5c0f19`, `13095da`). Teaching material written
against a syntax that is about to disappear should be rewritten, not translated: an
exercise whose point is `foreach i in 0 -> 5 ^ 2` has no counterpart worth preserving.

- **Keep**: the 39 solutions (245 lines), moved into the `tests/` golden-file harness
  that `TestRocketlangCode` already runs. They cover combinations the Go unit tests do
  not — control flow inside functions, hashes inside loops, errors across imports.
- **Delete**: the 39 task files, `exercises/expected/`, `exercises/run.sh`,
  `exercises/exercises_test.go`, and `wasm/exercises.json`, along with the
  `TestBundleMatchesExpectations` coupling between the repo and the playground bundle.
- **Later**: write fresh exercises for v1.0 and build the playground section properly.
  Everything is recoverable from git history.

`main_test.go` globs `tests/*.rl` non-recursively; moving the solutions into a
subdirectory requires widening that glob.

### Steps

Each step is a PR on `main` that keeps CI green and updates the corpus as it goes. No
intermediate releases; `v1.0.0` is tagged when the sequence completes.

| # | Step | Corpus impact |
|---|---|---|
| 1 | Library rename: the whole naming table, `select` → `filter`, `Math` consolidation, `OS.raise` deleted, `take` deleted, `skip_last`/`last_index_of`/`IO.write`/`IO.read_line`/`Time.now` added, `help`/`wat` | find-and-replace |
| 2 | `#` comments, `foreach` → `for`, scripts stop printing their final value | regex; deletes the trailing bare `nil` from fixtures |
| 3 | `begin` → `try`; the scoping rule; per-iteration loop binding; rescue-variable leak fixed | only files reading a `for` or `rescue` variable after its block |
| 4 | `Type` object, `is_a?(String)`, `.type` returns a `Type`; `HTTP` type renamed `HttpServer` | only files using `is_a?`/`type` |
| 5 | Ranges and slices: `..`, `..<`, `.by(n)`; retires `->`, `=>`, `^`, `[a:b]` | only files using ranges or slices |
| 6 | Lambdas `x -> expr` and trailing `do … end` | only files using `each`/`reduce` — 4 today |
| 7 | `#{}` interpolation | none forced; `+` concatenation keeps working |
| 8 | Parenless dot-calls | one regex, every file |
| 9 | Regenerate docs, `docusaurus docs:version v1.0.0`, rebuild the playground, update the VS Code extension grammar | — |

The ordering has one hard dependency: **step 5 must precede step 6**, because retiring
the range arrows is what frees `->` for lambdas. Step 8 is last because it touches every
line and should land only once everything else is settled.

Step 3 is the only step that changes *semantics* rather than surface syntax, so it needs
its own tests rather than relying on the golden-file corpus: the corpus asserts output,
and a scoping bug can preserve output while changing what a closure captures.

### External

- **VS Code extension** (`Flipez/rocket-lang-support`, separate repo) needs its grammar
  updated in lockstep or highlighting breaks for every user on the day v1.0 ships.
- **Published planets** are third-party `.rl` that will stop parsing. There is no
  compatibility shim; this is the accepted cost of the break.
- **Docs versioning** already keeps five versions. Hold `v0.23.0` well past the usual
  pruning point so the old syntax stays readable for anyone who has not migrated.

## Decisions considered and rejected

| Proposal | Outcome | Reason |
|---|---|---|
| Drop the `!` mutation suffix; make all methods non-mutating | rejected | `!` kept; the rule is tightened instead so a `!` method exists only as the counterpart of a same-named non-mutating one |
| `not` instead of `!`; delete the ternary | rejected | both kept; costs the identifier-suffix lexer rule |
| `begin` → `try` | **accepted** | `begin` introduces no scope, and `rescue` already works on every block, so the keyword meant nothing that `try` does not say more clearly |
| `begin` kept as an explicit local-scope block | rejected | speculative; no demand for it, and the scoping rule makes it unnecessary |
| Every block introduces a scope (C-family) | rejected | breaks assigning inside an `if` and reading the result after, which is common in scripts |
| Loop variable visible after the loop | rejected | it is a binding form, not an assignment; no language combines that with per-iteration binding |
| Word ranges (`0 to 5 by 2`) | rejected | `..`/`..<` chosen |
| `...` for exclusive ranges | rejected | `..<` is harder to misread |
| `{expr}` or `${expr}` interpolation | rejected | `#{}` is safest against literal text and needs no prefix |
| `length` instead of `size` | rejected | `size` is shorter, already the name, and not an abbreviation |
| Parens required on all calls (one call rule) | rejected | more uniform but noisier; ease of use won |
| Type names as reserved keywords | rejected | keywords are not values; `Type` objects give `x.type == String` for free |
| Ruby blocks with `do |x|` | rejected | pipes, and blocks as a concept separate from arguments |
| Dual-mode interpreter with a legacy flag | rejected | doubles the grammar and the test matrix indefinitely |
