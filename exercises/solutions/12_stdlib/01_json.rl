parsed = JSON.parse("{\"n\": 1}")

puts(parsed.get("n", 0).type())
puts(parsed.get("n", 0))
puts({"a": 1}.to_json())
