package main

import (
	"fmt"
	"unsafe"
)

func main() { // 没C好用, 建议直接使用CGO
	array := [...]int{0, 1, -2, 3, 4}
	fmt.Println("int offset: ", unsafe.Sizeof(array[0]))
	ptr := &array[0]
	fmt.Println(*ptr, "  ", ptr)
	// 直接进行指针偏移
	memoryAddress := uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		ptr = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Println(*ptr, "  ", ptr)
		memoryAddress = uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(array[0])
	}
	fmt.Println()
	ptr = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("one more: ", *ptr, "  ", ptr)
	fmt.Println()

}
