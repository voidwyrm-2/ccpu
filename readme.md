# CCPU
A CPU written in C<br>
It's assembler was going to be written in C, but C is hard so then I started writing it in Rust, but the borrow checker wants me dead so now it's written in Go<br>
~~Or maybe I'll rewrite it in Zig~~

See [architecture.md](./architecture.md) for, you know, the CPU architecture<br>
See [call_codes.md](./call_codes.md) for the CPU call codes for syscalls

The `casm/` folder is the assembler<br>
`cpu.c` and `lib.h` is the CPU itself
The `asm` folder contains all the assembly I wrote for the CPU(`.asm` files contain plaintext source code, `.bin` files contain the assembled code)

## Installation
**NOTE:** These instructions are meant for Unix systems<br>
**Installing Casm**
* Check if you have Go with `go version`
    > If not, either download it with Homebrew(`brew update && brew install go`) or from [Go's website](https://go.dev)
* Run `go install github.com/voidwyrm-2/ccpu@latest`

**Installing The CCPU**<br>
The CCPU must be built from source because I'm lazy and it's not that hard<br>
* Get a C compiler(I reccomend either gcc or clang)
    > If you already have a C compiler that you like, use that one, I only care that you have one
* Download this repo
* Run compile `cpu.c` with `-o ccpu`; use `ccpu -h` for help with running files