---
title: "Integer"
menu:
  docs:
    parent: "literals"
---

An integer can be positiv or negative and is always internally represented by a 64-Bit Integer.

To cast a negative integer a digit can be prefixed with a - eg. -456.


```js
a = 1;

b = a + 2;

is_true = 1 == 1;
is_false = 1 == 2;
```

## Literal Specific Methods

### plz_s(INTEGER)
> Returns `STRING`

Returns a string representation of the integer. Also takes an argument which represents the integer base to convert between different number systems


```js
🚀 > a = 456
=> 456
🚀 > a.plz_s()
=> "456"

🚀 > 1234.plz_s(2)
=> "10011010010"
🚀 > 1234.plz_s(8)
=> "2322"
🚀 > 1234.plz_s(10)
=> "1234"
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
