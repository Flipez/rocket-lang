import("util")

def add(a, b)
  c = []
  foreach i in a.size()
    c[i] = a[i] + b[i]
  end
  return c
end

def parseClaim(line)
  parts = line.split()
  start = parts[2][:-1].split(",")
  size = parts[3].split("x")
  return {
    "id": parts[0][1:],
    "start": [start[0].to_i(), start[1].to_i()],
    "size": [size[0].to_i(), size[1].to_i()],
  }
end

def newFabric(size)
  fabric = util.Make("ARRAY", size)
  foreach rIdx in fabric.size()
    fabric[rIdx] = util.Make("ARRAY", size)
  end
  return fabric
end

def fabricAddClaim(fabric, claim)
  foreach y in claim["size"][1]
    foreach x in claim["size"][0]
      r = fabric[claim["start"][1]+y]
      c = r[claim["start"][0]+x]
      c.push(claim["id"])
    end
  end
end


def part1(lines)
  fabric = newFabric(1000)
  foreach line in lines
    fabricAddClaim(fabric, parseClaim(line))
  end
  count = 0
  foreach row in fabric
    foreach col in row
      if (col.size() > 1)
        count = count + 1
      end
    end
  end
  return count
end

def part2(lines)
  claims = []
  fabric = newFabric(1000)
  foreach line in lines
    claim = parseClaim(line)
    claims.push(claim)
    fabricAddClaim(fabric, claim)
  end

  foreach claim in claims
    save = true
    foreach y in claim["size"][1]
      foreach x in claim["size"][0]
        buf = fabric[claim["start"][1]+y]
        if (buf[claim["start"][0]+x].size() > 1)
          save = false
        end
      end
    end
    if (save)
      return claim["id"]
    end
  end
end

input = IO.open("day3.txt").lines()

puts(part1(input))
puts(part2(input))
