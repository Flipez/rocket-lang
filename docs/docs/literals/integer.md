import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Integer

An integer can be positiv or negative and is always internally represented by a 64-Bit Integer.

To cast a negative integer a digit can be prefixed with a - eg. -456.



```js
a = 1;

b = a + 2;

is_true = 1 == 1;
is_false = 1 == 2;

```

## Literal Specific Methods

### abs()
> Returns `INTEGER`

Returns the absolute value, as an integer, keeping the base.


<CodeBlockSimple input='3.abs()
(0 - 5).abs()
"-0x10".to_integer().abs()
' output='3
5
0x10
' />


### acos()
> Returns `FLOAT`

Returns the arccosine, in radians, of the number, as a float.


<CodeBlockSimple input='1.acos()
' output='0.0
' />


### asin()
> Returns `FLOAT`

Returns the arcsine, in radians, of the number, as a float.


<CodeBlockSimple input='0.asin()
' output='0.0
' />


### atan()
> Returns `FLOAT`

Returns the arctangent, in radians, of the number, as a float.


<CodeBlockSimple input='0.atan()
' output='0.0
' />


### bit_length()
> Returns `INTEGER`

Returns how many bits are needed to represent the number, not counting the sign. A negative number needs as many as its complement, so `-256` is 8 bits just like `255`.


<CodeBlockSimple input='255.bit_length()
256.bit_length()
0.bit_length()
' output='8
9
0
' />


### ceil([INTEGER])
> Returns `INTEGER`

Rounds up to a multiple of ten, given a negative digit count. An integer is already exact, so no argument, or a count of zero or more, returns it unchanged.


<CodeBlockSimple input='555.ceil(0 - 1)
555.ceil(0 - 2)
555.ceil()
' output='560
600
555
' />


### copysign(NUMERIC)
> Returns `FLOAT`

Returns a float with the magnitude of the number and the sign of the argument.


<CodeBlockSimple input='3.copysign(0 - 1.0)
' output='-3.0
' />


### cos()
> Returns `FLOAT`

Returns the cosine, in radians, of the number, as a float.


<CodeBlockSimple input='0.cos()
' output='1.0
' />


### digits([INTEGER])
> Returns `ARRAY|ERROR`

Returns the digits of the number, least significant first, in base 10 or in a given base. A negative number has no digits and gives an error.


<CodeBlockSimple input='12345.digits()
12345.digits(7)
0.digits()
' output='[5, 4, 3, 2, 1]
[4, 6, 6, 0, 5]
[0]
' />


### divmod(INTEGER)
> Returns `ARRAY|ERROR`

Returns the quotient and the remainder as a two-element array. Both follow the `/` and `%` operators, which truncate toward zero. Ruby floors instead, so it answers `[-3, -1]` for the second example below.


<CodeBlockSimple input='11.divmod(4)
11.divmod(0 - 4)
' output='[2, 3]
[-2, 3]
' />


### downto(INTEGER, CALLABLE)
> Returns `INTEGER|ERROR`

Calls the callback with each integer from the receiver down to the given limit, inclusive at both ends, and returns the receiver. A limit above the receiver calls nothing.


<CodeBlockSimple input='3.downto(1, def(i) print(i) end)
' output='3
2
1
3
' />


### even?()
> Returns `BOOLEAN`

Returns `true` when the number divides by two exactly.


<CodeBlockSimple input='4.even?()
5.even?()
' output='true
false
' />


### exp()
> Returns `FLOAT`

Returns e raised to the number, as a float.


<CodeBlockSimple input='1.exp()
' output='2.718281828459045
' />


### floor([INTEGER])
> Returns `INTEGER`

Rounds down to a multiple of ten, given a negative digit count. No argument, or a count of zero or more, returns the number unchanged.


<CodeBlockSimple input='555.floor(0 - 1)
555.floor(0 - 2)
' output='550
500
' />


### gcd(INTEGER)
> Returns `INTEGER|ERROR`

Returns the greatest common divisor of the two numbers, always positive. Like the infix operators, it refuses two integers of different bases.


<CodeBlockSimple input='36.gcd(60)
3.gcd(0 - 7)
' output='12
1
' />


### lcm(INTEGER)
> Returns `INTEGER|ERROR`

Returns the least common multiple of the two numbers, always positive.


<CodeBlockSimple input='36.lcm(60)
3.lcm(0 - 7)
' output='180
21
' />


### log()
> Returns `FLOAT`

