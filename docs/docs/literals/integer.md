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


## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


<CodeBlockSimple input='"test".methods()
' output='["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "to_i", "strip", "downcase"]
' />


### to_f()
> Returns `FLOAT`







### to_i(INTEGER)
> Returns `INTEGER`







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


