import CodeBlockSimple from '@site/components/CodeBlockSimple'

# File




```js
input = open("examples/aoc/2021/day-1/input").lines()

```

## Literal Specific Methods

### close()
> Returns `BOOLEAN`

Closes the file pointer. Returns always `true`.





### content()
> Returns `STRING|ERROR`

Reads content of the file and returns it. Resets the position to 0 after read.





### lines()
> Returns `ARRAY|ERROR`

If successfull, returns all lines of the file as array elements, otherwise `nil`. Resets the position to 0 after read.





### position()
> Returns `INTEGER`

Returns the position of the current file handle. -1 if the file is closed.





### read(INTEGER)
> Returns `STRING|ERROR`

Reads the given amount of bytes from the file. Sets the position to the bytes that where actually read. At the end of file EOF error is returned.





### seek(INTEGER, INTEGER)
> Returns `INTEGER|ERROR`

Seek sets the offset for the next Read or Write on file to offset, interpreted according to whence. 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end.





### write(STRING)
> Returns `INTEGER|ERROR`

Writes the given string to the file. Returns number of written bytes on success.






## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns the names of the methods specific to this literal type, not including the generic methods listed on this page. The names are sorted, so the result is the same on every run. A type with no methods of its own returns an empty array.


<CodeBlockSimple input='1.0.methods().include?("round")
true.methods()
' output='true
[]
' />


### nil?()
> Returns `BOOLEAN`

Returns `true` only for `nil`. Reads better than comparing against `nil` at the end of a chain, and every type answers it.


<CodeBlockSimple input='nil.nil?()
1.nil?()
"".nil?()
[].first().nil?()
' output='true
false
false
true
' />


### to_f()
> Returns `FLOAT|NIL`

Converts an object to its float representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0.0`.


<CodeBlockSimple input='1.to_f()
"1.4".to_f()
"abc".to_f()
nil.to_f()
' output='1.0
1.4
nil
nil
' />


### to_i()
> Returns `INTEGER|NIL`

Converts an object to its integer representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0`. For strings a `0b`, `0o` or `0x` prefix selects binary, octal or hexadecimal and is matched case insensitively, a leading zero followed only by octal digits is octal, and anything else is decimal. The resulting integer keeps the base it was parsed with, and integers of differing bases cannot be combined directly.


<CodeBlockSimple input='true.to_i()
false.to_i()
1234.to_i()
"4".to_i()
"0".to_i()
"0125".to_i()
"0x2322".to_i()
"0b1010".to_i()
"test".to_i()
' output='1
0
1234
4
0
0o125
0x2322
0b1010
nil
' />


### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


<CodeBlockSimple input='a = {"test": 1234}
a.to_json()
' output='{"test": 1234}
"{\"test\":1234}"
' />


### to_s()
> Returns `STRING`

Converts an object to its string representation, or the empty string when it has none. Takes no arguments; an integer renders in its own base, so use `to_base` first to change it.


<CodeBlockSimple input='true.to_s()
1234.to_s()
"test".to_s()
1.4.to_s()
nil.to_s()
"0b1010".to_i().to_s()
' output='"true"
"1234"
"test"
"1.4"
""
"0b1010"
' />


### type()
> Returns `STRING`

Returns the type of the object.


<CodeBlockSimple input='"test".type()
' output='"STRING"
' />


### wat()
> Returns `STRING`

Returns the type's literal-specific methods with their usage information, as a single string, sorted by name. Types with no methods of their own list none.


<CodeBlockSimple input='true.wat()
' output='"BOOLEAN supports the following methods:\n"
' />


