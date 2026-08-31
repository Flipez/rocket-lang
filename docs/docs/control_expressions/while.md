---
title: "While"
menu:
  docs:
    parent: "controls"
---
# While
While loops will run as long as the condition is truthy.

Parentheses around conditions are optional.

Print numbers from 0 to 3:

```js
🚀 > a = 0
🚀 > while a != 4
  puts(a)
  a = a + 1
end

// which prints
0
1
2
3
=> nil
```

It is possible to use `next` or `break` inside a while loop.

```js
i = 0
while i < 10
  if i < 3
    i = i + 1
    next
  end
  puts(i)
  if i == 6
    break
  end
  i = i + 1
end

// which prints
3
4
5
6
```

## Return Value
Loops evaluate to `nil`, whether they run to completion or exit early through
`break`.

```js
def iterate(items)
  foreach item in items
    puts(item)
  end
end

a = [1, 2, 3, 4, 5]

b = iterate(a)

// b is nil
```

Build up a value explicitly if you need one out of a loop:

```js
def doubled(items)
  result = []
  foreach item in items
    result.append(item * 2)
  end
  return result
end
```

Until `0.24`, `foreach` returned the value it was iterating, so `b` above was
`[1, 2, 3, 4, 5]` — the same array that went in. That did not hold once the
loop hit a `break`, which produced `nil` instead, and `while` never returned
anything but `nil`.
