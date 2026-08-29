// An error is an ordinary value. It has one method of its own -- msg -- and
// everything else it answers comes from the methods every type has.
//
// Make this print:
//
//   ERROR
//   ["msg"]
//   true

begin
  nil.no_such_method()
rescue e
  puts(e.type())
  // TODO: print the methods ERROR has of its own
  // TODO: print whether msg() and to_s() give the same thing
end
