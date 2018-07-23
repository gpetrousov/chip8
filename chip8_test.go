package main

import (
	"chip8/lib"
	"fmt"
	"os"
	"testing"
)

func TestChip8(t *testing.T) {

	var (
		memory [4096]uint8
		lenROM uint16
	)
	emptyROM := true
	os.Args[1] = "roms/pong.rom"
	pc := uint16(0x200)
	opcode := uint16(0x0)

	// Test ROM loading into memory
	memory, lenROM = lib.LoadROMIntoMemory(memory)

	for _, v := range memory {
		if v != 0 {
			emptyROM = false
			break
		}
	}

	if emptyROM {
		t.Errorf("Expected ROM to have content")
	} else {
		fmt.Println("Memory loaded successfully")
		fmt.Println("Memory length:", lenROM)
		fmt.Println(memory)
	}

	// Test first opcode fetching
	opcode = lib.FetchOpcode(memory, pc)
	if opcode == 0x0 {
		t.Errorf("Expected non-zero opcode")
	} else {
		fmt.Println("Fetched first opcode successfully:", opcode)
	}

	// Test decoding
	for pc := uint16(0x200); pc < (uint16(0x200) + lenROM); pc += 1 {
		opcode = lib.FetchOpcode(memory, pc)
		lib.DecodeAndExecuteOpcode(opcode, pc)
	}

}
