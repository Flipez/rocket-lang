import CodeBlockSimple from '@site/components/CodeBlockSimple'

# HTTP




```js
def test()
  puts(request["body"])
  return("test")
end

HTTP.handle("/", test)

HTTP.listen(3000)

// Example request hash:
// {"protocol": "HTTP/1.1", "protocolMajor": 1, "protocolMinor": 1, "body": "servus", "method": "POST", "host": "localhost:3000", "contentLength": 6}

```

## Literal Specific Methods

### handle(STRING, FUNCTION)
> Returns `NIL|ERROR`

Adds a handle to the global HTTP server. Needs to be done before starting one via .listen().
Inside the function a variable called "request" will be populated which is a hash with information about the request.

Also a variable called "response" will be created which will be returned automatically as a response to the client.
The response can be adjusted to the needs. It is a HASH supports the following content:

- "status" needs to be an INTEGER (eg. 200, 400, 500). Default is 200.
- "body" needs to be a STRING. Default ""
- "headers" needs to be a HASH(STRING:STRING) eg. headers["Content-Type"] = "text/plain". Default is \{"Content-Type": "text/plain"\}




<CodeBlockSimple input='HTTP.handle("/", callback_func)
' />



### listen(INTEGER)
> Returns `NIL|ERROR`

Starts a blocking webserver on the given port.



<CodeBlockSimple input='HTTP.listen(3000)
' />




## Generic Literal Methods

### is_a?(STRING)
> Returns `BOOLEAN|ERROR`

Returns `true` when the value is of the given type or belongs to the given type group, so one question covers both. The name has to be one that exists: anything else is an error rather than a `false`, because a typo would otherwise read as a real answer. See [Types and type groups](../language/types).


<CodeBlockSimple input='nil.is_a?("NIL")
"a".is_a?("HASHABLE")
nil.is_a?("HASHABLE")
true.is_a?("INTEGERABLE")
"a".is_a?("INTEGER")
' output='true
true
false
true
false
' />


### methods()
> Returns `ARRAY`

Returns the names of the methods specific to this literal type, not including the generic methods listed on this page. The names are sorted, so the result is the same on every run. A type with no methods of its own returns an empty array.


<CodeBlockSimple input='1.0.methods().include?("round")
true.methods()
' output='true
[]
' />


### nil?()
> Returns `BOOLEAN`

Returns `true` only for `nil`. Reads better than comparing against `nil` at the end of a chain, and every type answers it.


<CodeBlockSimple input='nil.nil?()
1.nil?()
"".nil?()
[].first().nil?()
' output='true
false
false
true
' />


### to_f()
> Returns `FLOAT|NIL`

Converts an object to its float representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0.0`.


<CodeBlockSimple input='1.to_f()
"1.4".to_f()
"abc".to_f()
nil.to_f()
' output='1.0
1.4
nil
nil
' />


### to_i()
> Returns `INTEGER|NIL`

Converts an object to its integer representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0`. For strings a `0b`, `0o` or `0x` prefix selects binary, octal or hexadecimal and is matched case insensitively, a leading zero followed only by octal digits is octal, and anything else is decimal. The resulting integer keeps the base it was parsed with, and integers of differing bases cannot be combined directly.


<CodeBlockSimple input='true.to_i()
false.to_i()
1234.to_i()
"4".to_i()
"0".to_i()
"0125".to_i()
"0x2322".to_i()
"0b1010".to_i()
"test".to_i()
' output='1
0
1234
4
0
0o125
0x2322
0b1010
nil
' />


### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


<CodeBlockSimple input='a = {"test": 1234}
a.to_json()
' output='{"test": 1234}
"{\"test\":1234}"
' />


### to_s()
> Returns `STRING`

Converts an object to its string representation, or the empty string when it has none. Takes no arguments; an integer renders in its own base, so use `to_base` first to change it.


<CodeBlockSimple input='true.to_s()
1234.to_s()
"test".to_s()
1.4.to_s()
nil.to_s()
"0b1010".to_i().to_s()
' output='"true"
"1234"
"test"
"1.4"
""
"0b1010"
' />


### type()
> Returns `STRING`

Returns the type of the object.


<CodeBlockSimple input='"test".type()
' output='"STRING"
' />


### type_groups()
> Returns `ARRAY`

Returns the type groups the value belongs to, sorted. `ANY` is not listed: every value belongs to it, so it would say nothing while prefixing every answer. It exists for signatures, where `push(ANY)` means the argument accepts anything, and `is_a?("ANY")` still answers `true`. See [Types and type groups](../language/types) for what each group means.


<CodeBlockSimple input='1.type_groups()
nil.type_groups()
def() end.type_groups()
puts.type_groups()
' output='["COMPARABLE", "HASHABLE", "INTEGERABLE", "NUMERIC", "STRINGABLE"]
["STRINGABLE"]
["CALLABLE"]
["CALLABLE"]
' />


### wat()
> Returns `NIL`

Prints the type's literal-specific methods with their argument and return types, sorted by name, one per line. It returns `nil` rather than the listing: this exists to be read, and the REPL echoes a returned value through its escaped representation, which would put the whole thing on one line. Use `methods` when the names are wanted as data. A type with no methods of its own prints only the heading.


<CodeBlockSimple input='true.wat()
1.0.wat()
' output='BOOLEAN supports the following methods:
nil
FLOAT supports the following methods:
	abs()
	ceil([INTEGER])
	divmod(FLOAT)
	finite?()
	floor([INTEGER])
	infinite?()
	nan?()
	negative?()
	positive?()
	round([INTEGER])
	truncate([INTEGER])
	zero?()
nil
' />


