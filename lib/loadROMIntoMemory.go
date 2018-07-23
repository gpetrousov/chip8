package lib

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Load ROM into memory
func LoadROMIntoMemory(m [4096]uint8) ([4096]uint8, uint16) {
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	fmt.Println("ROM len:", len(bs))
	if len(bs) > 512 {
		fmt.Println("ERROR: Length ROM is greater that available memory.")
		os.Exit(2)
	}
	for romI, memI := 0, 512; romI < len(bs); memI, romI = memI+1, romI+1 {
		m[memI] = bs[romI]
	}
	return m, uint16(len(bs))
}