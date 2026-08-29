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
> Returns `ARRAY`

Returns the character codes of the string, always as an array with one entry per character. A single-character string gives a one-element array and an empty string gives an empty array, so the result never has to be checked for its type before use.


<CodeBlockSimple input='"".ascii()
"a".ascii()
"abc".ascii()
' output='[]
[97]
[97, 98, 99]
' />


### capitalize()
> Returns `STRING`

Returns a copy with the first character upcased and every following character downcased, as Ruby's `capitalize` does. A capital in the middle of the string is therefore lost.


<CodeBlockSimple input='a = "hello World!"
a.capitalize()
a
' output='"hello World!"
"Hello world!"
"hello World!"
' />


### capitalize!()
> Returns `STRING`

Upcases the first character and downcases the rest in place, and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "hello World!"
a.capitalize!()
a
' output='"hello World!"
"Hello world!"
"Hello world!"
' />


### chomp(STRING)
> Returns `STRING`

Returns a copy with one trailing line ending removed: `\r\n`, `\n` or `\r`. Given a string it removes one trailing occurrence of that string instead. Given `""` it removes every trailing `\n` and `\r\n`, which is the way to drop blank lines at the end of a file.


<CodeBlockSimple input='a = "line\n"
a.chomp()
a.chomp() == "line"
"abcdd".chomp("d")
"a\n\n\n".chomp("")
' output='"line\n"
"line"
true
"abcd"
"a"
' />


### chomp!(STRING)
> Returns `STRING`

Removes one trailing line ending in place and returns the string, so calls can be chained. Takes the same optional separator as `chomp`.


<CodeBlockSimple input='a = "line\n"
a.chomp!()
a
' output='"line\n"
"line"
"line"
' />


### chop()
> Returns `STRING`

Returns a copy with the last character removed. A trailing `\r\n` is removed as a unit so a line ending is never left half there. Chopping an empty string gives an empty string rather than an error.


<CodeBlockSimple input='a = "abcd"
a.chop()
a
"".chop()
' output='"abcd"
"abc"
"abcd"
""
' />


### chop!()
> Returns `STRING`

Removes the last character in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "abcd"
a.chop!()
a
' output='"abcd"
"abc"
"abc"
' />


### count(STRING)
> Returns `INTEGER`

Counts how often a given substring occurs in the string.


<CodeBlockSimple input='"test".count("t")
' output='2
' />


### downcase()
> Returns `STRING`

Returns a copy with all uppercase letters replaced by their lowercase counterparts.


<CodeBlockSimple input='a = "TEST"
a.downcase()
a
' output='"TEST"
"test"
"TEST"
' />


### downcase!()
> Returns `STRING`

Replaces uppercase characters with their lowercase counterparts in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "TEST"
a.downcase!()
a
' output='"TEST"
"test"
"test"
' />


### empty?()
> Returns `BOOLEAN`

Returns `true` when the string has no characters.


<CodeBlockSimple input='"".empty?()
" ".empty?()
' output='true
false
' />


### end_with?(STRING)
> Returns `BOOLEAN`

Returns `true` when the string ends with any of the given strings.


<CodeBlockSimple input='"test.rl".end_with?(".rl")
"test.rl".end_with?(".go")
"test.rl".end_with?(".go", ".rl")
' output='true
false
true
' />


### find(STRING)
> Returns `INTEGER`

Returns the character index of a given string if found. Otherwise returns `-1`


<CodeBlockSimple input='"test".find("e")
' output='1
' />


### format(STRING|INTEGER|FLOAT|BOOLEAN|ARRAY|HASH)
> Returns `STRING`

Formats according to a format specifier and returns the resulting string


<CodeBlockSimple input='"%s is %d".format("a", 1)
' output='"a is 1"
' />


### include?(STRING)
> Returns `BOOLEAN`

Returns `true` when the string contains the given substring.


<CodeBlockSimple input='"test".include?("es")
"test".include?("xy")
' output='true
false
' />


### lines()
> Returns `ARRAY`

