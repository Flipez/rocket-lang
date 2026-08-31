servers = [
  {"name": "web-1", "active": true,  "size": 512},
  {"name": "web-2", "active": false, "size": 1024},
  {"name": "db-1",  "active": true,  "size": 2048}
]

active = servers.filter(def(s) s.get("active", false) end)

puts(active.map(def(s) s.get("name", "") end).sort())
puts(active.reduce(0, def(sum, s) sum + s.get("size", 0) end))
puts(active.max_by(def(s) s.get("size", 0) end).get("name", ""))
