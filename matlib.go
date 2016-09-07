package matlib

/*
#cgo CFLAGS: -O3
#cgo LDFLAGS: -L. -lmatlib
#include "matlib.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// Matrix is a dense matrix structure that holds the data in column-major order
// in a linear array. Because of this lay
type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

// MakeMatrix is a convenience function for creating a matrix from a 2
// dimensional float slice. It is mainly used for testing purposes.
func MakeMatrix(data [][]float64) *Matrix {
	rows := len(data)
	cols := 1
	for i := range data {
		size := len(data[i])
		if size > cols {
			cols = size
		}
	}
	m := Zeros(rows, cols)
	for i := range data {
		for j := range data[i] {
			m.Set(i, j, data[i][j])
		}
	}
	return m
}

// Zeros creates a new matrix with all values as 0 of the give size.
func Zeros(rows, cols int) *Matrix {
	size := rows * cols
	m := Matrix{Rows: rows, Cols: cols}
	m.Data = make([]float64, size, size)
	return &m
}

// Eye returns the identity matrix of the given size.
func Eye(size int) *Matrix {
	eye := Zeros(size, size)
	for i := 0; i < size; i++ {
		eye.Set(i, i, 1)
	}
	return eye
}

// Get returns the value at the given row and column.
func (m *Matrix) Get(row, col int) float64 {
	i := row + m.Rows*col
	return m.Data[i]
}

// GetPtr returns a pointer to the matrix cell with the given
// row and column.
func (m *Matrix) GetPtr(row, col int) *float64 {
	i := row + m.Rows*col
	return &m.Data[i]
}

// Set sets the matrix cell at the given row and column to the given value.
func (m *Matrix) Set(row, col int, value float64) {
	i := row + m.Rows*col
	m.Data[i] = value
}

// Copy creates a copy of the matrix.
func (m *Matrix) Copy() *Matrix {
	c := Zeros(m.Rows, m.Cols)
	copy(c.Data, m.Data)
	return c
}

// Subtract calculates A - B = C where A is the matrix on which this method is
// called, B the method parameter, and C the return value. The matrix B can be
// smaller as A; C will have the same size as A.
func (m *Matrix) Subtract(b *Matrix) (*Matrix, error) {
	if b.Rows > m.Rows || b.Cols > m.Cols {
		return nil, errors.New("Matrix substraction failed: B is larger than A")
	}
	c := m.Copy()
	for row := 0; row < b.Rows; row++ {
		for col := 0; col < b.Cols; col++ {
			valA := m.Get(row, col)
			valB := b.Get(row, col)
			c.Set(row, col, valA-valB)
		}
	}
	return c, nil
}

// Multiply calculates the matrix-matrix-product C = A * B where A is the matrix
// on which the method is called, B the method parameter, and C the return value.
func (m *Matrix) Multiply(b *Matrix) (*Matrix, error) {
	if m.Cols != b.Rows {
		return nil, errors.New("Cannot multiply matrices: shapes do not match")
	}
	c := Zeros(m.Rows, b.Cols)
	aPtr := (*C.double)(unsafe.Pointer(&m.Data[0]))
	bPtr := (*C.double)(unsafe.Pointer(&b.Data[0]))
	cPtr := (*C.double)(unsafe.Pointer(&c.Data[0]))
	C.matlib_mmult(
		C.int(m.Rows),
		C.int(b.Cols),
		C.int(m.Cols),
		aPtr,
		bPtr,
		cPtr)
	return c, nil
}

// Invert calculates the inverse of the matrix.
func (m *Matrix) Invert() (*Matrix, error) {
	inverse := m.Copy()
	err := inverse.InvertInPlace()
	return inverse, err
}

// InvertInPlace calculates the inverse of the matrix which is stored directly
// in the original matrix.
func (m *Matrix) InvertInPlace() error {
	if m.Cols != m.Rows {
		return errors.New("The matrix is not square")
	}
	dataPtr := (*C.double)(unsafe.Pointer(&m.Data[0]))
	r := C.matlib_invert(C.int(m.Rows), dataPtr)
	info := int(r)
	if info > 0 {
		return errors.New("Matrix is singular")
	}
	if info < 0 {
		return errors.New("Invalid data input")
	}
	return nil
}
