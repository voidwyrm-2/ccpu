#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "lib.h"

const int INSTRUCTION_SIZE = 4;

int difference(int a, int b) {
    if (a > b) return a - b;
    return b - a;
}

void cpu(int* code) {
    int* registers = malloc(sizeof(int) * 256);
    int* memory = malloc(sizeof(int) * 1024);
    IntStack stack = NewIntStack();
    IntStack addrStack = NewIntStack();

    int pc = 0;
    while (1) {
        // printf("b1: %d, b2: %d, b3: %d, b4: %d\n", code[pc], code[pc + 1],
        //        code[pc + 2], code[pc + 3]);
        switch (code[pc]) {
            // add
            case 0:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] + registers[code[pc + 2]];
                break;
            // addi
            case 1:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0 ? 0
                                      : code[pc + 1] + registers[code[pc + 2]];
                break;

            // sub
            case 2:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] - registers[code[pc + 2]];
                break;
            // subi
            case 3:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0 ? 0
                                      : code[pc + 1] - registers[code[pc + 2]];
                break;

            // and
            case 4:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] & registers[code[pc + 2]];
                break;
            // andi
            case 5:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0 ? 0
                                      : code[pc + 1] & registers[code[pc + 2]];
                break;

            // or
            case 6:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] | registers[code[pc + 2]];
                break;
            // ori
            case 7:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0 ? 0
                                      : code[pc + 1] | registers[code[pc + 2]];
                break;

            // xor
            case 8:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] ^ registers[code[pc + 2]];
                break;
            // xori
            case 9:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0 ? 0
                                      : code[pc + 1] ^ registers[code[pc + 2]];
                break;

            // store
            case 10:
                memory[(code[pc + 2] << 8) + code[pc + 3]] =
                    registers[code[pc + 1]];
                break;
            // load
            case 11:
                registers[code[pc + 1]] =
                    code[pc + 1] == 0
                        ? 0
                        : memory[(code[pc + 2] << 8) + code[pc + 3]];
                break;
            // adr
            case 12:
                registers[code[pc + 1]] =
                    code[pc + 1] == 0 ? 0
                                      : (code[pc + 2] << 16) +
                                            (code[pc + 3] << 8) + code[pc + 4];
                pc += 5;
                continue;
                break;

            // push
            case 13:
                PushIntStack(&stack, registers[code[pc + 1]]);
                pc += 2;
                continue;
                break;
            // pop
            case 14:
                printf("");  // WHY DOES THE COMPILER YELL AT ME IF I
                             // DON'T INCLUDE THIS
                int* temp = malloc(sizeof(int));
                *temp = PopIntStack(&stack);
                if (*temp == -1) {
                    printf("(%d) cannot pop from stack as it's empty\n", pc);
                    return;
                }
                registers[code[pc + 1]] = code[pc + 1] == 0 ? 0 : *temp;
                free(temp);
                pc += 2;
                continue;
                break;

            // syscall/ecall/svc 0
            case 255:
                switch (registers[1]) {
                    case 0:
                        return;
                    case 1:
                        printf("");  // WHY DOES THE COMPILER YELL AT ME IF I
                                     // DON'T INCLUDE THIS
                        char* str = malloc(sizeof(char) * registers[3]);
                        int* i = malloc(sizeof(int));
                        *i = 0;

                        while (*i < registers[3]) {
                            str[*i] =
                                code[pc + difference(pc, registers[2]) + (*i)];
                            (*i)++;
                        }

                        printf("%s", str);
                        free(str);
                        free(i);
                        break;
                    default:
                        printf("invalid call code %d\n", registers[1]);
                }
                break;
            default:
                printf("invalid opcode %d\n", code[pc]);
                return;
        }
        pc += INSTRUCTION_SIZE;
    }

    // printf("r0: %d, r1: %d, r2: %d, r3: %d\n", registers[0], registers[1],
    //        registers[2], registers[3]);
}

int main(int argc, char* argv[]) {
    // int bytes[] = {0, 1, 0, 20, 0, 0, 0, 2, 0, 22, 0, 0, 1, 3, 2, 1, 0, 0};
    if (argc != 2) {
        printf("expected 'ccpu [file]'\n");
        return 0;
    }

    if (strcmp(argv[1], "-h") == 0) {
        printf("usage: ccpu [file]");
        return 0;
    }

    int* bytes = readfile(argv[1]);
    cpu(bytes);
    return 0;
}

// old 255 opcode, it's now syscall
/*
// opcode 255 is a flag that jumps over the amount of bytes
// specified by the byte immediately after the opcode
case 255:
    int* i = malloc(sizeof(int));
    int* shift = malloc(sizeof(int));
    int* acc = malloc(sizeof(int));
    *i = pc;
    shift = 0;
    acc = 0;
    while (code[*i] != 0) {
        acc += code[*i];
    }
    pc += *acc;
    break;
*/