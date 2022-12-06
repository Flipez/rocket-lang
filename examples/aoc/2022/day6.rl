def check(buff, window)
  foreach idx in buff.size()
    m = {}
    foreach i in window
      m[buff[idx+i]] = true
    end
    if (m.keys().size() == window)
      return idx+window
    end
  end
end

def part1(lines)
  return check(lines[0], 4)
end

def part2(lines)
  return check(lines[0], 14)
end

input = IO.open("day6.txt").lines()

puts(part1(input))
puts(part2(input))
