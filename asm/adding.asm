_main:
    addi 20, zero, r0
    addi 22, zero, r1
    add r0, r1, r2

_terminate:
    addi 0, zero, x0 ; return 0
    addi 0, zero, x1 ; tell the kernel we want to terminate the program
    syscall ; yell at the kernel again