Returns the natural logarithm of the number, as a float.


<CodeBlockSimple input='1.log()
' output='0.0
' />


### log10()
> Returns `FLOAT`

Returns the decimal logarithm of the number, as a float.


<CodeBlockSimple input='100.log10()
' output='2.0
' />


### log2()
> Returns `FLOAT`

Returns the binary logarithm of the number, as a float.


<CodeBlockSimple input='8.log2()
' output='3.0
' />


### negative?()
> Returns `BOOLEAN`

Returns `true` when the number is less than zero.


<CodeBlockSimple input='(0 - 1).negative?()
0.negative?()
' output='true
false
' />


### odd?()
> Returns `BOOLEAN`

Returns `true` when the number does not divide by two exactly.


<CodeBlockSimple input='5.odd?()
4.odd?()
' output='true
false
' />


### positive?()
> Returns `BOOLEAN`

Returns `true` when the number is greater than zero.


<CodeBlockSimple input='1.positive?()
0.positive?()
' output='true
false
' />


### pow(INTEGER, [INTEGER])
> Returns `INTEGER|ERROR`

Returns the number raised to the given power. A second argument takes the result modulo it, which keeps large exponents workable. A negative exponent is an error, since the result would not be an integer.


<CodeBlockSimple input='2.pow(3)
2.pow(3, 5)
' output='8
3
' />


### predecessor()
> Returns `INTEGER`

Returns the previous integer, keeping the base of the receiver.


<CodeBlockSimple input='1.predecessor()
0.predecessor()
' output='0
-1
' />


### remainder(NUMERIC)
> Returns `FLOAT`

Returns the IEEE 754 floating-point remainder of the number divided by the argument, as a float.


<CodeBlockSimple input='100.remainder(30.0)
' output='10.0
' />


### round([INTEGER])
> Returns `INTEGER`

Rounds to the nearest multiple of ten, halves away from zero, given a negative digit count. No argument, or a count of zero or more, returns the number unchanged.


<CodeBlockSimple input='555.round(0 - 1)
554.round(0 - 1)
' output='560
550
' />


### sin()
> Returns `FLOAT`

Returns the sine, in radians, of the number, as a float.


<CodeBlockSimple input='0.sin()
' output='0.0
' />


### sqrt()
> Returns `FLOAT`

Returns the square root of the number, as a float.


<CodeBlockSimple input='16.sqrt()
' output='4.0
' />


### successor()
> Returns `INTEGER`

Returns the next integer, keeping the base of the receiver.


<CodeBlockSimple input='1.successor()
' output='2
' />


### tan()
> Returns `FLOAT`

Returns the tangent, in radians, of the number, as a float.


<CodeBlockSimple input='0.tan()
' output='0.0
' />


### times(CALLABLE)
> Returns `INTEGER|ERROR`

Calls the callback with each integer from `0` up to one less than the receiver, and returns the receiver so calls can be chained. A count of zero or less calls nothing rather than erroring, which makes it safe to hand a computed count. The counter keeps the receiver's base.


<CodeBlockSimple input='3.times(def(i) print(i) end)
0.times(def(i) print("never") end)
' output='0
1
2
3
0
' />


### to_base(INTEGER)
> Returns `INTEGER`

Converts the integer into a integer with the given base


<CodeBlockSimple input='"0b1010".to_integer().to_base(8)
' output='0o12
' />


### to_character()
> Returns `STRING|ERROR`

Returns the character with this code point, as a string. The inverse of a single entry of `string.codepoints()`. A value outside the range of a character is an error.


<CodeBlockSimple input='65.to_character()
' output='"A"
' />


### truncate([INTEGER])
> Returns `INTEGER`

Drops the last digits, rounding toward zero, given a negative digit count. No argument, or a count of zero or more, returns the number unchanged.


<CodeBlockSimple input='555.truncate(0 - 1)
(0 - 555).truncate(0 - 1)
' output='550
-550
' />


### upto(INTEGER, CALLABLE)
> Returns `INTEGER|ERROR`

Calls the callback with each integer from the receiver up to the given limit, inclusive at both ends, and returns the receiver. A limit below the receiver calls nothing.


<CodeBlockSimple input='1.upto(3, def(i) print(i) end)
3.upto(1, def(i) print("never") end)
' output='1
2
3
1
3
' />


### zero?()
> Returns `BOOLEAN`

Returns `true` when the number is zero.


<CodeBlockSimple input='0.zero?()
1.zero?()
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
print.type_groups()
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


