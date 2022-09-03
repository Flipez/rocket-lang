# {{ .Title }}

{{ .Description }}

{{ if .Example }}
```js
{{ .Example }}
```
{{ end }}
## Literal Specific Methods
{{ range $method, $object := .LiteralMethods }}
### {{ $object.Layout.Usage $method }}
> Returns `{{ $object.Layout.DocsReturnPattern }}`

{{ $object.Layout.Description }}

{{ if $object.Layout.Example }}
```js
{{ $object.Layout.Example }}
```
{{ end }}
{{ end }}

## Generic Literal Methods
{{ range $method, $object := .DefaultMethods }}
### {{ $object.Layout.Usage $method }}
> Returns `{{ $object.Layout.DocsReturnPattern }}`

{{ $object.Layout.Description }}

```js
{{ $object.Layout.Example }}
```
{{ end }}
