package main

import (
	"fmt"
	"os"

	"github.com/igoracmelo/ch8/asm"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("<Usage>") // TODO:
		return
	}

	if os.Args[1] == "disasm" {
		// TODO: safety check

		instructions, err := asm.DisassembleFile(os.Args[2])
		if err != nil {
			panic(err)
		}

		for _, instr := range instructions {
			fmt.Println(instr)
		}
	}

}
