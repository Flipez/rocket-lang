m = [[1, 2], [3, 4]].to_m()

puts(m.t().to_a())
puts(m.to_a())

wide = [[1, 2, 3], [4, 5, 6]].to_m()
puts(wide.transpose().shape())
