# {{ .Name }}

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

{{ if $function.Layout.Example }}
```js
{{ $function.Layout.Example }}
```
{{ end }}
{{ end }}

## Properties
| Name | Value |
| ---- | ----- |
{{ range $propertyName, $property := .Properties -}}
| {{ $propertyName }} | {{ $property.Value.Value }} |
{{ end }}