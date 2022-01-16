input = open("examples/aoc/2021/day-1/input").lines()

// Part 1

increase = 0

a = []
foreach i, number in input {
  a.yoink(number.strip().plz_i())
}
input = a

foreach i, number in input {
  if (number > input[i-1])
    increase = increase + 1
  end
}

puts(increase)

// Part 2

increase = 0

foreach i, number in input {
  sum = number + input[i+1] + input[i+2]
  sum_two = input[i+1] + input[i+2] + input[i+3]
  
  if (sum_two > sum)
    increase = increase + 1
  end
}

puts(increase)