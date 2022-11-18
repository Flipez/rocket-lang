# Array




```js
a = [1, 2, 3, 4, 5]
puts(a[2])
puts(a[-2])
puts(a[:2])
puts(a[:-2])
puts(a[2:])
puts(a[-2:])
puts(a[1:-2])

// should output
[1, 2]
[1, 2, 3]
[3, 4, 5]
[4, 5]
[2, 3]
[1, 2, 8, 9, 5]

```

## Literal Specific Methods

### first()
> Returns `STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE`

Returns the first element of the array. Shorthand for `array[0]`


```js
["a", "b", 1, 2].first()

```

```js
"a"

```



### index(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FILE)
> Returns `INTEGER`

Returns the index of the given element in the array if found. Otherwise return `-1`.


```js
["a", "b", 1, 2].index(1)

```

```js
2

```



### last()
> Returns `STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE`

Returns the last element of the array.


```js
["a", "b", 1, 2].last()

```

```js
2

```



### reverse()
> Returns `ARRAY`

Reverses the elements of the array


```js
["a", "b", 1, 2].reverse()

```

```js
[2, 1, "b", "a"]

```



### size()
> Returns `INTEGER`

Returns the amount of elements in the array.


```js
["a", "b", 1, 2].size()

```

```js
4

```



### sort()
> Returns `ARRAY`

Sorts the array if it contains only one type of STRING, INTEGER or FLOAT


```js
[3.4, 3.1, 2.0].sort()

```

```js
[2.0, 3.1, 3.4]

```



### uniq()
> Returns `ARRAY|ERROR`

Returns a copy of the array with deduplicated elements. Raises an error if a element is not hashable.


```js
["a", 1, 1, 2].uniq()

```

```js
[1, 2, "a"]

```



### yeet()
> Returns `STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE`

Removes the last element of the array and returns it.


```js
a = [1,2,3]
a.yeet()
a

```

```js
[1, 2, 3]
3
[1, 2]

```



### yoink(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE)
> Returns `NIL`

Adds the given object as last element to the array.


```js
a = [1,2,3]
a.yoink("a")
a

```

```js
[1, 2, 3]
nil
[1, 2, 3, "a"]

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



