numbers = [1, 2, 3, 4]

puts(numbers.map(def(x) x * 2 end))
puts(numbers.filter(def(x) x % 2 == 0 end))
puts(numbers.reduce(0, def(sum, x) sum + x end))
