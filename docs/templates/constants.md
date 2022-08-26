# Constants

Constants are globally available and take precedence over other variables.
A created variable with the same name as a constant will never override the constant.

## List of Constants

| Name | Value |
| ---- | ----- |
{{ range $name, $object := .Constants -}}
| **{{ $name }}** | `{{ $object.Value }}` |
{{ end -}}
