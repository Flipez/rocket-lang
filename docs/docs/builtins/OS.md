import CodeBlockSimple from '@site/components/CodeBlockSimple'

# OS




## Module Function

### exit(INTEGER)
> Returns ``

Terminates the program with the given exit code.


<CodeBlockSimple input='OS.exit(1)
' output='exit status 1
' />


### raise(INTEGER, STRING)
> Returns ``

Terminates the program with the given exit code and prints the error message.


<CodeBlockSimple input='OS.raise(1, "broken")
' output='🔥 RocketLang raised an error: "broken"
exit status 1
' />



## Properties
| Name | Value |
| ---- | ----- |

