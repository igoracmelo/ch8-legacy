package asm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	instruction := OpcodeToInstruction(0x1ABC)
	assert.Equal(t, "1abc   JP $0ABC", instruction)
}

func Test2(t *testing.T) {
	instructions, err := DisassembleFile("/home/igor/Downloads/spaceinvaders.ch8")
	assert.NoError(t, err)

	for _, instr := range instructions {
		fmt.Println(instr)
	}
}
