// An operation that fails answers with an ERROR, which stops the program --
// unless it happens inside begin/rescue.
//
// rescue takes a name, and that name is bound to the error.
//
// Make this print:
//
//   caught: division by zero not allowed
//   still running

begin
  1 / 0
rescue e
  // TODO: print "caught: " followed by the error's message
end

puts("still running")
