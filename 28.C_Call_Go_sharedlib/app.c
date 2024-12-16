#include <stdio.h>
#include <stdlib.h>

extern int add(int, int);

int main(void) {
    int x = 10;
    int y = 20;
    int z = add(x, y);

    printf("%d\n", z);
    
    return EXIT_SUCCESS;
}