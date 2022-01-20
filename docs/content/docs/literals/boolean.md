---
title: "Boolean"
menu:
  docs:
    parent: "literals"
---
# Boolean

A Boolean can represent two values: `true` and `false` and can be used in control flows.


```js
true // Is the representation for truthyness
false // is it for a falsy value

a = true;
b = false;

is_true = a == a;
is_false = a == b;

is_true = a != b;
```

## Literal Specific Methods

### plz_s()
> Returns `STRING`

Converts a boolean into a String representation and returns `"true"` or `"false"` based on the value.


```js
🚀 > true.plz_s()
=> "true"
```



## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.

```js
🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]
```

### type()
> Returns `STRING`

Returns the type of the object.

```js
🚀 > "test".type()
=> "STRING"
```

### wat()
> Returns `STRING`

Returns the supported methods with usage information.

```js
🚀 > true.wat()
=> BOOLEAN supports the following methods:
				plz_s()
```
