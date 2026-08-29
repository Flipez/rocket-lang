// A hash takes callbacks too. select and reject receive the key and the value;
// transform_values receives only the value.
//
// Make this print:
//
//   1
//   20
//
// The first line is how many entries have a value above 1. The second is the
// value under "a" after multiplying every value by 10.

h = {"a": 2}
h["b"] = 1

// TODO: count the entries whose value is greater than 1
// TODO: multiply every value by 10, then read "a" back
