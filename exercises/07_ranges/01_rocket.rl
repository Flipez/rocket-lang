// A range is written with a rocket. -> stops before the end, => includes it.
//
//   foreach i in 0 -> 3     0 1 2
//   foreach i in 0 => 3     0 1 2 3
//
// Make this print 1, 2, 3, then 5, each on its own line.

foreach i in 1 -> 4
  puts(i)
end

// TODO: print just 5, using a range that includes its end
