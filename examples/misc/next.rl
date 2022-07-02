a = foreach i in 5
  if (i == 2)
    next("test")
  end
  puts(i)
end