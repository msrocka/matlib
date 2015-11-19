package goblapack

import (
	"math"
	"testing"
)

func TestSetGet(t *testing.T) {
	rows := 4
	cols := 5
	m := NewMatrix(rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			assertEqual(0, m.Get(row, col), t)
			m.Set(row, col, 42.2)
			assertEqual(42.2, m.Get(row, col), t)
		}
	}
}

func TestNoNativeLibError(t *testing.T) {
	if nativeLibError != nil {
		t.Error(nativeLibError)
	}
}

func TestMultiply(t *testing.T) {
	// A = [ 1 2 3; 4 5 6]
	a := &Matrix{Rows: 2, Cols: 3, Data: []float64{1, 4, 2, 5, 3, 6}}
	// B = [ 1 4 ; 2 5 ; 3 6 ]
	b := &Matrix{Rows: 3, Cols: 2, Data: []float64{1, 2, 3, 4, 5, 6}}
	c, err := a.Multiply(b)
	if err != nil {
		t.Error(err)
	}
	assertArraysEqual([]float64{14, 32, 32, 77}, c.Data, t)
}

func TestSubstract(t *testing.T) {
	// A = [ 1 2 3; 4 5 6]
	a := &Matrix{Rows: 2, Cols: 3, Data: []float64{1, 4, 2, 5, 3, 6}}
	b := Eye(2)
	c, err := a.Substract(b)
	if err != nil {
		t.Error(err)
	}
	assertArraysEqual([]float64{0, 4, 2, 4, 3, 6}, c.Data, t)
}

func TestInvertInPlace(t *testing.T) {
	m := NewMatrix(2, 2)
	m.Set(0, 0, 1.0)
	m.Set(0, 1, -0.5)
	m.Set(1, 1, 1.0)
	assertArraysEqual([]float64{1, 0.0, -0.5, 1.0}, m.Data, t)
	err := m.InvertInPlace()
	assertArraysEqual([]float64{1, 0.0, 0.5, 1.0}, m.Data, t)
	if err != nil {
		t.Error(err)
	}
}

func TestInvert(t *testing.T) {
	m := NewMatrix(2, 2)
	m.Set(0, 0, 1.0)
	m.Set(0, 1, -0.5)
	m.Set(1, 1, 1.0)
	inverse, err := m.Invert()
	assertArraysEqual([]float64{1, 0.0, -0.5, 1.0}, m.Data, t)
	assertArraysEqual([]float64{1, 0.0, 0.5, 1.0}, inverse.Data, t)
	if err != nil {
		t.Error(err)
	}
}

func assertArraysEqual(expexted, actual []float64, t *testing.T) {
	if len(expexted) != len(actual) {
		t.Error("array length is different:", len(expexted), "vs", len(actual))
	}
	for i, value := range expexted {
		assertEqual(value, actual[i], t)
	}
}

func assertEqual(expexted, actual float64, t *testing.T) {
	diff := math.Abs(expexted - actual)
	if diff > 1e-10 {
		t.Error("Expected", expexted, "but was", actual)
	}
}
