package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * 003. Write a Go program that keeps reading integers until it gets the word STOP as input
 * go run .\003.go
 */
func main() {
	var fd *os.File
	fd = os.Stdin // 直接将标准输入流打到fd中
	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		tmp := scanner.Text()
		if tmp == "STOP" {
			fmt.Println("exit loop")
			break
		}
		tmpNum, _ := strconv.ParseInt(tmp, 10, 64)
		if tmpNum != 0 {
			fmt.Println("int: ", tmpNum)
		}
	}
}
