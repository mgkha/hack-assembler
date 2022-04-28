package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var currentInstruction string
var fileReader *bufio.Reader

func Parse(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileReader = bufio.NewReader(f)
}

// Parse the given instruction string to remove
// - Empty lines / indentation
// - Line comments
// - In-line comments
func whiteSpaceParser(s string) string {
	r := regexp.MustCompile(`//.*|\s`)
	return r.ReplaceAllString(s, "")
}

func cInstructionParser(s string) string {
	if !strings.Contains(s, "=") {
		s = "null=" + s
	}
	if !strings.Contains(s, ";") {
		s = s + ";null"
	}
	return s
}

func HasMoreLines() bool {
	p, err := fileReader.Peek(1)
	return err != io.EOF || len(p) > 0
}

func Advance() {
	line, _, _ := fileReader.ReadLine()
	str := string(line)

	parsed := whiteSpaceParser(str)
	if len(parsed) > 0 {
		currentInstruction = parsed
		if InstructionType() == C_INSTRUCTION {
			currentInstruction = cInstructionParser(currentInstruction)
		}
		return
	}
	Advance()
}

func InstructionType() (instructionType INSTRUCTION_TYPE) {
	insA := regexp.MustCompile(`@.+`)
	insL := regexp.MustCompile(`\(.+\)`)
	insC := regexp.MustCompile(`(.+\=.+)|(.+\=.+;.+)|(.+;.+)`)

	switch {
	case insA.MatchString(currentInstruction):
		instructionType = A_INSTRUCTION
	case insC.MatchString(currentInstruction):
		instructionType = C_INSTRUCTION
	case insL.MatchString(currentInstruction):
		instructionType = L_INSTRUCTION
	case len(currentInstruction) == 0:
		instructionType = Undefined
	default:
		panic("error parsing instruction")
	}
	return
}

func Symbol() (symbol string) {
	if InstructionType() == A_INSTRUCTION {
		symbol = strings.Replace(currentInstruction, "@", "", 1)
	} else if InstructionType() == L_INSTRUCTION {
		re := regexp.MustCompile(`\((.*)\)`)
		m := re.FindAllStringSubmatch(currentInstruction, -1)
		symbol = m[0][1]
	}
	return
}

func Dest() (dest string) {
	if InstructionType() == C_INSTRUCTION {
		dest = regexp.MustCompile(`[=;]+`).Split(currentInstruction, -1)[0]
	}
	return
}

func Comp() (comp string) {
	if InstructionType() == C_INSTRUCTION {
		comp = regexp.MustCompile(`[=;]+`).Split(currentInstruction, -1)[1]
	}
	return
}

func Jump() (jump string) {
	if InstructionType() == C_INSTRUCTION {
		jump = regexp.MustCompile(`[=;]+`).Split(currentInstruction, -1)[2]
	}
	return
}
