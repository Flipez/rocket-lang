import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Boolean

A Boolean can represent two values: `true` and `false` and can be used in control flows.


```js
true // Is the representation for truthyness
false // is it for a falsy value

a = true;
b = false;

is_true = a == a;
is_false = a == b;

is_true = a != b;

```

## Literal Specific Methods

### plz_s()
> Returns `STRING`

Converts a boolean into a String representation and returns `"true"` or `"false"` based on the value.


<CodeBlockSimple input='true.plz_s()
' output='"true"
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


