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