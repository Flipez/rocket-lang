// A function is a value. It can be stored in a variable, put in an array, and
// passed to a method -- which is what a callback is.
//
// puts is a value too, so it can be a callback.
//
// Make this print:
//
//   1
//   2
//   84

numbers = [1, 2]

// TODO: print each element by giving each() the puts builtin

quadruple = def(x) x * 4 end
puts(quadruple(21))
