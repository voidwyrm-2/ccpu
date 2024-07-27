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
    IntStack stack = NewIntStack();
    IntStack addrStack = NewIntStack();

    int pc = 0;
    while (1) {
        // printf("r0: %d, r1: %d, r2: %d, r3: %d\n", registers[8],
        // registers[9],
        //        registers[10], registers[11]);
        switch (code[pc]) {
            // mov
            case 0:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0 ? 0 : (code[pc + 1] << 8) + code[pc + 2];
                break;

            // add
            case 1:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] + registers[code[pc + 2]];
                break;

            // sub
            case 2:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] - registers[code[pc + 2]];
                break;

            // and
            case 3:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] & registers[code[pc + 2]];
                break;

            // or
            case 4:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] | registers[code[pc + 2]];
                break;

            // xor
            case 5:
                registers[code[pc + 3]] =
                    code[pc + 3] == 0
                        ? 0
                        : registers[code[pc + 1]] ^ registers[code[pc + 2]];
                break;

            // store
            case 6:
                code[(code[pc + 2] << 16) + (code[pc + 3] << 8) +
                     code[pc + 4]] = registers[code[pc + 1]];
                pc += 5;
                continue;
                break;
            // load
            case 7:
                registers[code[pc + 1]] =
                    code[pc + 1] == 0
                        ? 0
                        : code[(code[pc + 2] << 16) + (code[pc + 3] << 8) +
                               code[pc + 4]];
                pc += 5;
                continue;
                break;
            // adr
            case 8:
                registers[code[pc + 1]] =
                    code[pc + 1] == 0 ? 0
                                      : (code[pc + 2] << 16) +
                                            (code[pc + 3] << 8) + code[pc + 4];
                pc += 5;
                continue;
                break;

            // push
            case 9:
                PushIntStack(&stack, registers[code[pc + 1]]);
                break;
            // pop
            case 10:
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
                break;

            // inc
            case 11:
                registers[code[pc + 1]]++;
                break;
            // dec
            case 12:
                registers[code[pc + 1]]--;
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
                    case 2:
                        printf("");  // WHY DOES THE COMPILER YELL AT ME IF I
                                     // DON'T INCLUDE THIS
                        str = malloc(sizeof(char) * (registers[3] + 1));
                        i = malloc(sizeof(int));
                        *i = 0;

                        fgets(str, sizeof(char) * (registers[3] + 1), stdin);

                        while (*i < registers[3]) {
                            code[pc + difference(pc, registers[2]) + (*i)] =
                                str[*i];
                            (*i)++;
                        }
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