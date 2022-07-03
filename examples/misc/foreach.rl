a = foreach i in 5
  if (i == 2)
    break()
  end
  puts(i)
end

puts(a.type())