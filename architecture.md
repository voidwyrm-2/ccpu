# Architecture
This is the architecture documentation for the CCPU<br>
Since this is a quick and dirty project it's on the simple side

Each instructions is 4 bytes long, except jump instructions, which are 6 bytes long<br>
The first byte is the opcode, rest are split 1:1:1, 1:2, or aren't split(depending on the instruction)
<!--0 000000-->

## Instruction Set
The opcode is zero-indexed and is just the previous opcode + 1 unless otherwise stated
Instruction | Arguments | Opcode
---- | ---- | ----
Add | [srcReg1], [srcReg2], [dstReg] |
Addi | [imm], [srcReg], [dstReg] |
Sub | [srcReg1], [srcReg2], [dstReg] |
Subi | [imm], [srcReg], [dstReg] |
And | [srcReg1], [srcReg2], [dstReg] |
Andi | [imm], [srcReg], [dstReg] |
Or | [srcReg1], [srcReg2], [dstReg] |
Ori | [imm], [srcReg], [dstReg] |
Xor | [srcReg1], [srcReg2], [dstReg] |
Xori | [imm], [srcReg], [dstReg] |
Store | [srcReg], [memoryAddress] |
Load | [dstReg], [memoryAddress] |
Push | [srcReg] |
Pop | [dstReg] |
Jeq | [srcReg1], [srcReg2], [label/immediate] |
Jne | [srcReg1], [srcReg2], [label/immediate] |
Jlt | [srcReg1], [srcReg2], [label/immediate] |
Jgt | [srcReg1], [srcReg2], [label/immediate] |
Jmp | [label/immediate] |
Jal | [label/immediate] |
Adr | [dstReg], [label/immediate]

**Pseudo-instructions**
* `noop`: the assembler converts it to `add 0, 0, 0`

## Conventions
The CCPU has 8 reserved registers:
* The zero register, which always contains zero and is accessed with `zero`
* The call code register(or register `x0`), which is read by the kernel when `syscall` is used
* Then we have the 6 argument registers(`x1`-`x7`), which are read by the kernel when `syscall` is used

Register calls prefixed with `X` or `X`(e.g. `add x0, x1, x2`) skips the zero register, so `x0` is actually register 1<br>
Register calls prefixed with `r` or `R`(e.g. `add r0, r1, r2`) skips the 8 reserved registers, so `r0` is actually register 8