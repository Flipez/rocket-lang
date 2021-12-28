import "helper.rl"

let input = open("examples/aoc/2021/day-1/input").lines()

// Part 1

let increase = 0

input = arrayInt(input)

foreach i, number in input {
  if (number > input[i-1]) {
    increase = increase + 1
  }
}

puts(increase + 1)

// Part 2

increase = 0

foreach i, number in input {
  let sum = number + input[i+1] + input[i+2]
  let sum_two = input[i+1] + input[i+2] + input[i+3]
  
  if (sum_two > sum) {
    increase = increase + 1
  }
}

puts(increase + 1)