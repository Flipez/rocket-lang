---
title: "{{ .Title }}"
menu:
  docs:
    parent: "literals"
---
# {{ .Title }}

{{ .Description }}

{{ if .Example }}
```js
{{ .Example }}
```
{{ end }}
## Literal Specific Methods
{{ range $method, $object := .LiteralMethods }}
### {{ $object.Usage $method }}
> Returns `{{ $object.ReturnPattern }}`

{{ $object.Description }}

{{ if $object.Example }}
```js
{{ $object.Example }}
```
{{ end }}
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