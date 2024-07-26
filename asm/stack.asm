_main:

_terminate:
    addi 0, zero, x0 ; return 0
    addi 0, zero, x1 ; tell the kernel we want to terminate the program
    syscall ; yell at the kernel again