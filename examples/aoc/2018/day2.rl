import "../util" as util
def countLetters(line)
  return [doubleLetters, trippleLetters]
end

def part1(lines)
  doubles = 0
  tripples = 0
  for line in lines
    line.trim!()
    letters = {}
    for letter in line
      if (letters[letter] == nil)
        letters[letter] = 1
      else
        letters[letter] = letters[letter] + 1
      end
    end
    doubleLetters = false
    trippleLetters = false
    for count in letters.values()
      if (count == 2)
        doubleLetters = true
      end
      if (count == 3)
        trippleLetters = true
      end
    end
    if (doubleLetters)
      doubles = doubles + 1
    end
    if (trippleLetters)
      tripples = tripples + 1
    end
  end
  return doubles * tripples
end

def part2(lines)
  for idx, id in lines
    for i in lines[idx:]
      this = id
      naxt = i
      differ = 0
      same = ""
      for idx, letter in this
        if (letter != naxt[idx])
          differ = differ + 1
        else
          same = same + letter
        end
      end
      if (differ == 1)
        return same
      end
    end
  end
end

input = IO.open("day2.txt").lines()

print(part1(input))
print(part2(input))
