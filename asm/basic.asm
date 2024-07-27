_terminate:
    mov 0, x0 ; return 0
    mov 0, x1 ; tell the kernel we want to terminate the program
    syscall ; yell at the kernel