Splits the string at newline escape sequence and return all chunks in an array. Shorthand for `string.split("\n")`.


<CodeBlockSimple input='"a\nb".lines()
' output='["a", "b"]
' />


### lstrip()
> Returns `STRING`

Returns a copy with leading whitespace removed. See `rstrip` for the trailing end and `strip` for both.


<CodeBlockSimple input='a = "  test  "
a.lstrip()
a
' output='"  test  "
"test  "
"  test  "
' />


### lstrip!()
> Returns `STRING`

Removes leading whitespace in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "  test  "
a.lstrip!()
a
' output='"  test  "
"test  "
"test  "
' />


### replace(STRING, STRING)
> Returns `STRING`

Returns a copy with every occurrence of the first string replaced by the second. This is Ruby's `gsub` with plain strings rather than Ruby's `replace`.


<CodeBlockSimple input='a = "test"
a.replace("t", "f")
a
' output='"test"
"fesf"
"test"
' />


### replace!(STRING, STRING)
> Returns `STRING`

Replaces every occurrence of the first string with the second in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "test"
a.replace!("t", "f")
a
' output='"test"
"fesf"
"fesf"
' />


### reverse()
> Returns `STRING`

Returns a copy of the string with all characters reversed.


<CodeBlockSimple input='a = "stressed"
a.reverse()
a
' output='"stressed"
"desserts"
"stressed"
' />


### reverse!()
> Returns `STRING`

Reverses the characters in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "stressed"
a.reverse!()
a
' output='"stressed"
"desserts"
"desserts"
' />


### rstrip()
> Returns `STRING`

Returns a copy with trailing whitespace removed. See `lstrip` for the leading end and `strip` for both.


<CodeBlockSimple input='a = "  test  "
a.rstrip()
a
' output='"  test  "
"  test"
"  test  "
' />


### rstrip!()
> Returns `STRING`

Removes trailing whitespace in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "  test  "
a.rstrip!()
a
' output='"  test  "
"  test"
"  test"
' />


### size()
> Returns `INTEGER`

Returns the amount of characters in the string.


<CodeBlockSimple input='"test".size()
' output='4
' />


### split(STRING)
> Returns `ARRAY`

Splits the string on a given seperator and returns all the chunks in an array. Default seperator is `" "`


<CodeBlockSimple input='"a b".split()
"a-b".split("-")
' output='["a", "b"]
["a", "b"]
' />


### start_with?(STRING)
> Returns `BOOLEAN`

Returns `true` when the string starts with any of the given strings.


<CodeBlockSimple input='"test.rl".start_with?("test")
"test.rl".start_with?("prod")
"test.rl".start_with?("prod", "test")
' output='true
false
true
' />


### strip()
> Returns `STRING`

Returns a copy with leading and trailing whitespace removed.


<CodeBlockSimple input='a = "  test  "
a.strip()
a
' output='"  test  "
"test"
"  test  "
' />


### strip!()
> Returns `STRING`

Removes leading and trailing whitespace in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "  test  "
a.strip!()
a
' output='"  test  "
"test"
"test"
' />


### swapcase()
> Returns `STRING`

Returns a copy with every uppercase character downcased and every lowercase character upcased.


<CodeBlockSimple input='a = "Hello World"
a.swapcase()
a
' output='"Hello World"
"hELLO wORLD"
"Hello World"
' />


### swapcase!()
> Returns `STRING`

Swaps the case of every character in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "Hello World"
a.swapcase!()
a
' output='"Hello World"
"hELLO wORLD"
"hELLO wORLD"
' />


### upcase()
> Returns `STRING`

Returns a copy with all lowercase letters replaced by their uppercase counterparts.


<CodeBlockSimple input='a = "test"
a.upcase()
a
' output='"test"
"TEST"
"test"
' />


### upcase!()
> Returns `STRING`

Replaces lowercase characters with their uppercase counterparts in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "test"
a.upcase!()
a
' output='"test"
"TEST"
"TEST"
' />



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


