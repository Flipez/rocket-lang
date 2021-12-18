# 🚀🇱🅰🆖

[![goreleaser Status](https://github.com/Flipez/rocket-lang/actions/workflows/release.yml/badge.svg)](https://github.com/Flipez/rocket-lang/actions/workflows/release.yml)
[![Test Status](https://github.com/Flipez/rocket-lang/actions/workflows/test.yml/badge.svg)](https://github.com/Flipez/rocket-lang/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/Flipez/rocket-lang/branch/master/graph/badge.svg)](https://codecov.io/gh/Flipez/rocket-lang)

RocketLang as of version 0.9.5 is the _full_ (as in the book was worked through) version of [MonkeyLang](https://monkeylang.org/) and is then being extended with various useful and not so useful features.

## Installation

Install RocketLang using

```
brew install flipez/homebrew-tap/rocket-lang
```
or download from [releases](https://github.com/Flipez/rocket-lang/releases).

## How To?

* `rocket-lang` without any arguments will start an interactive shell
* `rocket-lang FILE` will run the code in that file (no file extension check yet)
* Use _Javascript_ Highlighting in your editor for some convenience
* Checkout Code [Samples](examples/) for what is currently possible (and what not)

## Examples
### Variables
```js
let some_integer = 1;
let name = "RocketLang";
let array = [1, 2, 3, 4, 5];
let some_boolean = true;
```

Also expressions can be used
```js
let another_int = (10 / 2) * 5 + 30;
let an_array = [1 + 1, 2 * 2, 3];
```

### Functions
Implicit and explicit return statements are supported.
```js
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
```
### Closures
```js
let newGreeter = fn(greeting) {
  return fn(name) { puts(greeting + " " + name); }
};

let hello = newGreeter("Hello");

hello("dear, future Reader!");

```

### Builtin Functions
|Function|Argument(s)|Return Value(s)|Description|
|--|--|--|--|
|`len()`|`array`|`int`|Returns the size of the provided `array`.|
|`first()`|`array`|`any`, `null`|Returns the first object from `array` or `null` if `array` is empty|
|`last()`|`array`|`any`, `null`|Returns the last object from `array` or `null` if `array` is empty|
|`push()`|`array`;`any`|`array`|Returns new `array` advanced by one with given object as last element of given `array`|
|`pop()`|`array`|`array`|Returns the the given `array` as element one without the last object, returns the last object as second element|
|`puts()`|`string`|`null`|Prints the given `string` to stdout|
|`exit()`|`int`|-|Terminates program with exit code `int`|
|`raise()`|`int`; `string`|-|Prints `string` and terminates programm with exit code `int`|
|`open()`|`string`; `string`| `file`| Returns a `file` object for the given path.

### Data Types
#### Strings
```js
let a = "test_string;

let b = "test" + "_string";

let is_true = "test" == "test";
let is_false = "test" == "string";
```

#### Integer
```js
let a = 1;

let b = a + 2;

let is_true = 1 == 1;
let is_false = 1 == 2;
```

#### Boolean
```js
let a = true;
let b = false;

let is_true = a == a;
let is_false = a == b;

let is_true = a != b;
```

#### Hashes
```js
let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];
```
