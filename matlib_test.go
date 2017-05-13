package matlib

import (
	"math"
	"testing"
)

func TestMakeMatrix(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3}})
	if m.Rows != 4 || m.Cols != 3 {
		t.Error("Wrong matrix size: ", m.Rows, "x", m.Cols)
	}
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			switch col {
			case 0:
				assertEqual(1, m.Get(row, col), t)
			case 1:
				assertEqual(2, m.Get(row, col), t)
			case 3:
				assertEqual(2, m.Get(row, col), t)
			}
		}
	}
}

func TestSetGet(t *testing.T) {
	rows := 4
	cols := 5
	m := Zeros(rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			assertEqual(0, m.Get(row, col), t)
			m.Set(row, col, 42.2)
			assertEqual(42.2, m.Get(row, col), t)
		}
	}
}

func TestGetPtr(t *testing.T) {
	m := Zeros(5, 5)
	for i := 0; i < 5; i++ {
		ptr := m.GetPtr(i, i)
		*ptr = 1
	}
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if row == col {
				assertEqual(1, m.Get(row, col), t)
			} else {
				assertEqual(0, m.Get(row, col), t)
			}
		}
	}
}

func TestSubstract(t *testing.T) {
	// A = [ 1 2 3; 4 5 6]
	a := &Matrix{Rows: 2, Cols: 3, Data: []float64{1, 4, 2, 5, 3, 6}}
	b := Eye(2)
	c, err := a.Subtract(b)
	if err != nil {
		t.Error(err)
	}
	assertArraysEqual([]float64{0, 4, 2, 4, 3, 6}, c.Data, t)
}

func TestScaleColumns(t *testing.T) {
	a := MakeMatrix([][]float64{
		{1, 2, 3},
		{2, 4, 6}})
	b := a.ScaleColumns([]float64{2, 1, 0.5})
	assertArraysEqual([]float64{2, 4, 2, 4, 1.5, 3}, b.Data, t)
}

func TestScaledColumnSums(t *testing.T) {
	a := MakeMatrix([][]float64{
		{1, 2, 3},
		{2, 4, 6}})
	b := a.ScaledColumnSums([]float64{2, 1, 0.5})
	assertArraysEqual([]float64{5.5, 11}, b, t)
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
