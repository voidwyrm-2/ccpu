typedef enum Tokentypes {
    T_ILLEGAL,
    T_INSTRUCTION,
    T_REGCALL,
    T_IMMEDATE,
    T_LABEL,
    T_COMMA,
    T_NEWLINE,
    T_DIRECTIVE,
    T_NOOP,
    T_HALT,
    T_CALL,
    T_RET
} Tokentype;

typedef struct Token {
    int type;
    char* lit;
    unsigned int start;
    unsigned int end;
    unsigned int ln;
} Token;

int TokenLength(Token* tok) { return (tok->end - tok->start) + 1; }