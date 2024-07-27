_main:
    mov 10, r0
    store r0, bufn

_read:
    mov 2, x0 ; tell the 'kernel' we want to read from stdin
    adr x1, buf ; tell the 'kernel' where to write to
    mov 10, x2 ; tell the 'kernel' how how long our string is so it knows how much to read
    syscall ; yell at the 'kernel' to get off its lazy butt

_print:
    mov 1, x0 ; tell the 'kernel' we want to write to the screen
    adr x1, buf ; tell the 'kernel' where to read
    mov 11, x2 ; tell the 'kernel' how how long our string is so it knows how much to read
    syscall ; yell at the 'kernel' to get off its lazy butt

_terminate:
    mov 0, x0 ; tell the 'kernel' we want to terminate the program
    mov 0, x1 ; return 0, in effect
    syscall ; yell at the 'kernel' again

buf: .space 10 ; make a buffer for our text
bufn: .space 1