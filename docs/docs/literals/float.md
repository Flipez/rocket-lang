# Float




## Literal Specific Methods

### plz_f()
> Returns `FLOAT`

Returns self



### plz_i()
> Returns `INTEGER`

Converts the float into an integer.


```js
🚀 > a = 123.456
=> 123.456
🚀 > a.plz_i()
=> "123"
```


### plz_s()
> Returns `STRING`

Returns a string representation of the float.


```js
🚀 > a = 123.456
=> 123.456
🚀 > a.plz_s()
=> "123.456"
```



## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns an array of all supported methods names.

```js
🚀 > "test".methods()
=> [count, downcase, find, reverse!, split, lines, upcase!, strip!, downcase!, size, plz_i, replace, reverse, strip, upcase]
```

### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.

```js
🚀 > a = {"test": 1234}
=> {"test": 1234}
🚀 > a.to_json()
=> "{"test":1234}"
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

