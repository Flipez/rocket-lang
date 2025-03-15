import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Hash




```js
people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];

// reassign of values
h = {"a": 1, 2: true}
puts(h["a"])
puts(h[2])
h["a"] = 3
h["b"] = "moo"
puts(h["a"])
puts(h["b"])
puts(h[2])h = {"a": 1, 2: true}
puts(h["a"])
puts(h[2])
h["a"] = 3
h["b"] = "moo"

// should output
1
true
3
"moo"
true

```

## Literal Specific Methods

### get(INTEGER|STRING|BOOLEAN|ARRAY|HASH|FLOAT|ERROR|NIL, INTEGER|STRING|BOOLEAN|ARRAY|HASH|FLOAT|ERROR|NIL)
> Returns `INTEGER|STRING|BOOLEAN|ARRAY|HASH|FLOAT|ERROR|NIL`

Returns the value of the given key or the default


<CodeBlockSimple input='{"a": "1", "b": "2"}.get("a", 10)
{"a": "1", "b": "2"}.get("c", 10)
' output='1
10
' />


### include?(BOOLEAN|STRING|INTEGER|FLOAT|ARRAY|HASH)
> Returns `BOOLEAN`

Returns true or false wether the hash contains the given object as key


<CodeBlockSimple input='{"a": 1, 1: "b"}.include?(1)
{"a": 1, 1: "b"}.include?("c")
' output='true false' />


### keys()
> Returns `ARRAY`

Returns the keys of the hash.


<CodeBlockSimple input='{"a": "1", "b": "2"}.keys()
' output='["a", "b"]
' />


### values()
> Returns `ARRAY`

Returns the values of the hash.


<CodeBlockSimple input='{"a": "1", "b": "2"}.values()
' output='["1", "2"]
' />



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


