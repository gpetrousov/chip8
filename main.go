package main

import (
	"chip8/lib"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {

	var (
		// Memory
		memory [4096]uint8
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
		// Opcode (contains the actual code)
		// opcode uint16
		// Program Counter (index to program in ROM)
		pc uint16
		// Graphics (display)
		screen [64 * 32]uint8
	)

	chip8_fontset := []uint8{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}

	keys := map[uint8]uint8{
		0x0: 0,
		0x1: 0,
		0x2: 0,
		0x3: 0,
		0x4: 0,
		0x5: 0,
		0x6: 0,
		0x7: 0,
		0x8: 0,
		0x9: 0,
		0xA: 0,
		0xB: 0,
		0xC: 0,
		0xD: 0,
		0xE: 0,
		0xF: 0,
	}

	// Initialize PC
	pc = 0x200

	// Load fontset
	for i, v := range chip8_fontset {
		memory[i] = v
	}

	// Load ROM
	memory, _ = lib.LoadROMIntoMemory(memory)

	// Darwin instructions
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run() // Disable input buffering
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()              // Do not display entered characters on the screen
	// Linux instructions
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run() // Disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()              // Do not display entered characters on the screen

	// Routine to read from keyboard
	var b []byte = make([]byte, 1)
	go func() {
		for {
			os.Stdin.Read(b)
			fmt.Println("I got the byte", b, "("+string(b)+")")
			bounceKey(b, keys)
		}
	}()

	// MAin loop
	for {
		// Fetch opcode
		opcode := lib.FetchOpcode(memory, pc)

		// Decode
		decodeOpcode(opcode, stack, pc, sp, vReg, iReg, delayTimer, soundTimer, memory, screen, keys)

		// Execute

		// Store

		// Update timers

		debounceKeys(keys)
	}

}

// Decode and execute opcode, increase PC
func decodeOpcode(opcode uint16, stack [16]uint16, pc uint16, sp uint16, vReg [16]uint8, iReg uint16, delayTimer uint8, soundTimer uint8, memory [4096]uint8, screen [64 * 32]uint8, keys map[uint8]uint8) {

	// Zero opcodes
	switch opcode {
	case 0x00E0:
		fmt.Println("Clear screen")
		for i, _ := range screen {
			screen[i] = 0
		}
	case 0x00EE:
		fmt.Println("Return from a subroutine")
		pc = stack[sp]
		sp -= 1
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
		if vReg[(opcode&0x0F00)>>8] == uint8(opcode&0x00FF) {
			pc += 4
		} else {
			pc += 2
		}
	case 0x4000:
		fmt.Println("Skips the next instruction if VX doesn't equal NN. (Usually the next instruction is a jump to skip a code block")
		if vReg[(opcode&0x0F00)>>8] != uint8(opcode&0x00FF) {
			pc += 4
		} else {
			pc += 2
		}
	case 0x5000:
		fmt.Println("Skips the next instruction if VX equals VY. (Usually the next instruction is a jump to skip a code block")
		if vReg[(opcode&0x0F00)>>8] == vReg[(opcode&0x00F0)>>4] {
			pc += 4
		} else {
			pc += 2
		}
	case 0x6000:
		fmt.Println("Sets VX to NN")
		vReg[(opcode&0xF00)>>8] = uint8(opcode & 0x00FF)
		pc += 2
	case 0x7000:
		fmt.Println("Adds NN to VX.")
		vReg[(opcode&0x0F00)>>8] += uint8(opcode & 0x00FF)
		pc += 2

	// 8 Opcodes
	case 0x8000:
		switch opcode & 0x000F {
		case 0x0000:
			fmt.Println("Sets VX to the value of VY")
			vReg[(opcode&0x0F00)>>8] = vReg[(opcode&0x00F0)>>4]
			pc += 2
		case 0x0001:
			fmt.Println("Sets VX to VX or VY. (Bitwise OR operation)")
			vReg[(opcode&0x0F00)>>8] = vReg[(opcode&0x0F00)>>8] | vReg[(opcode&0x00F0)>>4]
			pc += 2
		case 0x0002:
			fmt.Println("Sets VX to VX and VY. (Bitwise AND operation)")
			vReg[(opcode&0x0F00)>>8] = vReg[(opcode&0x0F00)>>8] & vReg[(opcode&0x00F0)>>4]
			pc += 2
		case 0x0003:
			fmt.Println("Sets VX to VX xor VY.")
			vReg[(opcode&0x0F00)>>8] = vReg[(opcode&0x0F00)>>8] ^ vReg[(opcode&0x00F0)>>4]
			pc += 2
		case 0x0004:
			fmt.Println("Adds VX += VY. VF is set to 1 when there's a carry, and to 0 when there isn't.")
			if vReg[(opcode&0x00F0)>>4] > (0xFF - vReg[(opcode&0x0F00)>>8]) {
				vReg[0xF] = 1 //carry
			} else {
				vReg[0xF] = 0
				vReg[(opcode&0x0F00)>>8] += vReg[(opcode&0x00F0)>>4]
			}
			pc += 2
		case 0x0005:
			fmt.Println("VX -= VY. VF is set to 0 when there's a borrow, and 1 when there isn't.")
			if vReg[(opcode&0x00F0)>>4] > vReg[(opcode&0x0F00)>>8] {
				vReg[0xF] = 0 //borrow
			} else {
				vReg[0xF] = 1
				vReg[(opcode&0x0F00)>>8] -= vReg[(opcode&0x00F0)>>4]
			}
			pc += 2
		case 0x0006:
			fmt.Println("Shifts VY right by one and stores the result to VX (VY remains unchanged). VF is set to the value of the least significant bit of VY before the shift.[2]")
			vReg[(opcode&0x0F00)>>4] = vReg[(opcode&0x00F0)>>4] >> 1
			vReg[0xF] = vReg[(opcode&0x0F00)>>8] & 1 // Capture LSB of Vy
			pc += 2
		case 0x0007:
			fmt.Println("Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't. ")
			if vReg[(opcode&0x0F00)>>4] > (0xFF - vReg[(opcode&0x00F0)>>4]) {
				vReg[0xF] = 0 //borrow
			} else {
				vReg[(opcode&0x00F0)>>4] -= vReg[(opcode&0x0F00)>>8]
				vReg[0xF] = 1
			}
			pc += 2
		case 0x000E:
			fmt.Println("Shifts VY left by one and copies the result to VX. VF is set to the value of the most significant bit of VY before the shift")
			vReg[(opcode&0x0F00)>>8] = vReg[(opcode&0x00F0)>>4] << 1
			vReg[0xF] = vReg[(opcode&0x00F0)>>4] & 128 // Capture MSB of Vy
			pc += 2
		}

		// 9-D opcodes
	case 0x9000:
		fmt.Println("Skips the next instruction if VX doesn't equal VY. (Usually the next instruction is a jump to skip a code block)")
		if vReg[(opcode&0x0F00)>>8] != vReg[(opcode&0x00F0)>>4] {
			pc += 4
		} else {
			pc += 2
		}
	case 0xA000:
		fmt.Println("Sets I to the address NNN")
		iReg = opcode & 0x0FFF
		pc += 2
	case 0xB000:
		fmt.Println("Jumps to the address NNN plus V0")
		pc = (opcode & 0x0FFF) + uint16(vReg[0x0])
	case 0xC000:
		fmt.Println("Sets VX to the result of a bitwise AND operation on a random number (Typically: 0 to 255) and NN")
		rand.Seed(time.Now().UnixNano())
		vReg[(opcode&0x0F00)>>8] = uint8(uint16(rand.Intn(256)) & ((opcode & 0x00FF) >> 8))
		pc += 2
	case 0xD000:
		fmt.Println("Draws a sprite at coordinate (VX, VY)...")
		var yline uint16
		N := (opcode & 0x000F)
		X := (opcode & 0x0F00) >> 8
		Y := (opcode & 0x00F0) >> 4
		for yline = 0; yline < N; yline++ {
			data := memory[iReg+yline]               // this retreives the byte for a give line of pixels
			for xpix := uint8(0); xpix < 8; xpix++ { // each bit in data
				if (data & (0x80 >> xpix)) != 0 {
					if screen[vReg[X]+(vReg[Y]*64)] == 1 {
						vReg[0xF] = 1 //there has been a collision
					}
					screen[vReg[X]+(vReg[Y]*64)] ^= 1 //note: coordinate registers from opcode
				}
			}
		}

	// E opcodes
	case 0xE000:
		switch opcode & 0x00F0 {
		case 0x0090:
			fmt.Println("Skips the next instruction if the key stored in VX is pressed. (Usually the next instruction is a jump to skip a code block)")
			if keys[vReg[(opcode&0x0F00)>>8]] == 1 {
				pc += 4
			} else {
				pc += 2
			}
		case 0x00A0:
			fmt.Println("Skips the next instruction if the key stored in VX isn't pressed. (Usually the next instruction is a jump to skip a code block)")
			if keys[vReg[(opcode&0x0F00)>>8]] != 1 {
				pc += 4
			} else {
				pc += 2
			}
		}

	// F opcodes
	case 0xF000:
		switch opcode & 0x00FF {
		case 0x0007:
			fmt.Println("Sets VX to the value of the delay timer")
			vReg[(opcode&0x0F00)>>8] = delayTimer
			pc += 2
		case 0x000A:
			fmt.Println("A key press is awaited, and then stored in VX. (Blocking Operation. All instruction halted until next key event)")
		case 0x0015:
			fmt.Println("Sets the delay timer to VX")
			delayTimer = vReg[(opcode&0x0F00)>>8]
			pc += 2
		case 0x0018:
			fmt.Println("Sets the sound timer to VX")
			soundTimer = vReg[(opcode&0x0F00)>>8]
			pc += 2
		case 0x001E:
			fmt.Println("Adds VX to I.")
			iReg += uint16(vReg[(opcode&0x0F00)>>8])
			if iReg+uint16(vReg[(opcode&0x0F00)>>8]) > 255 {
				vReg[0xF] = 1 // Overflow
			}
			pc += 2
		case 0x0029:
			fmt.Println("Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font")
		case 0x0033:
			fmt.Println("Store the BCD representation of VX in memory[iReg], memory[iReg+1], memory[iReg+2]")
			// v := 123
			// hundreds := v / 100
			// tens := (v - (hundreds * 100)) / 10
			// ones := v % 10
			memory[iReg] = vReg[(opcode&0x0F00)>>8] / 100                     // Hundreds
			memory[iReg+1] = vReg[(opcode&0x0F00)>>8] - (memory[iReg]*100)/10 // Tens
			memory[iReg+2] = vReg[(opcode&0x0F00)>>8] % 10                    // Ones
			pc += 2

		case 0x0055:
			fmt.Println("Stores V0 to VX (including VX) in memory starting at address I. The offset from I is increased by 1 for each value written, but I itself is left unmodified")
			for i := uint16(0); i < ((opcode & 0x0F00) >> 8); i += 1 {
				memory[iReg+i] = vReg[i]
			}
		case 0x0065:
			fmt.Println("Fills V0 to VX (including VX) with values from memory starting at address I")
			for i := uint16(0); i < ((opcode & 0x0F00) >> 8); i += 1 {
				vReg[i] = memory[iReg+i]
			}
			pc += 2
		}
	default:
		fmt.Println("Unknown opcode:", opcode)
	}
}

// Update timers
func updateTimers(soundTimer uint8, delayTimer uint8) {
	if delayTimer > 0 {
		delayTimer -= 1
	}
	if soundTimer > 0 {
		if soundTimer == 1 {
			println("BEEP!")
		}
	}
	soundTimer -= 1
}

// Bounce key by updating it's state (key pressed)
func bounceKey(k []byte, keys map[uint8]uint8) {
	switch string(k) {
	case "1":
		keys[0x1] = 1
	case "2":
		keys[0x2] = 1
	case "3":
		keys[0x3] = 1
	case "4":
		keys[0xC] = 1

	case "q":
		keys[0x4] = 1
	case "w":
		keys[0x5] = 1
	case "e":
		keys[0x6] = 1
	case "r":
		keys[0xD] = 1

	case "a":
		keys[0x7] = 1
	case "s":
		keys[0x8] = 1
	case "d":
		keys[0x9] = 1
	case "f":
		keys[0xE] = 1

	case "z":
		keys[0xA] = 1
	case "x":
		keys[0x0] = 1
	case "c":
		keys[0xB] = 1
	case "v":
		keys[0xF] = 1

	default:
		fmt.Println("Unknown key pressed")
	}
}

// Set all keys to zero/debounce
func debounceKeys(k map[uint8]uint8) {
	for i, _ := range k {
		k[i] = 0
	}
}
