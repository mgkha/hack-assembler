package symbol

var symbolTable = map[string]uint16{
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"SCREEN": 16384,
	"KBD":    24576,
}

func AddEntry(symbol string, addresss uint16) {
	symbolTable[symbol] = addresss
}

func Contains(symbol string) bool {
	_, ok := symbolTable[symbol]
	return ok
}

func GetAddress(symbol string) uint16 {
	address := symbolTable[symbol]
	return address
}
