# Modules
> 👉 Modules were introduced in `0.11`

Modules are seperate RocketLang files can be imported using the `import` statement.
Functions and <mark>variables starting with a uppercase</mark> name are then available in the imported module.

For example take this module:

```js
// fixtures/module.rl
a = 1
A = 5

Sum = def (a, b)
    return a + b
end
```

You can import it with:

```js
import("fixtures/module")
```

This results in a variable `module` implicitly being assigned.
You can use it like so:

```js
🚀 > import("fixtures/module")
=> nil
🚀 > module.a
=> nil
🚀 > module.A
=> 5
🚀 > module.Sum(module.A, 2)
=> 7
```

You can also define a name for the variable in which the module will be available:

```js
🚀 > import("fixtures/module", "anotherModule")
=> nil
🚀 > anotherModule.A
=> 5
```