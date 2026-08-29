// This one is unusual, and worth understanding.
//
// An integer remembers the base it was written in, and prints in it. Two
// integers of different bases refuse to be combined, because 0x10 + 4 has no
// obvious answer -- so you convert first.
//
// Make this print:
//
//   0x10
//   16
//   0x14

hex = "0x10".to_i()

puts(hex)
puts(hex.to_base(10))
// TODO: add 4 to hex and print it, still in hexadecimal.
// Hint: give the 4 the same base first, with to_base.
