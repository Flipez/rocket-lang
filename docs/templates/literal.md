# String

## Literal Specific Methods
{{ range $method, $object := .Methods }}
### {{ $object.Usage $method }}
> Returns `{{ $object.ReturnPattern }}`

{{ $object.Description }}

```js
{{ $object.Example }}
```
{{ end }}
