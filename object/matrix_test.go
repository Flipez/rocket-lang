package object

import (
	"testing"
)

func TestMatrixCreation(t *testing.T) {
	tests := []struct {
		name      string
		data      [][]float64
		wantRows  int
		wantCols  int
		wantError bool
	}{
		{
			name:     "2x2 matrix",
			data:     [][]float64{{1, 2}, {3, 4}},
			wantRows: 2,
			wantCols: 2,
		},
		{
			name:     "3x2 matrix",
			data:     [][]float64{{1, 2}, {3, 4}, {5, 6}},
			wantRows: 3,
			wantCols: 2,
		},
		{
			name:     "1x3 matrix",
			data:     [][]float64{{1, 2, 3}},
			wantRows: 1,
			wantCols: 3,
		},
		{
			name:     "empty matrix",
			data:     [][]float64{},
			wantRows: 0,
			wantCols: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMatrix(tt.data)
			if m.Rows != tt.wantRows {
				t.Errorf("Rows = %d, want %d", m.Rows, tt.wantRows)
			}
			if m.Cols != tt.wantCols {
				t.Errorf("Cols = %d, want %d", m.Cols, tt.wantCols)
			}
		})
	}
}

func TestNewMatrixFromObjects(t *testing.T) {
	tests := []struct {
		name      string
		input     []Object
		wantRows  int
		wantCols  int
		wantError bool
	}{
		{
			name: "valid 2x2 matrix",
			input: []Object{
				NewArray([]Object{NewInteger(1), NewInteger(2)}),
				NewArray([]Object{NewInteger(3), NewInteger(4)}),
			},
			wantRows: 2,
			wantCols: 2,
		},
		{
			name: "mixed int and float",
			input: []Object{
				NewArray([]Object{NewInteger(1), NewFloat(2.5)}),
				NewArray([]Object{NewFloat(3.5), NewInteger(4)}),
			},
			wantRows: 2,
			wantCols: 2,
		},
		{
			name:      "empty array",
			input:     []Object{},
			wantError: true,
		},
		{
			name: "inconsistent row lengths",
			input: []Object{
				NewArray([]Object{NewInteger(1), NewInteger(2)}),
				NewArray([]Object{NewInteger(3)}),
			},
			wantError: true,
		},
		{
			name: "non-numeric element",
			input: []Object{
				NewArray([]Object{NewInteger(1), NewString("bad")}),
			},
			wantError: true,
		},
		{
			name: "row with empty array",
			input: []Object{
				NewArray([]Object{}),
			},
			wantError: true,
		},
		{
			name: "non-array row",
			input: []Object{
				NewInteger(1),
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrixFromObjects(tt.input)
			if tt.wantError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if m.Rows != tt.wantRows {
				t.Errorf("Rows = %d, want %d", m.Rows, tt.wantRows)
			}
			if m.Cols != tt.wantCols {
				t.Errorf("Cols = %d, want %d", m.Cols, tt.wantCols)
			}
		})
	}
}

func TestMatrixAdd(t *testing.T) {
	m1 := NewMatrix([][]float64{{1, 2}, {3, 4}})
	m2 := NewMatrix([][]float64{{5, 6}, {7, 8}})

	result, err := m1.Add(m2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := [][]float64{{6, 8}, {10, 12}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if result.Data[i][j] != expected[i][j] {
				t.Errorf("result[%d][%d] = %f, want %f", i, j, result.Data[i][j], expected[i][j])
			}
		}
	}

	// Test dimension mismatch
	m3 := NewMatrix([][]float64{{1, 2, 3}})
	_, err = m1.Add(m3)
	if err == nil {
		t.Error("expected dimension mismatch error")
	}
}

