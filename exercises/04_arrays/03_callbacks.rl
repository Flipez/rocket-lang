// A callback is a function literal: def(x) ... end. A function returns its
// last expression, so no return is needed.
//
// Make this print:
//
//   [2, 4, 6, 8]
//   [2, 4]
//   10

numbers = [1, 2, 3, 4]

puts(numbers.map(def(x) x * 2 end))
// TODO: print only the even numbers, using select
// TODO: print the total, using reduce with a starting value of 0
