numbers = [1, 2, 3, 4]

print(numbers.map(def(x) x * 2 end))
print(numbers.filter(def(x) x % 2 == 0 end))
print(numbers.reduce(0, def(sum, x) sum + x end))
