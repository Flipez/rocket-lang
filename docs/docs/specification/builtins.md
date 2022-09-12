---
title: "Builtin Functions"
menu:
  docs:
    parent: "specification"
toc: true
---
# Builtin Functions
## puts(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FILE)

Prints the string representation of the given object to STDOUT.

```js
🚀 > puts("test")
"test"

🚀 > puts([1,2,3])
[1, 2, 3]

🚀 > puts(["test",true,3])
["test", true, 3]
```