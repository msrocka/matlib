package goblapack

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

// Matrix is a dense matrix structure that holds the data in column-major order
// in a linear array. Because of this lay
type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

// NewMatrix creates a new matrix of the give size.
func NewMatrix(rows, cols int) *Matrix {
	size := rows * cols
	m := Matrix{Rows: rows, Cols: cols}
	m.Data = make([]float64, size, size)
	return &m
}

// Get returns the value at the given row and column.
func (m *Matrix) Get(row, col int) float64 {
	i := row + m.Rows*col
	return m.Data[i]
}

// Set sets the matrix cell at the given row and column to the given value.
func (m *Matrix) Set(row, col int, value float64) {
	i := row + m.Rows*col
	m.Data[i] = value
}

// Copy creates a copy of the matrix.
func (m *Matrix) Copy() *Matrix {
	c := NewMatrix(m.Rows, m.Cols)
	copy(c.Data, m.Data)
	return c
}

// Invert calculates the inverse of the matrix.
func (m *Matrix) Invert() *Matrix {
	inverse := m.Copy()
	inverse.InvertInPlace()
	return inverse
}

// InvertInPlace calculates the inverse of the matrix which is stored directly
// in the original matrix.
func (m *Matrix) InvertInPlace() error {
	if m.Cols != m.Rows {
		return errors.New("The matrix is not square")
	}
	handle, err := syscall.GetProcAddress(nativeLib, "goblapack_invert")
	if err != nil {
		return err
	}
	dataPtr := uintptr(unsafe.Pointer(&m.Data[0]))
	r, _, err := syscall.Syscall(uintptr(handle), 2, uintptr(m.Rows), dataPtr, 0)
	if err != nil {
		return err
	}
	info := int(r)
	if info > 0 {
		return errors.New("Matrix is singular")
	}
	if info < 0 {
		return errors.New("Invalid data input")
	}
	return nil
}

func m() {
	lib, err := syscall.LoadLibrary("goblapack.dll")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer syscall.FreeLibrary(lib)

	handle, err := syscall.GetProcAddress(lib, "goblapack_invert")
	data := []float64{0, 0.0, -0.5, 1.0}

	var nargs uintptr = 2
	r, _, err := syscall.Syscall(uintptr(handle), nargs, 2, uintptr(unsafe.Pointer(&data[0])), 0)

	fmt.Println(int32(r))

	fmt.Println(data)
}
