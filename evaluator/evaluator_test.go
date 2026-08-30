package evaluator

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/flipez/rocket-lang/lexer"
	"github.com/flipez/rocket-lang/object"
	"github.com/flipez/rocket-lang/parser"
	"github.com/flipez/rocket-lang/utilities"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 /2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"5 ➕ 5 ➕ 5 ➕ 5 - 10", 10},
		{"5 % 5", 0},
		{"5 % 4", 1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 <= 2", true},
		{"2 <= 2", true},
		{"3 <= 2", false},
		{"2 >= 1", true},
		{"2 >= 2", true},
		{"2.0 >= 3.0", false},
		{"3.0 >= 2.0", true},
		{"3.0 <= 3.0", true},
		{"2 >= 3", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"👍", true},
		{"👎", false},
		{"👍 == 👍", true},
		{"👍 == 👎", false},
		{"👍 != 👎", true},
		{"👍 != 👍", false},
		{"true ? true : false", true},
		{"false ? true : false", false},
		{"4 > 3 ? true : false", true},
		{"3 > 4 ? false : true", true},
		{"a = true ? (false ? 0 : true) : 0; a", true},
		{"[1] + [1] == [1, 1]", true},
		{"nil == nil", true},
		{"nil == 1", false},
		{"true and false", false},
		{"true or false", true},
		{"false and true", false},
		{"false or true", true},
		{"true && false", false},
		{"true || false", true},
		{"false && true", false},
		{"false || true", true},
		{"1 && true", true},
		{"(1 || true) == 1", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) \n 10 \nend", 10},
		{"if (false) \n 10 \nend", nil},
		{"if (1) \n 10 \nend", 10},
		{"if (1 < 2) \n 10 \nend", 10},
		{"if (1 > 2) \n 10 \nend", nil},
		{"if (1 > 2) \n 10 \n else \n 20 \nend", 20},
		{"if (1 < 2) \n 10 \n else \n 20 \nend", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"return 10;", 10},
		{"return 10; 9", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
		  if (10 > 1)
				if (10 > 1)
					return 10;
				end
				return 1;
			end
		`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5 + true;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + true; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; true + false; 5", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if (10 > 1) \n true + false \n", "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "test:1:1: identifier not found: foobar"},
		{
			`
			if (10 > 1)
				if (10 > 1)
					return true + false
        end
				return 1
			end
			`, "unknown operator: BOOLEAN + BOOLEAN",
		},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
		{`{"name": "Monkey"}[def(x) { x }];`, "unusable as hash key: FUNCTION"},
		{"🔥 != 👍", "test:1:0: identifier not found: IDENT"},
		{"5 % 0", "division by zero not allowed"},
		{"5 % 0 ? true : false", "division by zero not allowed"},
		{"(4 > 5 ? true).nope()", "test:1:15: undefined method `.nope()` for NIL"},
		{"if (5 % 0)\n puts(true)\nend", "division by zero not allowed"},
		{"a = {(5%0): true}", "division by zero not allowed"},
		{"a = {true: (5%0)}", "division by zero not allowed"},
		{"def test() \n puts(true) \nend; a = {test: true}", "unusable as hash key: FUNCTION"},
		{`import "fixtures/nope"`, "test:1:7: Import Error: no module named 'fixtures/nope' found"},
		{
			`import "../fixtures/parser_error"`,
			"Parse Error: [1:10: expected next token to be ), got EOF instead]",
		},
		{
			`import "../fixtures/module" only Nope`,
			`test:1:7: Import Error: '../fixtures/module' does not export 'Nope'; exported: 'A', 'Sum', 'lower'`,
		},
		{"def test() \n puts(true) \nend; test[1]", "index operator not supported: FUNCTION"},
		{"[1] - [1]", "unknown operator: ARRAY - ARRAY"},
		{"break(1.nope())", "test:1:8: undefined method `.nope()` for INTEGER"},
		{"next(1.nope())", "test:1:7: undefined method `.nope()` for INTEGER"},
		{"nil.nope()", "test:1:4: undefined method `.nope()` for NIL"},
		{"begin puts(nope) end", "test:1:12: identifier not found: nope"},
		{"begin puts(nope) rescue e e.nope() end", "test:1:28: undefined method `.nope()` for ERROR"},
		{"a = begin puts(nope) rescue e e.msg() end; a.nope()", "test:1:45: undefined method `.nope()` for STRING"},
		{`raise("custom error")`, "custom error"},
		{"foreach i in 'test' -> 3 \nputs(i)\nend", "test:1:8: range rocket start has to be an integer, got STRING"},
		{"foreach i in 0 -> 'test' \nputs(i)\nend", "test:1:8: unsupported range rocket value, got STRING"},
		{"foreach i in 0 -> 5 ^ 'test' \nputs(i)\nend", "test:1:8: range rocket step has to be an integer, got STRING"},
		{"[[1, 2]].to_m() + [[1], [2]].to_m()", "matrix addition failed: dimension mismatch: cannot add 1x2 and 2x1 matrices"},
		{"[[1, 2]].to_m() - [[1], [2]].to_m()", "matrix subtraction failed: dimension mismatch: cannot subtract 1x2 and 2x1 matrices"},
		{"[[1, 2]].to_m() * [[1, 2]].to_m()", "matrix multiplication failed: incompatible dimensions: cannot multiply 1x2 by 1x2"},
		{"[[1, 2]].to_m() % [[1, 2]].to_m()", "unknown operator: MATRIX % MATRIX"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestAssignStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"a = 5; a;", 5},
		{"a = 5 * 5; a;", 25},
		{"a = 5; b = a; b;", 5},
		{"a = 5; b = a; c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "def(x) \n x + 2; \nend;"

	evaluated := testEval(input)
	def, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(def.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", def.Parameters)
	}

	if def.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", def.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if def.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, def.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"identity = def(x) \n x; \nend; identity(5);", 5},
		{"identity = def(x) \n return x; \nend; identity(5);", 5},
		{"double = def(x) \n x * 2; \nend; double(5);", 10},
		{"add = def(x, y) \n x + y; \nend; add(5, 5);", 10},
		{"add = def(x, y) \n x + y; \nend; add(5 + 5, add(5, 5));", 20},
		{"def(x) \n x; \nend(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
	newAdder = def(x)
		def(y)
		  x + y
		end
	end;

	addTwo = newAdder(2);
	addTwo(2);`

	testIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"abc"[1]`, "b"},
		{`"abc"[-1]`, "c"},
		{`"abc"[4]`, nil},
		{`"abc"[:2]`, "ab"},
		{`"abc"[:-2]`, "a"},
		{`"abc"[2:]`, "c"},
		{`"abc"[-2:]`, "bc"},
		{`s="abc";s[1]="B";s[1]`, "B"},
		{`s="abc";s[-2]="B";s[-2]`, "B"},
		{`"test"[1]`, "e"},
		{`"test"[-1]`, "t"},
		{`"test"[7]`, nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		str, ok := tt.expected.(string)
		if ok {
			testStringObject(t, evaluated, str)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`puts("test")`, nil},
		{`OS.raise("Error")`, "too few arguments: got=1, want=2"},
		{`OS.raise("Error", 1)`, "wrong argument type on position 1: got=STRING, want=INTEGER"},
		{`OS.raise(1, 1)`, "wrong argument type on position 2: got=INTEGER, want=STRING"},
		{`OS.exit()`, "too few arguments: got=0, want=1"},
		{`OS.exit("Error")`, "wrong argument type on position 1: got=STRING, want=INTEGER"},
		{`IO.open()`, "too few arguments: got=0, want=1"},
		{`IO.open(1, "r", "0644")`, "wrong argument type on position 1: got=INTEGER, want=STRING"},
		{`IO.open("fixtures/module.rl", 1, "0644")`, "wrong argument type on position 2: got=INTEGER, want=STRING"},
		{`IO.open("fixtures/module.rl", "r", 1)`, "wrong argument type on position 3: got=INTEGER, want=STRING"},
		{`IO.open("fixtures/module.rl", "nope", "0644").read(1)`, "test:1:46: undefined method `.read()` for ERROR"},
		{"a = Time.unix(); Time.sleep(2); b = Time.unix(); b - a", 2},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, expected)
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"i = 0; [1][i]",
			1,
		},
		{
			"[1, 2, 3][1 + 1]",
			3,
		},
		{
			"myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"myArray = [1, 2, 3]; i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			3,
		},
		{
			`a=[1,2,3];a[1]=5;a[1]`,
			5,
		},
		{
			`a=[1,2,3];a[-1]=5;a[-1]`,
			5,
		},
		{
			`a=[[1,2], [3,4], [5,6]]; a[0][0] = 5 ; a[0][0]`,
			5,
		},
		{
			`a=[[[1,2,3],2], [3,4], [5,6]]; a[0][0][0] = 5 ; a[0][0][0]`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestHashLiterals(t *testing.T) {
	input := `two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval did not return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int{
		object.NewString("one").HashKey():   1,
		object.NewString("two").HashKey():   2,
		object.NewString("three").HashKey(): 3,
		object.NewInteger(4).HashKey():      4,
		object.TRUE.HashKey():               5,
		object.FALSE.HashKey():              6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]

		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
		{
			`h={"a": 1};h["a"]=5;h["a"]`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestMatrixIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"m = [[1, 2], [3, 4]].to_m(); m[0]",
			"[1.0, 2.0]",
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[1]",
			"[3.0, 4.0]",
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[0][0]",
			1.0,
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[0][1]",
			2.0,
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[1][0]",
			3.0,
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[1][1]",
			4.0,
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[-1]",
			"[3.0, 4.0]",
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[-1][0]",
			3.0,
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[-2]",
			"[1.0, 2.0]",
		},
		{
			"m = [[1, 2, 3], [4, 5, 6]].to_m(); m[0][2]",
			3.0,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case string:
			arr, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if arr.Inspect() != expected {
				t.Errorf("array has wrong value. got=%s, want=%s", arr.Inspect(), expected)
			}
		case float64:
			testFloatObject(t, evaluated, expected)
		default:
			t.Errorf("unexpected expected type: %T", expected)
		}
	}
}

