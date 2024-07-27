_main:
    mov 65, r0
    push r0
    inc r0
    inc r0
    pop r1

    store r0, buf1
    store r1, buf2
    mov 10, r2
    store r2, buf3

_print:
    mov 1, x0 ; tell the 'kernel' we want to write to the screen
    adr x1, buf1 ; tell the 'kernel' where to read
    mov 3, x2 ; tell the 'kernel' how how long our string is so it knows how much to read
    syscall ; yell at the 'kernel' to get off its lazy butt

_terminate:
    mov 0, x0 ; return 0
    mov 0, x1 ; tell the kernel we want to terminate the program
    syscall ; yell at the kernel

; yeah my store instruction doesn't have the fancy [dst, offset] syntax
buf1: .space 1 ; r1
buf2: .space 1 ; r0
buf3: .space 1 ; newline
buf4: .space 1 ; 0