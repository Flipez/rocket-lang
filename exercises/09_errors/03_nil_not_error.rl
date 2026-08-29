// Not every failure is an ERROR. A conversion that cannot succeed answers nil
// instead, so a failure can be told apart from a real zero -- and nil needs no
// rescue.
//
// Make this print:
//
//   nil
//   0
//   could not read a number

value = "abc".to_i()

puts(value)
puts("0".to_i())

// TODO: if value is nil, print "could not read a number"
