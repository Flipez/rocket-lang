import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Matrix

A matrix is a 2-dimensional array of numbers used for linear algebra operations.
Matrices are created by calling the to_m() method on nested arrays.

Matrix supports mathematical operations:
- Matrix multiplication: `m1 * m2`
- Element-wise addition: `m1 + m2`
- Element-wise subtraction: `m1 - m2`

Matrix indexing:
- `m[i]` returns row i as an array
- `m[i][j]` returns the element at row i, column j
- Negative indices are supported: `m[-1]` returns the last row



```js
m1 = [[1, 2], [3, 4]].to_m()
m2 = [[5, 6], [7, 8]].to_m()

result = m1 * m2
sum = m1 + m2
diff = m2 - m1

puts(result)

// should output
2x2 matrix
┌            ┐
│ 19.0  22.0 │
│ 43.0  50.0 │
└            ┘

```

## Literal Specific Methods

### col(INTEGER)
> Returns `ARRAY|ERROR`

Returns the specified column as an array (0-indexed).


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.col(1)
' output='[2.0, 5.0]
' />


### cols()
> Returns `INTEGER`

Returns the number of columns in the matrix.


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.cols()
' output='3
' />


### get(INTEGER, INTEGER)
> Returns `FLOAT|ERROR`

Returns the element at the specified row and column (0-indexed).


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.get(0, 2)
' output='3.0
' />


### row(INTEGER)
> Returns `ARRAY|ERROR`

Returns the specified row as an array (0-indexed).


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.row(0)
' output='[1.0, 2.0, 3.0]
' />


### rows()
> Returns `INTEGER`

Returns the number of rows in the matrix.


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.rows()
' output='2
' />


### set(INTEGER, INTEGER, NUMERIC)
> Returns `NIL|ERROR`

Sets the element at the specified row and column (0-indexed) and returns the matrix, so calls can be chained.


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.set(0, 2, 99)
m
' output='2x3 matrix
┌               ┐
│ 1.0  2.0  3.0 │
│ 4.0  5.0  6.0 │
└               ┘
2x3 matrix
┌                ┐
│ 1.0  2.0  99.0 │
│ 4.0  5.0   6.0 │
└                ┘
2x3 matrix
┌                ┐
│ 1.0  2.0  99.0 │
│ 4.0  5.0   6.0 │
└                ┘
' />


### shape()
> Returns `ARRAY`

Returns an array containing the dimensions [rows, cols] of the matrix.


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.shape()
' output='[2, 3]
' />


### size()
> Returns `INTEGER`

Returns the total number of elements in the matrix (rows * cols).


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.size()
' output='6
' />


### t()
> Returns `MATRIX`

Alias for transpose(). Returns the transposed matrix.


<CodeBlockSimple input='m = [[1, 2], [3, 4]].to_m()
m.t()
' output='2x2 matrix
┌          ┐
│ 1.0  3.0 │
│ 2.0  4.0 │
└          ┘
' />


### to_a()
> Returns `ARRAY`

Converts the matrix back to a nested array representation.


<CodeBlockSimple input='m = [[1, 2], [3, 4]].to_m()
m.to_a()
' output='[[1.0, 2.0], [3.0, 4.0]]
' />


### transpose()
> Returns `MATRIX`

Returns the transposed matrix (rows and columns swapped).


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.transpose()
' output='3x2 matrix
┌          ┐
│ 1.0  4.0 │
│ 2.0  5.0 │
│ 3.0  6.0 │
└          ┘
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


