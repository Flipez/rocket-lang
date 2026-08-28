def Classify(x)
  result = "unknown"
  if (x == 1)
    result = "one"
  else if (x == 2)
    result = "two"
  else
    result = "other"
  end
  return result
end

def After()
  return "after-ok"
end

puts(Classify(1))
puts(Classify(2))
puts(Classify(3))
puts(After())
