---
title: Home
type: docs
---
# Welcome to the Home of 🚀🇱🅰🆖

{{< columns >}}
## About

RocketLang as of version 0.9.5 is the full (as in the book was worked through) version of [MonkeyLang](https://monkeylang.org/) and
is then being extended with various useful and not so useful features.

<--->

## Latest Version

[![GitHub release](https://img.shields.io/github/release/flipez/rocket-lang.svg)](https://github.com/flipez/rocket-lang/releases/)
[![Average time to resolve an issue](http://isitmaintained.com/badge/resolution/flipez/rocket-lang.svg)](https://github.com/flipez/rocket-lang)
[![Percentage of issues still open](http://isitmaintained.com/badge/open/flipez/rocket-lang.svg)](https://github.com/flipez/rocket-lang)

{{< /columns >}}
## Quick Start

Get started with 🚀🇱🅰🆖 quickly with these examples:

```js
let input = open("examples/aoc/2021/day-1/input").lines()


let a = []
foreach i, number in input {
  a.yoink(number.strip().plz_i())
}
input = a

let increase = 0
foreach i, number in input {
  if (number > input[i-1]) {
    increase = increase + 1
  }
}
puts(increase + 1)

increase = 0
foreach i, number in input {
  let sum = number + input[i+1] + input[i+2]
  let sum_two = input[i+1] + input[i+2] + input[i+3]
  
  if (sum_two > sum) {
    increase = increase + 1
  }
}
puts(increase + 1)
```