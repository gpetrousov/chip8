package lib

import (
	"os"
	"testing"
)

func TestDecodeAndExecuteOpcode(t *testing.T) {
	var memory [4096]uint8
	var opcode uint16
	os.Args[1] = "../pong.rom"

	memory = LoadROMIntoMemory(memory)

	for pc := uint16(0x200); pc < 758; pc += 1 {
		opcode = FetchOpcode(memory, pc)
		DecodeAndExecuteOpcode(opcode)
	}

}
