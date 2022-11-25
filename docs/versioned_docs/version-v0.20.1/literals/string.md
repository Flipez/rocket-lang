import CodeBlockSimple from '@site/components/CodeBlockSimple'

# String




```js
a = "test_string";

b = "test" + "_string";

is_true = "test" == "test";
is_false = "test" == "string";

s = "abcdef"
puts(s[2])
puts(s[-2])
puts(s[:2])
puts(s[:-2])
puts(s[2:])
puts(s[-2:])
puts(s[1:-2])

s[2] = "C"
s[-2] = "E"
puts(s)

// should output
"c"
"e"
"ab"
"abcd"
"cdef"
"ef"
"bcd"
"abCdEf"

// you can also use single quotes
'test "string" with doublequotes'

// and you can scape a double quote in a double quote string
"te\"st" == 'te"st'

```

## Literal Specific Methods

### count(STRING)
> Returns `INTEGER`

Counts how often a given substring occurs in the string.





### downcase()
> Returns `STRING`

Returns the string with all uppercase letters replaced with lowercase counterparts.





### downcase!()
> Returns `NIL`

Replaces all upcase characters with lowercase counterparts.





### find(STRING)
> Returns `INTEGER`

Returns the character index of a given string if found. Otherwise returns `-1`





### format(STRING|INTEGER|FLOAT|BOOLEAN|ARRAY|HASH)
> Returns `STRING`

Formats according to a format specifier and returns the resulting string





### lines()
> Returns `ARRAY`

Splits the string at newline escape sequence and return all chunks in an array. Shorthand for `string.split("\n")`.





### plz_i(INTEGER)
> Returns `INTEGER`

Interprets the string as an integer with an optional given base. The default base is `10` and switched to `8` if the string starts with `0x`.





### replace(STRING, STRING)
> Returns `STRING`

Replaces the first string with the second string in the given string.





### reverse()
> Returns `STRING`

Returns a copy of the string with all characters reversed.





### reverse!()
> Returns `NIL`

Replaces all the characters in a string in reverse order.





### size()
> Returns `INTEGER`

Returns the amount of characters in the string.





### split(STRING)
> Returns `ARRAY`

Splits the string on a given seperator and returns all the chunks in an array. Default seperator is `" "`





### strip()
> Returns `STRING`

Returns a copy of the string with all leading and trailing whitespaces removed.





### strip!()
> Returns `NIL`

Removes all leading and trailing whitespaces in the string.





### upcase()
> Returns `STRING`

Returns the string with all lowercase letters replaced with uppercase counterparts.





### upcase!()
> Returns `NIL`

Replaces all lowercase characters with upcase counterparts.






## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


<CodeBlockSimple input='"test".methods()
' output='["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "plz_i", "strip", "downcase"]
' />


### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


<CodeBlockSimple input='a = {"test": 1234}
a.to_json()
' output='{"test": 1234}
"{\"test\":1234}"
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
  plz_s()"
' />


