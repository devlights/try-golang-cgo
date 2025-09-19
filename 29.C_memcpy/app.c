#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern void mymemcpy(char *dst, char *src, size_t n);

#define STR ("helloworld")

int main(void)
{
    size_t n = sizeof(STR);
    char   dst[n];

    memset(dst, 0, sizeof(dst));

    printf("[before] %s:%s\n", STR, dst);
    {
        mymemcpy(dst, STR, n);
    }
    printf("[after ] %s:%s\n", STR, dst);

    return EXIT_SUCCESS;
}