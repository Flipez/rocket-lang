# Modules
> 👉 Modules were introduced in `0.11`
>
> 👉 The `import` syntax changed in `0.24` and `export` became required

Modules are separate RocketLang files. A module makes a value public with
`export`; everything else stays private to the file.

## Exporting

`export` works in three forms and is only valid at the top level of a file:

```js
// fixtures/module.rl
a = 1
export A = 5
export lower = 7
Private = 99

export def Sum(a, b)
    return a + b
end
```

`a` and `Private` are both private. Capitalization means nothing — before
`0.24` an uppercase name was exported automatically, and that rule is gone.

You can also export a name that is already bound:

```js
Square = def(x) return x * x end
export Square
```

## Importing

```js
import "fixtures/module"
```

`import` is a statement, not an expression: it binds a name as a side effect
and cannot be used where a value is expected, so `x = import "lib"` is a parse
error.

This binds a variable named after the path's last segment:

```js
🚀 > import "fixtures/module"
=> nil
🚀 > module.A
=> 5
🚀 > module.Sum(module.A, 2)
=> 7
🚀 > module.a
=> ERROR: module 'fixtures/module' has no export 'a'
```

Accessing something the module does not export is an error, not `nil`.

### Choosing the name with `as`

```js
🚀 > import "fixtures/module" as anotherModule
=> nil
🚀 > anotherModule.A
=> 5
```

An import fails if its name is already taken, whether by a variable, an
earlier import, or a builtin module such as `Math`.

### Narrowing with `only`

`only` restricts what the namespace contains. It never puts names into the
current scope:

```js
🚀 > import "fixtures/module" as narrow only Sum
=> nil
🚀 > narrow.Sum(1, 2)
=> 3
🚀 > narrow.A
=> ERROR: module 'fixtures/module' has no export 'A'
```

Naming something the module does not export is an error:

```js
🚀 > import "fixtures/module" only Nope
=> ERROR: :1:7: Import Error: 'fixtures/module' does not export 'Nope'; exported: 'A', 'Sum', 'lower'
```

Every `Import Error` carries a `file:line:column:` prefix locating the failing
`import`. At the REPL there is no file, so that part is empty and the message
begins with a bare `:1:7:`. When the same import fails inside a script, the
script's path appears there instead.

### A module cannot be exported

`export` refuses a value that is a module, whether it names an already-bound
module or evaluates to one:

```js
🚀 > import "fixtures/module" as Inner
=> nil
🚀 > export Inner
=> ERROR: :1:7: Export Error: cannot export 'Inner': a module cannot be exported
```

To build on an imported module, export a function that calls through it
instead of trying to re-export the module itself:

```js
// math.rl
import "./stats" as Stats

export def Sum(a, b) return a + b end
export def Mean(numbers) return Stats.Mean(numbers) end
```

```js
import "./math"
math.Sum(1, 2)          // 3
math.Mean([1, 2])       // reaches Stats through the wrapper function
```

### Where an import may appear

An import is a statement, so it can go anywhere a statement can: at the top
level, inside a function body, or inside a branch. Importing conditionally is
the supported way to pick between implementations:

```js
if env == "prod"
  import "./config_prod" as config
else
  import "./config_stage" as config
end

puts(config.Url)
```

Only one branch runs, so only one binding is made, and it is visible after the
`if` ends.

An import inside a **loop body** does not work. Loops reuse one scope across
iterations, so the second iteration finds the name already bound:

```js
foreach name in ["a", "b"]
  import "./plugin" as p     // Import Error on the second iteration:
end                          // cannot bind module as 'p', name already in use
```

This is not a useful thing to write in any case. A path must be a string
literal, so every iteration would import the same module, and a module is
evaluated only once no matter how many times it is imported. Import what you
need once, outside the loop.

## Finding modules

A path starting with `./` or `../` resolves relative to the file doing the
importing, so a module can import its neighbours:

```js
// examples/aoc/2018/day2.rl
import "../util" as util
```

In the REPL and under `rocket-lang -e` there is no importing file, so a
relative path resolves against the working directory instead — the same place
a bare module name is found.

Any other path is looked up in the search paths: each entry of the
`ROCKETLANGPATH` environment variable in order, followed by the working
directory. Entries are separated by the platform's path list separator, `:` on
Unix and `;` on Windows.

The working directory is always searched, so setting `ROCKETLANGPATH` adds
places to look rather than replacing the default. Configured entries are tried
first, so they win when the same module name exists in both.

The path must be a string literal; it cannot be a variable or a computed
expression.

## Loading rules

A module file is evaluated **once**. Importing it again reuses the result, so
its side effects do not run a second time:

```js
import "fixtures/module"
import "fixtures/module" as sameThing
// the file was read and evaluated one time
```

Circular imports are an error rather than a hang. Each hop in the chain is
shown relative to the current working directory:

```js
🚀 > import "fixtures/cycle_a"
=> ERROR: .../fixtures/cycle_b.rl:1:7: Import Error: circular import
  fixtures/cycle_a.rl -> fixtures/cycle_b.rl -> fixtures/cycle_a.rl
```

The prefix names the module that closed the loop, and is shown here with its
leading directories elided. Note that the prefix is an absolute path while the
chain below it is relative.
