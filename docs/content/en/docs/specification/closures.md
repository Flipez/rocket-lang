---
title: "Closures"
menu:
  docs:
    parent: "specification"
toc: true
---
```js
newGreeter = fn(greeting) {
  return fn(name) { puts(greeting + " " + name); }
};

hello = newGreeter("Hello");

hello("dear, future Reader!");

```