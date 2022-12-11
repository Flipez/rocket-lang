def parse_input(lines)
  monkeys = []
  foreach block in lines.slices(7)
    items = []
    foreach field in block[1].split()[4:]
      if (field[-1] == ",")
        items.yoink(field[:-1].plz_i())
      else
        items.yoink(field.plz_i())
      end
    end

    monkeys.yoink({
      "items": items,
      "operation": {
        "left": block[2].split()[5],
        "middle": block[2].split()[6],
        "right": block[2].split()[7],
      },
      "test": {
        "num": block[3].split()[-1].plz_i(),
        true: block[4].split()[-1].plz_i(),
        false: block[5].split()[-1].plz_i(),
      },
    })
  end
  return monkeys
end

def inspect(monkey, item)
  nums = []
  if (monkey["operation"]["left"] == "old")
    nums.yoink(item)
  else
    nums.yoink(monkey["operation"]["left"].plz_i())
  end

  if (monkey["operation"]["right"] == "old")
    nums.yoink(item)
  else
    nums.yoink(monkey["operation"]["right"].plz_i())
  end

  if (monkey["operation"]["middle"] == "+")
    return nums[0] + nums[1]
  end
  if (monkey["operation"]["middle"] == "*")
    return nums[0] * nums[1]
  end

  return item
end

def test(monkey, item)
  return item % monkey["test"]["num"] == 0
end

def mbl(activities)
  activities.sort()
  return activities[-2] * activities[-1]
end

def part1(lines)
  monkeys = parse_input(lines)

  activities = []
  foreach i, m in monkeys
    activities.yoink(0)
  end

  foreach r in 20
    foreach idx, monkey in monkeys
      items = monkey["items"]
      monkeys[idx]["items"] = []
      while (items.size() > 0)
        if (items == nil)
          break
        end

        item = items[0]
        items = items[1:]

        item = inspect(monkey, item)
        item = (item / 3).plz_i()
    
        monkeys[monkey["test"][test(monkey, item)]]["items"].yoink(item)
        activities[idx] = activities[idx] + 1
      end
    end
  end

  return mbl(activities)
end

def part2(lines)
  monkeys = parse_input(lines)

  activities = []
  foreach m in monkeys
    activities.yoink(0)
  end

  relief = 1
  foreach monkey in monkeys
    relief = relief * monkey["test"]["num"]
  end

  foreach r in 10000
    foreach idx, monkey in monkeys
      items = monkey["items"]
      monkeys[idx]["items"] = []
      while (items.size() > 0)
        if (items == nil)
          break
        end

        item = items[0]
        items = items[1:]

        item = inspect(monkey, item)
        item = item % relief

        monkeys[monkey["test"][test(monkey, item)]]["items"].yoink(item)
        activities[idx] = activities[idx] + 1
      end
    end
  end

  return mbl(activities)
end

lines = IO.open("day11.txt").lines()

puts(part1(lines))
puts(part2(lines))
