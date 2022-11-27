import("util")

def part1(lines)
  expenses = []
  foreach line in lines
    expenses.yoink(line.plz_i())
  end
  foreach i, a in expenses
    foreach b in expenses[i+1:]
      if (a+b == 2020)
        return a*b
      end
    end
  end
end

def part2(lines)
  expenses = []
  foreach line in lines
    expenses.yoink(line.plz_i())
  end
  foreach i, a in expenses
    foreach b in expenses[i+1:]
      foreach c in expenses[i+2:]
        if (a+b+c == 2020)
          return a*b*c
        end
      end
    end
  end
end

input = IO.open("day1.txt").lines()

puts(part1(input))
puts(part2(input))
