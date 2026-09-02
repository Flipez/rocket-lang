# Methods

> 👉 The `!` convention became consistent across all types in `0.24`

Every value in RocketLang is an object, and objects are used by calling methods
on them. `methods()` lists what a value responds to, and `help()` lists the
same names with their arguments:

```js
🚀 > [1, 2, 3].methods()
=> ["all?", "any?", "append", "append!", "chunks", "compact", "compact!", "concat", "concat!", "contains?", "count", "each", "empty?", "filter", "filter!", "first", "flatten", "flatten!", "index_of", "insert", "insert!", "join", "last", "last_index_of", "map", "map!", "max", "max_by", "min", "min_by", "none?", "prepend", "prepend!", "reduce", "reject", "reject!", "remove", "remove!", "remove_at", "remove_at!", "remove_first", "remove_first!", "remove_last", "remove_last!", "reverse", "reverse!", "rotate", "rotate!", "size", "skip", "skip_last", "sort", "sort!", "sort_by", "sort_by!", "sum", "to_matrix", "unique", "unique!"]
```

Both are sorted by name. Before `0.24` they came out in a different order on
every run.

## Reading a signature

The documentation gives each method as a signature, for example
`fetch(HASHABLE, [ANY])`. Three things appear in the argument list:

| Notation | Meaning |
| -------- | ------- |
| `STRING` | a concrete type |
| `[STRING]` | may be left out |
| `STRING...` | one or more of them |

