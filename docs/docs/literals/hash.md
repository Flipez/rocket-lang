import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Hash

Hash keys are not restricted to strings. Any hashable value works, and a
single hash may mix key types freely: `STRING`, `INTEGER`, `FLOAT`,
`BOOLEAN`, `ARRAY` and `HASH`.

```js
h = {"a": 1, 2: true, 3.5: "float", true: "bool", [1, 2]: "array"}

puts(h["a"])     // 1
puts(h[2])       // true
puts(h[3.5])     // "float"
puts(h[true])    // "bool"
puts(h[[1, 2]])  // "array"
```

`NIL` and functions are not hashable and are rejected as keys, with
`unusable as hash key: NIL` and `expected index to be hashable`
respectively.

Reading a key that is not present returns `nil`.



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

### get(INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL, INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL)
> Returns `INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL`

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


