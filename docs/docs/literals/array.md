import CodeBlockSimple from '@site/components/CodeBlockSimple'

# Array




```js
a = [1, 2, 3, 4, 5]
puts(a[2])
puts(a[-2])
puts(a[:2])
puts(a[:-2])
puts(a[2:])
puts(a[-2:])
puts(a[1:-2])

// should output
[1, 2]
[1, 2, 3]
[3, 4, 5]
[4, 5]
[2, 3]
[1, 2, 8, 9, 5]

```

## Literal Specific Methods

### clear()
> Returns `ARRAY`

Removes every element and returns the array, so calls can be chained.


<CodeBlockSimple input='a = [1, 2]
a.clear()
a
' output='[1, 2]
[]
[]
' />


### compact()
> Returns `ARRAY|ERROR`

Returns a copy with every `nil` removed. The array itself is unchanged; use `compact!` to remove them in place.


<CodeBlockSimple input='a = [1, nil, 2, nil]
a.compact()
a
' output='[1, nil, 2, nil]
[1, 2]
[1, nil, 2, nil]
' />


### compact!()
> Returns `ARRAY|ERROR`

Removes every `nil` in place and returns the array, so calls can be chained.


<CodeBlockSimple input='a = [1, nil, 2, nil]
a.compact!()
a
' output='[1, nil, 2, nil]
[1, 2]
[1, 2]
' />


### concat(ARRAY)
> Returns `ARRAY`

Appends every element of the given array and returns the array, so calls can be chained.


<CodeBlockSimple input='a = [1]
a.concat([2, 3])
a
' output='[1]
[1, 2, 3]
[1, 2, 3]
' />


### count([ANY])
> Returns `INTEGER`

Without an argument this is `size`. With one it counts how often that element occurs, which is what `index` cannot tell you.


<CodeBlockSimple input='[1, 2, 2, 3].count()
[1, 2, 2, 3].count(2)
[1, 2, 2, 3].count(9)
' output='4
2
0
' />


### delete(ANY)
> Returns `ANY`

Removes every element equal to the argument and returns that argument, or `nil` when there was nothing to remove, so a removal can be told from a miss.


<CodeBlockSimple input='a = [1, 2, 1]
a.delete(1)
a
a.delete(9)
' output='[1, 2, 1]
1
[2]
nil
' />


### delete_at(INTEGER)
> Returns `ANY`

Removes the element at the given index and returns it. A negative index counts back from the end. An index that is not there gives `nil`, the same answer `first` gives for an empty array.


<CodeBlockSimple input='a = [1, 2, 3]
a.delete_at(1)
a
a.delete_at(9)
' output='[1, 2, 3]
2
[1, 3]
nil
' />


### drop(INTEGER)
> Returns `ARRAY|ERROR`

Returns everything after the first n elements as a new array. Dropping more than there are gives an empty array.


<CodeBlockSimple input='[1, 2, 3].drop(2)
[1, 2, 3].drop(9)
' output='[3]
[]
' />


### empty?()
> Returns `BOOLEAN`

Returns `true` when the array has no elements.


<CodeBlockSimple input='[].empty?()
[1].empty?()
' output='true
false
' />


### first([INTEGER])
> Returns `ANY`

Returns the first element of the array. Shorthand for `array[0]`


<CodeBlockSimple input='["a", "b", 1, 2].first()
' output='"a"
' />


### flatten([INTEGER])
> Returns `ARRAY|ERROR`

Returns a copy with nested arrays inlined, all the way down by default or to the given depth. The array itself is unchanged; use `flatten!` to inline in place.


<CodeBlockSimple input='a = [1, [2, [3]]]
a.flatten()
a.flatten(1)
a
' output='[1, [2, [3]]]
[1, 2, 3]
[1, 2, [3]]
[1, [2, [3]]]
' />


### flatten!([INTEGER])
> Returns `ARRAY|ERROR`

Inlines nested arrays in place and returns the array, so calls can be chained. Takes the same optional depth as `flatten`.


<CodeBlockSimple input='a = [1, [2, [3]]]
a.flatten!()
a
' output='[1, [2, [3]]]
[1, 2, 3]
[1, 2, 3]
' />


### include?(ANY)
> Returns `BOOLEAN`

Returns true or false wether the array contains the given element


<CodeBlockSimple input='[1,2,3].include?(4)
[1,2,3].include?(3)
' output='false
true
' />


### index(ANY)
> Returns `INTEGER`

Returns the index of the given element in the array if found. Otherwise return `-1`.


<CodeBlockSimple input='["a", "b", 1, 2].index(1)
' output='2
' />


### insert(INTEGER, ANY)
> Returns `ARRAY|ERROR`

Inserts an element at the given index and returns the array. A negative index counts back from the end, so `-1` appends. An index past the end is an error rather than a silent padding with `nil`.


<CodeBlockSimple input='a = [1, 2]
a.insert(1, 9)
a.insert(0 - 1, 8)
a
' output='[1, 2]
[1, 9, 2]
[1, 9, 2, 8]
[1, 9, 2, 8]
' />


