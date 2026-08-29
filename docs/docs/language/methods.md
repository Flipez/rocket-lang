# Methods

> 👉 The `!` convention became consistent across all types in `0.24`

Every value in RocketLang is an object, and objects are used by calling methods
on them. `methods()` lists what a value responds to, and `wat()` lists the same
names with their arguments:

```js
🚀 > [1, 2, 3].methods()
=> ["first", "include?", "index", "join", "last", "pop", "push", "reverse", "reverse!", "size", "slices", "sort", "sort!", "sum", "to_m", "uniq", "uniq!"]
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
listed. These names are **type groups**, not types — you never write them in
RocketLang, they only appear in signatures and in error messages:

| Group | Accepts | Used by |
| ----- | ------- | ------- |
| `ANY` | every value | `push`, `unshift`, `insert`, `include?` and `index` on an `ARRAY`, the fallback of `HASH.get` |
| `HASHABLE` | anything usable as a hash key: `STRING`, `INTEGER`, `FLOAT`, `BOOLEAN`, `ARRAY`, `HASH` | the key arguments of `HASH.get`, `fetch`, `delete` and `include?` |
| `NUMERIC` | `INTEGER` and `FLOAT` | `MATRIX.set` |

A group is checked by asking the value what it can do, not by comparing it
against a list, so a type added to the language joins the group it qualifies
for. An error names the group rather than spelling out its members:

```js
🚀 > {"a": 1}.get(nil, 0)
=> ERROR: wrong argument type on position 1: got=NIL, want=HASHABLE
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
| `uniq` | `uniq!` | `ARRAY` |
| `merge` | `merge!` | `HASH` |
| `capitalize` | `capitalize!` | `STRING` |
| `chomp` | `chomp!` | `STRING` |
| `chop` | `chop!` | `STRING` |
| `downcase` | `downcase!` | `STRING` |
| `lstrip` | `lstrip!` | `STRING` |
| `replace` | `replace!` | `STRING` |
| `rstrip` | `rstrip!` | `STRING` |
| `strip` | `strip!` | `STRING` |
| `swapcase` | `swapcase!` | `STRING` |
| `upcase` | `upcase!` | `STRING` |

A method that cannot sensibly be done in place has no `!` form. `size()` and
`split()` return something other than a string, so there is nothing for a
`size!()` to mean. Neither do the predicates, which answer a question rather
than change anything: `empty?`, `include?`, `start_with?`, `end_with?`,
`even?`, `odd?`, `zero?`, `positive?`, `negative?`, `nan?`, `finite?` and
`nil?`.

### A `!` method returns the value it changed

Since `0.24` every `!` method returns the object it just modified, rather than
`nil`. That makes them chainable:

```js
🚀 > "hello world".upcase!().reverse!()
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

## Mutating methods without a `!`

A few methods change the object without a `!`, because they have no pure
counterpart — the name would mean nothing else:

```js
🚀 > a = [1]
=> [1]
🚀 > a.push(2).push(3)
=> [1, 2, 3]
🚀 > a.pop()
=> 3
🚀 > a
=> [1, 2]
```

`push` returns the array, so pushes chain. `pop` returns the element it removed,
since that is the only thing worth having back. `set` on a `MATRIX` behaves like
`push`:

```js
🚀 > m = [[1, 2], [3, 4]].to_m()
🚀 > m.set(0, 0, 9).set(1, 1, 9).to_a()
=> [[9.0, 2.0], [3.0, 9.0]]
```

The rule these follow is what the method has to give back. One that puts
something in returns the receiver, so it chains; one that takes something out
returns what it took, so nothing is lost:

| Returns the receiver | Returns what it removed |
| -------------------- | ----------------------- |
| `push`, `unshift`, `insert`, `concat`, `clear` (`ARRAY`) | `pop`, `shift`, `delete`, `delete_at` (`ARRAY`) |
| `clear` (`HASH`) | `delete` (`HASH`) |
| `set` (`MATRIX`) | |

A method that removes something answers `nil` when there was nothing to remove,
so a removal can be told from a miss:

```js
🚀 > a = [1, 2, 1]
🚀 > a.delete(1)
=> 1
🚀 > a
=> [2]
🚀 > a.delete(9)
=> nil
```

## Ordering

`uniq` keeps the order in which elements first appear:

```js
🚀 > [5, 3, 1, 4, 2, 3].uniq()
=> [5, 3, 1, 4, 2]
```

The `keys()` and `values()` methods of a `HASH` are the exception: their order
is not defined and can differ between runs.
