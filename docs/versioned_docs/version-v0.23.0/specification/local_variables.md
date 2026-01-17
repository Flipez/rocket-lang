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
a = "test"
```
## Examples

```js
some_integer = 1;
name = "RocketLang";
array = [1, 2, 3, 4, 5];
some_boolean = true;
```

Also expressions can be used
```js
another_int = (10 / 2) * 5 + 30;
an_array = [1 + 1, 2 * 2, 3];
```

## Multiple Assignment (Array Unpacking)

Multiple variables can be assigned at once by unpacking an array:

```js
a, b, c = [1, 2, 3]
// a = 1, b = 2, c = 3
```

This works with any expression that returns an array:

```js
def get_coordinates()
  return [10, 20, 30]
end

x, y, z = get_coordinates()
// x = 10, y = 20, z = 30
```

If the array has more elements than variables, the extra elements are ignored:

```js
a, b = [1, 2, 3, 4]
// a = 1, b = 2 (3 and 4 are ignored)
```

If the array has fewer elements than variables, an error is raised:

```js
a, b, c = [1, 2]
// ERROR: not enough values to unpack (expected 3, got 2)
```
