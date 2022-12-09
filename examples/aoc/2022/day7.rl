def join(array, sep)
  if (array.size() == 0)
    return ""
  end
  if (array.size() == 1)
    return array[0]
  end
  str = array[0]
  foreach item in array[1:]
    if (str[-1] == sep)
      str = str + item
    else
      str = str + sep + item
    end
  end
  return str
end

def copy(array)
  a = []
  foreach item in array
    a.yoink(item)
  end
  return a
end

def parse_input(lines)
  blocks = []
  block = []
  foreach line in lines
    if (line[0] == "$")
      blocks.yoink(block)
      block = [line]
      next
    end
    block.yoink(line)
  end
  blocks.yoink(block)
  return blocks
end

def replay_history(history)
  fs = {"/": 0}
  path = ["/"]
  foreach block in history[2:]
    cmd = block[0]
    output = block[1:]

    fields = cmd.split()
    if (fields[1] == "cd")
      if (fields[2] == "..")
        path = path[:-1]
      else
        path.yoink(fields[2])
        fs[join(path, "/")] = 0
      end
    end

    if (fields[1] == "ls")
      foreach line in output
        tmp = copy(path)
        parts = line.split()
        if (parts[0] == "dir")
          tmp.yoink(parts[1])
          fs[join(tmp, "/")] = 0
        else
          p = join(path, "/")
          fs[p] = fs[p] + parts[0].plz_i()
          while (tmp.size() > 1)
            tmp = tmp[:-1]
            p = join(tmp, "/")
            fs[p] = fs[p] + parts[0].plz_i()
          end
        end
      end
    end
  end
  return fs
end

def part1(lines)
  fs = replay_history(parse_input(lines))

  sum = 0
  foreach path, size in fs
    if (size < 100000)
      sum = sum + size
    end
  end

  return sum
end

def part2(lines)
  fs = replay_history(parse_input(lines))

  unused = 70000000 - fs["/"]
  needed = 30000000 - unused

  min = 70000000
  foreach path, size in fs
    if (size >= needed && size < min)
      min = size
    end
  end

  return min
end


lines = IO.open("day7.txt").lines()

puts(part1(lines))
puts(part2(lines))
