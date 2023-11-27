---
title: "If / elif / else"
menu:
  docs:
    parent: "controls"
---
import CodeBlockSimple from '@site/components/CodeBlockSimple'

# If
With `if` and `else` keywords the flow of a program can be controlled.

<CodeBlockSimple input='a="test"
if (a.type() == "STRING")
puts("is a string")
else
puts("is not a string")
end' output='is a string' />

# Elif
`elif` allows providing of an additional consequence check  after `if` and before evaluating the alternative provided by `else`. There is no limit on how many `elif` statements can be used.

<CodeBlockSimple input='a = "test"
if (a.type() == "BOOLEAN")
puts("is a boolean")
elif (a.type() == "STRING")
puts("is a string")
else
puts("i have no idea")
end' output='is a string' />