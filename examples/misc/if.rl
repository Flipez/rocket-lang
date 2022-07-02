a = if (false)
  puts("true1")
else
  break("test")
  puts("false1")
end

a = 2
if (a == 1)
  puts("true2")
else if (a == 3)
  puts("false2")
else if (a == 2)
  puts(2)
end

if (true)
  puts("true3")
end

puts(a)