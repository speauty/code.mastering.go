package main

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc: ", mem.Alloc)
	fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("mem.NumGC: ", mem.NumGC)
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<")
}

/**
 * GODEBUG=gctrace=1 go run g_coll.go
 * gc 1 @0.004s 2%: 0.028+0.73+0.039 ms clock, 0.34+0.29/0.94/0+0.47 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 12 P
 */
func main() {
	var mem runtime.MemStats
	printStats(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("alloc failed")
		}
	}
	printStats(mem)
}
