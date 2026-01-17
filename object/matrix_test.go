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
