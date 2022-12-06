import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Float




## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Returns self


<CodeBlockSimple input='123.456.plz_f()
' output='123.456
' />


### plz_i()
> Returns `INTEGER`

Converts the float into an integer.


<CodeBlockSimple input='123.456.plz_i()
' output='123
' />


### plz_s()
> Returns `STRING`

Returns a string representation of the float.


<CodeBlockSimple input='123.456.plz_s()
' output='"123.456"
' />



## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


<CodeBlockSimple input='"test".methods()
' output='["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "plz_i", "strip", "downcase"]
' />


### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


<CodeBlockSimple input='a = {"test": 1234}
a.to_json()
' output='{"test": 1234}
"{\"test\":1234}"
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
  plz_s()"
' />


