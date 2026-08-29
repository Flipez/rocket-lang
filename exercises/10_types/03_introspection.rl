// A value can list what it responds to. methods() gives the names a type has of
// its own, sorted, and wat() prints them with their arguments. Try
// `1.5.wat()` in the playground.
//
// Make this print:
//
//   true
//   false
//   true
//
// The last line asks whether the list comes back sorted, which it does -- so
// the same program prints the same thing on every run.

puts(1.5.methods().include?("round"))
// TODO: does a float have an upcase method?
// TODO: is the list of methods already sorted? compare it with itself sorted
