parsed = JSON.parse("{\"n\": 1}")

print(parsed.get("n", 0).type())
print(parsed.get("n", 0))
print({"a": 1}.to_json())
