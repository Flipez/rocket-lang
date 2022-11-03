# Float




## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Returns self


```js
🚀 » 123.456.plz_f()
» 123.456

```


### plz_i()
> Returns `INTEGER`

Converts the float into an integer.


```js
🚀 » 123.456.plz_i()
» 123

```


### plz_s()
> Returns `STRING`

Returns a string representation of the float.


```js
🚀 » 123.456.plz_s()
» "123.456"

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

