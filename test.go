package main

/*
#cgo CFLAGS: -I.

#cgo LDFLAGS: -L./cmake-build-debug -ltest

#include "test.h"
*/
import "C" // 切勿换行再写这个

import (
    "fmt"
)

func main() {

    // function
    fmt.Println("calling cpp function...")
	fmt.Println(C.sum(C.int(1), C.int(1)))

	// class function wrapper
    fmt.Println("calling cpp class function by wrapper...")
	fmt.Println(C.p_sum(100, 100))
}

