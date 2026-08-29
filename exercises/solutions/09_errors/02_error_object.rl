begin
  nil.no_such_method()
rescue e
  puts(e.type())
  puts(e.methods())
  puts(e.msg() == e.to_s())
end
