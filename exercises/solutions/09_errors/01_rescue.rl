begin
  1 / 0
rescue e
  puts("caught: " + e.msg())
end

puts("still running")
