m1 = Matrix([[1, 2], [3, 4]])
puts(m1)

m2 = Matrix([[5, 6], [7, 8]])
result = m1 * m2
puts(result)

sum = m1 + m2
puts(sum)

diff = m2 - m1
puts(diff)

puts(result.to_a())
