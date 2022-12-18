def parseChanges(lines)
  changes = []
  foreach idx, line in lines
    changes.yoink(line.strip().to_i())
  end
  return changes
end

def part1(input)
  freq = 0
  foreach idx, change in parseChanges(input)
    freq = freq + change
  end
  return freq
end

def part2(input)
  freqs = [0]
  changes = parseChanges(input)
  while (true)
    foreach idx, change in changes
      lastFreq = freqs[-1] + change
      if (freqs.index(lastFreq) != -1)
        return lastFreq
      end
      freqs.yoink(lastFreq)
    end
  end
end

input = IO.open("day1.txt").lines()

puts(part1(input))
puts(part2(input))
