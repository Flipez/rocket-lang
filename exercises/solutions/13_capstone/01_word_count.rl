text = "the quick brown fox the"

counts = {}
foreach word in text.split()
  counts[word] = counts.get(word, 0) + 1
end

puts(counts.get("the", 0))
puts(counts.size())
puts(counts.keys().sort())
