package main

import (
	"fmt"
	"slices"
	"strconv"
)

func validForIdent(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

type Lexer struct {
	Text    string
	idx, ln int
	cchar   rune
}

func (l *Lexer) advance() {
	l.idx++
	if l.idx < len(l.Text) {
		l.cchar = rune(l.Text[l.idx])
	} else {
		l.cchar = -1
	}
	if l.cchar == '\n' {
		l.ln++
	}
}

func (l *Lexer) Lex() ([]Token, error) {
	var tokens []Token

	for l.cchar != -1 {
		switch l.cchar {
		case '\t', ' ', '\n':
			l.advance()
		case ',':
			tokens = append(tokens, NewToken(COMMA, ",", l.idx, l.idx, l.ln))
			l.advance()
		case '.':
			l.advance()
			tokens = append(tokens, l.collectInstruction(true))
		case ';':
			tokens = append(tokens, l.collectComment())
		case '"':
			l.advance()
			tokens = append(tokens, l.collectString())
		default:
			if validForIdent(l.cchar) {
				tokens = append(tokens, l.collectInstruction(false))
			} else if l.cchar >= '0' && l.cchar <= '9' {
				tokens = append(tokens, l.collectImmediate(false))
			} else {
				return []Token{}, fmt.Errorf("(line %d, col %d) illegal character '%c'", l.ln, l.idx+1, l.cchar)
			}
		}
	}

	return tokens, nil
}

func (l *Lexer) collectInstruction(returnAsDirective bool) Token {
	start := l.idx
	ident_str := ""

	if l.cchar == 'r' || l.cchar == 'R' {
		ident_str += string(l.cchar)
		l.advance()
		if l.cchar >= '0' && l.cchar <= '9' {
			t := l.collectImmediate(true)
			t.Lit = fmt.Sprint(Assert(strconv.Atoi(t.Lit)) + 8)
			return t
		}
	} else if l.cchar == 'x' || l.cchar == 'X' {
		ident_str += string(l.cchar)
		l.advance()
		if l.cchar >= '0' && l.cchar <= '9' {
			t := l.collectImmediate(true)
			t.Lit = fmt.Sprint(Assert(strconv.Atoi(t.Lit)) + 1)
			return t
		}
	}

	for l.cchar != -1 && validForIdent(l.cchar) {
		ident_str += string(l.cchar)
		l.advance()
	}

	if returnAsDirective {
		return NewToken(DIRECTIVE, ident_str, start, l.idx, l.ln)
	} else if l.cchar == ':' {
		l.advance()
		return NewToken(LABEL, ident_str, start, l.idx, l.ln)
	} else if slices.Contains(INSTRUCTIONS, ident_str) {
		return NewToken(INSTRUCTION, ident_str, start, l.idx, l.ln)
	} else if ident_str == "noop" {
		return NewToken(NOOP, "", start, l.idx, l.ln)
	} else if call, ok := REGISTER_IDS[ident_str]; ok {
		return NewToken(REGCALL, call, start, l.idx, l.ln)
	}

	return NewToken(IDENT, ident_str, start, l.idx, l.ln)
}

func (l *Lexer) collectImmediate(returnAsRegcall bool) Token {
	start := l.idx
	num_str := ""

	for l.cchar != -1 && l.cchar >= '0' && l.cchar <= '9' {
		num_str += string(l.cchar)
		l.advance()
	}

	if returnAsRegcall {
		return NewToken(REGCALL, num_str, start, l.idx, l.ln)
	}
	return NewToken(IMMEDIATE, num_str, start, l.idx, l.ln)
}

func (l *Lexer) collectComment() Token {
	start := l.idx
	comment_str := ""

	for l.cchar == ' ' {
		l.advance()
	}

	for l.cchar != -1 && l.cchar != '\n' {
		comment_str += string(l.cchar)
		l.advance()
	}
	return NewToken(COMMENT, comment_str, start, l.idx, l.ln)
}

func (l *Lexer) collectString() Token {
	start := l.idx
	string_str := ""

	ignore := false
	for l.cchar != -1 {
		if l.cchar == '\\' {
			ignore = true
		} else if ignore {
			if l.cchar == 'n' {
				string_str += "\n"
				l.advance()
				continue
			} else {
				l.advance()
				continue
			}
		} else if l.cchar == '"' && !ignore {
			break
		}

		string_str += string(l.cchar)
		ignore = false
		l.advance()
	}

	l.advance()

	return NewToken(STRING, string_str, start, l.idx, l.ln)
}

func NewLexer(text string) Lexer {
	l := Lexer{Text: text, idx: -1, cchar: -1, ln: 1}
	l.advance()
	return l
}
