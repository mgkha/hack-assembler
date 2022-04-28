package parser

type INSTRUCTION_TYPE uint8

const (
	Undefined INSTRUCTION_TYPE = iota
	A_INSTRUCTION
	C_INSTRUCTION
	L_INSTRUCTION
)

func (s INSTRUCTION_TYPE) String() string {
	switch s {
	case A_INSTRUCTION:
		return "A instruction"
	case C_INSTRUCTION:
		return "C instruction"
	case L_INSTRUCTION:
		return "Label declaration"
	}
	return "Unknown instruction"
}
