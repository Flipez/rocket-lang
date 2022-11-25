import CodeBlockSimple from '@site/components/CodeBlockSimple'

# JSON




## Module Function

### parse(STRING)
> Returns `HASH`

Takes a STRING and parses it to a HASH or ARRAY. Numbers are always FLOAT.


<CodeBlockSimple input='JSON.parse(&apos;{"test": 123}&apos;)
JSON.parse(&apos;["test", 123]&apos;)
' output='{"test": 123.0}
["test", 123.0]
' />



## Properties
| Name | Value |
| ---- | ----- |

