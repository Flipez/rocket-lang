// Time.format takes a unix timestamp and a layout. The layout is a date written
// out: 2006 is the year, 01 the month, 02 the day, 15:04:05 the time.
//
// It formats in the local timezone, so the exact date depends on where you are
// -- timestamp 0 is 1970-01-01 in London and 1969-12-31 in New York. This
// exercise asks only about the shape the layout produces, which is the same
// everywhere.
//
// Make this print:
//
//   4
//   10
//   true

puts(Time.format(0, "2006").size())
// TODO: format timestamp 0 as year-month-day, and print the length
// TODO: print whether Time.unix() is greater than 0
