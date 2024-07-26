#include <stdio.h>
#include <stdlib.h>

#define sizeofArray(arr, type) (unsigned long int)sizeof(arr) / sizeof(type)

/*
#define printArray(arr, type)                        \
    for (int i = 0; i < sizeofArray(arr, type); i++) \
    _Generic(arr[i],                                 \
        int: printf("%d\n", arr[i]),                 \
        char: printf("%c\n", arr[i]),                \
        char*: printf("%s\n", arr[i]))
*/

int *readfile(char *fname) {
    int c;

    FILE *file;

    file = fopen(fname, "r");

    int *out;

    if (file) {
        out = (int *)malloc(sizeof(file));
        int i = 0;
        while ((c = getc(file)) != EOF) {
            out[i] = c;
            i++;
        }
        fclose(file);
    } else {
        return 0;
    }

    return out;
}

typedef struct IntDLL {
    int value;
    void *prev;
    void *next;
    int isAssigned;
    int hasPrev;
    int hasNext;
} IntDLL;

IntDLL NewIntDLL() {
    IntDLL s = {.hasNext = 0, .hasPrev = 0, .value = 0, .isAssigned = 0};
    return s;
}

void AddToIntDLL(IntDLL *s, int value) {
    if (s->hasNext == 0) {
        if (s->isAssigned == 0) {
        }
        IntDLL *ns = malloc(sizeof(IntDLL));
        *ns = NewIntDLL();
        ns->value = value;
        ns->isAssigned = 1;
        ns->hasPrev = 1;
        ns->prev = s;

        s->next = ns;
        s->hasNext = 1;
    } else {
        AddToIntDLL(s->next, value);
    }
}

int RemoveFromIntDLL(IntDLL *s) {
    int out = 0;
    if (s->hasNext == 0) {
        if (s->hasPrev == 0 && s->isAssigned == 1) {
            s->isAssigned = 0;
            out = s->value;
            s->value = 0;
            return out;
        }
        return 0;
    }
    IntDLL *temp = malloc(sizeof(IntDLL));
    temp = s->next;
    if (temp->hasNext == 0) {
        out = temp->value;
        free(temp);
        s->next = malloc(sizeof(IntDLL));
        s->hasNext = 0;
        free(s->next);
        return out;
    } else {
        RemoveFromIntDLL(s->next);
    }
    return 0;
}

int IndexIntDLL(IntDLL *s, int index) {
    if (s->hasNext == 0 && index != 0) return 0;

    if (index == 0) return s->value;

    return IndexIntDLL(s->next, index - 1);
}

unsigned int LengthOfIntDLL(IntDLL *s) {
    int count = s->isAssigned == 1 ? 1 : 0;

    IntDLL *cdll = s;

    while (1) {
        if (cdll->hasNext == 1) {
            count++;
            cdll = cdll->next;
            continue;
        }
        break;
    }

    return count;
}

typedef struct IntStack {
    IntDLL stack;
    unsigned int size;
    int isEmpty;
} IntStack;

IntStack NewIntStack() {
    IntDLL dll = NewIntDLL();
    IntStack stack = {.stack = dll, .size = 0, .isEmpty = 1};
    return stack;
}

void PushIntStack(IntStack *s, int value) {
    s->size++;
    AddToIntDLL(&(s->stack), value);
}

int PopIntStack(IntStack *s) {
    if (s->size <= 0) {
        return -1;
    }
    s->size--;
    return RemoveFromIntDLL(&(s->stack));
}