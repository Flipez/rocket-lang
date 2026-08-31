h = {"a": 2}
h["b"] = 1

print(h.filter(def(k, v) v > 1 end).size())
print(h.transform_values(def(v) v * 10 end).get("a", 0))
