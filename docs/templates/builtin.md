# {{ .Title }}

{{ .Description }}

{{ with $example := .Example }}
{{ if $example ne "" }}
```js
{{ .Example }}
```
{{ end }}
{{ end }}
## Module Function
{{ range $name, $function := .Functions }}
### {{ $function.Layout.Usage $name }}
> Returns `{{ $function.Layout.DocsReturnPattern }}`

{{ $function.Layout.Description }}

{{ if $function.Layout.Input }}
```js
{{ $function.Layout.Input }}
```
{{ if $function.Layout.Output }}
```js
{{ $function.Layout.Output }}
```
{{ end}}
{{ end }}
{{ end }}

## Properties
| Name | Value |
| ---- | ----- |
{{ range $propertyName, $property := .Properties -}}
| {{ $propertyName }} | {{ $property.Value.Value }} |
{{ end }}