func TestMatrixIndexOutOfBounds(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"m = [[1, 2], [3, 4]].to_m(); m[2]",
			"row index 2 out of bounds [0, 2)",
		},
		{
			"m = [[1, 2], [3, 4]].to_m(); m[-3]",
			"row index -1 out of bounds [0, 2)",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expected {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expected, errObj.Message)
		}
	}
}

func TestNamedFunctionStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"def five() \n return 5 \nend five()", 5},
		{"def ten() \n return 10 \nend ten()", 10},
		{"def fifteen() \n return 15 \nend fifteen()", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestExport(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`import "../fixtures/module"; module.Sum(2, 3)`, 5},
		{`import "../fixtures/module"; module.A`, 5},
		{`import "../fixtures/module"; module.lower`, 7},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		number, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, number)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestImportExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`import "../fixtures/module"; module.A`,
			5,
		},
		{
			`import "../fixtures/module"; module.Sum(2, 3)`,
			5,
		},
		{
			`import "../fixtures/module" as module2; module2.A`,
			5,
		},
		{
			`import "../fixtures/module" only A; module.A`,
			5,
		},
		{
			`import "../fixtures/module"; m = module; m.Sum(2, 3)`,
			5,
		},
		{
			`import "../fixtures/wrapper"; wrapper.SumViaBase(2, 3)`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		number, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, number)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestModuleStrictness(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`import "../fixtures/module"; module.a`,
			"module '../fixtures/module' has no export 'a'",
		},
		{
			`import "../fixtures/module"; module.nope()`,
			"module '../fixtures/module' has no export 'nope'",
		},
		{
			`import "../fixtures/module"; module.Private`,
			"module '../fixtures/module' has no export 'Private'",
		},
		{
			`import "../fixtures/module" only Sum; module.A`,
			"module '../fixtures/module' has no export 'A'",
		},
		{
			`math = 5; import "../fixtures/module" as math`,
			"Import Error: cannot bind module as 'math', name already in use",
		},
		{
			`import "../fixtures/module" as Math`,
			"Import Error: cannot bind module as 'Math', name already in use",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		err, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("input %q: expected an error, got=%T (%+v)", tt.input, evaluated, evaluated)
			continue
		}

		if !strings.Contains(err.Message, tt.expected) {
			t.Errorf("input %q: expected message containing %q, got=%q", tt.input, tt.expected, err.Message)
		}
	}
}

