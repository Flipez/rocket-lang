import "../fixtures/module" as m
puts(m.A)
puts(m.Sum(2, 3))
puts(m.lower)

import "../fixtures/module" as narrow only Sum
puts(narrow.Sum(1, 1))

import "../fixtures/side_effect" as first
import "../fixtures/side_effect" as second
puts(first.X + second.X)
