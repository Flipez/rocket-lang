# Integer
An integer can be positiv or negative and is always internally represented by a 64-Bit Integer.

To cast a negative integer a digit can be prefixed with a `-` eg. `-456`.

## Literal Specific Methods
### .plz_s(_base=10_)
Returns a string representation of the integer.

```js
🚀 > a = 456
=> 456
🚀 > a.plz_s()
=> "456"
```

This also works with negative integers:

```js
🚀 > a = -345
=> -345
🚀 > a.plz_s()
=> "-345"
```

`.plz_s()` takes also an argument which represents the integer base to convert between different number systems:

```js
🚀 > 1234.plz_s(2)
=> 10011010010
🚀 > 1234.plz_s(8)
=> 2322
🚀 > 1234.plz_s(10)
=> 1234
```