### join([STRING])
> Returns `STRING`

Joins the elements into a string, with an optional separator between them. Every element has to have a string form; a function does not.


<CodeBlockSimple input='[1, 2, 3].join()
[1, 2, 3].join("-")
' output='"123"
"1-2-3"
' />


### last([INTEGER])
> Returns `ANY`

Returns the last element of the array.


<CodeBlockSimple input='["a", "b", 1, 2].last()
' output='2
' />


### max()
> Returns `ANY`

Returns the largest element, or `nil` for an empty array. Takes the same elements as `min`.


<CodeBlockSimple input='[3, 1, 2].max()
[].max()
' output='3
nil
' />


### min()
> Returns `ANY`

Returns the smallest element, or `nil` for an empty array. The elements have to be all strings, all integers or all floats, exactly as `sort` requires.


<CodeBlockSimple input='[3, 1, 2].min()
["b", "a"].min()
[].min()
' output='1
"a"
nil
' />


### pop()
> Returns `ANY`

Removes the last element of the array and returns it.


<CodeBlockSimple input='a = [1,2,3]
a.pop()
a
' output='[1, 2, 3]
3
[1, 2]
' />


### push(ANY)
> Returns `ARRAY`

Adds the given object as the last element and returns the array, so calls can be chained.


<CodeBlockSimple input='d = [1,2,3]
d.push("a")
d
' output='[1, 2, 3]
[1, 2, 3, "a"]
[1, 2, 3, "a"]
' />


### reverse()
> Returns `ARRAY`

Returns a new array with the elements in reverse order. The array itself is unchanged; use `reverse!` to reverse it in place.


<CodeBlockSimple input='a = ["a", "b", 1, 2]
a.reverse()
a
' output='["a", "b", 1, 2]
[2, 1, "b", "a"]
["a", "b", 1, 2]
' />


### reverse!()
> Returns `ARRAY`

Reverses the array in place and returns it, so calls can be chained.


<CodeBlockSimple input='a = ["a", "b", 1, 2]
a.reverse!()
a
' output='["a", "b", 1, 2]
[2, 1, "b", "a"]
[2, 1, "b", "a"]
' />


### rindex(ANY)
> Returns `INTEGER`

Returns the index of the last matching element, or `-1` when there is none. The mirror of `index`.


<CodeBlockSimple input='[1, 2, 2, 3].rindex(2)
[1, 2, 3].rindex(9)
' output='2
-1
' />


### rotate([INTEGER])
> Returns `ARRAY|ERROR`

Returns a copy with the first elements moved to the end, one by default. A negative count rotates the other way, and the count wraps. The array itself is unchanged; use `rotate!` to rotate in place.


<CodeBlockSimple input='a = [1, 2, 3]
a.rotate()
a.rotate(2)
a.rotate(0 - 1)
a
' output='[1, 2, 3]
[2, 3, 1]
[3, 1, 2]
[3, 1, 2]
[1, 2, 3]
' />


### rotate!([INTEGER])
> Returns `ARRAY|ERROR`

Rotates the array in place and returns it, so calls can be chained. Takes the same optional count as `rotate`.


<CodeBlockSimple input='a = [1, 2, 3]
a.rotate!()
a
' output='[1, 2, 3]
[2, 3, 1]
[2, 3, 1]
' />


### shift()
> Returns `ANY`

Removes the first element and returns it, or `nil` for an empty array. The mirror of `pop`, and like `pop` it changes the array without a `!`, because there is no pure version of taking something out.


<CodeBlockSimple input='a = [1, 2, 3]
a.shift()
a
' output='[1, 2, 3]
1
[2, 3]
' />


### size()
> Returns `INTEGER`

Returns the amount of elements in the array.


<CodeBlockSimple input='["a", "b", 1, 2].size()
' output='4
' />


### slices(INTEGER)
> Returns `ARRAY`

Returns the elements of the array in slices with the size of the given integer


<CodeBlockSimple input='[1,2,3,4,5,6,7,8].slices(3)
' output='[[1, 2, 3], [4, 5, 6], [7, 8]]
' />


### sort()
> Returns `ARRAY|ERROR`

Returns a new sorted array. The elements must all be STRING, all INTEGER or all FLOAT, otherwise an error is returned and the array is left untouched. Use `sort!` to sort in place.


<CodeBlockSimple input='b = [3.4, 3.1, 2.0]
b.sort()
b
' output='[3.4, 3.1, 2.0]
[2.0, 3.1, 3.4]
[3.4, 3.1, 2.0]
' />


### sort!()
> Returns `ARRAY|ERROR`

Sorts the array in place and returns it. A sort that cannot compare its elements returns an error and leaves the array unchanged rather than half-ordered.


<CodeBlockSimple input='b = [3.4, 3.1, 2.0]
b.sort!()
b
' output='[3.4, 3.1, 2.0]
[2.0, 3.1, 3.4]
[2.0, 3.1, 3.4]
' />


### sum()
> Returns `INTEGER`

