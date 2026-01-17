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


### set(INTEGER, INTEGER, FLOAT|INTEGER)
> Returns `NIL|ERROR`

Sets the element at the specified row and column (0-indexed).


<CodeBlockSimple input='m = [[1, 2, 3], [4, 5, 6]].to_m()
m.set(0, 2, 99)
m
' output='2x3 matrix
┌               ┐
│ 1.0  2.0  3.0 │
│ 4.0  5.0  6.0 │
└               ┘
nil
2x3 matrix
┌                ┐
│ 1.0   2.0  99.0 │
│ 4.0   5.0   6.0 │
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

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


<CodeBlockSimple input='"test".methods()
' output='["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "strip", "downcase"]
' />


### to_f()
> Returns `FLOAT`

If possible converts an object to its float representation. If not 0.0 is returned.


<CodeBlockSimple input='1.to_f()
"1.4".to_f()
nil.to_f()
' output='1.0
1.4
0.0
' />


### to_i()
> Returns `INTEGER`

If possible converts an object to its integer representation. If not 0 is returned.


<CodeBlockSimple input='true.to_i()
false.to_i()
1234.to_i()
"4".to_i()
"10011010010"to_i(2)
"2322".to_i(8)
"0x2322".to_i()
' output='1
0
1234
4
1234
1234
1234
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

If possible converts an object to its string representation. If not empty string is returned.


<CodeBlockSimple input='true.to_s()
1234.to_s()
1234.to_s(2)
1234.to_s(8)
1234.to_s(10)
"test".to_s()
1.4.to_s()
' output='"true"
"1234"
"10011010010"
"2322"
"1234"
"test"
"1.4"
' />


### type()
> Returns `STRING`

Returns the type of the object.


<CodeBlockSimple input='"test".type()
' output='"STRING"
' />


### wat()
> Returns `STRING`

Returns the supported methods with usage information.


<CodeBlockSimple input='true.wat()
' output='"BOOLEAN supports the following methods:
  to_s()"
' />


