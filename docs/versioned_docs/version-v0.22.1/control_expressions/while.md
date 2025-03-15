---
title: "While"
menu:
  docs:
    parent: "controls"
---
# While
While loops will run as long as the condition is truthy.

Print numbers from 0 to 3:

```js
🚀 > a = 0
🚀 > while (a != 4)
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
while (i < 10)
  if (i < 3)
    i = i + 1
    next
  end
  puts(i)
  if (i == 6)
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
Loops do return the variable they are iterating after the last loop.

```js
def iterate(items)
  foreach item in items
    puts(item)
  end
end

a = [1,2,3,4,5]

b = iterate(a)

// b is now [1,2,3,4,5]
```