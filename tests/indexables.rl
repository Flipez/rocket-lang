s = "abcdef"
print(s[2])
print(s[-2])
print(s[:2])
print(s[:-2])
print(s[2:])
print(s[-2:])
print(s[1:-2])

s[2] = "C"
s[-2] = "E"
print(s)

a = [1, 2, 3, 4, 5]
print(a[2])
print(a[-2])
print(a[:2])
print(a[:-2])
print(a[2:])
print(a[-2:])
print(a[1:-2])

a[2] = 8
a[-2] = 9
print(a)

h = {"a": 1, 2: true}
print(h["a"])
print(h[2])
h["a"] = 3
h["b"] = "moo"
print(h["a"])
print(h["b"])
print(h[2])
