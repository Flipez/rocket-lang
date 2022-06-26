---
title: "Functions"
menu:
  docs:
    parent: "specification"
toc: true
---
# Functions
Implicit and explicit return statements are supported.

```js
fibonacci = def (x)
  if (x == 0)
    0
  else
    if (x == 1)
      return 1;
    else
      fibonacci(x - 1) + fibonacci(x - 2);
    end
  end
end
```

> New in `0.11`:

Functions can now also be created as named functions:

```js
🚀 > def test()
  puts("test")
end

=> def ()
  puts(test)
end

🚀 > test()
"test"
```