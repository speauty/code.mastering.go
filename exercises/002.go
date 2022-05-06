package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/**
 * 002. Write a Go program that finds the average value of all of its float command-line arguments
 * go run .\002.go 4 6 9 -1
 * sum:  18  avg:  4.5
 */
func main() {
	args := os.Args
	argsLen := len(args)
	if argsLen < 2 {
		log.Fatalln("the args not found.")
		return
	}
	sum := float64(0)
	validCount := 0
	for i := 1; i < argsLen; i++ {
		tmp, _ := strconv.ParseFloat(args[i], 64)
		if tmp != 0 {
			validCount++
		}
		sum += tmp
	}
	fmt.Println("sum: ", sum, " avg: ", sum/float64(validCount))
}
