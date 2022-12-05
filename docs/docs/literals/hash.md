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


