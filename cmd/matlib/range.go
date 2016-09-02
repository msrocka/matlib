package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/GreenDelta/matlib"
)

var (
	errRangeFormat = errors.New("Range has invalid format.")
	errOutOfBounds = errors.New("Range is larger than matrix size")
)

type tRange struct {
	startRow int
	endRow   int
	startCol int
	endCol   int
}

func newRange() *tRange {
	return &tRange{startRow: -1, endRow: -1, startCol: -1, endCol: -1}
}

func (r *tRange) String() string {
	str := func(i int) string {
		if i < 0 {
			return ""
		}
		return strconv.Itoa(i)
	}
	return "[" + str(r.startRow) + ":" + str(r.endRow) + ", " +
		str(r.startCol) + ":" + str(r.endCol) + "]"
}

func parseRange(s string) (*tRange, error) {
	sub := strings.TrimSpace(s)
	if !strings.HasPrefix(sub, "[") || !strings.HasSuffix(sub, "]") {
		return nil, errRangeFormat
	}
	sub = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(sub, "["), "]"))
	if !strings.Contains(sub, ",") {
		return nil, errRangeFormat
	}
	parts := strings.Split(sub, ",")
	if len(parts) != 2 {
		return nil, errRangeFormat
	}
	r := newRange()
	var errRow, errCol error
	r.startRow, r.endRow, errRow = parseRangePart(parts[0])
	r.startCol, r.endCol, errCol = parseRangePart(parts[1])
	if errRow != nil || errCol != nil {
		return nil, errRangeFormat
	}
	return r, nil
}

func parseRangePart(part string) (int, int, error) {
	p := strings.TrimSpace(part)
	if p == ":" {
		return -1, -1, nil
	}
	if !strings.Contains(p, ":") {
		return -1, -1, errRangeFormat
	}
	parts := strings.Split(p, ":")
	if len(parts) != 2 {
		return -1, -1, errRangeFormat
	}

	atoi := func(sub string) (int, error) {
		s := strings.TrimSpace(sub)
		if s == "" {
			return -1, nil
		}
		return strconv.Atoi(s)
	}

	start, errS := atoi(parts[0])
	end, errE := atoi(parts[1])
	if errS != nil || errE != nil {
		return -1, -1, errRangeFormat
	}
	if start > end {
		end, start = start, end
	}
	return start, end, nil
}

func (r *tRange) apply(m *matlib.Matrix) error {
	if r.startRow < 0 {
		r.startRow = 0
	}
	if r.startCol < 0 {
		r.startCol = 0
	}
	if r.endRow < 0 {
		r.endRow = m.Rows
	}
	if r.endCol < 0 {
		r.endCol = m.Cols
	}
	if r.startRow > m.Rows || r.endRow > m.Rows {
		return errOutOfBounds
	}
	if r.startCol > m.Cols || r.endCol > m.Cols {
		return errOutOfBounds
	}
	return nil
}
