# Math




## Literal Specific Methods

### abs(FLOAT)
> Returns `FLOAT`

Returns the absolute value of the given number as FLOAT


```js
🚀 > Math.abs(-2)
=> 2.0
```


### acos(FLOAT)
> Returns `FLOAT`

Returns the arccosine, in radians, of the argument


```js
🚀 > Math.acos(1.0)
=> 0.0
```


### asin(FLOAT)
> Returns `FLOAT`

Returns the arcsine, in radians, of the argument


```js
🚀 > Math.asin(0.0)
=> 0.0
```


### atan(FLOAT)
> Returns `FLOAT`

Returns the arctangent, in radians, of the argument


```js
🚀 > Math.atan(0.0)
=> 0.0
```


### ceil(FLOAT)
> Returns `FLOAT`

Returns the least integer value greater or equal to the argument


```js
🚀 > Math.ceil(1.49)
=> 2.0
```


### copysign(FLOAT, FLOAT)
> Returns `FLOAT`

Returns a value with the magnitude of first argument and sign of second argument


```js
🚀 > Math.copysign(3.2, -1.0)
=> -3.2
```


### cos(FLOAT)
> Returns `FLOAT`

Returns the cosine of the radion argument


```js
🚀 > Math.cos(Pi/2)
=> 0.0
```


### exp(FLOAT)
> Returns `FLOAT`

Returns e**argument, the base-e exponential of argument


```js
🚀 > Math.exp(1.0)
=> 2.72
```


### floor(FLOAT)
> Returns `FLOAT`

Returns the greatest integer value less than or equal to argument


```js
🚀 > Math.floor(1.51)
=> 1.0
```


### log(FLOAT)
> Returns `FLOAT`

Returns the natural logarithm of argument


```js
🚀 > Math.log(2.7183)
=> 1.0
```


### log10(FLOAT)
> Returns `FLOAT`

Returns the decimal logarithm of argument


```js
🚀 > Math.log(100.0)
=> 2.0
```


### log2(FLOAT)
> Returns `FLOAT`

Returns the binary logarithm of argument


```js
🚀 > Math.log2(256.0)
=> 8.0
```


### pow(FLOAT, FLOAT)
> Returns `FLOAT`

Returns argument1**argument2, the base-argument1 exponential of argument2


```js
🚀 > Math.pow(2.0, 3.0)
=> 8.0
```


### remainder(FLOAT, FLOAT)
> Returns `FLOAT`

Returns the IEEE 754 floating-point remainder of argument1/argument2


```js
🚀 > Math.remainder(100.0, 30.0)
=> 10.0
```


### sin(FLOAT)
> Returns `FLOAT`

Returns the sine of the radion argument


```js
🚀 > Math.sin(Pi)
=> 0.0
```


### sqrt(FLOAT)
> Returns `FLOAT`

Returns the square root of argument


```js
🚀 > Math.sqrt(3.0 * 3.0 + 4.0 * 4.0)
=> 5.0
```


### tan(FLOAT)
> Returns `FLOAT`

Returns the tangent of the radion argument


```js
🚀 > Math.tan(0.0)
=> 0.0
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

