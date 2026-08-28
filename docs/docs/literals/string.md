import CodeBlockSimple from '@site/components/CodeBlockSimple'

# String

Strings can be written with double or single quotes, and the two differ in
how they treat escapes.

A **double-quoted** string processes escape sequences: `\"` for a quote,
`\n` for a newline, `\t` for a tab, `\r` for a carriage return.

```js
puts("test\"string")   // test"string
puts("a\tb")           // a<tab>b
```

A **single-quoted** string is raw: nothing is escaped and a backslash is an
ordinary character. This makes it convenient for text containing double
quotes.

```js
puts('test "string"')  // test "string"
puts('a\tb')           // a\tb, a literal backslash and t
```

Because a single-quoted string performs no escaping, it cannot contain a
single quote at all -- `'test \'string'` is a parse error. Use a
double-quoted string when you need one.



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

### ascii()
> Returns `INTEGER|ARRAY`

Returns the ascii representation of a char or string


<CodeBlockSimple input='"a".ascii()
"abc".ascii()
' output='97
[97, 98, 99]
' />


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

Returns the names of the methods specific to this literal type, not including the generic methods listed on this page. The order is unspecified, so sort the result if you need it stable. A type with no methods of its own returns an empty array.


<CodeBlockSimple input='"test".methods().sort()
true.methods()
' output='["ascii", "count", "downcase", "downcase!", "find", "format", "lines", "replace", "reverse", "reverse!", "size", "split", "strip", "strip!", "upcase", "upcase!"]
[]
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

Returns the type's literal-specific methods with their usage information, as a single string. Types with no methods of their own list none.


<CodeBlockSimple input='true.wat()
' output='"BOOLEAN supports the following methods:\n"
' />


