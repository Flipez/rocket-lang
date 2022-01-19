---
title: "Closures"
menu:
  docs:
    parent: "specification"
toc: true
---
```js
newGreeter = def (greeting) {
  return def (name) { puts(greeting + " " + name); }
};

hello = newGreeter("Hello");

hello("dear, future Reader!");

```