package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Matrix struct {
	Data [][]float64
	Rows int
	Cols int
}

func NewMatrix(data [][]float64) *Matrix {
	if len(data) == 0 {
		return &Matrix{Data: [][]float64{}, Rows: 0, Cols: 0}
	}
	rows := len(data)
	cols := len(data[0])
	return &Matrix{Data: data, Rows: rows, Cols: cols}
}

func NewMatrixFromObjects(elements []Object) (*Matrix, error) {
	if len(elements) == 0 {
		return nil, fmt.Errorf("cannot create matrix from empty array")
	}

	// Validate it's a 2D array
	firstRow, ok := elements[0].(*Array)
	if !ok {
		return nil, fmt.Errorf("matrix must be created from 2D array")
	}

	rows := len(elements)
	cols := len(firstRow.Elements)

	if cols == 0 {
		return nil, fmt.Errorf("matrix rows cannot be empty")
	}

	// Allocate matrix data
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}

	// Convert and validate all elements
	for i, rowObj := range elements {
		row, ok := rowObj.(*Array)
		if !ok {
			return nil, fmt.Errorf("row %d is not an array", i)
		}
		if len(row.Elements) != cols {
			return nil, fmt.Errorf("row %d has inconsistent length (expected %d, got %d)", i, cols, len(row.Elements))
		}

		for j, elem := range row.Elements {
			if !IsNumber(elem) {
				return nil, fmt.Errorf("element [%d][%d] is not a number", i, j)
			}

			// Convert to float64
			switch v := elem.(type) {
			case *Float:
				data[i][j] = v.Value
			case *Integer:
				data[i][j] = float64(v.Value)
			}
		}
	}

	return NewMatrix(data), nil
}

func (m *Matrix) Type() ObjectType { return MATRIX_OBJ }

func (m *Matrix) Inspect() string {
	var out bytes.Buffer

	out.WriteString("Matrix(")
	out.WriteString(fmt.Sprintf("%dx%d", m.Rows, m.Cols))
	out.WriteString(")[\n")

	for i := 0; i < m.Rows; i++ {
		out.WriteString("  [")
		row := make([]string, m.Cols)
		for j := 0; j < m.Cols; j++ {
			row[j] = fmt.Sprintf("%.4g", m.Data[i][j])
		}
		out.WriteString(strings.Join(row, ", "))
		out.WriteString("]")
		if i < m.Rows-1 {
			out.WriteString(",")
		}
		out.WriteString("\n")
	}

	out.WriteString("]")
	return out.String()
}

func (m *Matrix) ToArray() *Array {
	rows := make([]Object, m.Rows)
	for i := 0; i < m.Rows; i++ {
		cols := make([]Object, m.Cols)
		for j := 0; j < m.Cols; j++ {
			cols[j] = NewFloat(m.Data[i][j])
		}
		rows[i] = NewArray(cols)
	}
	return NewArray(rows)
}

func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.Rows != other.Rows || m.Cols != other.Cols {
		return nil, fmt.Errorf("dimension mismatch: cannot add %dx%d and %dx%d matrices", m.Rows, m.Cols, other.Rows, other.Cols)
	}

	result := make([][]float64, m.Rows)
	for i := 0; i < m.Rows; i++ {
		result[i] = make([]float64, m.Cols)
		for j := 0; j < m.Cols; j++ {
			result[i][j] = m.Data[i][j] + other.Data[i][j]
		}
	}

	return NewMatrix(result), nil
}

func (m *Matrix) Subtract(other *Matrix) (*Matrix, error) {
	if m.Rows != other.Rows || m.Cols != other.Cols {
		return nil, fmt.Errorf("dimension mismatch: cannot subtract %dx%d and %dx%d matrices", m.Rows, m.Cols, other.Rows, other.Cols)
	}

	result := make([][]float64, m.Rows)
	for i := 0; i < m.Rows; i++ {
		result[i] = make([]float64, m.Cols)
		for j := 0; j < m.Cols; j++ {
			result[i][j] = m.Data[i][j] - other.Data[i][j]
		}
	}

	return NewMatrix(result), nil
}

func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if m.Cols != other.Rows {
		return nil, fmt.Errorf("incompatible dimensions: cannot multiply %dx%d by %dx%d", m.Rows, m.Cols, other.Rows, other.Cols)
	}

	result := make([][]float64, m.Rows)
	for i := 0; i < m.Rows; i++ {
		result[i] = make([]float64, other.Cols)
		for j := 0; j < other.Cols; j++ {
			var sum float64
			for k := 0; k < m.Cols; k++ {
				sum += m.Data[i][k] * other.Data[k][j]
			}
			result[i][j] = sum
		}
	}

	return NewMatrix(result), nil
}

func init() {
	objectMethods[MATRIX_OBJ] = map[string]ObjectMethod{
		"to_a": {
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(ARRAY_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				m := o.(*Matrix)
				return m.ToArray()
			},
		},
	}
}

func (m *Matrix) InvokeMethod(method string, env Environment, args ...Object) Object {
	return objectMethodLookup(m, method, env, args)
}

func (m *Matrix) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"rows": m.Rows,
		"cols": m.Cols,
		"data": m.Data,
	})
}

func (m *Matrix) ToStringObj(_ *Integer) *String {
	return NewString(m.Inspect())
}
