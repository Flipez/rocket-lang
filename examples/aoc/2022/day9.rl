def sign(i)
  if (i == 0)
    return i
  end
  return i / Math.abs(i.plz_f())
end

def move_tail(head, tail)
  dist_y = head[0]-tail[0]
  dist_x = head[1]-tail[1]
  if (Math.abs(dist_y.plz_f()) >= 2 || Math.abs(dist_x.plz_f()) >= 2)
    tail[0] = tail[0] + sign(dist_y).plz_i()
    tail[1] = tail[1] + sign(dist_x).plz_i()
  end
  return tail
end

def pointString(pt)
  return "%d|%d".format(pt[0], pt[1])
end

def rope_simulation(moves, head, tail)
  visited = {}
  visited[pointString(tail[-1])] = true
  foreach move in moves
    parts = move.split()
    num = parts[1].plz_i()
    foreach i in num
      if (parts[0] == "U")
        head[0] = head[0] + 1
      end
      if (parts[0] == "D")
        head[0] = head[0] - 1
      end
      if (parts[0] == "R")
        head[1] = head[1] + 1
      end
      if (parts[0] == "L")
        head[1] = head[1] - 1
      end

      foreach index, tail_piece in tail
        if (index == 0)
          tail[index] = move_tail(head, tail_piece)
          next
        end
        tail[index] = move_tail(tail[index-1], tail_piece)
      end

      visited[pointString(tail[-1])] = true
    end
  end
  return visited.keys().size()
end

def part1(lines)
  head = [0,0]
  tail = [[0,0]]
  return rope_simulation(lines, head, tail)
end

def part2(lines)
  head = [0,0]
  tail = []
  foreach i in 9
    tail.yoink([0,0])
  end
  return rope_simulation(lines, head, tail)
end

lines = IO.open("day9.txt").lines()

puts(part1(lines))
puts(part2(lines))
