// A conversion that cannot succeed answers nil rather than inventing a zero,
// so a failure can be told from a real 0.
//
// Make this print:
//
//   125
//   nil
//   true

puts("125".to_i())
puts("abc".to_i())
// TODO: print whether "abc".to_i() is nil, using nil?
