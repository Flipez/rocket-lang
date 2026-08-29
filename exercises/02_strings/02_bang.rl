// A method ending in ! changes the value it is called on. The same name
// without the ! leaves it alone and returns a new value.
//
// The line below calls upcase(), which is why s is still lowercase after it.
// Make this print:
//
//   HELLO
//   HELLO
//
// Change one character.

s = "hello"

puts(s.upcase())
puts(s)
