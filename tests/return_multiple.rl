def basic()
  return 1, 2, 3
end

a, b, c = basic()
print(a)
print(b)
print(c)

def strings()
  return "hello", "world", "!"
end

x, y, z = strings()
print(x + " " + y + z)

def expressions()
  n = 10
  return n, n * 2, n * 3
end

one, two, three = expressions()
print(one)
print(two)
print(three)

def as_array()
  return 100, 200
end

arr = as_array()
print(arr)

nil
