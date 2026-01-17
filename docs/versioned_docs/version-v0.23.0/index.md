---
sidebar_position: 0
---
# Getting Started

RocketLang as of version 0.9.5 is the full (as in the book was worked through) version of [MonkeyLang](https://monkeylang.org/) and
is then being extended with various useful and not so useful features.

# Latest Version

[![GitHub release](https://img.shields.io/github/release/flipez/rocket-lang.svg)](https://github.com/flipez/rocket-lang/releases/)

# Quick Start
```js
input = open("examples/aoc/2021/day-1/input").lines()


a = []
foreach i, number in input
  a.push(number.strip().to_i())
end
input = a

increase = 0
foreach i, number in input
  if (number > input[i-1])
    increase = increase + 1
  end
end
puts(increase + 1)

increase = 0
foreach i, number in input
  sum = number + input[i+1] + input[i+2]
  sum_two = input[i+1] + input[i+2] + input[i+3]

  if (sum_two > sum)
    increase = increase + 1
  end
end
puts(increase + 1)
```

# Help
You can launch RocketLang with `-h` or `--help` to get an overview about the cli capabilities.

```zsh
$ rocket-lang -h
Usage: rocket-lang [flags] [program file] [arguments]

Available flags:
  -e, --exec string   Runs the given code.
  -v, --version       Prints the version and build date.
```