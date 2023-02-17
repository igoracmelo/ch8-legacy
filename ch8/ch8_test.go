package ch8

import (
	"bytes"
	"log"
	"testing"

	"github.com/igoracmelo/ch8/asm"
	"github.com/igoracmelo/ch8/assert"
	"github.com/igoracmelo/ch8/format"
)

func TestRun(t *testing.T) {
	args := []string{"ch8", "disasm", "/none"}
	buf := bytes.Buffer{}
	logger := log.New(&buf, "", 0)
	asm := asm.New(format.Format{DisableColors: true})

	Run(args, logger, asm)
	out := buf.String()
	assert.StringContains(t, out, "123")
}
