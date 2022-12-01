def parseCalories(invs)
  elves = []
  total = 0
  foreach line in invs
    if (line == "")
      elves.yoink(total)
      total = 0
      next
    end
    total = total + line.plz_i()
  end
  elves.yoink(total)
  elves.sort()
  elves.reverse()
  return elves
end

def part1(lines)
  elves = parseCalories(lines)
  return elves[0]
end

def part2(lines)
  elves = parseCalories(lines)
  return elves[0] + elves[1] + elves[2]
end

input = IO.open("day1.txt").lines()

puts(part1(input))
puts(part2(input))
