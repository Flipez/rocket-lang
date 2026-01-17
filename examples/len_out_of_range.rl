instructions = [
"move 1 from 2 to 1",
"move 3 from 1 to 3",
"move 2 from 2 to 1",
"move 1 from 1 to 2"]

stacks = [
  ["Z", "N"],
  ["M", "C", "D"],
  ["P"]
]

stacks2 = stacks

foreach instruction in instructions
  amount = instruction.split("from")[0].split("move")[-1].strip().to_i()
  from = instruction.split("to")[0].split("from")[-1].strip().to_i()
  to = instruction.split("to")[1].strip().to_i()

  foreach i in amount
    stacks[to - 1].push(stacks[from - 1].pop())
  end
end

result = ""
foreach stack in stacks
  result = result + stack[-1]
end
puts("Part 1: " + result)

foreach instruction in instructions
  amount = instruction.split("from")[0].split("move")[-1].strip().to_i()
  from = instruction.split("to")[0].split("from")[-1].strip().to_i()
  to = instruction.split("to")[1].strip().to_i()

  temp_stack = []

  foreach i in amount
    temp_stack.push(stacks2[from - 1].pop())
  end
  temp_stack.reverse()
  foreach item in temp_stack
    stacks2[to - 1].push(item)
  end

end

result = ""
foreach stack in stacks
  result = result + stack[-1]
end
puts("Part 2: " + result)

nil