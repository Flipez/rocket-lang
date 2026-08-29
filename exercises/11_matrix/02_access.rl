// get and set take a row and a column, both counted from 0. set changes the
// matrix and answers with it, so calls can be chained.
//
// Make this print:
//
//   2.0
//   [1.0, 2.0]
//   [[9.0, 2.0], [3.0, 4.0]]

m = [[1, 2], [3, 4]].to_m()

puts(m.get(0, 1))
// TODO: print the first row
// TODO: set the top-left value to 9, then print the whole matrix as an array
