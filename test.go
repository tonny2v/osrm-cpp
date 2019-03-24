package main

/*
#cgo CFLAGS: -I.

#cgo LDFLAGS: -L./cmake-build-debug -ltest

#include "test.h"
*/
import "C" // 切勿换行再写这个

import "fmt"

func main() {
	fmt.Println(C.add(C.int(200), C.int(1)))
}

