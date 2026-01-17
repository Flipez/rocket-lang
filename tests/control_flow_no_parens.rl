x = 5
if x < 3
  puts("x is less than 3")
elif x == 5
  puts("x equals 5")
else
  puts("x is greater than 5")
end

counter = 0
while counter < 3
  puts(counter)
  counter = counter + 1
end

a = "test"
if a.type() == "STRING" and a.size() > 0
  puts("a is a non-empty string")
end

nil
