package asm

import (
	"strings"
	"testing"

	"github.com/igoracmelo/ch8/assert"
	"github.com/igoracmelo/ch8/format"
)

func TestUnkownInstruction(t *testing.T) {
	a := New(format.Format{DisableColors: true})
	instr := a.OpcodeToInstruction(0xFFFF)
	assert.StringContains(t, strings.ToLower(instr), "unknown")
}
