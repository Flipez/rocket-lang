---
title: "If"
menu:
  docs:
    parent: "controls"
---
With `if` and `else` keywords the flow of a program can be controlled.

```js
🚀 > if (a.type() == "STRING") {
  puts("is a string")
} else {
  puts("is not a string")
}

// which prints
is a string
```

> 👉 Since `0.13` curly braces are completely optional (closing brace needs to be replaced with `end`)

```js
🚀 > if (a.type() == "STRING")
  puts("is a string")
else
  puts("is not a string")
end

// which prints
is a string
```