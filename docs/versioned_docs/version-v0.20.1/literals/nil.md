import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Nil

Nil is the representation of "nothing".
It will be returned if something returns nothing (eg. puts or an empty break/next) and can also be generated with 'nil'.



## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Returns zero float.


<CodeBlockSimple input='nil.plz_f()
' output='0.0
' />


### plz_i()
> Returns `INTEGER`

Returns zero integer.


<CodeBlockSimple input='nil.plz_i()
' output='0
' />


### plz_s()
> Returns `STRING`

Returns empty string.


<CodeBlockSimple input='nil.plz_s()
' output='""
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


