# Integer

An integer can be positiv or negative and is always internally represented by a 64-Bit Integer.

To cast a negative integer a digit can be prefixed with a - eg. -456.



```js
a = 1;

b = a + 2;

is_true = 1 == 1;
is_false = 1 == 2;

```

## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Converts the integer into a float.


```js
1234.plz_f()

```

```js
1234.0

```



### plz_i()
> Returns `INTEGER`

Returns self


```js
1234.plz_i()

```

```js
1234

```



### plz_s(INTEGER)
> Returns `STRING`

Returns a string representation of the integer. Also takes an argument which represents the integer base to convert between different number systems


```js
1234.plz_s()
1234.plz_s(2)
1234.plz_s(8)
1234.plz_s(10)

```

```js
"1234"
"10011010010"
"2322"
"1234"

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



