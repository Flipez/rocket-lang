import("util")

def part1(lines)
  valid = 0
  foreach line in lines
    parts = line.split()

    char = parts[1][:-1]
    pass = parts[2]

    min = parts[0].split("-")[0].to_i()
    max = parts[0].split("-")[1].to_i()

    charCount = pass.count(char)
    if (charCount >= min && charCount <= max)
      valid = valid + 1
    end
  end
  return valid
end

def part2(lines)
  valid = 0
  foreach line in lines
    parts = line.split()

    char = parts[1][:-1]
    pass = parts[2]

    firstPos = parts[0].split("-")[0].to_i()
    secondPos = parts[0].split("-")[1].to_i()

    a = pass[firstPos-1]
    b = pass[secondPos-1]
    if (a != b && (a == char || b == char))
      valid = valid + 1
    end
  end
  return valid
end

input = IO.open("day2.txt").lines()

puts(part1(input))
puts(part2(input))
