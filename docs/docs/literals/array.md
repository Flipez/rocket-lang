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
🚀 > ["a", "b", 1, 2].first()
=> "a"
```


### index(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FILE)
> Returns `INTEGER`

Returns the index of the given element in the array if found. Otherwise return `-1`.


```js
🚀 > ["a", "b", 1, 2].index(1)
=> 2
```


### last()
> Returns `STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE`

Returns the last element of the array.


```js
🚀 > ["a", "b", 1, 2].last()
=> 2
```


### size()
> Returns `INTEGER`

Returns the amount of elements in the array.


```js
🚀 > ["a", "b", 1, 2].size()
=> 4
```


### uniq()
> Returns `ARRAY|ERROR`

Returns a copy of the array with deduplicated elements. Raises an error if a element is not hashable.


```js
🚀 > ["a", 1, 1, 2].uniq()
=> [1, 2, "a"]
```


### yeet()
> Returns `STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE`

Removes the last element of the array and returns it.


```js
🚀 > a = [1,2,3]
=> [1, 2, 3]
🚀 > a.yeet()
=> 3
🚀 > a
=> [1, 2]
```


### yoink(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE)
> Returns `NIL`

Adds the given object as last element to the array.


```js
🚀 > a = [1,2,3]
=> [1, 2, 3]
🚀 > a.yoink("a")
=> nil
🚀 > a
=> [1, 2, 3, "a"]
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

