package main

import (
	"chip8/lib"
	"fmt"
	"os"
	"os/exec"
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
		lib.DecodeOpcode(opcode, stack, pc, sp, vReg, iReg, delayTimer, soundTimer, memory, screen, keys)

		// Execute

		// Store

		// Update timers

		debounceKeys(keys)
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
