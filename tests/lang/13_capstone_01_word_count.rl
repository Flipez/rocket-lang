text = "the quick brown fox the"

counts = {}
for word in text.split()
  counts[word] = counts.get(word, 0) + 1
end

print(counts.get("the", 0))
print(counts.size())
print(counts.keys().sort())
