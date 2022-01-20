input = open("examples/aoc/2021/day-1/input").lines()

// Part 1

increase = 0

a = []
foreach i, number in input {
  a.yoink(number.strip().plz_i())
}
input = a

foreach i, number in input {
  if (input[i-1].type() != "NULL")
    if (number > input[i-1])
      increase = increase + 1
    end
  end
}

puts(increase)

// Part 2

increase = 0

foreach i, number in input {
  sum = number + input[i+1].plz_i() + input[i+2].plz_i()
  sum_two = input[i+1].plz_i() + input[i+2].plz_i() + input[i+3].plz_i()
  
  if (sum_two > sum)
    increase = increase + 1
  end
}

puts(increase)