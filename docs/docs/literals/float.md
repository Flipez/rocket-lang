import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Float




## Literal Specific Methods

### abs()
> Returns `FLOAT`

Returns the absolute value, as a float.


<CodeBlockSimple input='3.14.abs()
(0.0 - 3.14).abs()
' output='3.14
3.14
' />


### ceil([INTEGER])
> Returns `FLOAT`

Returns the smallest float that is not less than the number. An optional digit count says how many decimal places to keep; a negative count rounds to a power of ten. The result stays a `FLOAT` -- in Ruby it would be an `INTEGER`.


<CodeBlockSimple input='1.5.ceil()
1.561.ceil(2)
555.5.ceil(0 - 1)
' output='2.0
1.57
560.0
' />


### divmod(FLOAT)
> Returns `ARRAY|ERROR`

Returns the quotient and the remainder as a two-element array. Truncated toward zero, so it agrees with `Integer#divmod` and with `/`.


<CodeBlockSimple input='11.0.divmod(4.0)
11.0.divmod(0.0 - 4.0)
' output='[2.0, 3.0]
[-2.0, 3.0]
' />


### finite?()
> Returns `BOOLEAN`

Returns `true` when the number is neither infinite nor not-a-number.


<CodeBlockSimple input='1.5.finite?()
' output='true
' />


### floor([INTEGER])
> Returns `FLOAT`

Returns the largest float that is not greater than the number. Takes the same optional digit count as `ceil`.


<CodeBlockSimple input='1.5.floor()
1.567.floor(2)
(0.0 - 1.5).floor()
' output='1.0
1.56
-2.0
' />


### infinite?()
> Returns `INTEGER|NIL`

Returns `1` for positive infinity, `-1` for negative infinity and `nil` otherwise, so the direction is not lost. Use `finite?` for the plain yes-or-no question.


<CodeBlockSimple input='1.5.infinite?()
' output='nil
' />


### nan?()
> Returns `BOOLEAN`

Returns `true` when the number is not a number.


<CodeBlockSimple input='1.5.nan?()
' output='false
' />


### negative?()
> Returns `BOOLEAN`

Returns `true` when the number is less than zero.


<CodeBlockSimple input='(0.0 - 1.5).negative?()
0.0.negative?()
' output='true
false
' />


### positive?()
> Returns `BOOLEAN`

Returns `true` when the number is greater than zero.


<CodeBlockSimple input='1.5.positive?()
0.0.positive?()
' output='true
false
' />


### round([INTEGER])
> Returns `FLOAT`

Returns the number rounded to the nearest value, halves away from zero. Takes the same optional digit count as `ceil`.


<CodeBlockSimple input='1.5.round()
1.567.round(2)
555.5.round(0 - 1)
' output='2.0
1.57
560.0
' />


### truncate([INTEGER])
> Returns `FLOAT`

Returns the number with its fractional part dropped, rounding toward zero. Unlike `floor` this moves a negative number up. Takes the same optional digit count as `ceil`.


<CodeBlockSimple input='1.567.truncate()
(0.0 - 1.5).truncate()
1.567.truncate(2)
' output='1.0
-1.0
1.56
' />


### zero?()
> Returns `BOOLEAN`

Returns `true` when the number is zero.


<CodeBlockSimple input='0.0.zero?()
1.5.zero?()
' output='true
false
' />



## Generic Literal Methods

### is_a?(STRING)
> Returns `BOOLEAN|ERROR`

Returns `true` when the value is of the given type or belongs to the given type group, so one question covers both. The name has to be one that exists: anything else is an error rather than a `false`, because a typo would otherwise read as a real answer. See [Types and type groups](/docs/language/types).


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


### type_groups()
> Returns `ARRAY`

Returns the type groups the value belongs to, sorted. `ANY` is not listed: every value belongs to it, so it would say nothing while prefixing every answer. It exists for signatures, where `push(ANY)` means the argument accepts anything, and `is_a?("ANY")` still answers `true`. See [Types and type groups](/docs/language/types) for what each group means.


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


