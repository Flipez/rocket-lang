import "../fixtures/module" as m
print(m.A)
print(m.Sum(2, 3))
print(m.lower)

import "../fixtures/module" as narrow only Sum
print(narrow.Sum(1, 1))

import "../fixtures/side_effect" as first
import "../fixtures/side_effect" as second
print(first.X + second.X)
