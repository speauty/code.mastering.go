package main

import "fmt"

func a() {
	fmt.Println("inside a")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("recover inside a")
		}
	}()
	fmt.Println("call b")
	b()
	fmt.Println("b exited")
	fmt.Println("a exiting")
}

func b() {
	fmt.Println("inside b")
	panic("panic in b")
	fmt.Println("b exiting")
}

func main() {
	a()
	fmt.Println("main end")
}
