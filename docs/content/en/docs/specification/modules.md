---
title: "Modules"
menu:
  docs:
    parent: "specification"
toc: true
---
{{< alert icon="👉" text="Modules were introduced in `0.11`" />}}

Modules are seperate RocketLang files can be imported using the `import` statement.
Functions and <mark>variables starting with a uppercase</mark> name are then available in the imported module.

For example take this module:

```js
// fixtures/module.rl
a = 1
A = 5

Sum = def (a, b) {
    return a + b
}


```

You can import it with:

```js
import("fixtures/module")
```

This results in a variable `module` implicitly being assigned.
You can use it like so:

```js
🚀 > import("fixtures/module")
=> null
🚀 > module.a
=> null
🚀 > module.A
=> 5
🚀 > module.Sum(module.A, 2)
=> 7
```
