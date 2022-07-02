foreach i in 5
  if (i == 2)
    next("next")
  end
  puts(i)
end

foreach i in 5
  if (i == 2)
    break("break")
  end
  puts(i)
end