// TestModuleMemberCallArgumentError is a regression test: when an argument
// to a module member call itself errors, evalExpressions returns a
// one-element slice holding just that error object. evalObjectCall must
// return it immediately -- mirroring the guard *ast.Call uses -- instead of
// letting applyFunction index straight into the slice by parameter
// position, which panics (or, for single-parameter functions, silently
// discards the error and returns a normal value).
func TestModuleMemberCallArgumentError(t *testing.T) {
	evaluated := testEval(`import "../fixtures/module" as m; m.Sum(5%0, 1)`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error, got=%T (%+v)", evaluated, evaluated)
	}

	if !strings.Contains(err.Message, "division by zero not allowed") {
		t.Errorf("expected message containing %q, got=%q", "division by zero not allowed", err.Message)
	}
}

// TestImportSearchPaths proves that a module reachable only through a
// SearchPaths entry -- not through the importer-relative "./"/"../" branch
// -- can be imported by its plain name. The fixture lives in its own
// directory so it cannot be found any other way; if the AddPath call below
// is removed, FindModule has nowhere else to look and this import fails.
func TestImportSearchPaths(t *testing.T) {
	if err := utilities.AddPath("../fixtures/searchpath_only"); err != nil {
		t.Errorf("error adding the search path: %s", err)
		return
	}

	evaluated := testEval(`import "only_via_searchpath"; only_via_searchpath.Marker`)
	testIntegerObject(t, evaluated, 99)
}

