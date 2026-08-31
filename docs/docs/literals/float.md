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


### acos()
> Returns `FLOAT`

Returns the arccosine, in radians, of the number, as a float.


<CodeBlockSimple input='1.0.acos()
' output='0.0
' />


### asin()
> Returns `FLOAT`

Returns the arcsine, in radians, of the number, as a float.


<CodeBlockSimple input='0.0.asin()
' output='0.0
' />


### atan()
> Returns `FLOAT`

Returns the arctangent, in radians, of the number, as a float.


<CodeBlockSimple input='0.0.atan()
' output='0.0
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


### copysign(NUMERIC)
> Returns `FLOAT`

Returns a float with the magnitude of the number and the sign of the argument.


<CodeBlockSimple input='3.2.copysign(0 - 1.0)
' output='-3.2
' />


### cos()
> Returns `FLOAT`

Returns the cosine, in radians, of the number, as a float.


<CodeBlockSimple input='0.0.cos()
' output='1.0
' />


### divmod(FLOAT)
> Returns `ARRAY|ERROR`

Returns the quotient and the remainder as a two-element array. Truncated toward zero, so it agrees with `Integer#divmod` and with `/`.


<CodeBlockSimple input='11.0.divmod(4.0)
11.0.divmod(0.0 - 4.0)
' output='[2.0, 3.0]
[-2.0, 3.0]
' />


### exp()
> Returns `FLOAT`

Returns e raised to the number, as a float.


<CodeBlockSimple input='1.0.exp()
' output='2.718281828459045
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


### log()
> Returns `FLOAT`

Returns the natural logarithm of the number, as a float.


<CodeBlockSimple input='1.0.log()
' output='0.0
' />


### log10()
> Returns `FLOAT`

Returns the decimal logarithm of the number, as a float.


<CodeBlockSimple input='100.0.log10()
' output='2.0
' />


### log2()
> Returns `FLOAT`

Returns the binary logarithm of the number, as a float.


<CodeBlockSimple input='8.0.log2()
' output='3.0
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


### pow(NUMERIC)
> Returns `FLOAT`

Returns the number raised to the given power, as a float. Unlike `Integer#pow` there is no modulus argument.


<CodeBlockSimple input='2.0.pow(3.0)
' output='8.0
' />


### remainder(NUMERIC)
> Returns `FLOAT`

Returns the IEEE 754 floating-point remainder of the number divided by the argument, as a float.


<CodeBlockSimple input='100.0.remainder(30.0)
' output='10.0
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


### sin()
> Returns `FLOAT`

Returns the sine, in radians, of the number, as a float.


<CodeBlockSimple input='0.0.sin()
' output='0.0
' />


### sqrt()
> Returns `FLOAT`

Returns the square root of the number, as a float.


<CodeBlockSimple input='16.0.sqrt()
' output='4.0
' />


### tan()
> Returns `FLOAT`

Returns the tangent, in radians, of the number, as a float.


<CodeBlockSimple input='0.0.tan()
' output='0.0
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


