import CodeBlockSimple from '@site/components/CodeBlockSimple'

# String

Strings can be written with double or single quotes, and the two differ in
how they treat escapes.

A **double-quoted** string processes escape sequences: `\"` for a quote,
`\n` for a newline, `\t` for a tab, `\r` for a carriage return.

```js
print("test\"string")   # test"string
print("a\tb")           # a<tab>b
```

A **single-quoted** string is raw: nothing is escaped and a backslash is an
ordinary character. This makes it convenient for text containing double
quotes.

```js
print('test "string"')  # test "string"
print('a\tb')           # a\tb, a literal backslash and t
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
print(s[2])
print(s[-2])
print(s[:2])
print(s[:-2])
print(s[2:])
print(s[-2:])
print(s[1:-2])

s[2] = "C"
s[-2] = "E"
print(s)

# should output
"c"
"e"
"ab"
"abcd"
"cdef"
"ef"
"bcd"
"abCdEf"

# you can also use single quotes
'test "string" with doublequotes'

# and you can scape a double quote in a double quote string
"te\"st" == 'te"st'

```

## Literal Specific Methods

### capitalize()
> Returns `STRING`

Returns a copy with the first character replaced by its uppercase counterpart and every other character replaced by its lowercase counterpart, as Ruby's `capitalize` does. A capital in the middle of the string is therefore lost.


<CodeBlockSimple input='a = "hello World!"
a.capitalize()
a
' output='"hello World!"
"Hello world!"
"hello World!"
' />


### capitalize!()
> Returns `STRING`

Replaces the first character with its uppercase counterpart and every other character with its lowercase counterpart, in place, and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "hello World!"
a.capitalize!()
a
' output='"hello World!"
"Hello world!"
"Hello world!"
' />


### codepoints()
> Returns `ARRAY`

Returns the character codes of the string, always as an array with one entry per character. A single-character string gives a one-element array and an empty string gives an empty array, so the result never has to be checked for its type before use.


<CodeBlockSimple input='"".codepoints()
"a".codepoints()
"abc".codepoints()
' output='[]
[97]
[97, 98, 99]
' />


### contains?(STRING)
> Returns `BOOLEAN`

Returns `true` when the string contains the given substring.


<CodeBlockSimple input='"test".contains?("es")
"test".contains?("xy")
' output='true
false
' />


### count(STRING)
> Returns `INTEGER`

Counts how often a given substring occurs in the string.


<CodeBlockSimple input='"test".count("t")
' output='2
' />


### empty?()
> Returns `BOOLEAN`

Returns `true` when the string has no characters.


<CodeBlockSimple input='"".empty?()
" ".empty?()
' output='true
false
' />


### ends_with?(STRING...)
> Returns `BOOLEAN`

Returns `true` when the string ends with any of the given strings.


<CodeBlockSimple input='"test.rl".ends_with?(".rl")
"test.rl".ends_with?(".go")
"test.rl".ends_with?(".go", ".rl")
' output='true
false
true
' />


### format(ANY...)
> Returns `STRING`

Formats according to a format specifier and returns the resulting string


<CodeBlockSimple input='"%s is %d".format("a", 1)
' output='"a is 1"
' />


### index_of(STRING)
> Returns `INTEGER`

Returns the character index of the first occurrence of a given string if found. Otherwise returns `-1`.


<CodeBlockSimple input='"test".index_of("e")
' output='1
' />


### last_index_of(STRING)
> Returns `INTEGER`

Returns the character index of the last occurrence of a given string if found. Otherwise returns `-1`.


<CodeBlockSimple input='"hello".last_index_of("l")
"hello".last_index_of("z")
' output='3
-1
' />


### lines()
> Returns `ARRAY`

Splits the string at newline escape sequence and return all chunks in an array. Shorthand for `string.split("\n")`.


<CodeBlockSimple input='"a\nb".lines()
' output='["a", "b"]
' />


### lowercase()
> Returns `STRING`

Returns a copy with all uppercase letters replaced by their lowercase counterparts.


<CodeBlockSimple input='a = "TEST"
a.lowercase()
a
' output='"TEST"
"test"
"TEST"
' />


### lowercase!()
> Returns `STRING`

Replaces uppercase characters with their lowercase counterparts in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "TEST"
a.lowercase!()
a
' output='"TEST"
"test"
"test"
' />


### remove_last()
> Returns `STRING`