Adds the elements up. They all have to be numbers.


<CodeBlockSimple input='[1, 2, 3].sum()
[1.5, 2.5].sum()
' output='6
3
' />


### take(INTEGER)
> Returns `ARRAY|ERROR`

Returns the first n elements as a new array. Asking for more than there are gives all of them. The result is a copy, so changing the original does not change it.


<CodeBlockSimple input='[1, 2, 3].take(2)
[1, 2, 3].take(9)
' output='[1, 2]
[1, 2, 3]
' />


### to_m()
> Returns `MATRIX|ERROR`

Converts a nested array (2D array) to a Matrix object.


<CodeBlockSimple input='[[1, 2], [3, 4]].to_m()
' output='2x2 matrix
┌          ┐
│ 1.0  2.0 │
│ 3.0  4.0 │
└          ┘
' />


### uniq()
> Returns `ARRAY|ERROR`

Returns a new array with duplicates removed, keeping the order of first appearance. Returns an error if an element is not hashable. Use `uniq!` to deduplicate in place.


<CodeBlockSimple input='c = ["a", 1, 1, 2]
c.uniq()
c
' output='["a", 1, 1, 2]
["a", 1, 2]
["a", 1, 1, 2]
' />


### uniq!()
> Returns `ARRAY|ERROR`

Removes duplicates from the array in place and returns it, keeping the order of first appearance.


<CodeBlockSimple input='c = ["a", 1, 1, 2]
c.uniq!()
c
' output='["a", 1, 1, 2]
["a", 1, 2]
["a", 1, 2]
' />


### unshift(ANY)
> Returns `ARRAY`

Adds an element to the front and returns the array, so calls can be chained. The mirror of `push`.


<CodeBlockSimple input='a = [2, 3]
a.unshift(1)
a
' output='[2, 3]
[1, 2, 3]
[1, 2, 3]
' />



## Generic Literal Methods

### methods()
> Returns `ARRAY`

Returns the names of the methods specific to this literal type, not including the generic methods listed on this page. The names are sorted, so the result is the same on every run. A type with no methods of its own returns an empty array.


<CodeBlockSimple input='1.0.methods().include?("round")
true.methods()
' output='true
[]
' />


### nil?()
> Returns `BOOLEAN`

Returns `true` only for `nil`. Reads better than comparing against `nil` at the end of a chain, and every type answers it.


<CodeBlockSimple input='nil.nil?()
1.nil?()
"".nil?()
[].first().nil?()
' output='true
false
false
true
' />


### to_f()
> Returns `FLOAT|NIL`

Converts an object to its float representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0.0`.


<CodeBlockSimple input='1.to_f()
"1.4".to_f()
"abc".to_f()
nil.to_f()
' output='1.0
1.4
nil
nil
' />


### to_i()
> Returns `INTEGER|NIL`

Converts an object to its integer representation, or `nil` when it cannot. A `nil` result is what distinguishes a failed conversion from a genuine `0`. For strings a `0b`, `0o` or `0x` prefix selects binary, octal or hexadecimal and is matched case insensitively, a leading zero followed only by octal digits is octal, and anything else is decimal. The resulting integer keeps the base it was parsed with, and integers of differing bases cannot be combined directly.


<CodeBlockSimple input='true.to_i()
false.to_i()
1234.to_i()
"4".to_i()
"0".to_i()
"0125".to_i()
"0x2322".to_i()
"0b1010".to_i()
"test".to_i()
' output='1
0
1234
4
0
0o125
0x2322
0b1010
nil
' />


### to_json()
> Returns `STRING|ERROR`

Returns the object as json notation.


<CodeBlockSimple input='a = {"test": 1234}
a.to_json()
' output='{"test": 1234}
"{\"test\":1234}"
' />


### to_s()
> Returns `STRING`

Converts an object to its string representation, or the empty string when it has none. Takes no arguments; an integer renders in its own base, so use `to_base` first to change it.


<CodeBlockSimple input='true.to_s()
1234.to_s()
"test".to_s()
1.4.to_s()
nil.to_s()
"0b1010".to_i().to_s()
' output='"true"
"1234"
"test"
"1.4"
""
"0b1010"
' />


### type()
> Returns `STRING`

Returns the type of the object.


<CodeBlockSimple input='"test".type()
' output='"STRING"
' />


### wat()
> Returns `NIL`

Prints the type's literal-specific methods with their argument and return types, sorted by name, one per line. It returns `nil` rather than the listing: this exists to be read, and the REPL echoes a returned value through its escaped representation, which would put the whole thing on one line. Use `methods` when the names are wanted as data. A type with no methods of its own prints only the heading.


<CodeBlockSimple input='true.wat()
1.0.wat()
' output='BOOLEAN supports the following methods:
nil
FLOAT supports the following methods:
	abs()
	ceil([INTEGER])
	divmod(FLOAT)
	finite?()
	floor([INTEGER])
	infinite?()
	nan?()
	negative?()
	positive?()
	round([INTEGER])
	truncate([INTEGER])
	zero?()
nil
' />


