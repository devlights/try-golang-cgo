#include <stdio.h>
#include <string.h>
#include "sample.h"

// 引数にコールバック用の関数ポインタを要求する関数
void func_with_callback(int id, void (*callback)(int id, void *data, size_t length)) {
    char data[] = "hello world";
    size_t szData = strnlen(data, 12);

    printf("[C ][func_with_callback] called\n");
    (*callback)(id, data, szData);
}
