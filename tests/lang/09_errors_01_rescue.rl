begin
  1 / 0
rescue e
  print("caught: " + e.message())
end

print("still running")
