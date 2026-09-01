import CodeBlockSimple from '@site/components/CodeBlockSimple'

# IO




## Module Function

### open(STRING, [STRING], [STRING])
> Returns `FILE`

Opens a file pointer to the file at the path, mode and permission can be set optionally.


<CodeBlockSimple input='IO.open("main.go", "r", "0644")
' output='<file:main.go>
' />


### read_line()
> Returns `STRING|NIL|ERROR`

Reads a single line from standard input, with the trailing newline removed. Returns `nil` at the end of input, which is not an error but the normal way a piped program ends. A genuine read failure, unlike end of input, returns an error.


<CodeBlockSimple input='IO.read_line()
' output='"hello"
' />


### write(ANY...)
> Returns `NIL`

Writes the given value(s) to standard output without appending a trailing newline. Strings are written as their raw value; use `print` when a trailing newline is wanted.


<CodeBlockSimple input='IO.write("Progress")
IO.write(".")
IO.write(".")
IO.write(".")
' output='Progress...
' />



## Properties
| Name | Value |
| ---- | ----- |

