package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/**
 * 001. Write a Go program that finds the sum of all of its numeric command-line arguments
 * go run .\001.go 1 2 4 5 6 7
 * sum:  25
 *
 * numeric? > 0?
 */
func main() {
	args := os.Args
	argsLen := len(args)
	if argsLen < 2 {
		log.Fatalln("the args not found.")
		return
	}
	sum := int64(0)
	for i := 1; i < argsLen; i++ {
		tmp, _ := strconv.ParseInt(args[i], 10, 8)
		sum += tmp
	}
	fmt.Println("sum: ", sum)
}
