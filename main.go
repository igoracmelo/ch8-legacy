package main

import (
	"log"
	"os"

	"github.com/igoracmelo/ch8/asm"
	"github.com/igoracmelo/ch8/ch8"
	"github.com/igoracmelo/ch8/format"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	a := asm.New(format.Format{})
	ch8.Run(os.Args, logger, a)
}
