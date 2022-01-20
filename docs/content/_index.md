---
title: Introduction
type: docs
bookToC: false
---

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
```js
input = open("examples/aoc/2021/day-1/input").lines()


a = []
foreach i, number in input {
  a.yoink(number.strip().plz_i())
}
input = a

increase = 0
foreach i, number in input {
  if (number > input[i-1])
    increase = increase + 1
  end
}
puts(increase + 1)

increase = 0
foreach i, number in input {
  sum = number + input[i+1] + input[i+2]
  sum_two = input[i+1] + input[i+2] + input[i+3]
  
  if (sum_two > sum)
    increase = increase + 1
  end
}
puts(increase + 1)
```