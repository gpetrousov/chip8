package lib

// Take memory and PC and return next opcode
func FetchOpcode(memory [4096]uint8, pc uint16) uint16 {
	opcode := uint16(memory[pc])<<8 | uint16(memory[pc+1])
	return opcode
}