Where a method takes a whole family of types, the family is named rather than
listed — `append(ANY)`, `get(HASHABLE, [ANY])`, `set(INTEGER, INTEGER, NUMERIC)`.
Those names are **type groups**; see
[Types and type groups](./types#type-groups) for what each one accepts and which
types belong to it.

## Methods that take a callback

A callback is a function literal, since there is no separate block syntax:

```js
🚀 > [1, 2, 3].map(def(x) x * 2 end)
=> [2, 4, 6]
🚀 > [1, 2, 3, 4].filter(def(x) x % 2 == 0 end)
=> [2, 4]
🚀 > [1, 2, 3].reduce(0, def(sum, x) sum + x end)
=> 6
```

A builtin is a value too, so it can be the callback:

```js
🚀 > [1, 2].each(print)
1
2
=> [1, 2]
```

Every one of them treats the callback's answer the same way:

| In the callback | Effect |
| --------------- | ------ |
| a value | used — what that means is the method's business |
| `break` | the walk ends here, and the answer covers what was walked |
| `next` | the element contributed nothing: a `nil` from `map`, a no from `filter` |
| an error | the walk ends and the error is passed on |

```js
🚀 > [1, 2, 3, 4].map(def(x) if x == 3 break end x end)
=> [1, 2]
🚀 > [1, 2, 3].map(def(x) if x == 2 next end x end)
=> [1, nil, 3]
```

`break` and `next` behave as they do in a `for`. Note that `return` inside
the callback returns from the *callback*, since it is an ordinary function —
there is no enclosing method to return from.

Only `false` and `nil` are no, so `filter` keeps a `0` and an empty string:

```js
🚀 > [1, 2].filter(def(x) 0 end)
=> [1, 2]
```

## Methods ending in `!`

A method whose name ends in `!` changes the value it is called on. The plain
method of the same name leaves it alone and returns a new value instead:

```js
🚀 > a = [3, 1, 2]
=> [3, 1, 2]
🚀 > a.sort()
=> [1, 2, 3]
🚀 > a
=> [3, 1, 2]
🚀 > a.sort!()
=> [1, 2, 3]
🚀 > a
=> [1, 2, 3]
```

`sort()` handed back a sorted copy and left `a` as it was. `sort!()` sorted `a`
itself.

The pairs are:

| Pure | In place | Type |
| ---- | -------- | ---- |
| `reverse` | `reverse!` | `ARRAY`, `STRING` |
| `compact` | `compact!` | `ARRAY`, `HASH` |
| `flatten` | `flatten!` | `ARRAY` |
| `rotate` | `rotate!` | `ARRAY` |
| `sort` | `sort!` | `ARRAY` |
| `unique` | `unique!` | `ARRAY` |
| `merge` | `merge!` | `HASH` |
| `append` | `append!` | `ARRAY` |
| `prepend` | `prepend!` | `ARRAY` |
| `insert` | `insert!` | `ARRAY` |
| `concat` | `concat!` | `ARRAY` |
| `remove_first` | `remove_first!` | `ARRAY` |
| `remove` | `remove!` | `ARRAY`, `HASH` |
| `remove_at` | `remove_at!` | `ARRAY` |
| `capitalize` | `capitalize!` | `STRING` |
| `lowercase` | `lowercase!` | `STRING` |
| `remove_last` | `remove_last!` | `ARRAY`, `STRING` |
| `replace` | `replace!` | `STRING` |
| `swap_case` | `swap_case!` | `STRING` |
| `trim` | `trim!` | `STRING` |
| `trim_end` | `trim_end!` | `STRING` |
| `trim_line_end` | `trim_line_end!` | `STRING` |
| `trim_start` | `trim_start!` | `STRING` |
| `uppercase` | `uppercase!` | `STRING` |
| `set` | `set!` | `MATRIX` |

A method that cannot sensibly be done in place has no `!` form. `size()` and
`split()` return something other than a string, so there is nothing for a
`size!()` to mean. Neither do the predicates, which answer a question rather
than change anything: `empty?`, `contains?`, `starts_with?`, `ends_with?`,
`even?`, `odd?`, `zero?`, `positive?`, `negative?`, `nan?`, `finite?` and
`nil?`.

### A `!` method returns the value it changed

Since `0.24` every `!` method returns the object it just modified, rather than
`nil`. That makes them chainable:

```js
🚀 > "hello world".uppercase!().reverse!()
=> "DLROW OLLEH"
```

Before `0.24` these returned `nil`, so the second call in that chain failed with
`undefined method '.reverse!()' for NIL`.

This is a deliberate difference from Ruby. Ruby's `String#upcase!` is documented
as returning "`self` if any changes were made, `nil` otherwise", which means the
same chain raises there:

```ruby
# Ruby
"ABC".upcase!            #=> nil
"ABC".upcase!.reverse!   #=> NoMethodError: undefined method 'reverse!' for nil
```

RocketLang returns the receiver whether or not anything changed, so a chain
never depends on whether the string happened to already be uppercase. The cost
is that you cannot use the return value to ask "did this change anything".

### A failed `!` method changes nothing

`sort!` needs the elements to be all strings, all integers or all floats. When
they are not, it returns an error and leaves the array as it was, rather than
half-ordered:

```js
🚀 > a = [1, "x", 2]
=> [1, "x", 2]
🚀 > a.sort!()
=> ERROR: Array does contain either an object not INTEGER, FLOAT or STRING or is mixed
🚀 > a
=> [1, "x", 2]
```

## `MATRIX#set` used to mutate without a `!`

`append`, `remove_last`, `MATRIX#set` and the rest of `ARRAY`, `HASH` and
`MATRIX` used to change the receiver without a `!` — the same gap that let
`remove_last` mean the opposite of `String#remove_last`. They are paired now
like everything else: the plain method leaves the receiver alone and returns
a new value, the `!` method changes the receiver and hands it back. A pop is
two calls where mutating `remove_last` used to be one:

```js
🚀 > a = [1, 2, 3]
=> [1, 2, 3]
🚀 > last = a.last()
=> 3
🚀 > a.remove_last!()
=> [1, 2]
🚀 > last
=> 3
```

`MATRIX#set` follows the same pair: `set` returns a new matrix with the
position changed and leaves the receiver alone, `set!` changes the receiver
in place and hands it back so calls chain:

```js
🚀 > m = [[1, 2], [3, 4]].to_matrix()
🚀 > m.set!(0, 0, 9).set!(1, 1, 9).to_array()
=> [[9.0, 2.0], [3.0, 9.0]]
🚀 > m.to_array()
=> [[9.0, 2.0], [3.0, 9.0]]
```

The pure form chains too, since each call hands back a new matrix for the
next one to act on — but the receiver is untouched:

```js
🚀 > m = [[1, 2], [3, 4]].to_matrix()
🚀 > m.set(0, 0, 9).set(1, 1, 9).to_array()
=> [[9.0, 2.0], [3.0, 9.0]]
🚀 > m.to_array()
=> [[1.0, 2.0], [3.0, 4.0]]
```

## Ordering

`unique` keeps the order in which elements first appear:

```js
🚀 > [5, 3, 1, 4, 2, 3].unique()
=> [5, 3, 1, 4, 2]
```

The `keys()` and `values()` methods of a `HASH` are the exception: their order
is not defined and can differ between runs.
