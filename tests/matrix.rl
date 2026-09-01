m1 = [[1, 2], [3, 4]].to_matrix()
print(m1)

m2 = [[5, 6], [7, 8]].to_matrix()
result = m1 * m2
print(result)

sum = m1 + m2
print(sum)

diff = m2 - m1
print(diff)

print(result.to_array())
