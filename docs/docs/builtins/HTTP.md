import CodeBlockSimple from '@site/components/CodeBlockSimple'

# HTTP




## Module Function

### new()
> Returns `HTTP`

Creates a new HTTP server. Handlers are registered on the returned object with .handle(), and .listen() then blocks serving them, so every handler has to be added first.

See [HTTP](../literals/http) for the server's own methods.




<CodeBlockSimple input='server = HTTP.new()
' />




## Properties
| Name | Value |
| ---- | ----- |

