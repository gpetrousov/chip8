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
		fmt.Println("ROM loaded successfully")
		fmt.Println("ROM length:", lenROM)
		fmt.Println(memory)
	}
}
