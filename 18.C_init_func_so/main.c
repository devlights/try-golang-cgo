#include <stdio.h>
#include <stdlib.h>

extern void my_printf(const char *);
extern void gofunc1(void);

int main(void) {
    my_printf("hello world");
    gofunc1();

    return EXIT_SUCCESS;
}