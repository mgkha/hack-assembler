package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	code "hack-assembler/lib/code"
	parser "hack-assembler/lib/parser"
	symbol "hack-assembler/lib/symbol"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fileName := os.Args[1]
	parser.Parse(fileName)

	outFilename := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".hack"

	outFile, err := os.Create(outFilename)
	check(err)
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	// first pass
	var i uint16 = 0
	for parser.HasMoreLines() {
		parser.Advance()
		if parser.InstructionType() == parser.L_INSTRUCTION {
			symbol.AddEntry(parser.Symbol(), i)
		} else {
			i++
		}
	}

	// second pass
	var memvar uint16 = 16
	parser.Parse(fileName)
	for parser.HasMoreLines() {
		parser.Advance()

		if parser.InstructionType() == parser.A_INSTRUCTION {

			// check if decimal value
			if regexp.MustCompile(`^[0-9]+$`).MatchString(parser.Symbol()) {
				d, _ := strconv.Atoi(parser.Symbol())
				fmt.Fprintf(writer, "%016b\n", d)
			} else {
				if !symbol.Contains(parser.Symbol()) {
					symbol.AddEntry(parser.Symbol(), memvar)
					memvar++
				}
				fmt.Fprintf(writer, "%016b\n", symbol.GetAddress(parser.Symbol()))
			}
		}
		if parser.InstructionType() == parser.C_INSTRUCTION {
			fmt.Fprintf(writer, "111%v%v%v\n", code.Comp(parser.Comp()), code.Dest(parser.Dest()), code.Jump(parser.Jump()))
		}
	}
	writer.Flush()
	fmt.Println("Compilation Finished")
}
