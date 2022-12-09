package asm

import (
	"fmt"
	"os"

	"github.com/igoracmelo/ch8/format"
)

func DisassembleFile(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(bytes) < 0x200 {
		return nil, fmt.Errorf("File is to small for being a CHIP8 game")
	}

	return Disassemble(bytes), err
}

func Disassemble(program []byte) []string {
	result := []string{}

	for i := 0; i+1 < len(program); i += 2 {
		opcode := uint16(program[i])<<8 + uint16(program[i+1])
		instruction := OpcodeToInstruction(opcode)
		result = append(result, fmt.Sprintf("%04X: %s", i+0x200, instruction))
	}

	return result
}

func OpcodeToInstruction(opcode uint16) string {
	ifmt := format.Format{}
	format := "%04x   " + ifmt.InstructionName("%s") + " %s" // <OPCODE>   <INSTRUCTION> <SUFFIX>

	// instructions that have a unique first nibble
	switch opcode & 0xF000 {
	case 0x0000:
		a := opcode & 0x0FFF
		return ifmt.Instruction(opcode, "SYS", ifmt.Address(a))

	case 0x1000:
		a := opcode & 0x0FFF
		return ifmt.Instruction(opcode, "JP", ifmt.Address(a))

	case 0x2000:
		a := opcode & 0x0FFF
		return ifmt.Instruction(opcode, "CALL", ifmt.Address(a))

	case 0x3000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return ifmt.Instruction(opcode, "SE", ifmt.Register(r1)+", "+ifmt.Byte(b))

	case 0x4000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return ifmt.Instruction(opcode, "SNE", ifmt.Register(r1)+", "+ifmt.Byte(b))

	case 0x5000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "SE", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x6000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return ifmt.Instruction(opcode, "LD", ifmt.Register(r1)+", "+ifmt.Byte(b))

	case 0x7000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return ifmt.Instruction(opcode, "ADD", ifmt.Register(r1)+", "+ifmt.Byte(b))

	case 0x9000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "SNE", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0xA000:
		a := opcode & 0x0FFF
		return ifmt.Instruction(opcode, "LD", ifmt.RegisterName("I")+", "+ifmt.Address(a))

	case 0xB000:
		a := opcode & 0x0FFF
		return ifmt.Instruction(opcode, "JP", ifmt.RegisterName("V0")+", "+ifmt.Address(a))

	case 0xC000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		b := uint8(opcode & 0x00FF)
		return ifmt.Instruction(opcode, "RND", ifmt.Register(r1)+", "+ifmt.Byte(b))

	case 0xD000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		n := uint8(opcode & 0x000F)
		return ifmt.Instruction(opcode, "DRW", ifmt.Register(r1)+", "+ifmt.Register(r2)+", "+ifmt.Nibble(n))
	}

	// 8xy#
	switch opcode & 0xF00F {
	case 0x8000:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "LD", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x8001:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "OR", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x8002:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "AND", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x8003:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "XOR", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x8004:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "ADD", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x8005:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "SUB", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x8006:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "SHR", ifmt.Register(r1)+" {, "+ifmt.Register(r2)+"}")

	case 0x8007:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "SUBN", ifmt.Register(r1)+", "+ifmt.Register(r2))

	case 0x800E:
		r1 := uint8((opcode & 0x0F00) >> 8)
		r2 := uint8((opcode & 0x00F0) >> 4)
		return ifmt.Instruction(opcode, "SHL", ifmt.Register(r1)+" {, "+ifmt.Register(r2)+"}")
	}

	// Ex## and Fx##
	switch opcode & 0xF0FF {
	case 0xE09E:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "SKP", ifmt.Register(r1))

	case 0xE0A1:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "SKNP", ifmt.Register(r1))

	case 0xF007:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "LD", ifmt.Register(r1)+", DT")

	case 0xF00A:
		r1 := (opcode & 0x0F00) >> 8
		return fmt.Sprintf(format, opcode, "LD", fmt.Sprintf(ifmt.Register(uint8(r1))+", K", r1))

	case 0xF015:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "LD", "DT, "+ifmt.Register(r1))

	case 0xF018:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "LD", "ST, "+ifmt.Register(r1))

	case 0xF01E:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "ADD", ifmt.RegisterName("I")+", "+ifmt.Register(r1))

	case 0xF029:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "LD", "F, "+ifmt.Register(r1))

	case 0xF033:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "LD", "B, "+ifmt.Register(r1))

	case 0xF055:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "LD", ifmt.AddressString("[I]")+", "+ifmt.Register(r1))

	case 0xF065:
		r1 := uint8((opcode & 0x0F00) >> 8)
		return ifmt.Instruction(opcode, "LD", ifmt.Register(r1)+", "+ifmt.AddressString("[I]"))
	}

	return ifmt.UnknownInstruction(opcode)
}
