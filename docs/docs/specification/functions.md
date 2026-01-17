---
title: "Functions"
menu:
  docs:
    parent: "specification"
toc: true
---
# Functions
Implicit and explicit return statements are supported.

```js
fibonacci = def (x)
  if x == 0
    0
  else
    if x == 1
      return 1;
    else
      fibonacci(x - 1) + fibonacci(x - 2);
    end
  end
end
```

> New in `0.11`:

Functions can now also be created as named functions:

```js
🚀 > def test()
  puts("test")
end

=> def ()
  puts(test)
end

🚀 > test()
"test"
```

## Multiple Return Values

Functions can return multiple values using comma-separated syntax:

```js
def get_coordinates()
  return 10, 20, 30
end

coords = get_coordinates()
// coords = [10, 20, 30]
```

This is syntax sugar that automatically wraps the values in an array. The following are equivalent:

```js
return 1, 2, 3
return [1, 2, 3]
```

The returned array can be unpacked into multiple variables (see [Multiple Assignment](./local_variables#multiple-assignment-array-unpacking)):

```js
x, y, z = get_coordinates()
// x = 10, y = 20, z = 30
```
