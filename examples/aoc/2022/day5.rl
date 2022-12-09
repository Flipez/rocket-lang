def chars(str)
  a = []
  foreach c in str
    a.yoink(c)
  end
  return a
end

def parse_input(lines)
  sep = lines.index("")
  stack_lines = lines[:sep]
  instr_lines = lines[sep:]

  stack_lines = stack_lines.reverse()

  labels = chars(stack_lines[0]).slices(4)
  stacks = []
  foreach i in labels.size()
    stacks.yoink([])
  end

  foreach line in stack_lines[1:]
    groups = chars(line).slices(4)
    foreach idx, group in groups
      i = group[1].ascii()
      if (i >= 65 && i <= 90)
        stacks[idx].yoink(group[1])
      end
    end
  end


  return [stacks, instr_lines]
end

def parse_instruction(instr)
  fields = instr.split()
  return [fields[1].plz_i(), fields[3].plz_i()-1, fields[5].plz_i()-1]
end

def part1(lines)
  result = parse_input(lines)
  stacks = result[0]
  instructions = result[1]
  foreach instruction in instructions[1:-1]
    result = parse_instruction(instruction)
    count = result[0]
    from = result[1]
    to = result[2]
    from_stack = stacks[from]
    to_stack = stacks[to]

    crates = from_stack[-count:].reverse()
    from_stack = from_stack[:-count]

    foreach crate in crates
      to_stack.yoink(crate)
    end

    stacks[from] = from_stack
    stacks[to] = to_stack
  end
  top = ""
  foreach _, stack in stacks
    top = top + stack[-1]
  end
  return top
end

def part2(lines)
  result = parse_input(lines)
  stacks = result[0]
  instructions = result[1]
  foreach instruction in instructions[1:-1]
    result = parse_instruction(instruction)
    count = result[0]
    from = result[1]
    to = result[2]
    from_stack = stacks[from]
    to_stack = stacks[to]
    
    crates = from_stack[-count:]
    from_stack = from_stack[:-count]

    foreach crate in crates
      to_stack.yoink(crate)
    end

    stacks[from] = from_stack
    stacks[to] = to_stack
  end
  top = ""
  foreach _, stack in stacks
    top = top + stack[-1]
  end
  return top
end


lines = IO.open("day5.txt").lines()

puts(part1(lines))
puts(part2(lines))
