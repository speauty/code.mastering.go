package main

import "C"

import "fmt"

//export PrintMsg
func PrintMsg() {
	fmt.Println("A Go func")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

// 构建C共享库 go build -o cshare/usedByC.o -buildmode=c-shared used_by_c.go
func main() {

}
