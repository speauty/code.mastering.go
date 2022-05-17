package main

import "fmt"

func GetPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	// 生命周期
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25
	ptrI := &i
	ptrJ := &j
	fmt.Println("ptrI memory: ", ptrI)
	fmt.Println("ptrJ memory: ", ptrJ)
	fmt.Println("ptrI value: ", *ptrI)
	fmt.Println("ptrJ value: ", *ptrJ)

	*ptrI = 123456
	*ptrI--
	fmt.Println("i value: ", i)

	GetPointer(ptrJ)
	fmt.Println("j value: ", j)

	// 居然没被释放, 也是怪了, 毕竟returnPointer.v是局部变量, 三色标记未理解透彻
	k := returnPointer(12)
	fmt.Println("k memory: ", k) // k memory:  0xc000022140
	fmt.Println("k value: ", *k) // k value:  144
}
