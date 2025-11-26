package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ITEM_COUNT (3)
#define KEY_LENGTH (32)

typedef struct Item {
	char key[KEY_LENGTH];
	int value;
} Item;

Item item_list[ITEM_COUNT];

// 初期化関数
void init_item_list(void) {
	for (size_t i = 0; i < ITEM_COUNT; i++) {
		Item *item = &item_list[i];

		char buf[KEY_LENGTH];
		int len = snprintf(buf, KEY_LENGTH, "item-%02zu", i);
		strncpy(item->key, buf, len);
		item->key[len] = '\0';

		item->value = (int)i;
	}
}

// 表示関数
void disp_item_list(void) {
	for (size_t i = 0; i < ITEM_COUNT; i++) {
		Item *item = &item_list[i];
		printf("[C ] key=%s, value=%d\n", item->key, item->value);
	}
}

// アクセサ関数
Item *get_item(size_t index) {
	return &item_list[index];
}

// 設定関数
void mod_item(Item *item, const char *key, int value) {
	strcpy(item->key, key);
	item->value = value;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	ITEM_COUNT = 3
)

func main() {
	//
	// cgoで配列を扱う場合は、できれば要素を取得する アクセサ関数 を用意したほうが安全
	//

	C.init_item_list()

	display("init")
	modify()
	display("modify")
}

func display(msg string) {
	fmt.Printf("======[%s]======\n", msg)
	C.disp_item_list()

	for i := range ITEM_COUNT {
		item := C.get_item(C.size_t(i))
		fmt.Printf("[Go] key=%s, value=%d\n", C.GoString(&item.key[0]), int(item.value))
	}
	fmt.Println("==================")
}

func modify() {
	for i := range ITEM_COUNT {
		func() {
			item := C.get_item(C.size_t(i))

			key := fmt.Sprintf("ITEM-%02d", i)
			keyPtr := C.CString(key)
			defer C.free(unsafe.Pointer(keyPtr))

			C.mod_item(item, keyPtr, C.int(0xFF-i))
		}()
	}
}
