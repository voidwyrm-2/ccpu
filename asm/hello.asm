_print:
    addi 1, zero, x0 ; tell the 'kernel' we want to write to the screen
    adr x1, helloworld ; tell the 'kernel' where to read
    addi 15, zero, x2 ; tell the 'kernel' how how long our string is so it knows how much to read
    ; the length of our string is 14 + the zero-terminator
    syscall ; yell at the 'kernel' to get off its lazy butt

_terminate:
    addi 0, zero, x0 ; tell the 'kernel' we want to terminate the program
    addi 0, zero, x1 ; return 0, in effect
    syscall ; yell at the 'kernel' again

helloworld: .asciz "Hello, Catdog!" ; make a zero-terminated string