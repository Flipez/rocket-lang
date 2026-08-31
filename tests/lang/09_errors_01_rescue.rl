begin
  1 / 0
rescue e
  puts("caught: " + e.message())
end

puts("still running")
