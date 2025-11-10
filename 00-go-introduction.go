package main // 声明包

import (
	"fmt"     // 标准库包
	"strings" // 另一个标准库包
	"sync"    // 用于 WaitGroup
)

func main() { // main 函数：程序的入口点
	var wg sync.WaitGroup

	wg.Add(6)

	go func() {
		defer wg.Done()
		// 这会并发运行
		fmt.Println("在 goroutine 中运行")
	}()
	go func() {
		defer wg.Done()
		// 这会并发运行
		fmt.Println("在 goroutine 中运行")
	}()
	go func() {
		defer wg.Done()
		// 这会并发运行
		fmt.Println("在 goroutine 中运行")
	}()
	go func() {
		defer wg.Done()
		// 这会并发运行
		fmt.Println("在 goroutine 中运行")
	}()
	go func() {
		defer wg.Done()
		// 这会并发运行
		fmt.Println("在 goroutine 中运行")
	}()
	go func() {
		defer wg.Done()
		// 这会并发运行
		fmt.Println("在 goroutine 中运行")
	}()

	wg.Wait()

	message := "Hello, World golang!"     // 声明变量
	fmt.Println(message)                  // 打印变量
	fmt.Println(strings.ToUpper(message)) // 调用函数
}
