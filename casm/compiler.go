package main

import (
	"fmt"
	"strconv"

	"slices"
)

func Errf(t Token, format string, a ...any) error {
	return fmt.Errorf("(line %d, col %d) %s", t.Ln, t.Start+1, fmt.Sprintf(format, a...))
}

func Assert[T any](v T, _ error) T {
	return v
}

func UnpackNumberIntoBytes(n int) []byte {
	var out []byte

	for {
		if n < 255 {
			break
		}
		n -= 255
		out = append(out, 255)
	}

	out = append(out, byte(n))

	return out
}

func interpret(tokens_uncleaned []Token) ([]byte, error) {
	var tokens []Token
	for _, t := range tokens_uncleaned {
		if !t.Istype(COMMENT) {
			tokens = append(tokens, t)
		}
	}

	var out []byte

	labels := make(map[string]int)

	byteAcc := 0
	for i, t := range tokens {
		if t.Istype(LABEL) {
			if _, ok := labels[t.Lit]; ok {
				return []byte{}, Errf(t, "label '%s' is already defined", t.Lit)
			}
			labels[t.Lit] = byteAcc
		} else if b, ok := INSTRUCTION_BYTES[t.Lit]; ok {
			byteAcc += b
		} else if t.Istoken(DIRECTIVE, "ascii") || t.Istoken(DIRECTIVE, "asciz") || t.Istoken(DIRECTIVE, "ascin") {
			byteAcc += len([]byte(tokens[i+1].Lit))
		} else if tokens[i].Istoken(IDENT, "syscall") || tokens[i].Istype(NOOP) {
			byteAcc += 4
		}
		//fmt.Println(byteAcc, i, t.Tostr())
	}

	i := 0
	for i < len(tokens) {
		if tokens[i].Istype(INSTRUCTION) {
			switch tokens[i].Lit {
			case "addi", "subi", "andi", "ori", "xori": // 3 inputs; 2 register calls, 1 immediate(1:1:1:1 bytes)
				if len(tokens)-i < 3 {
					return []byte{}, Errf(tokens[i], "expected '%s [src1], [src2], [dst]'", tokens[i].Lit)
				}

				if !tokens[i+1].Istype(IMMEDIATE) {
					return []byte{}, Errf(tokens[i+1], "expected immediate, but found '%s' instead", tokens[i+1].Lit)
				} else if !tokens[i+2].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+2], "expected ',', but found '%s' instead", tokens[i+2].Lit)
				} else if !tokens[i+3].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i+3], "expected register call, but found '%s' instead", tokens[i+3].Lit)
				} else if !tokens[i+4].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+4], "expected ',', but found '%s' instead", tokens[i+4].Lit)
				} else if !tokens[i+5].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i+5], "expected register call, but found '%s' instead", tokens[i+5].Lit)
				}

				out = append(out,
					byte(slices.Index(INSTRUCTIONS, tokens[i].Lit)),
					byte(Assert(strconv.Atoi(tokens[i+1].Lit))),
					byte(Assert(strconv.Atoi(tokens[i+3].Lit))),
					byte(Assert(strconv.Atoi(tokens[i+5].Lit))),
				)
				i += 6
			case "add", "sub", "and", "or", "xor": // 3 inputs; 3 register calls(1:1:1:1 bytes)
				if len(tokens)-i < 3 {
					return []byte{}, Errf(tokens[i], "expected '%s [src1], [src2], [dst]'", tokens[i].Lit)
				}

				if !tokens[i+1].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i+1], "expected register call, but found '%s' instead", tokens[i+1].Lit)
				} else if !tokens[i+2].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+2], "expected ',', but found '%s' instead", tokens[i+2].Lit)
				} else if !tokens[i+3].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i+3], "expected register call, but found '%s' instead", tokens[i+3].Lit)
				} else if !tokens[i+4].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+4], "expected ',', but found '%s' instead", tokens[i+4].Lit)
				} else if !tokens[i+5].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i+5], "expected register call, but found '%s' instead", tokens[i+5].Lit)
				}

				out = append(out,
					byte(slices.Index(INSTRUCTIONS, tokens[i].Lit)),
					byte(Assert(strconv.Atoi(tokens[i+1].Lit))),
					byte(Assert(strconv.Atoi(tokens[i+3].Lit))),
					byte(Assert(strconv.Atoi(tokens[i+5].Lit))),
				)
				i += 6
			case "store", "load": // 2 inputs(1:1:2 bytes)
				if len(tokens)-i < 2 {
					return []byte{}, Errf(tokens[i], "expected '%s [src/dst], [address]'", tokens[i].Lit)
				}

				if !tokens[i+1].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i], "expected register call, but found '%s' instead", tokens[i].Lit)
				} else if !tokens[i+2].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+2], "expected ',', but found '%s' instead", tokens[i+2].Lit)
				} else if !tokens[i+3].Istype(IMMEDIATE) {
					return []byte{}, Errf(tokens[i+1], "expected immediate, but found '%s' instead", tokens[i+1].Lit)
				}

				address := Assert(strconv.Atoi(tokens[i+3].Lit))

				p1 := address & 65280
				p2 := address & 255

				out = append(out,
					byte(slices.Index(INSTRUCTIONS, tokens[i].Lit)),
					byte(Assert(strconv.Atoi(tokens[i+1].Lit))),
					byte(p1),
					byte(p2),
				)
				i += 4
			case "adr": // 2 inputs(1:1:3 bytes)
				if len(tokens)-i < 2 {
					return []byte{}, Errf(tokens[i], "expected '%s [src/dst], [label/immediate]'", tokens[i].Lit)
				}

				if !tokens[i+1].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i], "expected register call, but found '%s' instead", tokens[i].Lit)
				} else if !tokens[i+2].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+2], "expected ',', but found '%s' instead", tokens[i+2].Lit)
				} else if !tokens[i+3].Istype(IDENT) && !tokens[i+3].Istype(IMMEDIATE) {
					return []byte{}, Errf(tokens[i+3], "expected immediate or label, but found '%s' instead", tokens[i+3].Lit)
				}

				jmpPos := 0
				if tokens[i+3].Istype(IDENT) {
					if pos, ok := labels[tokens[i+3].Lit]; ok {
						jmpPos = pos
					} else {
						return []byte{}, Errf(tokens[i+3], "label '%s' doesn't exist", tokens[i+3].Lit)
					}
				} else {
					jmpPos = Assert(strconv.Atoi(tokens[i+3].Lit))
				}

				p1 := jmpPos & 522240
				p2 := jmpPos & 65280
				p3 := jmpPos & 255

				out = append(out,
					byte(slices.Index(INSTRUCTIONS, tokens[i].Lit)),
					byte(Assert(strconv.Atoi(tokens[i+1].Lit))),
					byte(p1),
					byte(p2),
					byte(p3),
				)
				i += 4
			case "push", "pop": // 1 input(1:1 bytes)
				if len(tokens)-i < 2 {
					return []byte{}, Errf(tokens[i], "expected '%s [src/dst]'", tokens[i].Lit)
				}

				if !tokens[i+1].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i], "expected register call, but found '%s' instead", tokens[i].Lit)
				}

				out = append(out,
					byte(slices.Index(INSTRUCTIONS, tokens[i].Lit)),
					byte(Assert(strconv.Atoi(tokens[i+1].Lit))),
					0,
					0,
				)
				i += 2
			case "jeq", "jne", "jlt", "jgt": // 3 inputs(1:1:1:3 bytes)(yes I know that's for than 4 bytes shut up)
				if len(tokens)-i < 3 {
					return []byte{}, Errf(tokens[i], "expected '%s [src1] [src2] [label/immediate]'", tokens[i].Lit)
				}

				if !tokens[i+1].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i+1], "expected register call, but found '%s' instead", tokens[i+1].Lit)
				} else if !tokens[i+2].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+2], "expected ',', but found '%s' instead", tokens[i+2].Lit)
				} else if !tokens[i+3].Istype(REGCALL) {
					return []byte{}, Errf(tokens[i+3], "expected register call, but found '%s' instead", tokens[i+3].Lit)
				} else if !tokens[i+4].Istype(COMMA) {
					return []byte{}, Errf(tokens[i+4], "expected ',', but found '%s' instead", tokens[i+4].Lit)
				} else if !tokens[i+5].Istype(IDENT) && !tokens[i+5].Istype(IMMEDIATE) {
					return []byte{}, Errf(tokens[i+5], "expected label/immediate, but found '%s' instead", tokens[i+5].Lit)
				}

				jmpPos := 0
				if tokens[i+5].Istype(IDENT) {
					if pos, ok := labels[tokens[i+5].Lit]; ok {
						jmpPos = pos
					} else {
						return []byte{}, Errf(tokens[i+5], "label '%s' doesn't exist", tokens[i+5].Lit)
					}
				} else {
					jmpPos = Assert(strconv.Atoi(tokens[i+5].Lit))
				}

				p1 := jmpPos & 522240
				p2 := jmpPos & 65280
				p3 := jmpPos & 255

				out = append(out,
					byte(slices.Index(INSTRUCTIONS, tokens[i].Lit)),
					byte(Assert(strconv.Atoi(tokens[i+1].Lit))),
					byte(Assert(strconv.Atoi(tokens[i+3].Lit))),
					byte(p1),
					byte(p2),
					byte(p3),
				)
				i += 6
			case "jmp", "jal": // 1 input(1:3 bytes)
				if len(tokens)-i < 1 {
					return []byte{}, Errf(tokens[i], "expected '%s [label/immediate]'", tokens[i].Lit)
				}

				if !tokens[i+1].Istype(LABEL) && !tokens[i+1].Istype(IMMEDIATE) {
					return []byte{}, Errf(tokens[i], "expected label/immediate, but found '%s' instead", tokens[i].Lit)
				}

				jmpPos := 0
				if tokens[i+1].Istype(LABEL) {
					if pos, ok := labels[tokens[i+1].Lit]; ok {
						jmpPos = pos
					} else {
						return []byte{}, Errf(tokens[i+1], "label '%s' doesn't exist", tokens[i+1].Lit)
					}
				} else {
					jmpPos = Assert(strconv.Atoi(tokens[i+1].Lit))
				}

				p1 := jmpPos & 522240
				p2 := jmpPos & 65280
				p3 := jmpPos & 255

				out = append(out,
					byte(slices.Index(INSTRUCTIONS, tokens[i].Lit)),
					byte(p1),
					byte(p2),
					byte(p3),
				)
				i += 2
			default:
				fmt.Println(tokens[i].Tostr())
				return []byte{}, Errf(tokens[i], "invalid instruction '%s'", tokens[i].Lit)
			}
		} else if tokens[i].Istype(DIRECTIVE) {
			switch tokens[i].Lit {
			case "ascii", "asciz", "ascin":
				if i+1 >= len(tokens) {
					return []byte{}, Errf(tokens[i], "expected string")
				} else if !tokens[i+1].Istype(STRING) {
					return []byte{}, Errf(tokens[i+1], "expected string, but found '%s' instead", tokens[i+1].Lit)
				}

				//out = append(out, 255)
				//out = append(out, UnpackNumberIntoBytes(len(tokens[i+1].Lit))...)
				//out = append(out, 0)

				for _, b := range tokens[i+1].Lit {
					out = append(out, byte(b))
				}

				if tokens[i].Lit == "ascin" {
					out = append(out, '\n')
				}
				if tokens[i].Lit == "asciz" || tokens[i].Lit == "ascin" {
					out = append(out, 0)
				}
				i += 2
			default:
				return []byte{}, Errf(tokens[i], "invalid directive '%s'", tokens[i].Lit)
			}
		} else if tokens[i].Istype(NOOP) {
			out = append(out,
				byte(slices.Index(INSTRUCTIONS, "add")),
				0,
				0,
				0,
			)
			i += 1
		} else if tokens[i].Istoken(IDENT, "syscall") {
			out = append(out, 255, 255, 255, 255)
			i += 1
		} else if tokens[i].Istype(LABEL) {
			i += 1
			continue
		} else {
			fmt.Println(tokens[i].Tostr())
			return []byte{}, Errf(tokens[i], "invalid instruction '%s'", tokens[i].Lit)
		}
	}

	return out, nil
}