Returns a copy with the last character removed. A trailing `\r\n` is removed as a unit so a line ending is never left half there. Removing the last character of an empty string gives an empty string rather than an error.


<CodeBlockSimple input='a = "abcd"
a.remove_last()
a
"".remove_last()
' output='"abcd"
"abc"
"abcd"
""
' />


### remove_last!()
> Returns `STRING`

Removes the last character in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "abcd"
a.remove_last!()
a
' output='"abcd"
"abc"
"abc"
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


### size()
> Returns `INTEGER`

Returns the amount of characters in the string.


<CodeBlockSimple input='"test".size()
' output='4
' />


### split([STRING])
> Returns `ARRAY`

Splits the string on a given seperator and returns all the chunks in an array. Default seperator is `" "`


<CodeBlockSimple input='"a b".split()
"a-b".split("-")
' output='["a", "b"]
["a", "b"]
' />


### starts_with?(STRING...)
> Returns `BOOLEAN`

Returns `true` when the string starts with any of the given strings.


<CodeBlockSimple input='"test.rl".starts_with?("test")
"test.rl".starts_with?("prod")
"test.rl".starts_with?("prod", "test")
' output='true
false
true
' />


### swap_case()
> Returns `STRING`

Returns a copy with every uppercase character lowercased and every lowercase character uppercased.


<CodeBlockSimple input='a = "Hello World"
a.swap_case()
a
' output='"Hello World"
"hELLO wORLD"
"Hello World"
' />


### swap_case!()
> Returns `STRING`

Swaps the case of every character in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "Hello World"
a.swap_case!()
a
' output='"Hello World"
"hELLO wORLD"
"hELLO wORLD"
' />


### trim()
> Returns `STRING`

Returns a copy with leading and trailing whitespace removed.


<CodeBlockSimple input='a = "  test  "
a.trim()
a
' output='"  test  "
"test"
"  test  "
' />


### trim!()
> Returns `STRING`

Removes leading and trailing whitespace in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "  test  "
a.trim!()
a
' output='"  test  "
"test"
"test"
' />


### trim_end()
> Returns `STRING`

Returns a copy with trailing whitespace removed. See `trim_start` for the leading end and `trim` for both.


<CodeBlockSimple input='a = "  test  "
a.trim_end()
a
' output='"  test  "
"  test"
"  test  "
' />


### trim_end!()
> Returns `STRING`

Removes trailing whitespace in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "  test  "
a.trim_end!()
a
' output='"  test  "
"  test"
"  test"
' />


### trim_line_end([STRING])
> Returns `STRING`

Returns a copy with one trailing line ending removed: `\r\n`, `\n` or `\r`. Given a string it removes one trailing occurrence of that string instead. Given `""` it removes every trailing `\n` and `\r\n`, which is the way to drop blank lines at the end of a file.


<CodeBlockSimple input='a = "line\n"
a.trim_line_end()
a.trim_line_end() == "line"
"abcdd".trim_line_end("d")
"a\n\n\n".trim_line_end("")
' output='"line\n"
"line"
true
"abcd"
"a"
' />


### trim_line_end!([STRING])
> Returns `STRING`

Removes one trailing line ending in place and returns the string, so calls can be chained. Takes the same optional separator as `trim_line_end`.


<CodeBlockSimple input='a = "line\n"
a.trim_line_end!()
a
' output='"line\n"
"line"
"line"
' />


### trim_start()
> Returns `STRING`

Returns a copy with leading whitespace removed. See `trim_end` for the trailing end and `trim` for both.


<CodeBlockSimple input='a = "  test  "
a.trim_start()
a
' output='"  test  "
"test  "
"  test  "
' />


### trim_start!()
> Returns `STRING`

Removes leading whitespace in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "  test  "
a.trim_start!()
a
' output='"  test  "
"test  "
"test  "
' />


### uppercase()
> Returns `STRING`

Returns a copy with all lowercase letters replaced by their uppercase counterparts.


<CodeBlockSimple input='a = "test"
a.uppercase()
a
' output='"test"
"TEST"
"test"
' />


### uppercase!()
> Returns `STRING`

Replaces lowercase characters with their uppercase counterparts in place and returns the string, so calls can be chained.


<CodeBlockSimple input='a = "test"
a.uppercase!()
a
' output='"test"
"TEST"
"TEST"
' />



## Generic Literal Methods

### help()
> Returns `NIL`

