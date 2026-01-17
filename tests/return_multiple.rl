def basic()
  return 1, 2, 3
end

a, b, c = basic()
puts(a)
puts(b)
puts(c)

def strings()
  return "hello", "world", "!"
end

x, y, z = strings()
puts(x + " " + y + z)

def expressions()
  n = 10
  return n, n * 2, n * 3
end

one, two, three = expressions()
puts(one)
puts(two)
puts(three)

def as_array()
  return 100, 200
end

arr = as_array()
puts(arr)

nil
