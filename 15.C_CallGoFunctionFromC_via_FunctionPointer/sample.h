#ifndef SAMPLE_H
#define SAMPLE_H

#include <stdlib.h>

void func_with_callback(int, void (*callback)(int, void *, size_t));

#endif /* SAMPLE_H */
