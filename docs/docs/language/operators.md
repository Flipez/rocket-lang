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

## Division

Dividing two integers gives an integer, as in Ruby. The result is truncated
toward zero rather than rounded.

```js
puts(4 / 2)   // 2
puts(5 / 2)   // 2, not 2.5
```

Mixing an integer with a float promotes the result to a float, so make one
side a float when you want the fractional part:

```js
puts(4.0 / 2)   // 2.0
puts(4 / 2.0)   // 2.0
puts(5 / 2.0)   // 2.5
```

Truncation means a negative quotient rounds toward zero: `-5 / 2` is `-2`, not
`-3`. This keeps division consistent with `%`, whose sign follows the
dividend, so `(a / b) * b + a % b` always gives back `a`.

Dividing by zero is an error in both forms.

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

## Where one statement ends

Nothing terminates a statement — there is no required semicolon — so a line
break is the only separator there is. Almost always that is unremarkable:

```js
a = 1
b = 2
```

It only becomes a question for the three tokens that can either *open* an
expression or *continue* one:

| Token | Opens | Continues |
| ----- | ----- | --------- |
| `[` | an array literal | an index |
| `(` | a grouped expression | a call |
| `-` | a negative number | a subtraction |

**A line break in front of one of these selects the opening meaning**, so the
line starts a new statement:

```js
puts("a")
[1, 2].each(puts)   // two statements, as it reads
```

Without that rule the `[` would index the result of `puts`, which is `nil`.

Every other operator has only the one meaning. `* 3` is not an expression on
its own, so a line break in front of `*` cannot mean anything but continuation,
and the same goes for `.` — which is what lets a chain be split across lines:

```js
sorted = [3, 1, 2]
  .sort()
```

### Inside brackets a line break means nothing

An unclosed `(`, `[` or `{` is proof that the expression has not ended, so line
breaks inside one are insignificant:

```js
puts(4
  - 1)         // 3, not two statements

total = [
  1,
  2
]
```

### An operator at the end of a line always continues

The rule is about a line *starting* with an ambiguous token. Putting the
operator at the end of the previous line is never ambiguous and always works:

```js
sum = 1 +
  2            // 3
```
