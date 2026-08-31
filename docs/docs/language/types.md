# Types and type groups

> 👉 Type groups were introduced in `0.24`

Every value in RocketLang has a type, and `type()` reports it:

```js
🚀 > "abc".type()
=> "STRING"
🚀 > [1, 2].type()
=> "ARRAY"
```

## The types

| Type | How you get one | Own methods |
| ---- | --------------- | ----------- |
| `STRING` | [String](../literals/string) |
| `INTEGER` | [Integer](../literals/integer) |
| `FLOAT` | [Float](../literals/float) |
| `BOOLEAN` | none |
| `NIL` | none |
| `ARRAY` | [Array](../literals/array) |
| `HASH` | [Hash](../literals/hash) |
| `MATRIX` | [Matrix](../literals/matrix) |
| `FUNCTION` | none |
| `ERROR` | [Error](../literals/error) |
| `FILE` | [File](../literals/file) |
| `HTTP` | [HTTP](../literals/http) |
| `MODULE` | none |
| `BUILTIN_MODULE` | `Math`, `IO`, `JSON`, `OS`, `Time` | see [Builtins](../builtins/Math) |

Every type also answers the
[generic methods](../literals/string#generic-literal-methods)
`to_string`, `to_integer`, `to_float`, `to_json`, `type`, `type_groups`,
`methods`, `help`, `is_a?` and `nil?` — except a `MODULE`, which only exposes
what the module exports, so `lib.type()` is an error rather than `"MODULE"`.

## Type groups

A method that accepts a whole family of types names the family rather than
listing it. These names are **type groups**. You never write one in RocketLang:
they appear only in signatures and in error messages.

```js
🚀 > {"a": 1}.get(nil, 0)
=> ERROR: wrong argument type on position 1: got=NIL, want=HASHABLE
```

| Group | Means | Where it appears |
| ----- | ----- | ---------------- |
| `ANY` | any value at all | `push`, `unshift`, `insert`, `include?`, `index`, `rindex`, `count` and `delete` on an `ARRAY`; the fallback of `HASH.get` and `fetch`; `format`; `puts` |
| `HASHABLE` | can be used as a hash key | the key argument of `HASH.get`, `fetch`, `delete` and `include?`; the elements of `ARRAY.uniq`; what `HASH.transform_keys` answers |
| `COMPARABLE` | can be ordered against its own kind | the elements of `ARRAY.sort`, `min` and `max`; what `ARRAY.sort_by`, `min_by` and `max_by` answer |
| `STRINGABLE` | has a string form | the elements of `ARRAY.join` |
| `INTEGERABLE` | can be read as an integer | the elements of `ARRAY.sum` |
| `NUMERIC` | a number | the value argument of `MATRIX.set` |
| `CALLABLE` | a function, or a builtin such as `puts` — both are values | every callback: `ARRAY.each`, `map`, `select`, `reject`, `reduce`, `all?`, `sort_by`, `min_by`; `HASH.each`, `select`, `transform_values`, `transform_keys`; `INTEGER.times`, `upto`, `downto` |

### What belongs to what

`ANY` is not a row or a column here: every value belongs to it, so it says
nothing about any particular type. It exists for signatures, where
`push(ANY)` means the argument accepts anything.


| Type | `HASHABLE` | `COMPARABLE` | `STRINGABLE` | `INTEGERABLE` | `NUMERIC` | `CALLABLE` |
| -------- | --- | --- | --- | --- | --- | --- |
| `STRING` | ✅ | ✅ | ✅ | ✅ | | |
| `INTEGER` | ✅ | ✅ | ✅ | ✅ | ✅ | |
| `FLOAT` | ✅ | ✅ | ✅ | ✅ | ✅ | |
| `BOOLEAN` | ✅ | | ✅ | ✅ | | |
| `ARRAY` | ✅ | | ✅ | | | |
| `HASH` | ✅ | | ✅ | | | |
| `MATRIX` | | | ✅ | | | |
| `NIL` | | | ✅ | | | |
| `ERROR` | | | ✅ | | | |
| `FILE` | | | ✅ | | | |
| `HTTP` | | | ✅ | | | |
| `FUNCTION` | | | | | | ✅ |
| `MODULE` | | | | | | |

Two rows are worth a second look:

- A `BOOLEAN` is `INTEGERABLE`, and a `STRING` is too when it parses. That is
  wider than being a number, which is why `["12"].sum()` is `12` and
  `[true].sum()` is `1`.
- An `ARRAY` and a `HASH` are `HASHABLE`, so they can be hash keys:
  `{[1]: "a"}` is a valid hash.
- A `FUNCTION` is `CALLABLE` and nothing else — not even `STRINGABLE`, which is
  why `[def() end].join()` fails. A builtin such as `puts` is `CALLABLE` too,
  so `[1, 2].each(puts)` works.

A group is decided by asking the value what it can do, not by comparing it
against a list of type names. A type added to the language therefore joins every
group it qualifies for without anyone maintaining a list — which is how
`push` once came to accept a `FUNCTION` but reject a `FLOAT`.

### Asking a value

`type_groups()` lists the groups a value belongs to, and `is_a?` answers for one
of them:

```js
🚀 > 1.type_groups()
=> ["COMPARABLE", "HASHABLE", "INTEGERABLE", "NUMERIC", "STRINGABLE"]
🚀 > def() end.type_groups()
=> ["CALLABLE"]
🚀 > "a".is_a?("HASHABLE")
=> true
🚀 > nil.is_a?("HASHABLE")
=> false
```

`is_a?` takes a type name as readily as a group name, so it covers both
questions:

```js
🚀 > "a".is_a?("STRING")
=> true
🚀 > "a".is_a?("INTEGER")
=> false
```

`ANY` is missing from those lists on purpose — every value belongs to it, so it
would only prefix every answer with the same word. `is_a?("ANY")` still answers
`true`.

A name that is neither is an error rather than a `false`, because a typo would
otherwise read as a real answer:

```js
🚀 > "a".is_a?("HASHBALE")
=> ERROR: unknown type or type group: HASHBALE
```

The names are exact, and `type()` answers in upper case, so `is_a?` asks in it:
`is_a?("string")` is an error too.

## Where groups show up

### In a signature

The documentation gives each method as a signature. A group sits where a type
would:

```
push(ANY)
get(HASHABLE, ANY)
set(INTEGER, INTEGER, NUMERIC)
```

See [Methods](./methods#reading-a-signature) for the rest of the notation,
including how an optional or repeatable argument is written.

### In a requirement on elements

A signature can only describe the *arguments*. Some methods require something of
the **elements** they are given, and they name the same groups:

```js
🚀 > [def() end].join()
=> ERROR: element 0 is not STRINGABLE, got FUNCTION
🚀 > [1, nil].uniq()
=> ERROR: element 1 is not HASHABLE, got NIL
🚀 > [1, nil].sum()
=> ERROR: element 1 is not INTEGERABLE, got NIL
🚀 > [1, nil].sort()
=> ERROR: element 1 is not COMPARABLE, got NIL
```

Ordering carries one requirement no group can express, because it is about the
collection rather than any single value: the elements have to be the **same**
comparable type. `1` and `2.5` are each `COMPARABLE` and still cannot be sorted
together:

```js
🚀 > [1, 2.5].sort()
=> ERROR: elements must all be one COMPARABLE type, got INTEGER at 0 and FLOAT at 1
```
