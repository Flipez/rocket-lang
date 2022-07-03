"// A solution for part 1 of https://adventofcode.com/2015/day/1"
"// Last Updated: 2022.01.18"
"// RocketLang Version: 0.14.1"
"// ------------------------------------"
part_one = def (input) 
  "// There is a foreach loop now but I wanted to stick as close to the original version as possible"
  calc(input, input.size(), 0)
end

calc = def (input, idx, floor)
  char_to_value = {
    "(": 1,
    ")": -1
  }
  
  if (idx == 0)
    return floor
  end
  
  new_idx = idx - 1
  delta = char_to_value[input[new_idx]]
  if (delta == nil)
    raise(1, new_idx.plz_s())
  end
  
  calc(input, new_idx, floor + delta)
end

"// Test some inputs to check that our code is correct..."
if (part_one("(())") != 0)
  puts("Assertion ((( => 0 failed")
end

"// ... for multiple known results."
if (part_one(")())())") != -3)
  puts("Assertion )())()) => -3 failed")
end

"// We can now read data from files so using 'real' input data is much easier"
real_input = open("examples/aoc/2015/day1.input").content().strip()

puts("Solution Day 1 Part 1: ")
puts(part_one(real_input))
