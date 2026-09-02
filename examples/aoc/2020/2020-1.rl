import "../util" as util

def part1(lines)
  expenses = []
  for line in lines
    expenses.append!(line.to_integer())
  end
  for i, a in expenses
    for b in expenses[i+1:]
      if (a+b == 2020)
        return a*b
      end
    end
  end
end

def part2(lines)
  expenses = []
  for line in lines
    expenses.append!(line.to_integer())
  end
  for i, a in expenses
    for b in expenses[i+1:]
      for c in expenses[i+2:]
        if (a+b+c == 2020)
          return a*b*c
        end
      end
    end
  end
end

input = IO.open("day1.txt").lines()

print(part1(input))
print(part2(input))
