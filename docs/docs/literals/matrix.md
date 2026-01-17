import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Matrix

A matrix is a 2-dimensional array of numbers used for linear algebra operations.
Matrices are created using the Matrix() constructor with nested arrays.

Matrix supports mathematical operations:
- Matrix multiplication: `m1 * m2`
- Element-wise addition: `m1 + m2`
- Element-wise subtraction: `m1 - m2`



```js
m1 = Matrix([[1, 2], [3, 4]])
m2 = Matrix([[5, 6], [7, 8]])

result = m1 * m2
sum = m1 + m2
diff = m2 - m1

puts(result)

// should output
Matrix(2x2)[
  [19, 22],
  [43, 50]
]

```

## Literal Specific Methods

### to_a()
> Returns `ARRAY`

Converts the matrix back to a nested array representation.


<CodeBlockSimple input='m = Matrix([[1, 2], [3, 4]])
m.to_a()
' output='[[1.0, 2.0], [3.0, 4.0]]
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


### to_i(INTEGER)
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


### to_s(INTEGER)
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


