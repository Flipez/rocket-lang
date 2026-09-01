---
title: "Foreach"
menu:
  docs:
    parent: "controls"
---
# Foreach
For loops allow to iterate over different sets of data and perform actions based on them.

```js
// read a file with numbers in it (file content will always be represented by strings)
// .lines() splits the lines of the file into an array
input = open("examples/aoc/2021/day-1/input").lines()

// define temporary array
a = []

foreach i, number in input
  // read each line into temporary array and cast it into an integer
  a.append(number.trim().to_integer())
end

// assign temporary array to input array
input = a
```

## Return Value
Loops evaluate to `nil`, whether they run to completion or exit early through
`break`.

```js
def iterate(items)
  foreach item in items
    print(item)
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

## Scope of the loop variable

The loop variable is created in the loop's own scope, so it does not exist
after the loop finishes:

```js
foreach j in [1, 2, 3]
end

print(j)  // ERROR: identifier not found: j
```

If a variable of that name already exists outside the loop, though, the loop
assigns to *that* variable rather than creating a new one, and it keeps the
last value the loop gave it:

```js
i = 100

foreach i in [1, 2, 3]
end

print(i)  // 3, not 100
```

This follows from how assignment works generally: assigning to a name that
already exists further out updates it in place instead of shadowing it. Use a
loop variable name that is not already taken if you need to keep the outer
value.

## Using an integer
Count form zero to a given number (excluding):

```js
🚀 > foreach i in 5
  print(i)
end

0
1
2
3
4
=> nil
```

## Using a string
Iterate over a string:

```js
🚀 > foreach i in "test"
  print(i)
end

t
e
s
t
=> nil
```

## Using break and next
It is possible to use `next` or `break` inside a loop.

```js
foreach i in 5
  if (i == 2)
    next
  end
  print(i)
end

foreach i in 5
  if (i == 2)
    break
  end
  print(i)
end

// Returns
0
1
3
4
0
1
nil
```

## Using range

You can use the so called `rocket range` operator to create an individual range with optional stepping:

```js
foreach i in 0 -> 5
  print(i)
end

// outputs
0
1
2
3
4
```

There is also an inclusive alternative:

```js
foreach i in 0 => 5
  print(i)
end

// outputs
0
1
2
3
4
5
```

### Stepping

You can specify stepping to change the default of `1`

```js
foreach i in 0 -> 5 ^ 2
  print(i)
end

// outputs
0
2
4
```

### Reverse

Ranges do support going from a higher value to a lower one

```js
foreach i in 5 -> 0 ^ 2
  print(i)
end

// outputs
5
3
1
```