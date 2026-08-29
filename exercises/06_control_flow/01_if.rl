// if takes elif and else. There is no `unless` and no `case`.
//
// Only false and nil are false, so 0 and "" are true. That catches people out.
//
// Make this print:
//
//   big
//   zero is truthy

n = 10

if n > 100
  puts("huge")
elif n > 5
  // TODO: print "big"
else
  puts("small")
end

if 0
  puts("zero is truthy")
else
  puts("zero is falsy")
end
