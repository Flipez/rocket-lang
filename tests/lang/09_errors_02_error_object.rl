begin
  nil.no_such_method()
rescue e
  puts(e.type())
  puts(e.methods())
  puts(e.message() == e.to_string())
end
