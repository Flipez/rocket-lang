---
title: "String"
menu:
  docs:
    parent: "literals"
---
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
```

## Literal Specific Methods

### count(STRING|INTEGER)
> Returns `INTEGER`

Counts how often a given string or integer occurs in the string. Converts given integers to strings automatically.


```js
🚀 > "test".count("t")
=> 2
🚀 > "test".count("f")
=> 0
🚀 > "test1".count("1")
=> 1
🚀 > "test1".count(1)
=> 1
```


### downcase()
> Returns `STRING`

Returns the string with all uppercase letters replaced with lowercase counterparts.


```js
🚀 > "TeST".downcase()
=> test
```


### downcase!()
> Returns `NULL`

Replaces all upcase characters with lowercase counterparts.


```js

🚀 > a = "TeST"
=> TeST
🚀 > a.downcase!()
=> null
🚀 > a
=> test
```


### find(STRING|INTEGER)
> Returns `INTEGER`

Returns the character index of a given string if found. Otherwise returns `-1`


```js
🚀 > "test".find("e")
=> 1
🚀 > "test".find("f")
=> -1
```


### lines()
> Returns `ARRAY`

Splits the string at newline escape sequence and return all chunks in an array. Shorthand for `string.split("\n")`.


```js
🚀 > "test\ntest2".lines()
=> ["test", "test2"]
```


### plz_i(INTEGER)
> Returns `INTEGER`

Interprets the string as an integer with an optional given base. The default base is `10` and switched to `8` if the string starts with `0x`.


```js
🚀 > "1234".plz_i()
=> 1234

🚀 > "1234".plz_i(8)
=> 668

🚀 > "0x1234".plz_i(8)
=> 668

🚀 > "0x1234".plz_i()
=> 668

🚀 > "0x1234".plz_i(10)
=> 0
```


### replace(STRING, STRING)
> Returns `STRING`

Replaces the first string with the second string in the given string.


```js
🚀 > "test".replace("t", "f")
=> "fesf"
```


### reverse()
> Returns `STRING`

Returns a copy of the string with all characters reversed.


```js
🚀 > "stressed".reverse()
=> "desserts"
```


### reverse!()
> Returns `NULL`

Replaces all the characters in a string in reverse order.


```js
🚀 > a = "stressed"
=> "stressed"
🚀 > a.reverse!()
=> null
🚀 > a
=> "desserts"
```


### size()
> Returns `INTEGER`

Returns the amount of characters in the string.


```js
🚀 > "test".size()
=> 4
```


### split(STRING)
> Returns `ARRAY`

Splits the string on a given seperator and returns all the chunks in an array. Default seperator is `" "`


```js
🚀 > "a,b,c,d".split(",")
=> ["a", "b", "c", "d"]

🚀 > "test and another test".split()
=> ["test", "and", "another", "test"]
```


### strip()
> Returns `STRING`

Returns a copy of the string with all leading and trailing whitespaces removed.


```js
🚀 > " test ".strip()
=> "test"
```


### strip!()
> Returns `NULL`

Removes all leading and trailing whitespaces in the string.


```js

🚀 > a = " test "
=> " test "
🚀 > a.strip!()
=> null
🚀 > a
=> "test"
```


### upcase()
> Returns `STRING`

Returns the string with all lowercase letters replaced with uppercase counterparts.


```js
🚀 > "test".upcase()
=> TEST
```


### upcase!()
> Returns `NULL`

Replaces all lowercase characters with upcase counterparts.


```js

🚀 > a = "test"
=> test
🚀 > a.upcase!()
=> null
🚀 > a
=> TEST
```



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
