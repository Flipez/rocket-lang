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

