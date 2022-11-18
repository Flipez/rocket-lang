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
true.plz_s()

```

```js
"true"

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



