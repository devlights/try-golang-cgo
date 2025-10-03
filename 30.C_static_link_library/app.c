#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern int cgo_func(int x, int y, int *z);

int main(void) {
    int x = 0x00FF;
    int y = 0xFF00;
    int z;

    cgo_func(x, y, &z);

    printf("%d\n", z);
    
    return EXIT_SUCCESS;
}