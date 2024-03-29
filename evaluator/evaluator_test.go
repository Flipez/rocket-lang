package evaluator

import (
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
		{"import(true)", "test:1:7: Import Error: invalid import path '&{%!s(bool=true)}'"},
		{"import(5%0)", "division by zero not allowed"},
		{`import("fixtures/nope")`, "Import Error: no module named 'fixtures/nope' found"},
		{
			`import("../fixtures/parser_error")`,
			"Parse Error: [1:10: expected next token to be ), got EOF instead]",
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
		{`OS.raise("Error")`, "to few arguments: got=1, want=2"},
		{`OS.raise("Error", 1)`, "wrong argument type on position 1: got=STRING, want=INTEGER"},
		{`OS.raise(1, 1)`, "wrong argument type on position 2: got=INTEGER, want=STRING"},
		{`OS.exit()`, "to few arguments: got=0, want=1"},
		{`OS.exit("Error")`, "wrong argument type on position 1: got=STRING, want=INTEGER"},
		{`IO.open()`, "to few arguments: got=0, want=1"},
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

func TestImportExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`import("../fixtures/module"); module.A`,
			5,
		},
		{
			`import("../fixtures/module"); module.Sum(2, 3)`,
			5,
		},
		{
			`import("../fixtures/module"); module.a`,
			nil,
		},
		{
			`import("../fixtures/module", "module2"); module2.A`,
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

func TestImportSearchPaths(t *testing.T) {
	if err := utilities.AddPath("../stubs"); err != nil {
		t.Errorf("error adding the stubs path: %s", err)
		return
	}

	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`import("../fixtures/module"); module.A`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		number, _ := tt.expected.(int)

		testIntegerObject(t, evaluated, number)
	}
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
	imports := make(map[string]struct{})
	p := parser.New(l, imports)
	program, _ := p.ParseProgram()
	env := object.NewEnvironment()

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
