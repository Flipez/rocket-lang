---
title: "Hash"
menu:
  docs:
    parent: "literals"
---
# Hash




```js
people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];
```

## Literal Specific Methods

### keys()
> Returns `ARRAY`

Returns the keys of the hash.


```js
🚀 > {"a": "1", "b": "2"}.keys()
=> ["a", "b"]
```


### values()
> Returns `ARRAY`

Returns the values of the hash.


```js
🚀 > {"a": "1", "b": "2"}.values()
=> ["2", "1"]
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