func TestMatrixSubtract(t *testing.T) {
	m1 := NewMatrix([][]float64{{5, 6}, {7, 8}})
	m2 := NewMatrix([][]float64{{1, 2}, {3, 4}})

	result, err := m1.Subtract(m2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := [][]float64{{4, 4}, {4, 4}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if result.Data[i][j] != expected[i][j] {
				t.Errorf("result[%d][%d] = %f, want %f", i, j, result.Data[i][j], expected[i][j])
			}
		}
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := NewMatrix([][]float64{{1, 2}, {3, 4}})
	m2 := NewMatrix([][]float64{{5, 6}, {7, 8}})

	result, err := m1.Multiply(m2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// [1*5+2*7, 1*6+2*8] = [19, 22]
	// [3*5+4*7, 3*6+4*8] = [43, 50]
	expected := [][]float64{{19, 22}, {43, 50}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if result.Data[i][j] != expected[i][j] {
				t.Errorf("result[%d][%d] = %f, want %f", i, j, result.Data[i][j], expected[i][j])
			}
		}
	}

	// Test incompatible dimensions
	m3 := NewMatrix([][]float64{{1}, {2}, {3}})
	_, err = m1.Multiply(m3)
	if err == nil {
		t.Error("expected incompatible dimensions error")
	}
}

func TestMatrixToArray(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}, {3, 4}})
	arr := m.ToArray()

	if arr.Type() != ARRAY_OBJ {
		t.Errorf("ToArray() type = %s, want ARRAY", arr.Type())
	}

	if len(arr.Elements) != 2 {
		t.Errorf("array length = %d, want 2", len(arr.Elements))
	}

	row0 := arr.Elements[0].(*Array)
	if len(row0.Elements) != 2 {
		t.Errorf("row 0 length = %d, want 2", len(row0.Elements))
	}

	val := row0.Elements[0].(*Float).Value
	if val != 1 {
		t.Errorf("arr[0][0] = %f, want 1", val)
	}
}

func TestMatrixType(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}})
	if m.Type() != MATRIX_OBJ {
		t.Errorf("Type() = %s, want MATRIX", m.Type())
	}
}

func TestMatrixInspect(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}, {3, 4}})
	inspect := m.Inspect()
	if inspect == "" {
		t.Error("Inspect() returned empty string")
	}
}

func TestMatrixMarshalJSON(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}, {3, 4}})
	jsonBytes, err := m.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
	}

	jsonStr := string(jsonBytes)
	// Check that JSON contains expected fields
	if !contains(jsonStr, "\"rows\":2") {
		t.Errorf("JSON missing rows field, got: %s", jsonStr)
	}
	if !contains(jsonStr, "\"cols\":2") {
		t.Errorf("JSON missing cols field, got: %s", jsonStr)
	}
	if !contains(jsonStr, "\"data\":[[1,2],[3,4]]") {
		t.Errorf("JSON missing data field, got: %s", jsonStr)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || indexOfSubstring(s, substr) >= 0))
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func TestMatrixShape(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()
	shape := m.InvokeMethod("shape", *env)

	arr, ok := shape.(*Array)
	if !ok {
		t.Fatalf("shape() should return Array, got %T", shape)
	}

	if len(arr.Elements) != 2 {
		t.Errorf("shape array length = %d, want 2", len(arr.Elements))
	}

	rows := arr.Elements[0].(*Integer).Value
	cols := arr.Elements[1].(*Integer).Value

	if rows != 2 {
		t.Errorf("rows = %d, want 2", rows)
	}
	if cols != 3 {
		t.Errorf("cols = %d, want 3", cols)
	}
}

func TestMatrixRows(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()
	result := m.InvokeMethod("rows", *env)

	rows, ok := result.(*Integer)
	if !ok {
		t.Fatalf("rows() should return Integer, got %T", result)
	}

	if rows.Value != 2 {
		t.Errorf("rows = %d, want 2", rows.Value)
	}
}

func TestMatrixCols(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()
	result := m.InvokeMethod("cols", *env)

	cols, ok := result.(*Integer)
	if !ok {
		t.Fatalf("cols() should return Integer, got %T", result)
	}

	if cols.Value != 3 {
		t.Errorf("cols = %d, want 3", cols.Value)
	}
}

func TestMatrixSize(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()
	result := m.InvokeMethod("size", *env)

	size, ok := result.(*Integer)
	if !ok {
		t.Fatalf("size() should return Integer, got %T", result)
	}

	if size.Value != 6 {
		t.Errorf("size = %d, want 6", size.Value)
	}
}

func TestMatrixTranspose(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	transposed := m.Transpose()

	if transposed.Rows != 3 {
		t.Errorf("transposed.Rows = %d, want 3", transposed.Rows)
	}
	if transposed.Cols != 2 {
		t.Errorf("transposed.Cols = %d, want 2", transposed.Cols)
	}

	expected := [][]float64{{1, 4}, {2, 5}, {3, 6}}
	for i := 0; i < transposed.Rows; i++ {
		for j := 0; j < transposed.Cols; j++ {
			if transposed.Data[i][j] != expected[i][j] {
				t.Errorf("transposed[%d][%d] = %f, want %f", i, j, transposed.Data[i][j], expected[i][j])
			}
		}
	}
}

