---
title: "Local Variables"
menu:
  docs:
    parent: "specification"
toc: true
---

# Local Variables
Variables represent the object as a memorable name.

```js
let a = "test"
```

Variables can also be overridden by ommiting the `let` statement.

```js
a = 1
```

> New in `0.10`:
Ommiting the `let` will also create a variable if not present making `let` completely optional.

## Examples

```js
let some_integer = 1;
let name = "RocketLang";
let array = [1, 2, 3, 4, 5];
let some_boolean = true;
```

Also expressions can be used
```js
let another_int = (10 / 2) * 5 + 30;
let an_array = [1 + 1, 2 * 2, 3];
```