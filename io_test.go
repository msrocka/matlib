package matlib

import (
	"os"
	"testing"
)

func TestWriteRead(t *testing.T) {

	file := os.TempDir() + "/_matlib_test_io.bin"

	// create and save
	m := Zeros(42, 24)
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			if row == col {
				m.Set(row, col, 24)
			} else {
				m.Set(row, col, 42)
			}
		}
	}
	WriteMatrix(m, file)
	defer os.Remove(file)
	t.Log("Saved matrix file", file)

	// load and compare
	clone, err := ReadMatrix(file)
	if err != nil {
		t.Error(err)
		return
	}
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			if row == col {
				if clone.Get(row, col) != 24 {
					t.Error("Value (", row, col, ") should be 24")
				}
			} else {
				if clone.Get(row, col) != 42 {
					t.Error("Value (", row, col, ") should be 42")
				}
			}
		}
	}
}

func TestReadColumn(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3}})
	file := os.TempDir() + "/_matlib_test_read_column.bin"
	WriteMatrix(m, file)
	defer os.Remove(file)

	col, _ := ReadColumn(file, 0)
	assertArraysEqual([]float64{1, 1, 1, 1}, col, t)
	col, _ = ReadColumn(file, 1)
	assertArraysEqual([]float64{2, 2, 2, 2}, col, t)
	col, _ = ReadColumn(file, 2)
	assertArraysEqual([]float64{3, 3, 3, 3}, col, t)

}

func TestReadRow(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3}})
	file := os.TempDir() + "/_matlib_test_read_row.bin"
	WriteMatrix(m, file)
	defer os.Remove(file)

	for row := 0; row < m.Rows; row++ {
		vals, err := ReadRow(file, row)
		if err != nil {
			t.Fatal(err)
		}
		for i, val := range [3]float64{1., 2., 3.} {
			if vals[i] != val {
				t.Fatal("ReadRow mapping failed")
			}
		}
	}
}

func TestReadDiag(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25}})
	file := os.TempDir() + "/_matlib_test_read_diag.bin"
	WriteMatrix(m, file)
	defer os.Remove(file)

	diag, err := ReadDiag(file)
	if err != nil {
		t.Fatal(err)
	}

	for i, val := range [5]float64{1., 7., 13., 19., 25.} {
		if diag[i] != val {
			t.Fatal("ReadDiag failed")
		}
	}
}

func TestMemMap(t *testing.T) {
	file := os.TempDir() + "/_matlib_test_memmap.bin"
	m := MakeMatrix([][]float64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3}})
	WriteMatrix(m, file)
	defer os.Remove(file)

	mapped, err := MemMap(file)
	if err != nil {
		t.Fatal("failed to read matrix", err)
	}

	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			if m.Get(row, col) != mapped.Get(row, col) {
				t.Fatal("memory mapping failed")
			}
		}
	}
}
