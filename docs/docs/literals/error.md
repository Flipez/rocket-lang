import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Error

An Error is created by RocketLang if unallowed or invalid code is run.
An error does often replace the original return value of a function or identifier.
The documentation of those functions does indicate ERROR as a potential return value.

A program can rescue from errors within a block or alter it's behavior within other blocks like 'if' or 'def'.

It is possible for the user to create errors using 'raise(STRING)' which will return an ERROR object with STRING as the message.



```js
def test()
  puts(nope)
rescue e
  puts("Got error: '" + e.msg() + "'")
end

test()

=> "Got error in if: 'identifier not found: error'"

if (true)
  nope()
rescue your_name
  puts("Got error in if: '" + your_name.msg() + "'")
end

=> "Got error in if: 'identifier not found: nope'"

begin
  puts(nope)
rescue e
  puts("rescue")
end

=> "rescue"

```

## Literal Specific Methods

### msg()
> Returns `STRING`

Returns the error message

:::caution
Please note that performing `.msg()` on a ERROR object does result in a STRING object which then will no longer be treated as an error!
:::






## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


<CodeBlockSimple input='"test".methods()
' output='["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "to_i", "strip", "downcase"]
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