// TestImportOnlyDoesNotLeakNames proves that `only` binds solely the
// namespace name -- the imported members never become bare identifiers in
// the importing scope.
func TestImportOnlyDoesNotLeakNames(t *testing.T) {
	leaked := testEval(`import "../fixtures/module" only Sum; Sum`)

	err, ok := leaked.(*object.Error)
	if !ok {
		t.Fatalf("expected `Sum` to be unbound, got=%T (%+v)", leaked, leaked)
	}
	if !strings.Contains(err.Message, "identifier not found") {
		t.Errorf("expected message containing %q, got=%q", "identifier not found", err.Message)
	}

	scoped := testEval(`import "../fixtures/module" only Sum; module.Sum(2, 3)`)
	testIntegerObject(t, scoped, 5)
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != object.NIL {
		t.Errorf("object is not object.NIL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func testEval(input string) object.Object {
	l := lexer.New(input, "test")
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

// testEvalWithRebind mirrors testEval but enables the REPL's import-rebind
// relaxation on the top-level environment, the way repl.go does before
// evaluating each entered line.
func testEvalWithRebind(input string) object.Object {
	l := lexer.New(input, "test")
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	env.AllowRebind()

	return Eval(program, env)
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s", result.Value, expected)
		return false
	}

	return true
}

func testIntegerObject(t *testing.T, obj object.Object, expected int) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%T, want=%t", result.Value, expected)
		return false
	}

	return true
}

func TestCircularImport(t *testing.T) {
	evaluated := testEval(`import "../fixtures/cycle_a"`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error object, got=%T (%+v)", evaluated, evaluated)
	}

	if !strings.Contains(err.Message, "circular import") {
		t.Errorf("expected a circular import error, got=%q", err.Message)
	}

	if !strings.Contains(err.Message, "../fixtures/cycle_a.rl -> ../fixtures/cycle_b.rl -> ../fixtures/cycle_a.rl") {
		t.Errorf("expected chain cycle_a -> cycle_b -> cycle_a (in order), got=%q", err.Message)
	}
}

func TestRelativeImport(t *testing.T) {
	evaluated := testEval(`import "../fixtures/sibling/parent"; parent.LeafValue()`)

	testIntegerObject(t, evaluated, 3)
}

func TestExportingAModuleIsAnError(t *testing.T) {
	evaluated := testEval(`import "../fixtures/export_module"`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error object, got=%T (%+v)", evaluated, evaluated)
	}

	if !strings.Contains(err.Message, "Export Error: cannot export 'Inner': a module cannot be exported") {
		t.Errorf("expected message about exporting a module, got=%q", err.Message)
	}
}

// TestExportPropagatesValueError proves that `export Name = <expr>` returns
// the underlying evaluation error unchanged when <expr> itself errors,
// instead of masking it or continuing on to mark the export.
func TestExportPropagatesValueError(t *testing.T) {
	evaluated := testEval(`export Foo = 5 % 0`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error, got=%T (%+v)", evaluated, evaluated)
	}
	if !strings.Contains(err.Message, "division by zero not allowed") {
		t.Errorf("expected message containing %q, got=%q", "division by zero not allowed", err.Message)
	}
}

// TestExportAssignedModuleIsAnError covers the `export Name = <expr>` form
// where <expr> evaluates to a module, as distinct from the bare `export
// Name` form already covered by TestExportingAModuleIsAnError.
func TestExportAssignedModuleIsAnError(t *testing.T) {
	evaluated := testEval(`import "../fixtures/module" as Inner; export Alias = Inner`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error object, got=%T (%+v)", evaluated, evaluated)
	}
	if !strings.Contains(err.Message, "Export Error: cannot export 'Alias': a module cannot be exported") {
		t.Errorf("expected message about exporting a module, got=%q", err.Message)
	}
}

// TestExportUndefinedNameIsAnError covers the bare `export Name` form when
// Name was never defined in scope.
func TestExportUndefinedNameIsAnError(t *testing.T) {
	evaluated := testEval(`export NeverDefined`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error object, got=%T (%+v)", evaluated, evaluated)
	}
	if !strings.Contains(err.Message, "Export Error: 'NeverDefined' is not defined") {
		t.Errorf("expected message about the name not being defined, got=%q", err.Message)
	}
}

func TestModuleCachedAcrossImports(t *testing.T) {
	l := lexer.New(`import "../fixtures/module"; import "../fixtures/module" as m2`, "test")
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	evaluated := Eval(program, env)
	if object.IsError(evaluated) {
		t.Fatalf("unexpected error evaluating program: %+v", evaluated)
	}

	first, ok := env.Get("module")
	if !ok {
		t.Fatalf("expected 'module' to be bound in env")
	}
	firstMod, ok := first.(*object.Module)
	if !ok {
		t.Fatalf("expected 'module' to be *object.Module, got=%T", first)
	}

	second, ok := env.Get("m2")
	if !ok {
		t.Fatalf("expected 'm2' to be bound in env")
	}
	secondMod, ok := second.(*object.Module)
	if !ok {
		t.Fatalf("expected 'm2' to be *object.Module, got=%T", second)
	}

	if firstMod != secondMod {
		t.Errorf("expected both imports of the same file to bind the same *object.Module instance, got distinct pointers %p and %p", firstMod, secondMod)
	}
}

