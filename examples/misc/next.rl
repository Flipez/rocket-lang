a = foreach i in 5
  if (i == 2)
    next
  end
  puts(i)
end