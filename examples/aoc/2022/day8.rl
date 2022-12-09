def parse_input(lines)
  grid = []
  foreach line in lines
    row = []
    foreach char in line
      row.yoink(char.plz_i())
    end
    grid.yoink(row)
  end
  return grid
end

def max(nums)
  i = 0
  foreach num in nums
    if (num > i)
      i = num
    end
  end
  return i
end

def part1(lines)
  grid = parse_input(lines)

  row_len = grid.size()
  col_len = grid[0].size()
  visible = 2*row_len + 2*col_len - 4

  y = 1
  while (y < col_len-1)
    x = 1
    while (x < row_len-1)
      tree = grid[y][x]

      // look from left edge
      if (max(grid[y][:x]) < tree)
        visible = visible + 1
        x = x + 1
        next
      end

      // look from upper edge
      trees = []
      foreach i in y
        trees.yoink(grid[i][x])
      end
      if (max(trees) < tree)
        visible = visible + 1
        x = x + 1
        next
      end

      // look from right edge
      if (max(grid[y][x+1:]) < tree)
        visible = visible + 1
        x = x + 1
        next
      end

      // look from bottom edge
      trees = []
      foreach i in col_len
        if (i <= y)
          next
        end
        trees.yoink(grid[i][x])
      end
      if (max(trees) < tree)
        visible = visible + 1
        x = x + 1
        next
      end

      x = x + 1
    end
    y = y + 1
  end

  return visible
end

def part2(lines)
  grid = parse_input(lines)

  row_len = grid.size()
  col_len = grid[0].size()
  highest = 0

  y = 1
  while (y < col_len-1)
    x = 1
    while (x < row_len-1)
      tree = grid[y][x]
      scores = []

      // look up
      count = 0
      i = y - 1
      while (i >= 0)
        count = count + 1
        if (grid[i][x] >= tree)
          break
        end
        i = i - 1
      end
      scores.yoink(count)

      // look left
      count = 0
      i = x - 1
      while (i >= 0)
        count = count + 1
        if (grid[y][i] >= tree)
          break
        end
        i = i - 1
      end
      scores.yoink(count)

      // look down
      count = 0
      i = y + 1
      while (i < col_len)
        count = count + 1
        if (grid[i][x] >= tree)
          break
        end
        i = i + 1
      end
      scores.yoink(count)

      // look right
      count = 0
      i = x + 1
      while (i < row_len)
        count = count + 1
        if (grid[y][i] >= tree)
          break
        end
        i = i + 1
      end
      scores.yoink(count)

      score = scores[0]*scores[1]*scores[2]*scores[3]
      if (score > highest)
        highest = score
      end

      x = x + 1
    end
    y = y + 1
  end

  return highest
end

lines = IO.open("day8.txt").lines()

puts(part1(lines))
puts(part2(lines))
