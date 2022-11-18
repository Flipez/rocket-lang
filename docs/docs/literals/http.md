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
- "headers" needs to be a HASH(STRING:STRING) eg. headers["Content-Type"] = "text/plain". Default is {"Content-Type": "text/plain"}



```js
HTTP.handle("/", callback_func)

```



### listen(INTEGER)
> Returns `NIL|ERROR`

Starts a blocking webserver on the given port.


```js
HTTP.listen(3000)

```




## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.


```js
"test".methods()

```

```js
["upcase", "find", "format", "reverse", "split", "replace", "strip!", "count", "reverse!", "lines", "downcase!", "upcase!", "size", "plz_i", "strip", "downcase"]

```



### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


```js
a = {"test": 1234}
a.to_json()

```

```js
{"test": 1234}
"{\"test\":1234}"

```



### type()
> Returns `STRING`

Returns the type of the object.


```js
"test".type()

```

```js
"STRING"

```



### wat()
> Returns `STRING`

Returns the supported methods with usage information.


```js
true.wat()

```

```js
"BOOLEAN supports the following methods:
  plz_s()"

```



