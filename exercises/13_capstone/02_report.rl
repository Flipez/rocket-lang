// The same data, with callbacks instead of a loop.
//
// Each record is a hash. Pick the ones that are active, sort them by size, and
// total the sizes.
//
// Make this print:
//
//   ["db-1", "web-1"]
//   2560
//   db-1

servers = [
  {"name": "web-1", "active": true,  "size": 512},
  {"name": "web-2", "active": false, "size": 1024},
  {"name": "db-1",  "active": true,  "size": 2048}
]

// TODO: the names of the active servers, sorted
// TODO: the total size of the active servers, using reduce with a start of 0
// TODO: the name of the largest active server, using max_by
