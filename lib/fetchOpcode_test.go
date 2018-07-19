package lib

import (
	"os"
	"testing"
)

func TestFetchOpcode(t *testing.T) {

	pc := uint16(0x200)
	opcode := uint16(0x0)
	var memory [4096]uint8
	os.Args[1] = "../pong.rom"

	memory = LoadROMIntoMemory(memory)
	opcode = FetchOpcode(memory, pc)

	if opcode == 0x0 {
		t.Errorf("Expected non-zero opcode")
	}
}
