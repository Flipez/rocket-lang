def values(lines)
  vals = [0]
  foreach line in lines
    vals.yoink(0)
    fields = line.split()
    if (fields[0] == "addx")
      vals.yoink(fields[1].plz_i())
    end
  end
  return vals
end

def part1(lines)
  x = 1
  sum = 0
  foreach idx, val in values(lines)
    if (idx == 20 || idx == 60 || idx == 100 || idx == 140 || idx == 180 || idx == 220)
      sum = sum + (idx * x)
    end
    x = x + val
  end
  return sum
end

def part2(lines)
  x = 1
  out = ""
  foreach idx, val in values(lines)
    x = x + val
    v = idx%40
    if (v == 0)
      puts(out)
      out = ""
    end
    if (v == x-1 || v == x || v == x+1)
      out = out + "#"
    else
      out = out + " "
    end
  end
  return nil
end

lines = IO.open("day10.txt").lines()

puts(part1(lines))
puts(part2(lines))
