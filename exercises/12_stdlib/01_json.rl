// JSON.parse turns text into a hash or an array. Every object answers to_json.
//
// One thing to watch: JSON has one number type, so a number comes back as a
// FLOAT even when it was written without a decimal point.
//
// Make this print:
//
//   FLOAT
//   1.0
//   {"a":1}

parsed = JSON.parse("{\"n\": 1}")

// TODO: print the type of the value under "n"
// TODO: print the value itself
puts({"a": 1}.to_json())
