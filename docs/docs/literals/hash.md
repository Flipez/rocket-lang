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

## Calling a value stored in a hash

> 👉 Introduced in `0.24`

A value under a name can be read with a dot, and a **callable** value under a
name can be *called* with one:

```js
h = {"double": def(x) return x * 2 end}

puts(h.double(21))     // 42
puts(h["double"](21))  // the same thing
```

That makes a hash of functions read like the object it already is. The functions
close over the locals of whatever built the hash, so the state is private and
each one is independent:

```js
def new_account(owner, balance)
  return {
    "owner":    owner,
    "deposit":  def(n) balance = balance + n return balance end,
    "describe": def() return owner + ": " + balance.to_string() end
  }
end

a = new_account("robert", 100)
b = new_account("someone", 0)

puts(a.deposit(50))   // 150
puts(a.describe())    // "robert: 150"
puts(b.describe())    // "someone: 0"
puts(a.owner)         // "robert" -- plain data, read the same way
```

A real hash method always wins, so a hash of data cannot take over `size` or
`keys`. The stored value is still reachable by index:

```js
h = {"size": def() return 99 end}

puts(h.size())      // 1  -- the hash method
puts(h["size"]())   // 99 -- the stored function
```

A name holding something that cannot be called says so, and a name that is not
there at all reports a missing method as before:

```js
h = {"n": 1}

h.n()      // ERROR: `n` is not callable for HASH, it is INTEGER
h.other()  // ERROR: undefined method `.other()` for HASH
```

This is not a class. The hash is still a `HASH`, so `type()` says `HASH`,
`methods()` lists the hash's own methods rather than the stored ones, and there
is no `self` or inheritance. What it adds is the call syntax for a pattern the
language could already express.



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


### each(CALLABLE)
> Returns `HASH|ERROR`

Calls the callback once per entry with the key and the value, and returns the hash so calls can be chained. The order the entries arrive in is **not** defined and differs between runs, the same caveat `keys` carries -- use it for a side effect per entry, not to build something ordered.


<CodeBlockSimple input='h = {"a": 1}
h.each(def(key, value) puts(key + "=" + value.to_string()) end)
' output='{"a": 1}
a=1
{"a": 1}
' />


### empty?()
> Returns `BOOLEAN`

Returns `true` when the hash has no entries.


<CodeBlockSimple input='{}.empty?()
{"a": 1}.empty?()
' output='true
false
' />


### fetch(HASHABLE, [ANY])
> Returns `ANY`

Returns the value for `key`, and raises when the key is absent. Use
`get` when a missing key is an expected case.



<CodeBlockSimple input='h = {"a": 1}
h.fetch("a")
' output='{"a": 1}
1
' />


### filter(CALLABLE)
> Returns `HASH|ERROR`

Returns a new hash of the entries the callback said yes to. The callback receives the key and the value. Only `false` and `nil` are no. Use `filter!` to filter in place.


<CodeBlockSimple input='h = {"a": 1}
h["b"] = 2
h.filter(def(k, v) v > 1 end).size()
h.size()
' output='{"a": 1}
2
1
2
' />


### filter!(CALLABLE)
> Returns `HASH|ERROR`

Keeps only the entries the callback said yes to and returns the hash, so calls can be chained.


<CodeBlockSimple input='h = {"a": 1}
h["b"] = 2
h.filter!(def(k, v) v > 1 end).size()
h.size()
' output='{"a": 1}
2
1
1
' />


### get(HASHABLE, [ANY])
> Returns `ANY`

Returns the value for `key`, or `nil` when the key is absent. Pass a
second argument to get that instead of `nil`. Use `fetch` when a missing
key is a mistake rather than an expected case.



<CodeBlockSimple input='h = {"a": 1}
h.get("a")
h.get("z")
h.get("z", 0)
' output='{"a": 1}
1
nil
0
' />


### has_key?(HASHABLE)
> Returns `BOOLEAN`

Returns `true` when the hash has an entry under the given key. The argument has to be usable as a key, which is what `HASHABLE` in the signature means: a `STRING`, `INTEGER`, `FLOAT`, `BOOLEAN`, `ARRAY` or `HASH`.


<CodeBlockSimple input='{"a": 1, 1: "b"}.has_key?(1)
{"a": 1, 1: "b"}.has_key?("c")
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


### reject(CALLABLE)
> Returns `HASH|ERROR`

Returns a new hash of the entries the callback said no to -- the mirror of `filter`. Use `reject!` to filter in place.


<CodeBlockSimple input='h = {"a": 1}
h["b"] = 2
h.reject(def(k, v) v > 1 end).size()
' output='{"a": 1}
2
1
' />


### reject!(CALLABLE)
> Returns `HASH|ERROR`

Drops the entries the callback said yes to and returns the hash, so calls can be chained.


<CodeBlockSimple input='h = {"a": 1}
h["b"] = 2
h.reject!(def(k, v) v > 1 end).size()
' output='{"a": 1}
2
1
' />


### remove(HASHABLE)
> Returns `ANY`

Removes the entry for a key and returns its value, or `nil` when the key was not there, so a removal can be told from a miss.


<CodeBlockSimple input='h = {"a": 1}
h.remove("a")
h.size()
h.remove("a")
' output='{"a": 1}
1
0
nil
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


### transform_keys(CALLABLE)
> Returns `HASH|ERROR`

Returns a new hash with each key replaced by what the callback answered for it. A new key still has to be `HASHABLE`, and two keys answering the same collapse into one entry -- which of them survives is not defined. Use `transform_keys!` to replace them in place.


<CodeBlockSimple input='h = {"a": 1}
h.transform_keys(def(k) k.uppercase() end).get("A", 0)
' output='{"a": 1}
1
' />


### transform_keys!(CALLABLE)
> Returns `HASH|ERROR`

Replaces each key with what the callback answered and returns the hash, so calls can be chained.


<CodeBlockSimple input='h = {"a": 1}
h.transform_keys!(def(k) k.uppercase() end).get("A", 0)
' output='{"a": 1}
1
' />


### transform_values(CALLABLE)
> Returns `HASH|ERROR`

Returns a new hash with each value replaced by what the callback answered for it. The keys are untouched. Use `transform_values!` to replace them in place.


<CodeBlockSimple input='h = {"a": 1}
h.transform_values(def(v) v * 10 end).get("a", 0)
h.get("a", 0)
' output='{"a": 1}
10
1
' />


### transform_values!(CALLABLE)
> Returns `HASH|ERROR`

Replaces each value with what the callback answered and returns the hash, so calls can be chained.


<CodeBlockSimple input='h = {"a": 1}
h.transform_values!(def(v) v * 10 end).get("a", 0)
' output='{"a": 1}
10
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

### help()
> Returns `NIL`

Prints the type's literal-specific methods with their argument and return types, sorted by name, one per line. It returns `nil` rather than the listing: this exists to be read, and the REPL echoes a returned value through its escaped representation, which would put the whole thing on one line. Use `methods` when the names are wanted as data. A type with no methods of its own prints only the heading.


<CodeBlockSimple input='true.help()
1.0.help()
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
puts.type_groups()
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


