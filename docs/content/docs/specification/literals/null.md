# Null

`NULL`represents the absence of a value and can be received when a method or a control flow do not return a value.

At the moment `null`can not be created by the user as an object and only received.

```js
🚀 > a = if (false) {}
=> null
```

In a control flow it behances falsy.

```js
🚀 > if (a) { puts("true") } else { puts("false") }
false
```