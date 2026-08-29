foreach n in 1 => 15
  if n % 15 == 0
    puts("fizzbuzz")
  elif n % 3 == 0
    puts("fizz")
  elif n % 5 == 0
    puts("buzz")
  else
    puts(n)
  end
end
