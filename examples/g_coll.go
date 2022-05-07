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
 * linux可视信息
 * gc 1 @0.004s 2%: 0.028+0.73+0.039 ms clock, 0.34+0.29/0.94/0+0.47 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 12 P
 * gc 执行次数 程序耗时 垃圾回收时间占比 STW清扫时间+并发标记扫描时间+STW标记时间垃圾回收占CPU时间 GC前堆大小+GC后堆大小+存活堆大小 整体堆大小 整体栈大小 总大小 使用处理器数量
 *
 * windows可视信息
 * gc 1 @0.005s 0%: 0+0.53+0 ms clock, 0+0/0.53/0.53+0 ms cpu, 47->48->0 MB, 48 MB goal, 12 P
 * gc 执行次数 程序耗时 垃圾回收时间占比 STW清扫时间+并发标记扫描时间+STW标记时间垃圾回收占CPU时间 GC前堆大小+GC后堆大小+存活堆大小 整体堆大小 使用处理器数量
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
