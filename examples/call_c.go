package main

// #cgo CFLAGS: -I ${SRCDIR}/clib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdio.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"unsafe"
)

/**
 * 生成目标文件 gcc -c clib/*.c
 * 生成静态库 ar rs callC.a *.o
 * 还是不怎么好使, 看起来
 */
func main() {
	fmt.Println("Going to call a C function")
	C.cHello()

	msg := C.CString("this is a additional msg fro cgo module")
	defer C.free(unsafe.Pointer(msg))
	C.printMsg(msg)
}
