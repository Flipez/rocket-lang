# Nil

Nil is the representation of "nothing".
It will be returned if something returns nothing (eg. puts or an empty break/next) and can also be generated with 'nil'.



## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Returns zero float.


```js
🚀 » nil.plz_f()
» 0.0

```


### plz_i()
> Returns `INTEGER`

Returns zero integer.


```js
🚀 » nil.plz_i()
» 0

```


### plz_s()
> Returns `STRING`

Returns empty string.


```js
🚀 » nil.plz_s()
» ""

```



## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.

```js
🚀 »  "test".methods()
» ["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "plz_i", "strip", "downcase"]

```

### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.

```js
🚀 » a = {"test": 1234}
» {"test": 1234}
🚀 » a.to_json()
» "{\"test\":1234}"

```

### type()
> Returns `STRING`

Returns the type of the object.

```js
🚀 » "test".type()
» "STRING"

```

### wat()
> Returns `STRING`

Returns the supported methods with usage information.

```js
🚀 » true.wat()
» "BOOLEAN supports the following methods:
        plz_s()"

```

