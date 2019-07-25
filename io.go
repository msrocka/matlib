package matlib

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"os"
	"unsafe"

	"golang.org/x/exp/mmap"
)

// LoadColumn reads the given (zero-based) column from the given matrix file.
func LoadColumn(file string, column int) ([]float64, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows, _, err := readShape(f)
	if err != nil {
		return nil, err
	}

	offset := column*rows*8 + 8
	_, err = f.Seek(int64(offset), 0)
	if err != nil {
		return nil, err
	}

	buf := bufio.NewReader(f)
	data := make([]float64, rows)
	bin8 := make([]byte, 8)
	for row := 0; row < rows; row++ {
		val, err := readFloat(bin8, buf)
		if err != nil {
			return nil, err
		}
		data[row] = val
	}
	return data, nil
}

// Load reads a matrix from the given file.
func Load(file string) (*Matrix, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := bufio.NewReader(f)

	rows, cols, err := readShape(buf)
	if err != nil {
		return nil, err
	}

	m := Zeros(rows, cols)
	bin8 := make([]byte, 8)
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			val, err := readFloat(bin8, buf)
			if err != nil {
				return nil, err
			}
			m.Set(row, col, val)
		}
	}
	return m, nil
}

func MemMap(file string) (*Matrix, error) {

	r, err := mmap.Open(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	rows, cols, err := readShapeAt(r)
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, 8*rows*cols)
	n, err := r.ReadAt(bytes, 8)
	if err != nil {
		return nil, err
	}
	if n != len(bytes) {
		return nil, errors.New(
			"Failed to read matrix; not enough data")
	}

	data := (*[]float64)(unsafe.Pointer(&bytes))
	m := &Matrix{
		Rows: rows,
		Cols: cols,
		Data: *data,
	}

	/*
		m := Zeros(rows, cols)
		bin8 := make([]byte, 8)
		var offset int64 = 8
		for col := 0; col < cols; col++ {
			for row := 0; row < rows; row++ {
				val, err := readFloatAt(bin8, r, offset)
				offset += 8
				if err != nil {
					return nil, err
				}
				m.Set(row, col, val)
			}
		}
	*/
	return m, nil
}

func readShape(reader io.Reader) (int, int, error) {
	bin4 := make([]byte, 4)
	rows, err := readInt(bin4, reader)
	if err != nil {
		return -1, -1, err
	}
	cols, err := readInt(bin4, reader)
	if err != nil {
		return -1, -1, err
	}
	return rows, cols, nil
}

func readShapeAt(reader io.ReaderAt) (int, int, error) {
	bin4 := make([]byte, 4)
	rows, err := readIntAt(bin4, reader, 0)
	if err != nil {
		return -1, -1, err
	}
	cols, err := readIntAt(bin4, reader, 4)
	if err != nil {
		return -1, -1, err
	}
	return rows, cols, nil
}

func readFloat(bin8 []byte, r io.Reader) (float64, error) {
	n, err := r.Read(bin8)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, errors.New("Failed to read float: n != 8")
	}
	bits := binary.LittleEndian.Uint64(bin8)
	float := math.Float64frombits(bits)
	return float, err
}

func readFloatAt(bin8 []byte, r io.ReaderAt, offset int64) (float64, error) {
	n, err := r.ReadAt(bin8, offset)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, errors.New("Failed to read float: n != 8")
	}
	bits := binary.LittleEndian.Uint64(bin8)
	float := math.Float64frombits(bits)
	return float, err
}

func readInt(bin4 []byte, r io.Reader) (int, error) {
	n, err := r.Read(bin4)
	if err != nil {
		return 0, err
	}
	if n != 4 {
		return 0, errors.New("Failed to read int: n != 4")
	}
	return int(binary.LittleEndian.Uint32(bin4)), nil
}

func readIntAt(bin4 []byte, r io.ReaderAt, offset int64) (int, error) {
	n, err := r.ReadAt(bin4, offset)
	if err != nil {
		return 0, err
	}
	if n != 4 {
		return 0, errors.New("Failed to read int: n != 4")
	}
	return int(binary.LittleEndian.Uint32(bin4)), nil
}

// Save writes the matrix to the given file.
func Save(m *Matrix, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	bin4 := make([]byte, 4)
	bin8 := make([]byte, 8)
	buf := bufio.NewWriter(f)

	// rows
	binary.LittleEndian.PutUint32(bin4, uint32(m.Rows))
	_, err = buf.Write(bin4)
	if err != nil {
		return err
	}

	// columns
	binary.LittleEndian.PutUint32(bin4, uint32(m.Cols))
	_, err = buf.Write(bin4)
	if err != nil {
		return err
	}

	for col := 0; col < m.Cols; col++ {
		for row := 0; row < m.Rows; row++ {
			bits := math.Float64bits(m.Get(row, col))
			binary.LittleEndian.PutUint64(bin8, bits)
			_, err := buf.Write(bin8)
			if err != nil {
				return err
			}
		}
	}
	return buf.Flush()
}
