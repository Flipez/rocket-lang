scores = {
  "A": 1,
  "B": 2,
  "C": 3,
  "X": 1,
  "Y": 2,
  "Z": 3,
  "draw": 3,
  "win": 6,
}

win = {
  "A": "Y",
  "B": "Z",
  "C": "X",
}

draw = {
  "A": "X",
  "B": "Y",
  "C": "Z",
}

defeat = {
  "A": "Z",
  "B": "X",
  "C": "Y",
}

def part1(lines)
  total = 0
  foreach line in lines
    fields = line.split()
    total = total + scores[fields[1]]
    if (draw.include?(fields[0]) && draw[fields[0]] == fields[1])
      total = total + scores["draw"]
    end
    if (win.include?(fields[0]) && win[fields[0]] == fields[1])
      total = total + scores["win"]
    end
  end
  return total
end

def part2(lines)
  total = 0
  foreach line in lines
    fields = line.split()
    if (fields[1] == "X")
      total = total + scores[defeat[fields[0]]]
    end
    if (fields[1] == "Y")
      total = total + scores["draw"]
      total = total + scores[draw[fields[0]]]
    end
    if (fields[1] == "Z")
      total = total + scores["win"]
      total = total + scores[win[fields[0]]]
    end
  end
  return total
end

input = IO.open("day2.txt").lines()

puts(part1(input))
puts(part2(input))
