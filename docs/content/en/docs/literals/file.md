---
title: "File"
menu:
  docs:
    parent: "literals"
---




## Literal Specific Methods

### close()
> Returns `BOOLEAN`

Closes the file pointer. Returns always `true`.



### content()
> Returns `STRING|ERROR`

Reads content of the file and returns it. Resets the position to 0 after read.



### lines()
> Returns `ARRAY|ERROR`

If successfull, returns all lines of the file as array elements, otherwise `null`. Resets the position to 0 after read.



### position()
> Returns `INTEGER`

Returns the position of the current file handle. -1 if the file is closed.



### read(INTEGER)
> Returns `STRING`

Reads the given amount of bytes from the file.



### seek(INTEGER, INTEGER)
> Returns `BOOLEAN`

Seeks the file handle relative from the given position.



### write(STRING)
> Returns `BOOLEAN|NULL`

Writes the given string to the file. Returns `true` on success, `false` on failure and `null` if pointer is invalid.




## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.

```js
🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]
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
