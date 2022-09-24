---
title: "If"
menu:
  docs:
    parent: "controls"
---
# If
With `if` and `else` keywords the flow of a program can be controlled.

```js
🚀 > if (a.type() == "STRING")
  puts("is a string")
else
  puts("is not a string")
end

// which prints
is a string
```

## else if
You can also chain an `if` statement after an `else` statement which makes the nested `end` optional.

```js
if (type == "BOOLEAN")
  a.yoink(true)
else if (type == "STRING")
  a.yoink("")
else if (type == "INTEGER")
  a.yoink(0)
else if (type == "FLOAT")
  a.yoink(0.0)
else if (type == "ARRAY")
  a.yoink([])
else if (type == "HASH")
  a.yoink({})
end
```