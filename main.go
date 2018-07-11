package main

import "fmt"

func main() {

	var (
		// Memory
		// memory [4096]uint8
		// // Registers
		// vReg [16]byte //uint8
		// // Address Register
		// iReg uint16
		// // Stack
		// stack [16]uint16
		// // Stack pointer
		// sp uint16
		// // Delay timer
		// delayTimer uint8
		// // Sound timer
		// soundTimer uint8
		// // Opcode (contains the actual code)
		// opcode uint16
		// Program Counter (index to program in ROM)
		pc uint16
	// // Graphics (display)
	// gfx [64 * 32]uint8
	)

	// Initialize
	pc = 0x200

	// MAin loop
	for {
		// Fetch
		// Decode
		// Execute
		// Store
		// Update timers

	}

}

// Take memory and PC and return next opcode to execute
func fetchOpcode(memory []byte, pc uint16) uint16 {
	opcode := uint16(memory[pc]<<8 | memory[pc+1])
	return opcode
}

func decodeOpcode(opcode uint16) {

	// Standard opcodes
	switch opcode {
	case 0x00E0:
		{
			// Clears the screen.
			fmt.Println("Clear screen")
		}
	case 0x00EE:
		{
			// Return from subroutine
			fmt.Println("Return from a subroutine")
		}
	}

	switch opcode & 0xF000 {
	case 0x1000:
		{
			// Jumps to address NNN
			fmt.Println("Jumps to address NNN")
		}
	default:
		fmt.Println("Unknown opcode")
	}
}

func clearScreen() {

}
