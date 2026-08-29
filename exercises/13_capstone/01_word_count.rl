// Everything from the earlier topics, on one small problem.
//
// Count how often each word appears, then answer three questions about it.
// split() with no argument splits on spaces, and get() with a default is what
// makes counting a one-liner.
//
// Make this print:
//
//   2
//   4
//   ["brown", "fox", "quick", "the"]

text = "the quick brown fox the"

counts = {}
foreach word in text.split()
  // TODO: add one to this word's count, defaulting to 0 when it is new
end

puts(counts.get("the", 0))
puts(counts.size())
puts(counts.keys().sort())
