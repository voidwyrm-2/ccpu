# Architecture
This is the architecture documentation for the CCPU<br>
Since this is a quick and dirty project it's on the simple side

Each instructions is 4 bytes long, except jump instructions, which are 6 bytes long<br>
The first byte is the opcode, rest are split 1:1:1, 1:2, or aren't split(depending on the instruction)
<!--0 000000-->

## Instruction Set
The opcode is zero-indexed and is just the previous opcode + 1 unless otherwise stated
Instruction | Arguments | Notes | Opcode
---- | ---- | ---- | ----
Mov | [immediate], [dstReg] |
Add | [srcReg1], [srcReg2], [dstReg] |
Sub | [srcReg1], [srcReg2], [dstReg] |
And | [srcReg1], [srcReg2], [dstReg] |
Or | [srcReg1], [srcReg2], [dstReg] |
Xor | [srcReg1], [srcReg2], [dstReg] |
Store | [srcReg], [memoryAddress] |
Load | [dstReg], [memoryAddress] |
Adr | [dstReg], [memoryAddress] |
Push | [srcReg] |
Pop | [dstReg] |
Inc | [dstReg] |
Dec | [dstReg] |
Jeq | [srcReg1], [srcReg2], [label/immediate] |
Jne | [srcReg1], [srcReg2], [label/immediate] |
Jlt | [srcReg1], [srcReg2], [label/immediate] |
Jgt | [srcReg1], [srcReg2], [label/immediate] |
Jmp | [label/immediate] |
Jal | [label/immediate] |

**Pseudo-instructions**
* `noop`: the assembler converts it to `add 0, 0, 0`

**Directives**
* `ascii`, `asciz`, and `ascin`: they take a string as input and create an ASCII string, a zero-terminated ASCII string, and a newline- and zero-terminated string, respectively
* `byte`: takes an immediate as input and puts it in the final compiled file
* `space`: takes an immediate as input and generates that amount zero bytes 

## Conventions
The CCPU has 8 reserved registers:
* The zero register, which always contains zero and is accessed with `zero`
* The call code register(or register `x0`), which is read by the kernel when `syscall` is used
* And then there's the 6 argument registers(`x1`-`x7`), which are read by the kernel when `syscall` is used

Register calls prefixed with `X` or `X`(e.g. `add x0, x1, x2`) skips the zero register, so `x0` is actually register 1<br>
Register calls prefixed with `r` or `R`(e.g. `add r0, r1, r2`) skips the 8 reserved registers, so `r0` is actually register 8