// TestImportRebindNoOpWhenIdentical proves that re-entering the exact same
// (non-`only`) import line under RebindAllowed is a silent no-op, and the
// module is still fully usable afterwards.
func TestImportRebindNoOpWhenIdentical(t *testing.T) {
	evaluated := testEvalWithRebind(`import "../fixtures/module"; import "../fixtures/module"; module.Sum(2, 3)`)

	testIntegerObject(t, evaluated, 5)
}

// TestImportRebindNoOpWhenIdenticalOnly is a regression test: re-entering an
// `only` import with the exact same narrowing must also be treated as a
// no-op, not as "name already in use". This case regressed once already
// because a naive rebind check compared the *object.Module pointers
// directly, and an `only` import always constructs a fresh *object.Module,
// so it never pointer-matched the existing binding.
func TestImportRebindNoOpWhenIdenticalOnly(t *testing.T) {
	evaluated := testEvalWithRebind(`import "../fixtures/module" only Sum; import "../fixtures/module" only Sum; module.Sum(2, 3)`)

	testIntegerObject(t, evaluated, 5)
}

// TestImportRebindStillErrorsOnNarrowingChange proves that RebindAllowed
// only relaxes the check for an import that would bind the exact same
// namespace as what's already bound. Any change to the narrowing --
// widening, narrowing, or swapping to a different `only` set -- must still
// error.
func TestImportRebindStillErrorsOnNarrowingChange(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			"full import then only",
			`import "../fixtures/module"; import "../fixtures/module" only Sum`,
		},
		{
			"only then full import",
			`import "../fixtures/module" only Sum; import "../fixtures/module"`,
		},
		{
			"different only lists",
			`import "../fixtures/module" only Sum; import "../fixtures/module" only A`,
		},
	}

	for _, tt := range tests {
		evaluated := testEvalWithRebind(tt.input)

		err, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("%s: expected an error, got=%T (%+v)", tt.name, evaluated, evaluated)
			continue
		}
		if !strings.Contains(err.Message, "name already in use") {
			t.Errorf("%s: expected message containing %q, got=%q", tt.name, "name already in use", err.Message)
		}
	}
}

// TestImportRebindStillErrorsOnNonModuleName proves that RebindAllowed does
// not relax the check when the binding name is already held by a plain
// (non-module) value -- there is no "same binding" to compare against.
func TestImportRebindStillErrorsOnNonModuleName(t *testing.T) {
	evaluated := testEvalWithRebind(`module = 5; import "../fixtures/module"`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error, got=%T (%+v)", evaluated, evaluated)
	}
	if !strings.Contains(err.Message, "name already in use") {
		t.Errorf("expected message containing %q, got=%q", "name already in use", err.Message)
	}
}

// TestImportRebindDifferentNamesBothUsable proves that importing the same
// path twice under two different names is not a rebind case at all -- both
// bindings are independent and both remain usable.
func TestImportRebindDifferentNamesBothUsable(t *testing.T) {
	evaluated := testEvalWithRebind(`import "../fixtures/module"; import "../fixtures/module" as m2; module.Sum(2, 3) + m2.Sum(1, 1)`)

	testIntegerObject(t, evaluated, 7)
}

// TestImportRebindReachesNestedScope proves that the rebind relaxation
// enabled on the top-level (REPL) environment is visible from an enclosed
// scope, the way a while-loop body inherits it in the real REPL. This
// exercises RebindAllowed's outer-delegation branch, mirroring how
// Registry() delegates to outer.
func TestImportRebindReachesNestedScope(t *testing.T) {
	evaluated := testEvalWithRebind(`
import "../fixtures/module"
i = 0
while i < 2
	import "../fixtures/module"
	i = i + 1
end
module.Sum(2, 3)
`)

	testIntegerObject(t, evaluated, 5)
}

