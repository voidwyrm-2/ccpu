_print:
    mov 1, x0 ; tell the 'kernel' we want to write to the screen
    adr x1, helloworld ; tell the 'kernel' where to read
    mov 15, x2 ; tell the 'kernel' how how long our string is so it knows how much to read
    ; the length of our string is 14 + the zero-terminator
    syscall ; yell at the 'kernel' to get off its lazy butt

_terminate:
    mov 0, x0 ; tell the 'kernel' we want to terminate the program
    mov 0, x1 ; return 0, in effect
    syscall ; yell at the 'kernel' again

helloworld: .ascin "Hello, Catdog!" ; make a newline- and zero- terminated string