---
title: "JSON"
menu:
  docs:
    parent: "literals"
---
# JSON




```js
🚀 > JSON.parse('{"test": 123}')
=> {"test": 123.0}
```

## Literal Specific Methods

### parse(STRING)
> Returns `HASH`

Takes a STRING and parses it to a HASH or ARRAY. Numbers are always FLOAT.


```js
🚀 > JSON.parse('{"test": 123}')
=> {"test": 123.0}
🚀 > JSON.parse('["test", 123]')
=> ["test", 123.0]
```



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
