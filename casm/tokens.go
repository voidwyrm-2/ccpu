package main

import "fmt"

var INSTRUCTIONS = []string{
	"add",
	"addi",
	"sub",
	"subi",
	"and",
	"andi",
	"or",
	"ori",
	"xor",
	"xori",
	"store",
	"load",
	"adr",
	"push",
	"pop",
	"jeq",
	"jne",
	"jlt",
	"jgt",
	"jmp",
	"jal",
}

var INSTRUCTION_BYTES = map[string]int{
	"add":   4,
	"addi":  4,
	"sub":   4,
	"subi":  4,
	"and":   4,
	"andi":  4,
	"or":    4,
	"ori":   4,
	"xor":   4,
	"xori":  4,
	"store": 4,
	"load":  4,
	"adr":   5,
	"push":  2,
	"pop":   2,
	"jeq":   6,
	"jne":   6,
	"jlt":   6,
	"jgt":   6,
	"jmp":   4,
	"jal":   4,
}

var REGISTER_IDS = map[string]string{
	"zero": "0",
	"cc":   "1",
}

const (
	INSTRUCTION TokenType = iota
	REGCALL
	IMMEDIATE
	DIRECTIVE
	STRING
	LABEL
	IDENT
	COMMENT
	COMMA
	NOOP
)

type (
	TokenType uint32
	Token     struct {
		Type           TokenType
		Lit            string
		Start, End, Ln int
	}
)

func (t Token) Istype(_type TokenType) bool {
	return t.Type == _type
}

func (t Token) Isvalue(value string) bool {
	return t.Lit == value
}

func (t Token) Istoken(_type TokenType, value string) bool {
	return t.Type == _type && t.Lit == value
}

func (t Token) Range() int {
	return (t.End - t.Start) + 1
}

func (t Token) Tostr() string {
	return fmt.Sprintf("{%s: '%s', %d, %d..%d(%d)}", []string{
		"INSTRUCTION",
		"REGCALL",
		"IMMEDIATE",
		"DIRECTIVE",
		"STRING",
		"LABEL",
		"IDENT",
		"COMMENT",
		"COMMA",
	}[t.Type], t.Lit, t.Ln, t.Start, t.End, t.Range())
}

func NewToken(_type TokenType, literal string, start, end, ln int) Token {
	return Token{Type: _type, Lit: literal, Start: start, End: end, Ln: ln}
}
