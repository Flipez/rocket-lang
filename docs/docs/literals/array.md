import CodeBlockSimple from '@site/components/CodeBlockSimple'

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


<CodeBlockSimple input='["a", "b", 1, 2].first()
' output='"a"
' />


### include?(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FILE)
> Returns `BOOLEAN`

Returns true or false wether the array contains the given element


<CodeBlockSimple input='[1,2,3].include?(4)
[1,2,3].include?(3)
' output='false
true
' />


### index(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FILE)
> Returns `INTEGER`

Returns the index of the given element in the array if found. Otherwise return `-1`.


<CodeBlockSimple input='["a", "b", 1, 2].index(1)
' output='2
' />


### join(STRING)
> Returns `STRING`







### last()
> Returns `STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE`

Returns the last element of the array.


<CodeBlockSimple input='["a", "b", 1, 2].last()
' output='2
' />


### reverse()
> Returns `ARRAY`

Reverses the elements of the array


<CodeBlockSimple input='["a", "b", 1, 2].reverse()
' output='[2, 1, "b", "a"]
' />


### size()
> Returns `INTEGER`

Returns the amount of elements in the array.


<CodeBlockSimple input='["a", "b", 1, 2].size()
' output='4
' />


### slices(INTEGER)
> Returns `ARRAY`

Returns the elements of the array in slices with the size of the given integer


<CodeBlockSimple input='[1,2,3,4,5,6,7,8].slices(3)
' output='[[1, 2, 3], [4, 5, 6], [7, 8]]
' />


### sort()
> Returns `ARRAY`

Sorts the array if it contains only one type of STRING, INTEGER or FLOAT


<CodeBlockSimple input='[3.4, 3.1, 2.0].sort()
' output='[2.0, 3.1, 3.4]
' />


### uniq()
> Returns `ARRAY|ERROR`

Returns a copy of the array with deduplicated elements. Raises an error if a element is not hashable.


<CodeBlockSimple input='["a", 1, 1, 2].uniq()
' output='[1, 2, "a"]
' />


### yeet()
> Returns `STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE`

Removes the last element of the array and returns it.


<CodeBlockSimple input='a = [1,2,3]
a.yeet()
a
' output='[1, 2, 3]
3
[1, 2]
' />


### yoink(STRING|ARRAY|HASH|BOOLEAN|INTEGER|NIL|FUNCTION|FILE)
> Returns `NIL`

Adds the given object as last element to the array.


<CodeBlockSimple input='a = [1,2,3]
a.yoink("a")
a
' output='[1, 2, 3]
nil
[1, 2, 3, "a"]
' />



## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


<CodeBlockSimple input='"test".methods()
' output='["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "strip", "downcase"]
' />


### to_f()
> Returns `FLOAT`

If possible converts an object to its float representation. If not 0.0 is returned.


<CodeBlockSimple input='1.to_f()
"1.4".to_f()
nil.to_f()
' output='1.0
1.4
0.0
' />


### to_i(INTEGER)
> Returns `INTEGER`

If possible converts an object to its integer representation. If not 0 is returned.


<CodeBlockSimple input='true.to_i()
false.to_i()
1234.to_i()
"4".to_i()
"10011010010"to_i(2)
"2322".to_i(8)
"0x2322".to_i()
' output='1
0
1234
4
1234
1234
1234
' />


### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


<CodeBlockSimple input='a = {"test": 1234}
a.to_json()
' output='{"test": 1234}
"{\"test\":1234}"
' />


### to_s(INTEGER)
> Returns `STRING`

If possible converts an object to its string representation. If not empty string is returned.


<CodeBlockSimple input='true.to_s()
1234.to_s()
1234.to_s(2)
1234.to_s(8)
1234.to_s(10)
"test".to_s()
1.4.to_s()
' output='"true"
"1234"
"10011010010"
"2322"
"1234"
"test"
"1.4"
' />


### type()
> Returns `STRING`

Returns the type of the object.


<CodeBlockSimple input='"test".type()
' output='"STRING"
' />


### wat()
> Returns `STRING`

Returns the supported methods with usage information.


<CodeBlockSimple input='true.wat()
' output='"BOOLEAN supports the following methods:
  to_s()"
' />


