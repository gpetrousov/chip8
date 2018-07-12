package main

import "fmt"

func main() {

	var (
		// Memory
		// memory [4096]uint8
		// Registers
		vReg [16]uint8
		// Address Register
		iReg uint16
		// Stack
		stack [16]uint16
		// Stack pointer
		sp uint16
		// Delay timer
		delayTimer uint8
		// Sound timer
		soundTimer uint8
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

func decodeOpcode(opcode uint16, stack [16]uint16, pc uint16, sp uint16, vReg []uint8, iReg uint16, delayTimer uint8, soundTimer uint8) {

	// Zero opcodes
	switch opcode {
	case 0x00E0:
		fmt.Println("Clear screen")
		clearScreen()
	case 0x00EE:
		fmt.Println("Return from a subroutine")
	}

	// 1-7 opcodes
	switch opcode & 0xF000 {
	case 0x1000:
		fmt.Println("Jumps to address NNN")
		pc = opcode & 0x0FFF
	case 0x2000:
		fmt.Println("Calls subroutine at NNN")
		stack[sp] = pc
		sp += 1
		pc = opcode & 0x0FFF
	case 0x3000:
		fmt.Println("Skips the next instruction if VX equals NN. (Usually the next instruction is a jump to skip a code block")
	case 0x4000:
		fmt.Println("Skips the next instruction if VX doesn't equal NN. (Usually the next instruction is a jump to skip a code block")
	case 0x5000:
		fmt.Println("Skips the next instruction if VX equals VY. (Usually the next instruction is a jump to skip a code block")
	case 0x6000:
		fmt.Println("Sets VX to NN")
		registerIndex := opcode & 0x0F00
		registerValue := opcode & 0x00FF
		vReg[registerIndex] = uint8(registerValue)
	case 0x7000:
		fmt.Println("Adds NN to VX. (Carry flag is not changed)")
		registerIndex := opcode & 0x0F00
		addValue := opcode & 0x00FF
		vReg[registerIndex] += uint8(addValue)

	// 8 Opcodes
	case 0x8000:
		switch opcode & 0x000F {
		case 0x0000:
			fmt.Println("Sets VX to the value of VY")
			vReg[opcode&0x0F00] = vReg[opcode&0x00F0]
		case 0x0001:
			fmt.Println("Sets VX to VX or VY. (Bitwise OR operation)")
			vReg[opcode&0x0F00] = vReg[opcode&0x0F00] | vReg[opcode&0x00F0]
		case 0x0002:
			fmt.Println("Sets VX to VX and VY. (Bitwise AND operation)")
			vReg[opcode&0x0F00] = vReg[opcode&0x0F00] & vReg[opcode&0x00F0]
		case 0x0003:
			fmt.Println("Sets VX to VX xor VY.")
			vReg[opcode&0x0F00] = vReg[opcode&0x0F00] ^ vReg[opcode&0x00F0]
		case 0x0004:
			fmt.Println("Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.")
		case 0x0005:
			fmt.Println("VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.")
		case 0x0006:
			fmt.Println("Shifts VY right by one and stores the result to VX (VY remains unchanged). VF is set to the value of the least significant bit of VY before the shift.[2]")
		case 0x0007:
			fmt.Println("Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't. ")
		case 0x000E:
			fmt.Println("Shifts VY left by one and copies the result to VX. VF is set to the value of the most significant bit of VY before the shift")
		}

		// 9-D opcodes
	case 0x9000:
		fmt.Println("Skips the next instruction if VX doesn't equal VY. (Usually the next instruction is a jump to skip a code block)")
	case 0xA000:
		fmt.Println("Sets I to the address NNN")
		iReg = opcode & 0x0FFF
	case 0xB000:
		fmt.Println("Jumps to the address NNN plus V0")
	case 0xC000:
		fmt.Println("Sets VX to the result of a bitwise and operation on a random number (Typically: 0 to 255) and NN")
	case 0xD000:
		fmt.Println("Draws a sprite at coordinate (VX, VY)...")

	// E opcodes
	case 0xE000:
		switch opcode & 0x00F0 {
		case 0x0090:
			fmt.Println("Skips the next instruction if the key stored in VX is pressed. (Usually the next instruction is a jump to skip a code block)")
		case 0x00A0:
			fmt.Println("Skips the next instruction if the key stored in VX isn't pressed. (Usually the next instruction is a jump to skip a code block)")
		}

	// F opcodes
	case 0xF000:
		switch opcode & 0x00FF {
		case 0x0007:
			fmt.Println("Sets VX to the value of the delay timer")
			registerIndex := opcode & 0x0F00
			vReg[registerIndex] = delayTimer
		case 0x000A:
			fmt.Println("A key press is awaited, and then stored in VX. (Blocking Operation. All instruction halted until next key event)")
		case 0x0015:
			fmt.Println("Sets the delay timer to VX")
			registerIndex := opcode & 0x0F00
			delayTimer = vReg[registerIndex]
		case 0x0018:
			fmt.Println("Sets the sound timer to VX")
			registerIndex := opcode & 0x0F00
			soundTimer = vReg[registerIndex]
		case 0x001E:
			fmt.Println("Adds VX to I.[3]")
			registerIndex := opcode & 0x0F00
			iReg += uint16(vReg[registerIndex])
		case 0x0029:
			fmt.Println("Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font")
		case 0x0033:
			fmt.Println("Something about BCD, TL;DR")
		case 0x0055:
			fmt.Println("Stores V0 to VX (including VX) in memory starting at address I. The offset from I is increased by 1 for each value written, but I itself is left unmodified")
		case 0x0065:
			fmt.Println("Fills V0 to VX (including VX) with values from memory starting at address I")
		}
	default:
		fmt.Println("Unknown opcode: %v", opcode)
	}
}

func clearScreen() {

}
