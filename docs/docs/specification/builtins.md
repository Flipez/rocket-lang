---
title: "Builtin Functions"
menu:
  docs:
    parent: "specification"
toc: true
---
# Builtin Functions
## print(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FILE)

Prints the string representation of the given object to STDOUT.

```js
🚀 > print("test")
test

🚀 > print([1,2,3])
[1, 2, 3]

🚀 > print(["test",true,3])
["test", true, 3]
```

## raise(STRING)

Returns an ERROR object carrying the given message -- the same kind of value
any other builtin returns when something goes wrong, just returned on
purpose. A `rescue` block (see [Error](../literals/error)) can catch it and
carry on. Left uncaught, it aborts the rest of the program and becomes its
result, so the process exits with status 1.

`OS.raise(code, message)` used to exist alongside this and looked like a
duplicate, but was not one: it let the caller choose the process's exit
code, where an uncaught `raise` always exits 1. It was removed anyway,
because the same effect is reachable as `print(message); OS.exit(code)`.

```js
🚀 > raise("broken")
ERROR: broken

🚀 > def test()
       raise("broken")
     rescue e
       print("caught: " + e.message())
     end
🚀 > test()
caught: broken
```