func TestMatrixTransposeMethod(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}, {3, 4}})
	env := NewEnvironment()

	// Test transpose()
	result := m.InvokeMethod("transpose", *env)
	transposed, ok := result.(*Matrix)
	if !ok {
		t.Fatalf("transpose() should return Matrix, got %T", result)
	}

	if transposed.Rows != 2 || transposed.Cols != 2 {
		t.Errorf("transposed shape = %dx%d, want 2x2", transposed.Rows, transposed.Cols)
	}

	if transposed.Data[0][1] != 3 || transposed.Data[1][0] != 2 {
		t.Errorf("transpose values incorrect")
	}

	// Test t() alias
	result2 := m.InvokeMethod("t", *env)
	transposed2, ok := result2.(*Matrix)
	if !ok {
		t.Fatalf("t() should return Matrix, got %T", result2)
	}

	if transposed2.Rows != 2 || transposed2.Cols != 2 {
		t.Errorf("t() shape = %dx%d, want 2x2", transposed2.Rows, transposed2.Cols)
	}
}

func TestMatrixGet(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})

	// Valid get
	val, err := m.Get(0, 2)
	if err != nil {
		t.Errorf("Get(0, 2) unexpected error: %v", err)
	}
	if val != 3 {
		t.Errorf("Get(0, 2) = %f, want 3", val)
	}

	// Out of bounds row
	_, err = m.Get(2, 0)
	if err == nil {
		t.Error("Get(2, 0) should return error for out of bounds row")
	}

	// Out of bounds col
	_, err = m.Get(0, 3)
	if err == nil {
		t.Error("Get(0, 3) should return error for out of bounds col")
	}

	// Negative indices
	_, err = m.Get(-1, 0)
	if err == nil {
		t.Error("Get(-1, 0) should return error for negative row")
	}
}

func TestMatrixSet(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})

	// Valid set
	err := m.Set(0, 2, 99)
	if err != nil {
		t.Errorf("Set(0, 2, 99) unexpected error: %v", err)
	}
	if m.Data[0][2] != 99 {
		t.Errorf("After Set, m[0][2] = %f, want 99", m.Data[0][2])
	}

	// Out of bounds
	err = m.Set(2, 0, 1)
	if err == nil {
		t.Error("Set(2, 0, 1) should return error for out of bounds")
	}
}

func TestMatrixRow(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})

	// Valid row
	row, err := m.Row(0)
	if err != nil {
		t.Errorf("Row(0) unexpected error: %v", err)
	}
	if len(row.Elements) != 3 {
		t.Errorf("Row(0) length = %d, want 3", len(row.Elements))
	}
	if row.Elements[0].(*Float).Value != 1 {
		t.Errorf("Row(0)[0] = %f, want 1", row.Elements[0].(*Float).Value)
	}

	// Out of bounds
	_, err = m.Row(2)
	if err == nil {
		t.Error("Row(2) should return error for out of bounds")
	}
}

func TestMatrixCol(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})

	// Valid col
	col, err := m.Col(1)
	if err != nil {
		t.Errorf("Col(1) unexpected error: %v", err)
	}
	if len(col.Elements) != 2 {
		t.Errorf("Col(1) length = %d, want 2", len(col.Elements))
	}
	if col.Elements[0].(*Float).Value != 2 {
		t.Errorf("Col(1)[0] = %f, want 2", col.Elements[0].(*Float).Value)
	}
	if col.Elements[1].(*Float).Value != 5 {
		t.Errorf("Col(1)[1] = %f, want 5", col.Elements[1].(*Float).Value)
	}

	// Out of bounds
	_, err = m.Col(3)
	if err == nil {
		t.Error("Col(3) should return error for out of bounds")
	}
}

func TestMatrixInspectNonIntegerValues(t *testing.T) {
	// Test that non-integer values are formatted with %.4g
	m := NewMatrix([][]float64{{1.2345, 2.6789}, {3.14159, 4.0}})
	inspect := m.Inspect()

	// Should contain formatted values
	if inspect == "" {
		t.Error("Inspect() returned empty string")
	}

	// Check that it contains the dimension header
	if !contains(inspect, "2x2 matrix") {
		t.Error("Inspect() should contain dimension header")
	}
}

func TestMatrixGetOutOfBoundsColumn(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}, {3, 4}})

	// Test column out of bounds
	_, err := m.Get(0, 2)
	if err == nil {
		t.Error("Get(0, 2) should return error for column out of bounds")
	}

	// Test negative column
	_, err = m.Get(0, -1)
	if err == nil {
		t.Error("Get(0, -1) should return error for negative column")
	}
}

