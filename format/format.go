package format

import (
	"fmt"
	"strings"
)

type Format struct {
	DisableColors bool
}

func (c *Format) surround(text string, color string) string {
	if c.DisableColors {
		return text
	}

	noColor := "\033[0m"
	return fmt.Sprintf("%s%s%s", color, text, noColor)
}

func (c *Format) colorNibble(n uint8) string {
	n = n & 0x0F
	color := fmt.Sprintf("\033[0;%dm", (n%7)+31)
	return c.surround(fmt.Sprintf("%01x", n), color)
}

func (c *Format) Instruction(opcode uint16, name string, suffix string) string {
	if len(name) < 4 {
		name += strings.Repeat(" ", 4-len(name))
	}
	return fmt.Sprintf("%s        %s %s", c.Opcode(opcode), c.InstructionName(name), suffix)
}

func (c *Format) InstructionName(text string) string {
	return c.surround(text, "\033[1;37m")
}

func (c *Format) UnknownInstruction(opcode uint16) string {
	return fmt.Sprintf("%s        %s", c.Opcode(opcode), c.Unknown())
}

func (c *Format) Unknown() string {
	return c.surround("(Unknown)", "\033[1;30m")
}

func (c *Format) Opcode(opcode uint16) string {
	return c.surround(fmt.Sprintf("%04x", opcode), "\033[1;30m")

	// s := ""
	// for i := 0; i < 4; i++ {
	// 	n := (opcode >> i) & 0x0F
	// 	s += c.colorNibble(uint8(n))
	// }

	// return s
}

func (c *Format) Address(addr uint16) string {
	text := fmt.Sprintf("$%03X", addr)
	return c.AddressString(text)
}

func (c *Format) AddressString(text string) string {
	return c.surround(text, "\033[0;32m")
}

func (c *Format) Register(r uint8) string {
	text := fmt.Sprintf("V%01X", r)
	return c.RegisterName(text)
}

func (c *Format) RegisterName(text string) string {
	return c.surround(text, "\033[0;36m")
}

func (c *Format) Byte(b uint8) string {
	text := fmt.Sprintf("#%02X", b)
	return c.surround(text, "\033[0;34m")
}

func (c *Format) Nibble(b uint8) string {
	text := fmt.Sprintf("#%01X", b)
	return c.surround(text, "\033[0;34m")
}
