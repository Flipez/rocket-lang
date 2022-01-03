# Functions

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

> New in `0.11`:

Functions can now also be created as named functions:

```js
🚀 > fn test() { puts("test")}
=> fn() {
puts(test)
}

🚀 > test()
"test"
```