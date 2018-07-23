package lib

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Load ROM into memory
func LoadROMIntoMemory(m []uint8) uint16 {
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	if len(bs) > 512 {
		fmt.Println("ERROR: Length ROM is greater that available memory.")
		os.Exit(2)
	}
	for romI, memI := 0, 512; romI < len(bs); memI, romI = memI+1, romI+1 {
		m[memI] = bs[romI]
	}
	return uint16(len(bs))
}
