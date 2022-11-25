import CodeBlockSimple from '@site/components/CodeBlockSimple'

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

{{ if and $object.Layout.Input $object.Layout.Output}}
<CodeBlockSimple input='{{$object.Layout.Input}}' output='{{$object.Layout.Output}}' />
{{ else }}
{{ if $object.Layout.Input }}
<CodeBlockSimple input='{{$object.Layout.Input}}' />
{{ end }}
{{ end }}
{{ end }}

## Generic Literal Methods
{{ range $method, $object := .DefaultMethods }}
### {{ $object.Layout.Usage $method }}
> Returns `{{ $object.Layout.DocsReturnPattern }}`

{{ $object.Layout.Description }}

{{ if and $object.Layout.Input $object.Layout.Output}}
<CodeBlockSimple input='{{$object.Layout.Input}}' output='{{$object.Layout.Output}}' />
{{ else }}
{{ if $object.Layout.Input }}
<CodeBlockSimple input='{{$object.Layout.Input}}' />
{{ end }}
{{ end }}
{{ end }}
