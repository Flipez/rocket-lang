// break leaves a loop, next skips to the following turn. Both words work in a
// foreach, in a while, and inside a callback given to each, map or select.
//
// Make this print:
//
//   1
//   2
//   1
//   3

numbers = [1, 2, 3, 4]

// break inside a callback ends the walk, so this prints 1 then 2.
numbers.each(def(n)
  if n == 3
    break
  end
  puts(n)
end)

// Skip the 2 and stop before the 4, so this prints 1 then 3.
foreach n in numbers
  // TODO: skip when n is 2
  // TODO: stop when n is 4
  puts(n)
end