// testEvalWithoutFile evaluates input with no source filename, which is what
// the REPL and `rocket-lang -e` both do.
func testEvalWithoutFile(input string) object.Object {
	l := lexer.New(input, "")
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

// TestRelativeImportWithoutSourceFile covers the REPL and `-e`, where there is
// no importing file to resolve `./` and `../` against. Those anchor on the
// working directory instead, so that the explicit spelling of an import
// resolves wherever the bare name already does.
func TestRelativeImportWithoutSourceFile(t *testing.T) {
	evaluated := testEvalWithoutFile(`import "../fixtures/module"; module.A`)

	testIntegerObject(t, evaluated, 5)
}

// TestLoopsReturnNil covers the change from foreach handing back the value it
// was iterating. That value was the caller's own input, and it silently became
// nil as soon as the loop broke early, so the two loop forms disagreed with
// each other and foreach disagreed with itself.
func TestLoopsReturnNil(t *testing.T) {
	inputs := []string{
		`foreach i in 5 end`,
		`foreach i in [7, 8] end`,
		`foreach i in "hi" end`,
		`foreach i in 5
			if i == 2
				break
			end
		end`,
		`x = 0
		while x < 2
			x = x + 1
		end`,
		`x = 0
		while x < 5
			x = x + 1
			if x == 2
				break
			end
		end`,
	}

	for _, input := range inputs {
		evaluated := testEval(input)

		if _, ok := evaluated.(*object.Nil); !ok {
			t.Errorf("input %q: expected nil, got=%T (%+v)", input, evaluated, evaluated)
		}
	}
}

// TestImportRejectsUnusableImplicitBinding covers a silent failure that
// predates planets: an implicit binding taken from the path could contain
// characters that cannot be referenced afterwards. "my-lib" bound fine, but
// my-lib.X parses as subtraction, so the module was unreachable and the import
// reported nothing. Planet names routinely contain hyphens, which makes this
// common rather than exotic.
func TestImportRejectsUnusableImplicitBinding(t *testing.T) {
	evaluated := testEval(`import "../fixtures/hyphen-name"`)

	err, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected an error, got=%T (%+v)", evaluated, evaluated)
	}
	if !strings.Contains(err.Message, "cannot bind module as 'hyphen-name'") {
		t.Errorf("expected the unusable-name error, got=%q", err.Message)
	}
	if !strings.Contains(err.Message, "use 'as'") {
		t.Errorf("the error does not point at the fix: %q", err.Message)
	}
}

// TestImportWithAsAcceptsAnyPath is the counterpart: an explicit alias makes a
// path that cannot supply a usable name importable.
func TestImportWithAsAcceptsAnyPath(t *testing.T) {
	testIntegerObject(t, testEval(`import "../fixtures/hyphen-name" as hyphen; hyphen.Value`), 7)
}

// TestFunctionArityIsChecked covers a crash: extendFunctionEnv iterated the
// parameters and indexed into the arguments, so calling a function with fewer
// arguments than it has parameters ran off the end of the slice and panicked,
// killing the process. Too many arguments were silently discarded. Builtins
// validated both cases already; user-defined functions validated neither.
func TestFunctionArityIsChecked(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// The panic.
		{`def f(a, b) return a end f(1)`, "too few arguments: got=1, want=2"},
		{`def f(a, b) return a end f()`, "too few arguments: got=0, want=2"},
		// The quiet half: these used to return a value and say nothing.
		{`def f(a) return a end f(1, 2, 3)`, "too many arguments: got=3, want=1"},
		{`def f() return 1 end f(1)`, "too many arguments: got=1, want=0"},
		// A named function says which one, which matters when the call is into
		// a library rather than a function on screen.
		{`def named(a) return a end named()`, "named: too few arguments"},
		// An anonymous function has no name to report.
		{`f = def(a) return a end f()`, "too few arguments: got=0, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		err, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("input %q: expected an error, got=%T (%+v)", tt.input, evaluated, evaluated)
			continue
		}

		if !strings.Contains(err.Message, tt.expected) {
			t.Errorf("input %q: expected %q, got=%q", tt.input, tt.expected, err.Message)
		}
	}
}

// TestFunctionArityAcceptsExactMatch guards the other direction: the check must
// not reject a correct call.
func TestFunctionArityAcceptsExactMatch(t *testing.T) {
	testIntegerObject(t, testEval(`def add(a, b) return a + b end add(2, 3)`), 5)
	testIntegerObject(t, testEval(`def none() return 7 end none()`), 7)
	testIntegerObject(t, testEval(`f = def(a) return a end f(9)`), 9)
}

