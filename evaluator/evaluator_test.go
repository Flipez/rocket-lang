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
		expected int64
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
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (true)\n 10 end", 10},
		{"if (false) \n 10 end", nil},
		{"if (1) { 10 }", 10},
		{"if (1) \n 10 end", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 < 2) \n 10 \n else \n 20 end", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
		  if (10 > 1) {
				if (10 > 1) {
					return 10;
				}
				return 1;
			}
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
		{"if (10 > 1) { true + false }", "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "identifier not found: foobar"},
		{
			`
			if (10 > 1) {
				if (10 > 1) {
					return true + false
				}

				return 1
			}
			`, "unknown operator: BOOLEAN + BOOLEAN",
		},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
		{`{"name": "Monkey"}[def(x) { x }];`, "unusable as hash key: FUNCTION"},
		{"🔥 != 👍", "identifier not found: IDENT"},
		{"5 % 0", "division by zero not allowed"},
		{"5 % 0 ? true : false", "division by zero not allowed"},
		{"(4 > 5 ? true).nope()", "undefined method `.nope()` for NULL"},
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
		expected int64
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
	input := "def(x) { x + 2; };"

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
		expected int64
	}{
		{"identity = def(x) { x; }; identity(5);", 5},
		{"identity = def(x) { return x; }; identity(5);", 5},
		{"double = def(x) { x * 2; }; double(5);", 10},
		{"add = def(x, y) { x + y; }; add(5, 5);", 10},
		{"add = def(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"def(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
	newAdder = def(x) {
		def(y) { x + y };
	};

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
		{
			`"abc"[1]`,
			"b",
		},
		{
			`"abc"[-1]`,
			"c",
		},
		{
			`"abc"[4]`,
			nil,
		},
		{
			`"abc"[:2]`,
			"ab",
		},
		{
			`"abc"[:-2]`,
			"a",
		},
		{
			`"abc"[2:]`,
			"c",
		},
		{
			`"abc"[-2:]`,
			"bc",
		},
		{
			`s="abc";s[1]="B";s[1]`,
			"B",
		},
		{
			`s="abc";s[-2]="B";s[-2]`,
			"B",
		},
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
		{`raise("Error")`, "wrong number of arguments. got=1, want=2"},
		{`raise("Error", 1)`, "first argument to `raise` must be INTEGER, got=STRING"},
		{`raise(1, 1)`, "second argument to `raise` must be STRING, got=INTEGER"},
		{`exit()`, "wrong number of arguments. got=0, want=1"},
		{`exit("Error")`, "argument to `exit` must be INTEGER, got=STRING"},
		{`open()`, "wrong number of arguments. got=0, want=1"},
		{`open(1)`, "argument to `file` not supported, got=INTEGER"},
		{`open("fixtures/module.rl", 1)`, "argument mode to `file` not supported, got=INTEGER"},
		{`open("fixtures/module.rl", "r", 1)`, "argument perm to `file` not supported, got=INTEGER"},
		{`open("fixtures/module.rl", "nope", "0644").read(1)`, "undefined method `.read()` for ERROR"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
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
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
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

	expected := map[object.HashKey]int64{
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
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestNamedFunctionStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"def five() { return 5 } five()", 5},
		{"def ten() { return 10 } ten()", 10},
		{"def fifteen() { return 15 } fifteen()", 15},
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
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		number, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(number))
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

		testIntegerObject(t, evaluated, int64(number))
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != object.NULL {
		t.Errorf("object is not object.NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func testEval(input string) object.Object {
	l := lexer.New(input)
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

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
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
