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

## open(STRING, STRING, STRING)
> Returns FILE

Opens a file pointer to the file at the path of argument one with the mode of (optional) argument two and a file permission as optional argument three.

```js
🚀 > open("main.go", "r", "0644")
=> <file:main.go>
```

Available file modes are `r`, `w`, `wa`, `rw` and `rwa`,