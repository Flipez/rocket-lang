---
title: "Builtin Functions"
menu:
  docs:
    parent: "specification"
toc: true
---
# Builtin Functions
## exit(INTEGER)

Terminates the program with the given exit code.

```js
🚀 > exit(1)
exit status 1
```

## raise(INTEGER, STRING)

Terminates the program with the given exit code and prints the error message.

```js
🚀 > raise(1, "broken")
🔥 RocketLang raised an error: "broken"
exit status 1
```

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