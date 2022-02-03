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

foreach i, number in input {
  // read each line into temporary array and cast it into an integer
  a.yoink(number.strip().plz_i())
}

// assign temporary array to input array
input = a
```


Count form zero to a given number (excluding):

```js
🚀 > foreach i in 5 { puts(i) }
0
1
2
3
4
=> 5
```

Iterate over a string:

```js
🚀 > foreach i in "test" { puts(i) }
"t"
"e"
"s"
"t" 
=> "test"
```
