---
sidebar_position: 1
---
# Install Guide

[![GitHub release](https://img.shields.io/github/release/flipez/rocket-lang.svg)](https://github.com/flipez/rocket-lang/releases/)

## Installation
### macOS
```
brew install flipez/homebrew-tap/rocket-lang
```

### APT
```
echo "deb [trusted=yes] https://apt.rocket-lang.org/ /" > \
/etc/apt/sources.list.d/fury.list
```

### RPM
```
[fury]
name=RocketLang Repo
baseurl=https://rpm.rocket-lang.org
enabled=1
gpgcheck=0
```
### Manual
Download from [releases](https://github.com/Flipez/rocket-lang/releases).

There is also a [Visual Studio Code Extension](https://marketplace.visualstudio.com/items?itemName=Flipez.rocket-lang-support) available. Just search for `rocket-lang` in the extension menu.

## Running a program

```
rocket-lang program.rl        # run a file
rocket-lang -e 'print("hi")'  # run the code given
rocket-lang                   # start the REPL
```

### Exit codes

> 👉 Before `0.24` everything exited `0`, including a crash

| | Exit code |
| --- | --- |
| the program ran | `0` |
| an error nothing handled | `1` |
| the program did not parse | `1` |
| the file could not be read | `1` |
| `OS.exit(n)` | `n` |

So `rocket-lang build.rl && deploy.sh` no longer runs the deploy after a crash.

An error the program handles is not a failure — `begin`/`rescue` succeeding
exits `0`, and so does a conversion answering `nil`, which is what `nil` is for:

```js
begin
  1 / 0
rescue e
  print("handled")
end
# exits 0
```

Diagnostics — a parse error, an unreadable file — go to standard error. The
program's own output, and the value it ends on, go to standard output.
