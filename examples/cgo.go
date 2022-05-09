package main

//#include <stdio.h>
//void callC() {
// printf("calling C code!\n");
//}
import "C"

import "fmt"

/**
 * 在注释中写C代码, 和之后的import "C"不能存在空行间隔, 否则会出现错误 "could not determine kind of name for C.callC"
 */
func main() {
	fmt.Println("A Go statement")
	C.callC()
	fmt.Println("Another Go statement")

}
