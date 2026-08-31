m1 = [[1, 2], [3, 4]].to_matrix()
puts(m1)

m2 = [[5, 6], [7, 8]].to_matrix()
result = m1 * m2
puts(result)

sum = m1 + m2
puts(sum)

diff = m2 - m1
puts(diff)

puts(result.to_a())
