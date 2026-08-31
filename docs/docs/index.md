---
sidebar_position: 0
---
# Getting Started

RocketLang started as a complete implementation of
[MonkeyLang](https://monkeylang.org/) — the language built in
*Writing an Interpreter in Go* — and has been extended since with features both
useful and not so useful.

# Latest Version

[![GitHub release](https://img.shields.io/github/release/flipez/rocket-lang.svg)](https://github.com/flipez/rocket-lang/releases/)

# Quick Start

```js
// Values carry methods.
name = "rocket-lang"
puts(name.uppercase())        // ROCKET-LANG
puts(name.split("-").size()) // 2

// Arrays and hashes, both with methods of their own.
crew = ["ada", "grace", "alan"]
puts(crew.size())            // 3
puts(crew.contains?("ada"))   // true

ages = {"ada": 36, "grace": 45}
puts(ages["grace"])          // 45

// Functions are values, so they can be passed around.
double = def(n)
  return n * 2
end
puts(double(21))             // 42

// Blocks close with `end`, and parentheses around a condition are optional.
foreach i, member in crew
  if i > 0
    puts(i.to_string() + ": " + member)
  end
end

// Errors are values you can catch.
begin
  puts(1 / 0)
rescue e
  puts("caught: " + e.msg())
end
```

Split a program across files with [modules](./language/modules), and pull in
someone else's with [planets](./language/planets):

```js
import "./helpers" as helpers
```

For longer programs, the [examples
directory](https://github.com/flipez/rocket-lang/tree/main/examples) has
solutions to several Advent of Code puzzles.

# Help

Launch RocketLang with `-h` or `--help` for an overview of the CLI.

```zsh
$ rocket-lang -h
Usage: rocket-lang [flags] [program file] [arguments]
       rocket-lang planet <command>

Available flags:
  -e, --exec string   Runs the given code.
  -v, --version       Prints the version and build date.
```

Run a file, evaluate a snippet, or start a REPL with no arguments at all:

```zsh
$ rocket-lang program.rl
$ rocket-lang -e 'puts("hi")'
$ rocket-lang
```
