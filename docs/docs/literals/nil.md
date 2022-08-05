---
title: "Nil"
menu:
  docs:
    parent: "literals"
---
# Nil

Nil is the representation of "nothing".
	It will be returned if something returns nothing (eg. puts or an empty break/next) and can also be generated with 'nil'.


## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Returns zero float.



### plz_i()
> Returns `INTEGER`

Returns zero integer.



### plz_s()
> Returns `STRING`

Returns empty string.




## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.

```js
🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]
```

### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.

```js
🚀 > a = {"test": 1234}
=> {"test": 1234}
🚀 > a.to_json()
=> "{"test":1234}"
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
