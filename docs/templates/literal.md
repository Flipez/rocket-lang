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

{{ if $object.Layout.Input }}
```js
{{ $object.Layout.Input }}
```
{{ if $object.Layout.Output }}
```js
{{ $object.Layout.Output }}
```
{{ end}}
{{ end }}
{{ end }}

## Generic Literal Methods
{{ range $method, $object := .DefaultMethods }}
### {{ $object.Layout.Usage $method }}
> Returns `{{ $object.Layout.DocsReturnPattern }}`

{{ $object.Layout.Description }}

{{ if $object.Layout.Input }}
```js
{{ $object.Layout.Input }}
```
{{ if $object.Layout.Output }}
```js
{{ $object.Layout.Output }}
```
{{ end}}
{{ end }}
{{ end }}
