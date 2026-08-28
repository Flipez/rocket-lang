# Operators

## Available operators

| Operator | Description |
| -------- | ----------- |
| `=` | Assignment |
| `? :` | Ternary conditional |
| `and`, `&&` | Logical and |
| `or`, `\|\|` | Logical or |
| `==`, `!=` | Equal, not equal |
| `<`, `>`, `<=`, `>=` | Comparison |
| `+`, `-` | Addition, subtraction |
| `*`, `/` | Multiplication, division |
| `%` | Modulo |
| `-`, `!` | Prefix negation, prefix not |
| `()` | Call |
| `.` | Member access |
| `[]` | Index |

`and` and `&&` are the same operator, as are `or` and `||`. Pick one style and
keep to it within a file; the word forms read more naturally in conditions,
and the symbol forms are shorter in dense expressions.

## Precedence

Operators lower in this table bind more tightly, so they are applied first.

| | Operators | Example |
| - | --------- | ------- |
| 1 | `=` | `a = 1 + 2` assigns `3` |
| 2 | `? :` | `a == 1 ? "y" : "n"` |
| 3 | `and`, `&&`, `or`, `\|\|` | see the note below |
| 4 | `==`, `!=` | `2 + 3 == 5` is `true` |
| 5 | `<`, `>`, `<=`, `>=` | `1 < 2 and 3 > 2` is `true` |
| 6 | `+`, `-` | |
| 7 | `*`, `/` | `1 + 2 * 3` is `7` |
| 8 | `%` | `2 * 10 % 4` is `4` |
| 9 | prefix `-`, `!` | `-2 * 3` is `-6` |
| 10 | `()`, `.` | |
| 11 | `[]` | |

Parentheses override all of it: `(1 + 2) * 3` is `9`.

Note that `%` binds more tightly than `*` and `/`, which is not true in every
language. `2 * 10 % 4` groups as `2 * (10 % 4)`, giving `4` rather than `0`.

### `and` and `or` share one level

Unlike most languages, `and` does **not** bind more tightly than `or`. They sit
at the same precedence and group left to right, so:

```js
true or false and false
```

is `(true or false) and false`, which is `false` — not `true or (false and
false)`, which would be `true`. Parenthesise whenever a condition mixes the
two:

```js
puts(true or (false and false))  // true
puts((true or false) and false)  // false
```