// TestCallableHashMembers covers calling a callable stored in a hash as though
// it were a method. A hash of functions closing over a constructor's locals is
// already an object with private state; before this it had to be called as
// h["deposit"](50), which was the only thing that did not read like one.
func TestCallableHashMembers(t *testing.T) {
	constructor := `
def new_account(owner, balance)
  return {
    "owner":    owner,
    "deposit":  def(n) balance = balance + n return balance end,
    "describe": def() return owner + ": " + balance.to_s() end
  }
end
`

	tests := []struct {
		input    string
		expected any
	}{
		{constructor + `a = new_account("r", 100); a.deposit(50)`, 150},
		{constructor + `a = new_account("r", 100); a.deposit(50); a.describe()`, "r: 150"},
		// Plain data under a name still reads as data, which already worked.
		{constructor + `a = new_account("r", 100); a.owner`, "r"},

		// The state is the constructor's locals, so instances are independent.
		{constructor + `a = new_account("r", 1); b = new_account("s", 2); a.deposit(10); b.describe()`, "s: 2"},

		// A real Hash method wins over a key of the same name, so a hash of
		// data cannot hijack size() or keys(). The key is still reachable by
		// index.
		{`h = {"size": def() return 99 end}; h.size()`, 1},
		{`h = {"size": def() return 99 end}; h["size"]()`, 99},
		{`h = {"keys": def() return "hijacked" end}; h.keys().size()`, 1},

		// A builtin is callable, so it can be stored and called too.
		{`h = {"f": def(x) return x * 2 end}; h.f(21)`, 42},

		// Nesting works, because each step is an ordinary value.
		{`inner = {"go": def() return "deep" end}; outer = {"inner": inner}; outer.inner.go()`, "deep"},

		// Arity and errors come from the function, reported as anywhere else.
		{`h = {"f": def(x) return x end}; h.f()`, "too few arguments: got=0, want=1"},

		// A name holding something uncallable says so, rather than claiming the
		// method does not exist.
		{`h = {"f": 1}; h.f()`, "`f` is not callable for HASH, it is INTEGER"},
		{`h = {"f": nil}; h.f()`, "`f` is not callable for HASH, it is NIL"},
		{`h = {"f": [1]}; h.f()`, "`f` is not callable for HASH, it is ARRAY"},

		// A name that is not there at all keeps the old message, which is the
		// more useful one: the method does not exist.
		{`h = {}; h.nope()`, "undefined method `.nope()` for HASH"},
		{`h = {"f": 1}; h.other()`, "undefined method `.other()` for HASH"},

		// No other type gained a fallback.
		{`a = [1]; a.nope()`, "undefined method `.nope()` for ARRAY"},
		{`"a".nope()`, "undefined method `.nope()` for STRING"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			integer, ok := evaluated.(*object.Integer)
			if !ok {
				t.Errorf("input %q: not an Integer. got=%T (%+v)", tt.input, evaluated, evaluated)
				continue
			}
			if integer.Value != expected {
				t.Errorf("input %q: got=%d, want=%d", tt.input, integer.Value, expected)
			}
		case string:
			if str, ok := evaluated.(*object.String); ok {
				if str.Value != expected {
					t.Errorf("input %q: got=%q, want=%q", tt.input, str.Value, expected)
				}
				continue
			}

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("input %q: not a String or Error. got=%T (%+v)", tt.input, evaluated, evaluated)
				continue
			}
			if !strings.Contains(errObj.Message, expected) {
				t.Errorf("input %q: error %q does not contain %q", tt.input, errObj.Message, expected)
			}
		}
	}
}

// TestImportFromVirtualFileSystem runs a real multi-file program with no files
// on disk. That is what the playground needs -- a browser has no filesystem for
// an import to resolve against -- and it is also how a module test should be
// written: the fixtures are in the test, not beside it.
func TestImportFromVirtualFileSystem(t *testing.T) {
	const root = "/play"

	tests := []struct {
		name      string
		files     map[string][]byte
		expected  string
		wantError bool
	}{
		{
			name: "a bare name resolves through the search path",
			files: map[string][]byte{
				root + "/util.rl": []byte(`export def double(x) return x * 2 end`),
				root + "/main.rl": []byte(`import "util"
puts(util.double(21))`),
			},
			expected: "42\n",
		},
		{
			name: "a ./ path resolves against the importing file",
			files: map[string][]byte{
				root + "/util.rl": []byte(`export def shout(s) return s.upcase() end`),
				root + "/main.rl": []byte(`import "./util" as helper
puts(helper.shout("hi"))`),
			},
			expected: "HI\n",
		},
		{
			name: "only narrows what the namespace holds",
			files: map[string][]byte{
				root + "/util.rl": []byte(`export a = 1
export b = 2`),
				root + "/main.rl": []byte(`import "util" only a
puts(util.a)`),
			},
			expected: "1\n",
		},
		{
			name: "a module can import another module",
			files: map[string][]byte{
				root + "/inner.rl": []byte(`export def value() return "deep" end`),
				root + "/middle.rl": []byte(`import "inner"
export def reach() return inner.value() end`),
				root + "/main.rl": []byte(`import "middle"
puts(middle.reach())`),
			},
			expected: "deep\n",
		},
		{
			name: "a missing module is reported, not a crash",
			files: map[string][]byte{
				root + "/main.rl": []byte(`import "nope"`),
			},
			expected:  "no module named 'nope' found",
			wantError: true,
		},
		{
			name: "a circular import is caught",
			files: map[string][]byte{
				root + "/a.rl":    []byte(`import "b"`),
				root + "/b.rl":    []byte(`import "a"`),
				root + "/main.rl": []byte(`import "a"`),
			},
			expected:  "circular import",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			previousFS := utilities.SetFileSystem(utilities.MapFileSystem{Files: tt.files})
			previousPaths := utilities.SearchPaths
			utilities.SearchPaths = []string{root}

			defer func() {
				utilities.SetFileSystem(previousFS)
				utilities.SearchPaths = previousPaths
			}()

			printed, result := runProgramFile(t, root+"/main.rl")

			if !tt.wantError {
				if printed != tt.expected {
					t.Errorf("printed %q, want %q", printed, tt.expected)
				}

				return
			}

			if !object.IsError(result) {
				t.Fatalf("expected an error, got %s (printed %q)", result.Inspect(), printed)
			}
			if !strings.Contains(result.(*object.Error).Message, tt.expected) {
				t.Errorf("error %q does not contain %q", result.(*object.Error).Message, tt.expected)
			}
		})
	}
}

