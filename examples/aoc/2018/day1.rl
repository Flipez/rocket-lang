def parseChanges(lines)
  changes = []
  for idx, line in lines
    changes.append!(line.trim().to_integer())
  end
  return changes
end

def part1(input)
  freq = 0
  for idx, change in parseChanges(input)
    freq = freq + change
  end
  return freq
end

def part2(input)
  freqs = [0]
  changes = parseChanges(input)
  while (true)
    for idx, change in changes
      lastFreq = freqs[-1] + change
      if (freqs.index_of(lastFreq) != -1)
        return lastFreq
      end
      freqs.append!(lastFreq)
    end
  end
end

input = IO.open("day1.txt").lines()

print(part1(input))
print(part2(input))
