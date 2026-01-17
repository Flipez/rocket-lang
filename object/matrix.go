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

	// Format all values and determine column widths
	formattedValues := make([][]string, m.Rows)
	colWidths := make([]int, m.Cols)

	for i := 0; i < m.Rows; i++ {
		formattedValues[i] = make([]string, m.Cols)
		for j := 0; j < m.Cols; j++ {
			val := m.Data[i][j]
			var formatted string
			if val == float64(int64(val)) && val >= -1e15 && val <= 1e15 {
				// Integer value - show with .0
				formatted = fmt.Sprintf("%.1f", val)
			} else {
				// Non-integer value - use smart formatting
				formatted = fmt.Sprintf("%.4g", val)
			}
			formattedValues[i][j] = formatted
			if len(formatted) > colWidths[j] {
				colWidths[j] = len(formatted)
			}
		}
	}

	// Calculate total width needed for the matrix content
	totalWidth := 0
	for _, width := range colWidths {
		totalWidth += width
	}
	totalWidth += (m.Cols - 1) * 2 // spaces between columns

	// Header with dimensions
	out.WriteString(fmt.Sprintf("%dx%d matrix\n", m.Rows, m.Cols))

	// Top border
	out.WriteString("┌")
	out.WriteString(strings.Repeat(" ", totalWidth+2))
	out.WriteString("┐")

	// Matrix rows
	for i := 0; i < m.Rows; i++ {
		out.WriteString("\n│ ")
		for j := 0; j < m.Cols; j++ {
			// Right-align each value within its column width
			out.WriteString(fmt.Sprintf("%*s", colWidths[j], formattedValues[i][j]))
			if j < m.Cols-1 {
				out.WriteString("  ")
			}
		}
		out.WriteString(" │")
	}

	// Bottom border
	out.WriteString("\n└")
	out.WriteString(strings.Repeat(" ", totalWidth+2))
	out.WriteString("┘")

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

func (m *Matrix) Transpose() *Matrix {
	result := make([][]float64, m.Cols)
	for i := 0; i < m.Cols; i++ {
		result[i] = make([]float64, m.Rows)
		for j := 0; j < m.Rows; j++ {
			result[i][j] = m.Data[j][i]
		}
	}

	return NewMatrix(result)
}

func (m *Matrix) Get(row, col int) (float64, error) {
	if row < 0 || row >= m.Rows {
		return 0, fmt.Errorf("row index %d out of bounds [0, %d)", row, m.Rows)
	}
	if col < 0 || col >= m.Cols {
		return 0, fmt.Errorf("column index %d out of bounds [0, %d)", col, m.Cols)
	}
	return m.Data[row][col], nil
}

func (m *Matrix) Set(row, col int, value float64) error {
	if row < 0 || row >= m.Rows {
		return fmt.Errorf("row index %d out of bounds [0, %d)", row, m.Rows)
	}
	if col < 0 || col >= m.Cols {
		return fmt.Errorf("column index %d out of bounds [0, %d)", col, m.Cols)
	}
	m.Data[row][col] = value
	return nil
}

func (m *Matrix) Row(index int) (*Array, error) {
	if index < 0 || index >= m.Rows {
		return nil, fmt.Errorf("row index %d out of bounds [0, %d)", index, m.Rows)
	}
	elements := make([]Object, m.Cols)
	for j := 0; j < m.Cols; j++ {
		elements[j] = NewFloat(m.Data[index][j])
	}
	return NewArray(elements), nil
}

func (m *Matrix) Col(index int) (*Array, error) {
	if index < 0 || index >= m.Cols {
		return nil, fmt.Errorf("column index %d out of bounds [0, %d)", index, m.Cols)
	}
	elements := make([]Object, m.Rows)
	for i := 0; i < m.Rows; i++ {
		elements[i] = NewFloat(m.Data[i][index])
	}
	return NewArray(elements), nil
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
		"shape": {
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(ARRAY_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				m := o.(*Matrix)
				return NewArray([]Object{NewInteger(m.Rows), NewInteger(m.Cols)})
			},
		},
		"rows": {
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(INTEGER_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				m := o.(*Matrix)
				return NewInteger(m.Rows)
			},
		},
		"cols": {
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(INTEGER_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				m := o.(*Matrix)
				return NewInteger(m.Cols)
			},
		},
		"size": {
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(INTEGER_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				m := o.(*Matrix)
				return NewInteger(m.Rows * m.Cols)
			},
		},
		"transpose": {
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(MATRIX_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				m := o.(*Matrix)
				return m.Transpose()
			},
		},
		"t": {
			Layout: MethodLayout{
				ReturnPattern: Args(Arg(MATRIX_OBJ)),
			},
			method: func(o Object, _ []Object, _ Environment) Object {
				m := o.(*Matrix)
				return m.Transpose()
			},
		},
		"get": {
			Layout: MethodLayout{
				ArgPattern:    Args(Arg(INTEGER_OBJ), Arg(INTEGER_OBJ)),
				ReturnPattern: Args(Arg(FLOAT_OBJ, ERROR_OBJ)),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				m := o.(*Matrix)
				row := int(args[0].(*Integer).Value)
				col := int(args[1].(*Integer).Value)
				value, err := m.Get(row, col)
				if err != nil {
					return NewErrorFormat("%s", err.Error())
				}
				return NewFloat(value)
			},
		},
		"set": {
			Layout: MethodLayout{
				ArgPattern:    Args(Arg(INTEGER_OBJ), Arg(INTEGER_OBJ), Arg(FLOAT_OBJ, INTEGER_OBJ)),
				ReturnPattern: Args(Arg(NIL_OBJ, ERROR_OBJ)),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				m := o.(*Matrix)
				row := int(args[0].(*Integer).Value)
				col := int(args[1].(*Integer).Value)
				var value float64
				if IsNumber(args[2]) {
					if f, ok := args[2].(*Float); ok {
						value = f.Value
					} else if i, ok := args[2].(*Integer); ok {
						value = float64(i.Value)
					}
				}
				err := m.Set(row, col, value)
				if err != nil {
					return NewErrorFormat("%s", err.Error())
				}
				return NIL
			},
		},
		"row": {
			Layout: MethodLayout{
				ArgPattern:    Args(Arg(INTEGER_OBJ)),
				ReturnPattern: Args(Arg(ARRAY_OBJ, ERROR_OBJ)),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				m := o.(*Matrix)
				index := int(args[0].(*Integer).Value)
				row, err := m.Row(index)
				if err != nil {
					return NewErrorFormat("%s", err.Error())
				}
				return row
			},
		},
		"col": {
			Layout: MethodLayout{
				ArgPattern:    Args(Arg(INTEGER_OBJ)),
				ReturnPattern: Args(Arg(ARRAY_OBJ, ERROR_OBJ)),
			},
			method: func(o Object, args []Object, _ Environment) Object {
				m := o.(*Matrix)
				index := int(args[0].(*Integer).Value)
				col, err := m.Col(index)
				if err != nil {
					return NewErrorFormat("%s", err.Error())
				}
				return col
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
