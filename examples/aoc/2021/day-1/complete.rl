input = IO.open("examples/aoc/2021/day-1/input").lines()

// Part 1

increase = 0

a = []
foreach i, number in input
  a.push(number.strip().to_i())
end
input = a

foreach i, number in input
  if (input[i-1] != nil)
    if (number > input[i-1])
      increase = increase + 1
    end
  end
end

puts(increase)

// Part 2

increase = 0

foreach i, number in input
  sum = number + input[i+1].to_i() + input[i+2].to_i()
  sum_two = input[i+1].to_i() + input[i+2].to_i() + input[i+3].to_i()
  
  if (sum_two > sum)
    increase = increase + 1
  end
end

puts(increase)