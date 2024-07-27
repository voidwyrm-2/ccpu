package main

import "fmt"

var INSTRUCTIONS = []string{
	"mov",
	"add",
	"sub",
	"and",
	"or",
	"xor",
	"store",
	"load",
	"adr",
	"push",
	"pop",
	"inc",
	"dec",
	"jeq",
	"jne",
	"jlt",
	"jgt",
	"jmp",
	"jal",
}

var INSTRUCTION_BYTES = map[string]int{
	"mov":   4,
	"add":   4,
	"sub":   4,
	"and":   4,
	"or":    4,
	"xor":   4,
	"store": 5,
	"load":  5,
	"adr":   5,
	"push":  4,
	"pop":   4,
	"inc":   4,
	"dec":   4,
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
