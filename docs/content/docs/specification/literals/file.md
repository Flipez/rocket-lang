# File




## Literal Specific Methods

### close()
> Returns `BOOLEAN`

Closes the file pointer. Returns always `true`.



### lines()
> Returns `ARRAY|NULL`

If successfull, returns all lines of the file as array elements, otherwise `null`.



### read()
> Returns `STRING`

Reads content of the file and returns it.



### rewind()
> Returns `BOOLEAN`

Resets the read pointer back to position `0`. Always returns `true`.



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
