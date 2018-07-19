package lib

import (
	"os"
	"testing"
)

func TestLoadROMIntoMemory(t *testing.T) {
	var memory [4096]uint8
	emptyROM := true
	os.Args[1] = "../pong.rom"

	memory = LoadROMIntoMemory(memory)

	for _, v := range memory {
		if v != 0 {
			emptyROM = false
		}
	}
	if emptyROM {
		t.Errorf("Expected ROM to have content")
	}
}
