# Time




## Module Function

### format(INTEGER, STRING)
> Returns `STRING`

Formats the given unix timestamp with the given layout.

[Go date and time formats](https://gosamples.dev/date-time-format-cheatsheet/) are natively supported.
You can also use some but not all [formats present in many other languages](https://apidock.com/ruby/Time/strftime) which are not fully supported.
Take a look at [the source](https://github.com/Flipez/rocket-lang/blob/main/stdlib/time.go) to see which formatters are supported.


```js
🚀 » Time.format(Time.unix(), "%a %%b %b %e %H:%M:%S %Y")
» "Mon %Oct Oct 31 00:28:37 2022"
🚀 » Time.format(Time.unix(), "%a %b %e %H:%M:%S %Y")
» "Mon Oct 31 00:28:43 2022"
```


### parse(STRING, STRING)
> Returns `STRING`

Parses a given string with the given format to a unix timestamp.

[Go date and time formats](https://gosamples.dev/date-time-format-cheatsheet/) are natively supported.
You can also use some but not all [formats present in many other languages](https://apidock.com/ruby/Time/strftime) which are not fully supported.
Take a look at [the source](https://github.com/Flipez/rocket-lang/blob/main/stdlib/time.go) to see which formatters are supported.


```js
🚀 » a = "2022-03-23"
» "2022-03-23"
🚀 » format = "2006-01-02"
» "2006-01-02"
🚀 » Time.parse(a, format)
» 1647993600
```


### sleep(INTEGER)
> Returns `NIL`

Stops the RocketLang routine for at least the stated duration in seconds


```js
🚀 > Time.sleep(2)
```


### unix()
> Returns `INTEGER`

Returns the current time as unix timestamp


```js
🚀 > Time.Unix()
```



## Properties
| Name | Value |
| ---- | ----- |
| ANSIC | Mon Jan _2 15:04:05 2006 |
| Kitchen | 3:04PM |
| Layout | 01/02 03:04:05PM '06 -0700 |
| RFC1123 | Mon, 02 Jan 2006 15:04:05 MST |
| RFC1123Z | Mon, 02 Jan 2006 15:04:05 -0700 |
| RFC3339 | 2006-01-02T15:04:05Z07:00 |
| RFC3339Nano | 2006-01-02T15:04:05.999999999Z07:00 |
| RFC822 | 02 Jan 06 15:04 MST |
| RFC822Z | 02 Jan 06 15:04 -0700 |
| RFC850 | Monday, 02-Jan-06 15:04:05 MST |
| RubyDate | Mon Jan 02 15:04:05 -0700 2006 |
| Stamp | Jan _2 15:04:05 |
| StampMicro | Jan _2 15:04:05.000000 |
| StampMilli | Jan _2 15:04:05.000 |
| StampNano | Jan _2 15:04:05.000000000 |
| UnixDate | Mon Jan _2 15:04:05 MST 2006 |
