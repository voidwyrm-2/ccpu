#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../lib.h"
#include "tokens.h"

#define TOKEN_LINE_SIZE sizeof(Token) * 7

Token* Lex(char* text) {
    Token* tokens = (Token*)malloc(TOKEN_LINE_SIZE * 5);

    return tokens;
}

Token CollectString(int* i, char** text) {
    char* acc = malloc(sizeof(char) * 20);

    while (i < sizeofArray(text, char*)) {
    }
}

char* usage = "usage: casm [path] [-h|--help]";
char* help = "usage: casm [path] [-h|--help]\n\n   The CAsm assembler";

int main(int argc, char** argv) {
    // printf("%d\n", argc);
    if (argc < 2) {
        printf("%s\n", usage);
        return 0;
    }

    for (int i = 0; i < argc; i++) {
        if (strlen(argv[i]) < 1 || i == 0) continue;
        // printf("strcmp for '%s': %d\n", argv[i], strcmp(argv[i], "-h"));

        if (strcmp(argv[i], "-h") == 0 || strcmp(argv[i], "--help") == 0) {
            printf("%s\n", usage);
            return 0;
        } else if (argv[i][0] == '-') {
            printf("unknown flag '%s'", argv[i]);
            return 0;
        }
        /*
        else if (strcmp(argv[i], "-c")) {
            printf("%d", argc);
            return 0;
        }
        */
    }

    char* content = readfile(argv[1]);
    if (strcmp(content, "\0") == 0) {
        printf("file '%s' does not exist\n", argv[1]);
        return 0;
    }

    printf("%s\n", content);

    return 0;
}