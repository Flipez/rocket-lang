import("util")
def countLetters(line)
  return [doubleLetters, trippleLetters]
end

def part1(lines)
  doubles = 0
  tripples = 0
  foreach line in lines
    line.strip!()
    letters = {}
    foreach letter in line
      if (letters[letter] == nil)
        letters[letter] = 1
      else
        letters[letter] = letters[letter] + 1
      end
    end
    doubleLetters = false
    trippleLetters = false
    foreach count in letters.values()
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
  foreach idx, id in lines
    foreach i in lines[idx:]
      this = id
      naxt = i
      differ = 0
      same = ""
      foreach idx, letter in this
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

puts(part1(input))
puts(part2(input))
