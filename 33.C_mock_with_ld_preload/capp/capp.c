#include <stdio.h>
#include <stdlib.h>

extern int c_func(int x, int y);

int main(void) {
    int x = 10;
    int y = 20;
    int z = c_func(x, y);
    
    printf("c_func: %d\n", z);

    return EXIT_SUCCESS;
}