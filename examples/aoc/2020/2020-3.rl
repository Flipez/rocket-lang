import("util")

def countTrees(rows, right, down)
  row = 0
  col = 0
  trees = 0
  while (row+down <= rows.size()-1)
    col = col + right
    row = row + down
    if (col >= rows[row].size())
      col = col % rows[row].size()
    end
    buf = rows[row]
    if (buf[col] == "#")
      trees = trees + 1
    end
  end
  return trees
end

def part1(lines)
  return countTrees(lines, 3, 1)  
end

def part2(lines)
  slopes = [
    [1, 1],
    [3, 1],
    [5, 1],
    [7, 1],
    [1, 2]
  ]
  totalTrees = 0
  foreach slope in slopes
    trees = countTrees(lines, slope[0], slope[1])
    if (totalTrees == 0)
      totalTrees = trees
    else
      totalTrees = totalTrees * trees
    end
  end
  return totalTrees
end

input = IO.open("day3.txt").lines()

puts(part1(input))
puts(part2(input))
