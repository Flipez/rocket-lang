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

for instruction in instructions
  amount = instruction.split("from")[0].split("move")[-1].trim().to_integer()
  from = instruction.split("to")[0].split("from")[-1].trim().to_integer()
  to = instruction.split("to")[1].trim().to_integer()

  for i in amount
    item = stacks[from - 1].last()
    stacks[from - 1].remove_last!()
    stacks[to - 1].append!(item)
  end
end

result = ""
for stack in stacks
  result = result + stack[-1]
end
print("Part 1: " + result)

for instruction in instructions
  amount = instruction.split("from")[0].split("move")[-1].trim().to_integer()
  from = instruction.split("to")[0].split("from")[-1].trim().to_integer()
  to = instruction.split("to")[1].trim().to_integer()

  temp_stack = []

  for i in amount
    item = stacks2[from - 1].last()
    stacks2[from - 1].remove_last!()
    temp_stack.append!(item)
  end
  temp_stack.reverse!()
  for item in temp_stack
    stacks2[to - 1].append!(item)
  end

end

result = ""
for stack in stacks
  result = result + stack[-1]
end
print("Part 2: " + result)

nil