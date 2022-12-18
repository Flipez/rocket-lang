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