func TestMatrixSetOutOfBoundsColumn(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}, {3, 4}})

	// Test column out of bounds
	err := m.Set(0, 2, 99)
	if err == nil {
		t.Error("Set(0, 2, 99) should return error for column out of bounds")
	}
}

func TestMatrixMethodGet(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()

	// Valid get
	result := m.InvokeMethod("get", *env, NewInteger(0), NewInteger(2))
	floatResult, ok := result.(*Float)
	if !ok {
		t.Fatalf("get method should return Float, got %T", result)
	}
	if floatResult.Value != 3.0 {
		t.Errorf("get(0, 2) = %f, want 3.0", floatResult.Value)
	}

	// Out of bounds get
	result = m.InvokeMethod("get", *env, NewInteger(2), NewInteger(0))
	_, ok = result.(*Error)
	if !ok {
		t.Errorf("get method with out of bounds should return Error, got %T", result)
	}
}

func TestMatrixMethodSet(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()

	// Valid set with integer
	result := m.InvokeMethod("set", *env, NewInteger(0), NewInteger(2), NewInteger(99))
	if result.Type() != NIL_OBJ {
		t.Errorf("set method should return NIL, got %s", result.Type())
	}
	if m.Data[0][2] != 99.0 {
		t.Errorf("After set, m[0][2] = %f, want 99.0", m.Data[0][2])
	}

	// Valid set with float
	result = m.InvokeMethod("set", *env, NewInteger(1), NewInteger(1), NewFloat(3.14))
	if result.Type() != NIL_OBJ {
		t.Errorf("set method should return NIL, got %s", result.Type())
	}
	if m.Data[1][1] != 3.14 {
		t.Errorf("After set, m[1][1] = %f, want 3.14", m.Data[1][1])
	}

	// Out of bounds set
	result = m.InvokeMethod("set", *env, NewInteger(2), NewInteger(0), NewInteger(1))
	_, ok := result.(*Error)
	if !ok {
		t.Errorf("set method with out of bounds should return Error, got %T", result)
	}
}

func TestMatrixMethodRow(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()

	// Valid row
	result := m.InvokeMethod("row", *env, NewInteger(1))
	arr, ok := result.(*Array)
	if !ok {
		t.Fatalf("row method should return Array, got %T", result)
	}
	if len(arr.Elements) != 3 {
		t.Errorf("row(1) length = %d, want 3", len(arr.Elements))
	}
	if arr.Elements[0].(*Float).Value != 4.0 {
		t.Errorf("row(1)[0] = %f, want 4.0", arr.Elements[0].(*Float).Value)
	}

	// Out of bounds row
	result = m.InvokeMethod("row", *env, NewInteger(2))
	_, ok = result.(*Error)
	if !ok {
		t.Errorf("row method with out of bounds should return Error, got %T", result)
	}
}

func TestMatrixMethodCol(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	env := NewEnvironment()

	// Valid col
	result := m.InvokeMethod("col", *env, NewInteger(2))
	arr, ok := result.(*Array)
	if !ok {
		t.Fatalf("col method should return Array, got %T", result)
	}
	if len(arr.Elements) != 2 {
		t.Errorf("col(2) length = %d, want 2", len(arr.Elements))
	}
	if arr.Elements[0].(*Float).Value != 3.0 {
		t.Errorf("col(2)[0] = %f, want 3.0", arr.Elements[0].(*Float).Value)
	}
	if arr.Elements[1].(*Float).Value != 6.0 {
		t.Errorf("col(2)[1] = %f, want 6.0", arr.Elements[1].(*Float).Value)
	}

	// Out of bounds col
	result = m.InvokeMethod("col", *env, NewInteger(3))
	_, ok = result.(*Error)
	if !ok {
		t.Errorf("col method with out of bounds should return Error, got %T", result)
	}
}

func TestMatrixToStringObj(t *testing.T) {
	m := NewMatrix([][]float64{{1, 2}, {3, 4}})
	strObj := m.ToStringObj(nil)

	if strObj.Type() != STRING_OBJ {
		t.Errorf("ToStringObj() type = %s, want STRING", strObj.Type())
	}

	// Should contain the matrix representation
	if !contains(strObj.Value, "2x2 matrix") {
		t.Error("ToStringObj() should contain dimension header")
	}
}
