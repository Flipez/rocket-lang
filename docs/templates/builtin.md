import CodeBlockSimple from '@site/components/CodeBlockSimple'

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

{{ if and $function.Layout.Input $function.Layout.Output}}
<CodeBlockSimple input='{{$function.Layout.Input}}' output='{{$function.Layout.Output}}' />
{{ else }}
{{ if $function.Layout.Input }}
<CodeBlockSimple input='{{$function.Layout.Input}}' />
{{ end }}
{{ end }}
{{ end }}

## Properties
| Name | Value |
| ---- | ----- |
{{ range $propertyName, $property := .Properties -}}
| {{ $propertyName }} | {{ $property.Value.Value }} |
{{ end }}
