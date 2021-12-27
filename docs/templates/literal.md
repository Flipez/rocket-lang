# String

## Literal Specific Methods
{{ range $method, $object := .StringMethods }}
### {{ $object.Usage $method }}
> Returns `{{ $object.ReturnPattern }}`

{{ $object.Description }}

```js
{{ $object.Example }}
```
{{ end }}

## Generic Literal Methods
{{ range $method, $object := .DefaultMethods }}
### {{ $object.Usage $method }}
> Returns `{{ $object.ReturnPattern }}`

{{ $object.Description }}

```js
{{ $object.Example }}
```
{{ end }}