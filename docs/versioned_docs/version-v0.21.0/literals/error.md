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


