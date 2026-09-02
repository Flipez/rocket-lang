input = IO.open("input").lines()

depth = 0
hor = 0
aim = 0

for i, line in input
  command = line.split(" ")[0]
  value = line.trim().split(" ")[1].to_integer()
  if (command == "forward")
    hor = hor + value
    depth = depth + (value * aim)
  end

  if (command == "down")
    aim = aim + value
  end

  if (command == "up")
    aim = aim - value
  end
end

print(hor * depth)