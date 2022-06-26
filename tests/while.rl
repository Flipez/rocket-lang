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
  a = 0
  while (a != 3)
    puts(a)
    a = a % 0
  end
end

test_one()
test_two()