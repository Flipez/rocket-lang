a = if (false)
  puts("true1")
else
  break
  puts("false1")
end

a = 2
if (a == 1)
  puts("true2")
elif (a == 3)
  puts("false2")
elif (a == 2)
  puts(2)
end

if (true)
  puts("true3")
end

puts(a)