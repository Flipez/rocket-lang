// A hash maps keys to values. Any hashable value works as a key, not just
// strings -- integers, floats, booleans, arrays and hashes all do.
//
// Make this print:
//
//   1
//   true
//   0
//
// The last line is what get() answers when the key is not there: the default
// you pass it.

h = {"a": 1, 2: true}

puts(h["a"])
puts(h[2])
// TODO: look up the missing key "z" with get, defaulting to 0
