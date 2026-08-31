h = {"a": 2}
h["b"] = 1

puts(h.filter(def(k, v) v > 1 end).size())
puts(h.transform_values(def(v) v * 10 end).get("a", 0))
