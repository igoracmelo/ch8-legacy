package ch8

import (
	"log"

	"github.com/igoracmelo/ch8/asm"
)

func Run(args []string, logger *log.Logger, a *asm.Asm) {
	if len(args) == 1 {
		logger.Println("<Usage>") // TODO:
		return
	}

	if args[1] == "disasm" {
		// TODO: safety check

		instructions, err := a.DisassembleFile(args[2])
		if err != nil {
			// TODO:
			panic(err)
		}

		for _, instr := range instructions {
			logger.Println(instr)
		}
	}
}
