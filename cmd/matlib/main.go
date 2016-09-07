package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GreenDelta/matlib"
)

func main() {

	if len(os.Args) < 2 {
		help()
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "-h", "help":
		help()
	case "-i", "invert":
		invert()
	case "-s", "show":
		show()
	default:
		fmt.Println("Unknown command:", cmd, " (try help)")
	}

}

func help() {
	text := `
matlib

Usage: matlib <command> <args>

-h, help                     prints this help
-i, invert <input> <output>  inverts the matrix in the input file and writes it
                             to the output file
-s, show <input> <range>     shows the matrix content of the input file of the
                             given range. the range contains the start and end
                             rows and columns in Pythons slice notation, 
                             e.g. [3:8,4:6] -> row 3 to 8 and column 4 to 6    
`
	fmt.Println(text)
}

func invert() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments: invert <input> <output>")
		return
	}
	m, err := matlib.Load(os.Args[2])
	if err != nil {
		fmt.Println("Failed to read matrix from", os.Args[2], err.Error())
		return
	}
	err = m.InvertInPlace()
	if err != nil {
		fmt.Println("Failed to invert matrix", os.Args[2], err.Error())
		return
	}
	err = matlib.Save(m, os.Args[3])
	if err != nil {
		fmt.Println("Failed to write inverse to", os.Args[3], err.Error())
	}
}

func show() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments: show <input> <range>")
		return
	}
	r, err := parseRange(os.Args[3])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	m, err := matlib.Load(os.Args[2])
	if err != nil {
		fmt.Println("Failed to read matrix from", os.Args[2], err.Error())
		return
	}

	err = r.apply(m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for row := r.startRow; row < r.endRow; row++ {
		line := ""
		for col := r.startCol; col < r.endCol; col++ {
			val := m.Get(row, col)
			line = line + strconv.FormatFloat(val, 'E', 4, 64) + "\t"
		}
		fmt.Println(line)
	}
}
