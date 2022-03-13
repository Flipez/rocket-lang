---
title: "HTTP"
menu:
  docs:
    parent: "literals"
---
# HTTP




```js
def test() {
  puts(request["body"])
  return("test")
}

HTTP.handle("/", test)

HTTP.listen(3000)

// Example request hash:
// {"protocol": "HTTP/1.1", "protocolMajor": 1, "protocolMinor": 1, "body": "servus", "method": "POST", "host": "localhost:3000", "contentLength": 6}
```

## Literal Specific Methods

### handle(STRING, FUNCTION)
> Returns `NULL|ERROR`

Adds a handle to the global HTTP server. Needs to be done before starting one via .listen().
Inside the function a variable called "request" will be populated which is a hash with information about the request.


```js
🚀 > HTTP.handle("/", callback_func)
```


### listen(INTEGER)
> Returns `NULL|ERROR`

Starts a blocking webserver on the given port.


```js
🚀 > HTTP.listen(3000)
```


### new()
> Returns `HTTP`






## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.

```js
🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]
```

### type()
> Returns `STRING`

Returns the type of the object.

```js
🚀 > "test".type()
=> "STRING"
```

### wat()
> Returns `STRING`

Returns the supported methods with usage information.

```js
🚀 > true.wat()
=> BOOLEAN supports the following methods:
				plz_s()
```
