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
  print("test")
end

=> def ()
  print(test)
end

🚀 > test()
test
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

## Arguments

A call has to supply exactly as many arguments as the function has parameters.
Too few or too many is an error:

```js
def add(a, b)
  return a + b
end

add(1, 2)     // 3
add(1)        // ERROR: add: too few arguments: got=1, want=2
add(1, 2, 3)  // ERROR: add: too many arguments: got=3, want=2
```

A named function reports its own name, which helps when the call is into a
module rather than a function on screen. Like any other error it can be caught:

```js
begin
  add(1)
rescue e
  print(e.message())   // add: too few arguments: got=1, want=2
end
```
