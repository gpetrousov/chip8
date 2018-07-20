package lib

import (
	"fmt"
)

func DecodeAndExecuteOpcode(opcode uint16) {

	// Zero opcodes
	switch opcode {
	case 0x00E0:
		fmt.Println("Clear screen")
	case 0x00EE:
		fmt.Println("Return from a subroutine")
	}

	// 1-7 opcodes
	switch opcode & 0xF000 {
	case 0x1000:
		fmt.Println("Jumps to address NNN")
	case 0x2000:
		fmt.Println("Calls subroutine at NNN")
	case 0x3000:
		fmt.Println("Skips the next instruction if VX equals NN. (Usually the next instruction is a jump to skip a code block")
	case 0x4000:
		fmt.Println("Skips the next instruction if VX doesn't equal NN. (Usually the next instruction is a jump to skip a code block")
	case 0x5000:
		fmt.Println("Skips the next instruction if VX equals VY. (Usually the next instruction is a jump to skip a code block")
	case 0x6000:
		fmt.Println("Sets VX to NN")
	case 0x7000:
		fmt.Println("Adds NN to VX.")

	// 8 Opcodes
	case 0x8000:
		switch opcode & 0x000F {
		case 0x0000:
			fmt.Println("Sets VX to the value of VY")
		case 0x0001:
			fmt.Println("Sets VX to VX or VY. (Bitwise OR operation)")
		case 0x0002:
			fmt.Println("Sets VX to VX and VY. (Bitwise AND operation)")
		case 0x0003:
			fmt.Println("Sets VX to VX xor VY.")
		case 0x0004:
			fmt.Println("Adds VX += VY. VF is set to 1 when there's a carry, and to 0 when there isn't.")
		case 0x0005:
			fmt.Println("VX -= VY. VF is set to 0 when there's a borrow, and 1 when there isn't.")
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
	case 0xB000:
		fmt.Println("Jumps to the address NNN plus V0")
	case 0xC000:
		fmt.Println("Sets VX to the result of a bitwise AND operation on a random number (Typically: 0 to 255) and NN")
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
		case 0x000A:
			fmt.Println("A key press is awaited, and then stored in VX. (Blocking Operation. All instruction halted until next key event)")
		case 0x0015:
			fmt.Println("Sets the delay timer to VX")
		case 0x0018:
			fmt.Println("Sets the sound timer to VX")
		case 0x001E:
			fmt.Println("Adds VX to I.")
		case 0x0029:
			fmt.Println("Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font")
		case 0x0033:
			fmt.Println("Store the BCD representation of VX in memory[iReg], memory[iReg+1], memory[iReg+2]")

		case 0x0055:
			fmt.Println("Stores V0 to VX (including VX) in memory starting at address I. The offset from I is increased by 1 for each value written, but I itself is left unmodified")
		case 0x0065:
			fmt.Println("Fills V0 to VX (including VX) with values from memory starting at address I")
		default:
			fmt.Println("Unknown opcode:", opcode)
		}
	}
}
