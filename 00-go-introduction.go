package main // 声明包

import (
	"fmt"     // 标准库包
	"strings" // 另一个标准库包
)

func main() { // main 函数：程序的入口点

	message := "Hello, World golang!"     // 声明变量
	fmt.Println(message)                  // 打印变量
	fmt.Println(strings.ToUpper(message)) // 调用函数
}
