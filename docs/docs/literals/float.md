# Float




## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Returns self


```js
123.456.plz_f()

```

```js
123.456

```



### plz_i()
> Returns `INTEGER`

Converts the float into an integer.


```js
123.456.plz_i()

```

```js
123

```



### plz_s()
> Returns `STRING`

Returns a string representation of the float.


```js
123.456.plz_s()

```

```js
"123.456"

```




## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


```js
"test".methods()

```

```js
["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "plz_i", "strip", "downcase"]

```



### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


```js
a = {"test": 1234}
a.to_json()

```

```js
{"test": 1234}
"{\"test\":1234}"

```



### type()
> Returns `STRING`

Returns the type of the object.


```js
"test".type()

```

```js
"STRING"

```



### wat()
> Returns `STRING`

Returns the supported methods with usage information.


```js
true.wat()

```

```js
"BOOLEAN supports the following methods:
  plz_s()"

```



