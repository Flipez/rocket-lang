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
  a.push(number.strip().to_i())
end

// assign temporary array to input array
input = a
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

## Using an integer
Count form zero to a given number (excluding):

```js
🚀 > foreach i in 5
  puts(i)
end

0
1
2
3
4
=> 5
```

## Using a string
Iterate over a string:

```js
🚀 > foreach i in "test" 
  puts(i)
end

"t"
"e"
"s"
"t" 
=> "test"
```

## Using break and next
It is possible to use `next` or `break` inside a loop.

```js
foreach i in 5
  if (i == 2)
    next
  end
  puts(i)
end

foreach i in 5
  if (i == 2)
    break
  end
  puts(i)
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
  puts(i)
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
  puts(i)
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
  puts(i)
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
  puts(i)
end

// outputs
5
3
1
```