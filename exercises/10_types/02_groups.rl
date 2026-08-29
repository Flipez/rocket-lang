// A method that accepts a family of types names the family rather than listing
// it. Those names are type groups, and a value can be asked about them.
//
// is_a? takes a type name as readily as a group name. A name that is neither is
// an error, not a false -- a typo should not look like an answer.
//
// Make this print:
//
//   true
//   false
//   ["CALLABLE"]

puts("a".is_a?("HASHABLE"))
// TODO: is nil HASHABLE?
// TODO: print the groups a function belongs to
