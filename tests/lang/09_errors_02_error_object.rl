begin
  nil.no_such_method()
rescue e
  print(e.type())
  print(e.methods())
  print(e.message() == e.to_string())
end
