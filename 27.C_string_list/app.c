#include <stdio.h>
#include <stdlib.h>

extern void gofunc(int argc, char** argv);

int main(int argc, char **argv) {
    gofunc(argc, argv);
    return EXIT_SUCCESS;
}