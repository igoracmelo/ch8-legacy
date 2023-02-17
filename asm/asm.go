package asm

import (
	"fmt"
	"os"

	"github.com/igoracmelo/ch8/format"
)

type Asm struct {
	fmt format.Format
}

func New(fmt format.Format) *Asm {
	return &Asm{
		fmt: fmt,
	}
}

func (asm *Asm) DisassembleFile(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(bytes) < 0x200 {
		return nil, fmt.Errorf("File is to small for being a CHIP8 game")
	}

	return asm.Disassemble(bytes), err
}

func (asm *Asm) Disassemble(program []byte) []string {
	result := []string{}

	for i := 0; i+1 < len(program); i += 2 {
		opcode := uint16(program[i])<<8 + uint16(program[i+1])
		instruction := asm.OpcodeToInstruction(opcode)
		result = append(result, fmt.Sprintf("%04X: %s", i+0x200, instruction))
	}

	return result
}

func (asm *Asm) OpcodeToInstruction(opcode uint16) string {
	// naming conventions:
	// a = address
	// r1 = register 1
	// r2 = register 2
	// b = byte
	// n = nibble

	// instructions that have a unique first nibble
	switch opcode & 0xF000 {
	case 0x0000:
		a := opcode & 0x0FFF
		return asm.fmt.Instruction(opcode, "SYS", asm.fmt.Address(a))

	case 0x1000:
		a := opcode & 0x0FFF
		return asm.fmt.Instruction(opcode, "JP", asm.fmt.Address(a))

	case 0x2000:
		a := opcode & 0x0FFF
		return asm.fmt.Instruction(opcode, "CALL", asm.fmt.Address(a))

	case 0x3000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return asm.fmt.Instruction(opcode, "SE", asm.fmt.Register(r1)+", "+asm.fmt.Byte(b))

	case 0x4000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return asm.fmt.Instruction(opcode, "SNE", asm.fmt.Register(r1)+", "+asm.fmt.Byte(b))

	case 0x5000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "SE", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x6000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return asm.fmt.Instruction(opcode, "LD", asm.fmt.Register(r1)+", "+asm.fmt.Byte(b))

	case 0x7000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return asm.fmt.Instruction(opcode, "ADD", asm.fmt.Register(r1)+", "+asm.fmt.Byte(b))

	case 0x9000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "SNE", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0xA000:
		a := opcode & 0x0FFF
		return asm.fmt.Instruction(opcode, "LD", asm.fmt.RegisterName("I")+", "+asm.fmt.Address(a))

	case 0xB000:
		a := opcode & 0x0FFF
		return asm.fmt.Instruction(opcode, "JP", asm.fmt.RegisterName("V0")+", "+asm.fmt.Address(a))

	case 0xC000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return asm.fmt.Instruction(opcode, "RND", asm.fmt.Register(r1)+", "+asm.fmt.Byte(b))

	case 0xD000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		n := uint8(opcode & 0x000F)
		return asm.fmt.Instruction(opcode, "DRW", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2)+", "+asm.fmt.Nibble(n))
	}

	// 8xy#: x and y are registers
	switch opcode & 0xF00F {
	case 0x8000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "LD", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x8001:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "OR", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x8002:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "AND", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x8003:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "XOR", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x8004:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "ADD", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x8005:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "SUB", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x8006:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "SHR", asm.fmt.Register(r1)+" {, "+asm.fmt.Register(r2)+"}")

	case 0x8007:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "SUBN", asm.fmt.Register(r1)+", "+asm.fmt.Register(r2))

	case 0x800E:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return asm.fmt.Instruction(opcode, "SHL", asm.fmt.Register(r1)+" {, "+asm.fmt.Register(r2)+"}")
	}

	// Ex## and Fx##
	switch opcode & 0xF0FF {
	case 0xE09E:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "SKP", asm.fmt.Register(r1))

	case 0xE0A1:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "SKNP", asm.fmt.Register(r1))

	case 0xF007:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", asm.fmt.Register(r1)+", DT")

	case 0xF00A:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", asm.fmt.Register(r1)+", K")

	case 0xF015:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", "DT, "+asm.fmt.Register(r1))

	case 0xF018:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", "ST, "+asm.fmt.Register(r1))

	case 0xF01E:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "ADD", asm.fmt.RegisterName("I")+", "+asm.fmt.Register(r1))

	case 0xF029:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", "F, "+asm.fmt.Register(r1))

	case 0xF033:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", "B, "+asm.fmt.Register(r1))

	case 0xF055:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", asm.fmt.AddressString("[I]")+", "+asm.fmt.Register(r1))

	case 0xF065:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return asm.fmt.Instruction(opcode, "LD", asm.fmt.Register(r1)+", "+asm.fmt.AddressString("[I]"))
	}

	return asm.fmt.UnknownInstruction(opcode)
}