Prints the type's literal-specific methods with their argument and return types, sorted by name, one per line. It returns `nil` rather than the listing: this exists to be read, and the REPL echoes a returned value through its escaped representation, which would put the whole thing on one line. Use `methods` when the names are wanted as data. A type with no methods of its own prints only the heading.


<CodeBlockSimple input='true.help()
1.0.help()
' output='BOOLEAN supports the following methods:
nil
FLOAT supports the following methods:
	abs()
	acos()
	asin()
	atan()
	ceil([INTEGER])
	copysign(NUMERIC)
	cos()
	divmod(FLOAT)
	exp()
	finite?()
	floor([INTEGER])
	infinite?()
	log()
	log10()
	log2()
	nan?()
	negative?()
	positive?()
	pow(NUMERIC)
	remainder(NUMERIC)
	round([INTEGER])
	sin()
	sqrt()
	tan()
	truncate([INTEGER])
	zero?()
nil
' />


### is_a?(STRING)
> Returns `BOOLEAN|ERROR`

Returns `true` when the value is of the given type or belongs to the given type group, so one question covers both. The name has to be one that exists: anything else is an error rather than a `false`, because a typo would otherwise read as a real answer. See [Types and type groups](../language/types).


<CodeBlockSimple input='nil.is_a?("NIL")
"a".is_a?("HASHABLE")
nil.is_a?("HASHABLE")
true.is_a?("INTEGERABLE")
"a".is_a?("INTEGER")
' output='true
true
false
true
false
' />


### methods()
> Returns `ARRAY`

Returns the names of the methods specific to this literal type, not including the generic methods listed on this page. The names are sorted, so the result is the same on every run. A type with no methods of its own returns an empty array.


<CodeBlockSimple input='1.0.methods().contains?("round")
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


### to_float()
> Returns `FLOAT|NIL`

Converts an object to its float representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0.0`.


<CodeBlockSimple input='1.to_float()
"1.4".to_float()
"abc".to_float()
nil.to_float()
' output='1.0
1.4
nil
nil
' />


### to_integer()
> Returns `INTEGER|NIL`

Converts an object to its integer representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0`. For strings a `0b`, `0o` or `0x` prefix selects binary, octal or hexadecimal and is matched case insensitively, a leading zero followed only by octal digits is octal, and anything else is decimal. The resulting integer keeps the base it was parsed with, and integers of differing bases cannot be combined directly.


<CodeBlockSimple input='true.to_integer()
false.to_integer()
1234.to_integer()
"4".to_integer()
"0".to_integer()
"0125".to_integer()
"0x2322".to_integer()
"0b1010".to_integer()
"test".to_integer()
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


### to_string()
> Returns `STRING`

Converts an object to its string representation, or the empty string when it has none. Takes no arguments; an integer renders in its own base, so use `to_base` first to change it.


<CodeBlockSimple input='true.to_string()
1234.to_string()
"test".to_string()
1.4.to_string()
nil.to_string()
"0b1010".to_integer().to_string()
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


### type_groups()
> Returns `ARRAY`

Returns the type groups the value belongs to, sorted. `ANY` is not listed: every value belongs to it, so it would say nothing while prefixing every answer. It exists for signatures, where `append(ANY)` means the argument accepts anything, and `is_a?("ANY")` still answers `true`. See [Types and type groups](../language/types) for what each group means.


<CodeBlockSimple input='1.type_groups()
nil.type_groups()
def() end.type_groups()
print.type_groups()
' output='["COMPARABLE", "HASHABLE", "INTEGERABLE", "NUMERIC", "STRINGABLE"]
["STRINGABLE"]
["CALLABLE"]
["CALLABLE"]
' />


### wat()
> Returns `NIL`

An alias of `help`, kept as an easter egg. This is the only alias in
RocketLang; every other method has exactly one name.



<CodeBlockSimple input='true.wat()
1.0.wat()
' output='BOOLEAN supports the following methods:
nil
FLOAT supports the following methods:
	abs()
	acos()
	asin()
	atan()
	ceil([INTEGER])
	copysign(NUMERIC)
	cos()
	divmod(FLOAT)
	exp()
	finite?()
	floor([INTEGER])
	infinite?()
	log()
	log10()
	log2()
	nan?()
	negative?()
	positive?()
	pow(NUMERIC)
	remainder(NUMERIC)
	round([INTEGER])
	sin()
	sqrt()
	tan()
	truncate([INTEGER])
	zero?()
nil
' />