// runProgramFile reads a file through the installed FileSystem and evaluates it
// the way main does, returning what it printed and what it evaluated to.
func runProgramFile(t *testing.T, path string) (string, object.Object) {
	t.Helper()

	source, err := utilities.ReadFile(path)
	if err != nil {
		t.Fatalf("%s: %s", path, err)
	}

	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = write

	collected := make(chan string, 1)
	go func() {
		out, _ := io.ReadAll(read)
		collected <- string(out)
	}()

	l := lexer.New(string(source), path)
	p := parser.New(l)
	program := p.ParseProgram()

	var result object.Object
	if len(p.Errors()) == 0 {
		result = Eval(program, object.NewEnvironment())
	}

	write.Close()
	os.Stdout = original

	if len(p.Errors()) > 0 {
		t.Fatalf("%s does not parse: %s", path, strings.Join(p.Errors(), "; "))
	}

	return <-collected, result
}

// TestNonASCIIIndexing covers indexing, slicing and assigning into a string
// outside ASCII. Those three walked bytes while size(), reverse() and the rest
// counted characters, so "тест"[0] answered a single byte -- half a character.
func TestNonASCIIIndexing(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		// Indexing.
		{`s = "тест"; s[0]`, "т"},
		{`s = "тест"; s[1]`, "е"},
		{`s = "тест"; s[-1]`, "т"},
		{`s = "тест"; s[3]`, "т"},
		{`s = "こんにちは"; s[0]`, "こ"},
		{`s = "café"; s[3]`, "é"},

		// Slicing.
		{`s = "тест"; s[:2]`, "те"},
		{`s = "тест"; s[2:]`, "ст"},
		{`s = "тест"; s[1:3]`, "ес"},

		// Assigning into one.
		{`s = "тест"; s[0] = "Т"; s`, "Тест"},
		{`s = "abc"; s[1] = "ä"; s`, "aäc"},
		{`s = "тест"; s[0] = "Т"; s.size()`, 4},

		// The bounds are in characters too, so the length in the message
		// matches what size() answers.
		{`s = "тест"; s[4] = "x"`, "index out of range, got 4 but string is only 4 long"},

		// ASCII is unchanged, which the documented example depends on.
		{`s = "abcdef"; s[2]`, "c"},
		{`s = "abcdef"; s[-2]`, "e"},
		{`s = "abcdef"; s[:2]`, "ab"},
		{`s = "abcdef"; s[2:]`, "cdef"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, expected)
		case string:
			if str, ok := evaluated.(*object.String); ok {
				testStringObject(t, str, expected)
				continue
			}

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("input %q: not a String or Error, got %T", tt.input, evaluated)
				continue
			}
			if !strings.Contains(errObj.Message, expected) {
				t.Errorf("input %q: error %q does not contain %q", tt.input, errObj.Message, expected)
			}
		}
	}
}

// TestStringOperationsAgreeOnLength walks every string method and operation
// that counts, and checks they all count characters. Any one of them counting
// bytes puts it out of step with the others, which is how this began.
func TestStringOperationsAgreeOnLength(t *testing.T) {
	for _, sample := range []string{`"тест"`, `"こんにちは"`, `"café"`, `"plain"`, `"a👍b"`} {
		size := testEval(sample + ".size()")
		ascii := testEval(sample + ".ascii().size()")
		reversed := testEval(sample + ".reverse().size()")
		sliced := testEval(sample + "[0:1].size()")

		if size.Inspect() != ascii.Inspect() {
			t.Errorf("%s: size() is %s but ascii() has %s entries", sample, size.Inspect(), ascii.Inspect())
		}
		if size.Inspect() != reversed.Inspect() {
			t.Errorf("%s: size() is %s but reverse() is %s long", sample, size.Inspect(), reversed.Inspect())
		}
		if sliced.Inspect() != "1" {
			t.Errorf("%s: a one-character slice is %s characters long", sample, sliced.Inspect())
		}
	}
}
