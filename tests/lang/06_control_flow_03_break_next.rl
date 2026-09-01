numbers = [1, 2, 3, 4]

numbers.each(def(n)
  if n == 3
    break
  end
  print(n)
end)

for n in numbers
  if n == 2
    next
  end
  if n == 4
    break
  end
  print(n)
end
