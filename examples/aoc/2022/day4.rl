def parse_line(line)
  parts = line.split(",")
  leftPair = parts[0].split("-")
  rightPair = parts[1].split("-")
  return [
    [leftPair[0].plz_i(), leftPair[1].plz_i()],
    [rightPair[0].plz_i(), rightPair[1].plz_i()]
  ]
end

def between(i, min, max)
  if (min <= i && i <= max)
    return true
  end
  return false
end

def full_overlap(nums)
  if (between(nums[1][0], nums[0][0], nums[0][1]) && between(nums[1][1], nums[0][0], nums[0][1]))
    return true
  end

  if (between(nums[0][0], nums[1][0], nums[1][1]) && between(nums[0][1], nums[1][0], nums[1][1]))
    return true
  end

  return false
end

def partial_overlap(nums)
  if (between(nums[1][0], nums[0][0], nums[0][1]) && nums[1][1] >= nums[0][1])
    return true
  end

  if (between(nums[0][0], nums[1][0], nums[1][1]) && nums[0][1] >= nums[1][1])
    return true
  end

  return false
end

def part1(lines)
  total = 0
  foreach line in lines
    nums = parse_line(line)

    if (full_overlap(nums))
      total = total + 1
    end
  end
  return total
end

def part2(lines)
  total = 0
  foreach line in lines
    nums = parse_line(line)

    if (full_overlap(nums))
      total = total + 1
      next
    end

    if (partial_overlap(nums))
      total = total + 1
    end
  end
  return total
end

input = IO.open("day4.txt").lines()

puts(part1(input))
puts(part2(input))
