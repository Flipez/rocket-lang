# JSON




## Module Function

### parse(STRING)
> Returns `HASH`

Takes a STRING and parses it to a HASH or ARRAY. Numbers are always FLOAT.


```js
JSON.parse('{"test": 123}')
JSON.parse('["test", 123]')

```

```js
{"test": 123.0}
["test", 123.0]

```




## Properties
| Name | Value |
| ---- | ----- |

