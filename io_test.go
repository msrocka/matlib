package matlib

import (
	"os"
	"testing"
)

func TestSaveLoad(t *testing.T) {

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
	Save(m, file)
	t.Log("Saved matrix file", file)

	// load and compare
	clone, err := Load(file)
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

func TestLoadColumn(t *testing.T) {
	m := MakeMatrix([][]float64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3}})
	file := os.TempDir() + "/_matlib_test_load_column.bin"
	Save(m, file)

	col, _ := LoadColumn(file, 0)
	assertArraysEqual([]float64{1, 1, 1, 1}, col, t)
	col, _ = LoadColumn(file, 1)
	assertArraysEqual([]float64{2, 2, 2, 2}, col, t)
	col, _ = LoadColumn(file, 2)
	assertArraysEqual([]float64{3, 3, 3, 3}, col, t)
}
