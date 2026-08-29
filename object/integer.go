package object

import (
	"encoding/json"
	"fmt"
	"unicode"
)

type Integer struct {
	Value int
	Base  int
}

func NewInteger(i int) *Integer {
	return &Integer{Value: i, Base: 10}
}

func NewIntegerWithBase(i, b int) *Integer {
	return &Integer{Value: i, Base: b}
}

func (i *Integer) Inspect() string {
	switch i.Base {
	case 2:
		return fmt.Sprintf("0b%b", i.Value)
	case 8:
		return fmt.Sprintf("0o%o", i.Value)
	case 16:
		return fmt.Sprintf("0x%x", i.Value)
	}
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func init() {
	objectMethods[INTEGER_OBJ] = map[string]ObjectMethod{
		"abs": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				i := o.(*Integer)
				if i.Value < 0 {
					return NewIntegerWithBase(-i.Value, i.Base)
				}

				return i
			},
		},
		"base": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				return NewInteger(o.(*Integer).Base)
			},
		},
		"to_base": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				return NewIntegerWithBase(o.(*Integer).Value, args[0].(*Integer).Value)
			},
		},
		"chr": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(STRING_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				value := o.(*Integer).Value
				if value < 0 || value > unicode.MaxRune {
					return NewErrorFormat("%d is out of the range of a character", value)
				}

				return NewString(string(rune(value)))
			},
		},
		"digits": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					OptArg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				value := o.(*Integer).Value
				if value < 0 {
					return NewError("digits of a negative number are not defined")
				}

				base := 10
				if len(args) > 0 {
					base = args[0].(*Integer).Value
				}
				if base < 2 {
					return NewErrorFormat("base has to be 2 or greater, got %d", base)
				}

				// Least significant digit first, as in Ruby: 12345.digits is
				// [5, 4, 3, 2, 1].
				if value == 0 {
					return NewArrayWithObjects(NewInteger(0))
				}

				digits := make([]Object, 0)
				for value > 0 {
					digits = append(digits, NewInteger(value%base))
					value /= base
				}

				return NewArray(digits)
			},
		},
		"divmod": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(ARRAY_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				i := o.(*Integer)
				other := args[0].(*Integer)

				if err := requireSameBase(i, other); err != nil {
					return err
				}
				if other.Value == 0 {
					return NewError("division by zero not allowed")
				}

				// Matches this language's own / and %, which truncate toward
				// zero. Ruby floors instead, so 11.divmod(-4) is [-3, -1] there
				// and [-2, 3] here -- but here it agrees with 11 / -4 and
				// 11 % -4, which matters more than agreeing with Ruby.
				return NewArrayWithObjects(
					NewIntegerWithBase(i.Value/other.Value, i.Base),
					NewIntegerWithBase(i.Value%other.Value, i.Base),
				)
			},
		},
		"gcd": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				i := o.(*Integer)
				other := args[0].(*Integer)

				if err := requireSameBase(i, other); err != nil {
					return err
				}

				return NewIntegerWithBase(greatestCommonDivisor(i.Value, other.Value), i.Base)
			},
		},
		"lcm": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				i := o.(*Integer)
				other := args[0].(*Integer)

				if err := requireSameBase(i, other); err != nil {
					return err
				}

				divisor := greatestCommonDivisor(i.Value, other.Value)
				if divisor == 0 {
					return NewIntegerWithBase(0, i.Base)
				}

				return NewIntegerWithBase(absInt(i.Value/divisor*other.Value), i.Base)
			},
		},
		"pow": ObjectMethod{
			Layout: MethodLayout{
				ArgPattern: Args(
					Arg(INTEGER_OBJ),
					OptArg(INTEGER_OBJ),
				),
				ReturnPattern: Args(
					Arg(INTEGER_OBJ, ERROR_OBJ),
				),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				i := o.(*Integer)
				exponent := args[0].(*Integer).Value

				// Ruby answers a Rational here. There is no such type, and
				// silently handing back a float would break the rule that a
				// numeric method returns the type it was given.
				if exponent < 0 {
					return NewError("negative exponent is not supported")
				}

				modulus := 0
				if len(args) > 1 {
					modulus = args[1].(*Integer).Value
					if modulus == 0 {
						return NewError("modulus of zero not allowed")
					}
				}

				return NewIntegerWithBase(integerPow(i.Value, exponent, modulus), i.Base)
			},
		},
		"succ": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				i := o.(*Integer)

				return NewIntegerWithBase(i.Value+1, i.Base)
			},
		},
		"pred": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				i := o.(*Integer)

				return NewIntegerWithBase(i.Value-1, i.Base)
			},
		},
		"bit_length": ObjectMethod{
			Layout: MethodLayout{
				ReturnPattern: Args(
					Arg(INTEGER_OBJ),
				),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				// A negative number needs as many bits as its complement, so
				// -256 is 8 bits just like 255.
				value := o.(*Integer).Value
				if value < 0 {
					value = ^value
				}

				length := 0
				for value > 0 {
					length++
					value >>= 1
				}

				return NewInteger(length)
			},
		},
	}

	// Rounding an integer only does something with a negative digit count,
	// which rounds to a power of ten: 555.floor(-1) is 550.
	integerRounding("ceil", func(value, factor int) int {
		return ceilTo(value, factor)
	})
	integerRounding("floor", func(value, factor int) int {
		return floorTo(value, factor)
	})
	integerRounding("round", func(value, factor int) int {
		return roundTo(value, factor)
	})
	integerRounding("truncate", func(value, factor int) int {
		return value / factor * factor
	})

	integerPredicate("even?", func(value int) bool { return value%2 == 0 })
	integerPredicate("odd?", func(value int) bool { return value%2 != 0 })
	integerPredicate("zero?", func(value int) bool { return value == 0 })
	integerPredicate("positive?", func(value int) bool { return value > 0 })
	integerPredicate("negative?", func(value int) bool { return value < 0 })
}

