package goblapack

import (
	"os"
	"testing"
)

func TestSaveLoad(t *testing.T) {

	file := os.TempDir() + "/_goblapack_test_io.bin"

	// create and save
	m := NewMatrix(42, 24)
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
