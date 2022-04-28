package symbol

var symbolTable = map[string]uint16{}

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