// integerPredicate registers a method returning a BOOLEAN about the value.
func integerPredicate(name string, test func(value int) bool) {
	objectMethods[INTEGER_OBJ][name] = ObjectMethod{
		Layout: MethodLayout{
			ReturnPattern: Args(Arg(BOOLEAN_OBJ)),
		},
		method: func(o Object, _ []Object, _ Environment) Object {
			if test(o.(*Integer).Value) {
				return TRUE
			}

			return FALSE
		},
	}
}

// integerRounding registers ceil, floor, round or truncate. Without an
// argument, or with a digit count of zero or more, an integer is already as
// precise as it gets and is returned unchanged.
func integerRounding(name string, apply func(value, factor int) int) {
	objectMethods[INTEGER_OBJ][name] = ObjectMethod{
		Layout: MethodLayout{
			ArgPattern:    Args(OptArg(INTEGER_OBJ)),
			ReturnPattern: Args(Arg(INTEGER_OBJ)),
		},
		method: func(o Object, args []Object, _ Environment) Object {
			i := o.(*Integer)

			digits := 0
			if len(args) > 0 {
				digits = args[0].(*Integer).Value
			}
			if digits >= 0 {
				return i
			}

			factor := 1
			for range -digits {
				factor *= 10
			}

			return NewIntegerWithBase(apply(i.Value, factor), i.Base)
		},
	}
}

// requireSameBase mirrors the check the infix operators make: 0x10 + 4 is an
// error, so 0x10.gcd(4) is one too.
func requireSameBase(left, right *Integer) Object {
	if left.Base != right.Base {
		return NewError("infix operation with unequal base not allowed")
	}

	return nil
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func greatestCommonDivisor(a, b int) int {
	a, b = absInt(a), absInt(b)
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// integerPow raises base to exponent by squaring. A non-zero modulus reduces
// at every step, which is what keeps a.pow(b, m) usable for large exponents.
func integerPow(base, exponent, modulus int) int {
	result := 1
	if modulus != 0 {
		result %= modulus
		base %= modulus
	}

	for exponent > 0 {
		if exponent&1 == 1 {
			result *= base
			if modulus != 0 {
				result %= modulus
			}
		}

		exponent >>= 1
		if exponent == 0 {
			break
		}

		base *= base
		if modulus != 0 {
			base %= modulus
		}
	}

	return result
}

// ceilTo, floorTo and roundTo round value to a multiple of factor. They are
// written for integers so that no float rounding creeps in.
func ceilTo(value, factor int) int {
	truncated := value / factor * factor
	if value > 0 && truncated != value {
		truncated += factor
	}

	return truncated
}

func floorTo(value, factor int) int {
	truncated := value / factor * factor
	if value < 0 && truncated != value {
		truncated -= factor
	}

	return truncated
}

// roundTo rounds half away from zero, matching Ruby's default and Go's
// math.Round.
func roundTo(value, factor int) int {
	remainder := value % factor
	truncated := value - remainder

	if absInt(remainder)*2 >= factor {
		if value < 0 {
			return truncated - factor
		}

		return truncated + factor
	}

	return truncated
}

func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(i, method, env, args)
}

func (i *Integer) GetIterator(start, step int, inclusive bool) Iterator {
	val := int(i.Value)
	if val < start {
		step *= -1
		if inclusive {
			val--
		}
	} else if inclusive {
		val++
	}

	return &integerIterator{max: val, step: step, current: start}
}

func (i *Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

func (i *Integer) ToStringObj() *String {
	return NewString(i.Inspect())
}

func (i *Integer) ToIntegerObj() Object {
	return i
}

func (i *Integer) ToFloatObj() Object {
	return NewFloat(float64(i.Value))
}

type integerIterator struct {
	current, max, step int
}

func (i *integerIterator) Next() (Object, Object, bool) {
	if (i.step < 0 && i.current <= i.max) || (i.step > 0 && i.current >= i.max) {
		return nil, NewInteger(0), false
	}

	obj := NewInteger(i.current)
	i.current += i.step
	return obj, obj, true
}
