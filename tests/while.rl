a = 0
while (a != 3)
  puts(a)
  a = a + 1
end


def test_one()
  a = 0
  while (a != 3)
    puts(a)
    if (a == 1)
      return a
    end
    a = a + 1
  end
end

def test_two()
  i = 0
  while (i < 10)
    if (i < 3)
      i = i + 1
      next
    end
    puts(i)
    if (i == 6)
      break
    end
    i = i + 1
  end
end

def test_three()
  a = 0
  while (a != 3)
    puts(a)
    a = a % 0
  end
end

test_one()
test_two()
test_three()
