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

### clear()
> Returns `HASH`

Removes every entry and returns the hash, so calls can be chained.


<CodeBlockSimple input='h = {"a": 1}
h.clear()
h.size()
' output='{"a": 1}
{}
0
' />


### compact()
> Returns `HASH|ERROR`

Returns a new hash without the entries whose value is `nil`. The hash itself is unchanged; use `compact!` to remove them in place.


<CodeBlockSimple input='h = {"a": nil}
h.compact().size()
h.size()
' output='{"a": nil}
0
1
' />


### compact!()
> Returns `HASH|ERROR`

Removes the entries whose value is `nil` in place and returns the hash, so calls can be chained.


<CodeBlockSimple input='h = {"a": nil}
h.compact!().size()
h.size()
' output='{"a": nil}
0
0
' />


### delete(INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL)
> Returns `INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL`

Removes the entry for a key and returns its value, or `nil` when the key was not there, so a removal can be told from a miss.


<CodeBlockSimple input='h = {"a": 1}
h.delete("a")
h.size()
h.delete("a")
' output='{"a": 1}
1
0
nil
' />


### empty?()
> Returns `BOOLEAN`

Returns `true` when the hash has no entries.


<CodeBlockSimple input='{}.empty?()
{"a": 1}.empty?()
' output='true
false
' />


### fetch(INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL, [INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL])
> Returns `INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL`

Returns the value for a key. Without a fallback a missing key is an error, which is the difference from `get`, where a fallback is required.


<CodeBlockSimple input='h = {"a": 1}
h.fetch("a")
h.fetch("z", 0)
' output='{"a": 1}
1
0
' />


### get(INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL, INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL)
> Returns `INTEGER|STRING|BOOLEAN|ARRAY|HASH|MATRIX|FLOAT|ERROR|NIL`

Returns the value of the given key or the default


<CodeBlockSimple input='{"a": "1", "b": "2"}.get("a", 10)
{"a": "1", "b": "2"}.get("c", 10)
' output='"1"
10
' />


### include?(BOOLEAN|STRING|INTEGER|FLOAT|ARRAY|HASH)
> Returns `BOOLEAN`

Returns true or false wether the hash contains the given object as key


<CodeBlockSimple input='{"a": 1, 1: "b"}.include?(1)
{"a": 1, 1: "b"}.include?("c")
' output='true
false
' />


### invert()
> Returns `HASH|ERROR`

Returns a new hash with keys and values swapped. Values that repeat collapse into one entry, and which one survives is not defined. A value that cannot be a key is an error.


<CodeBlockSimple input='h = {"a": 1}
h.invert().get(1, "missing")
' output='{"a": 1}
"a"
' />


### keys()
> Returns `ARRAY`

Returns the keys of the hash. The order is **not** defined and can differ between runs of the same program, so sort the result if you need it stable.


<CodeBlockSimple input='{"a": "1"}.keys()
{"a": "1", "b": "2"}.keys().sort()
' output='["a"]
["a", "b"]
' />


### merge(HASH)
> Returns `HASH|ERROR`

Returns a new hash with the entries of both. The argument wins where a key is in both. The hash itself is unchanged; use `merge!` to merge in place.


<CodeBlockSimple input='h = {"a": 1}
h.merge({"b": 2}).size()
h.size()
h.merge({"a": 9}).get("a", 0)
' output='{"a": 1}
2
1
9
' />


### merge!(HASH)
> Returns `HASH|ERROR`

Merges the given hash in place and returns the hash, so calls can be chained.


<CodeBlockSimple input='h = {"a": 1}
h.merge!({"b": 2}).size()
h.size()
' output='{"a": 1}
2
2
' />


### size()
> Returns `INTEGER`

Returns the number of entries.


<CodeBlockSimple input='h = {"a": 1}
h.size()
{}.size()
' output='{"a": 1}
1
0
' />


### values()
> Returns `ARRAY`

Returns the values of the hash, in the same unspecified order as `keys`.


<CodeBlockSimple input='{"a": "1"}.values()
{"a": "1", "b": "2"}.values().sort()
' output='["1"]
["1", "2"]
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
> Returns `NIL`

Prints the type's literal-specific methods with their argument and return types, sorted by name, one per line. It returns `nil` rather than the listing: this exists to be read, and the REPL echoes a returned value through its escaped representation, which would put the whole thing on one line. Use `methods` when the names are wanted as data. A type with no methods of its own prints only the heading.


<CodeBlockSimple input='true.wat()
1.0.wat()
' output='BOOLEAN supports the following methods:
nil
FLOAT supports the following methods:
	abs()
	ceil([INTEGER])
	divmod(FLOAT)
	finite?()
	floor([INTEGER])
	infinite?()
	nan?()
	negative?()
	positive?()
	round([INTEGER])
	truncate([INTEGER])
	zero?()
nil
' />


