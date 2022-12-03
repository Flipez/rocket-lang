def convert(str)
  alpha = "abcdefghijklmnopqrstuvwxyz"
  alpha = alpha + alpha.upcase()
  return alpha.find(str)+1
end

def part1(lines)
  total = 0
  foreach line in lines
    half = line.size()/2
    left = line[:half]
    right = line[half:]
    foreach item in left
      if (right.count(item) > 0)
        total = total + convert(item)
        break
      end
    end
  end
  return total
end

def shortest(strs)
  s = strs[0]
  l = s.size()
  foreach str in strs[1:]
    if (str.size() < l)
      s = str
      l = s.size()
    end
  end
  return s
end

def part2(lines)
  total = 0
  foreach chunk in lines.slices(3)
    foreach item in shortest(chunk)
      if (chunk[0].count(item) > 0 && chunk[1].count(item) > 0 && chunk[2].count(item) > 0)
        total = total + convert(item)
        break
      end
    end
  end
  return total
end

input = IO.open("day3.txt").lines()

puts(part1(input))
puts(part2(input))
