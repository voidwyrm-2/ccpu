_main:
    mov 20, r0
    mov 22, r1
    add r0, r1, r2

_terminate:
    mov 0, x0 ; return 0
    mov 0, x1 ; tell the kernel we want to terminate the program
    syscall ; yell at the kernel