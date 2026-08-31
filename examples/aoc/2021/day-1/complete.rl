input = IO.open("input").lines()

// Part 1

increase = 0

a = []
foreach i, number in input
  a.append(number.trim().to_integer())
end
input = a

foreach i, number in input
  if (input[i-1] != nil)
    if (number > input[i-1])
      increase = increase + 1
    end
  end
end

print(increase)

// Part 2

increase = 0

// A three-wide window rises exactly when the value entering it is larger than
// the one leaving, so the sums need not be computed. Reading past the end used
// to work only because nil.to_integer() was 0.
foreach i, number in input
  if (i + 3 < input.size())
    if (input[i+3] > number)
      increase = increase + 1
    end
  end
end

print(increase)