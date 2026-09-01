m = [[1, 2], [3, 4]].to_matrix()

print(m.transpose().to_array())
print(m.to_array())

wide = [[1, 2, 3], [4, 5, 6]].to_matrix()
print(wide.transpose().shape())
