// // 1.Go: If/Else
// package main

// import "fmt"

// func main() {
// 	temperature := 25

// 	if temperature > 30 {
// 		fmt.Println("天气很热！")
// 	} else if temperature > 20 {
// 		fmt.Println("天气温暖。")
// 	} else {
// 		fmt.Println("天气寒冷。")
// 	}

// 	// Go 需要显式的布尔表达式
// 	if temperature > 0 {
// 		fmt.Println("温度是正数")
// 	}
// }

// // 2.Go: For 循环
// package main

// import "fmt"

// func main() {
// 	// 传统 for 循环（Go 的唯一循环结构）
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}

// 	// 基于 range 的 for 循环（类似 for...of）
// 	arr := []int{1, 2, 3, 4, 5}
// 	fmt.Println("arr:", arr)
// 	for i, val := range arr {
// 		fmt.Printf("索引 %d: %d\n", i, val)
// 	}

// 	// 遍历 map（类似 for...in）
// 	obj := map[string]int{"a": 1, "b": 2, "c": 3}
// 	for key, val := range obj {
// 		fmt.Printf("%s: %d\n", key, val)
// 	}

// 	// 遍历切片（只要索引）
// 	for i := range arr {
// 		fmt.Printf("索引: %d\n", i)
// 	}

// 	// 遍历切片（只要值）
// 	for _, val := range arr {
// 		fmt.Printf("值: %d\n", val)
// 	}
// }

// // 3.Go: 函数定义
// package main

// import "fmt"

// // 函数声明
// func greet(name string) string {
// 	return fmt.Sprintf("Hello, %s!", name)
// }

// // 带多个返回值的函数
// // divide 执行除法运算
// // 参数: a - 被除数, b - 除数
// // 返回: 商(float64) 和 错误(error)
// // 当除数为0时返回错误，否则返回nil
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("除零错误") // 返回零值和错误
// 	}
// 	return a / b, nil // 返回计算结果和nil表示无错误
// }

// // 带命名返回值的函数
// // 命名返回值特点：
// // 1. 在函数签名中声明返回值的变量名 (result int)
// // 2. 该变量会在函数开始时自动初始化为类型的零值（int 类型的零值为 0）
// // 3. 可以在函数体内直接使用和修改这个变量
// // 4. 支持裸返回（naked return）：return 语句后不需要显式指定返回值
// // 参数说明：a, b 为两个整数输入参数
// // 返回值：result 为两数相乘的结果
// func multiply(a, b int) (result int) {
// 	// 直接对命名返回值 result 赋值
// 	// 这里 result 相当于函数内的一个局部变量
// 	result = a * b
// 	// 裸返回：return 后不需要指定返回值，会自动返回 result 的当前值
// 	// 等价于：return result
// 	// 注意：裸返回虽然简洁，但在复杂函数中可能降低可读性，建议谨慎使用
// 	return
// }

// // 带可变参数的函数（类似 JavaScript 的剩余参数 ...args）
// // 可变参数特点：
// // 1. 使用 ...type 语法声明，必须是函数的最后一个参数
// // 2. 在函数内部，可变参数被当作 []int 切片（数组）使用
// // 3. 调用时可以传入任意数量（0个或多个）的参数，如：sum(1, 2, 3) 或 sum()
// // 4. 也可以使用切片展开语法：nums := []int{1,2,3}; sum(nums...)
// // 参数说明：numbers 接收任意数量的整数参数
// // 返回值：所有数字的总和
// func sum(numbers ...int) int {
// 	// 初始化累加变量
// 	total := 0
// 	// 使用 range 遍历可变参数切片
// 	// range 返回两个值：索引和元素值
// 	// _ 表示忽略索引，只使用元素值 num
// 	for _, num := range numbers {
// 		// 累加每个数字
// 		total += num
// 	}
// 	// 返回累加结果
// 	return total
// }

// // 函数作为变量（类似函数表达式）
// var add = func(a, b int) int {
// 	return a + b
// }

// // main 函数是 Go 程序的入口点
// // 每个可执行的 Go 程序都必须包含一个 main 包和一个 main 函数
// func main() {
// 	// 调用 greet 函数，传入字符串参数 "Bob"
// 	// 直接打印返回的问候语
// 	fmt.Println(greet("Bob"))

// 	// 调用 divide 函数，演示多返回值的处理
// 	// result 接收计算结果，err 接收错误信息（如果有）
// 	result, err := divide(10, 2)
// 	// Go 的错误处理惯例：先检查 err 是否为 nil
// 	if err != nil {
// 		// 如果 err 不为 nil，说明发生了错误，打印错误信息
// 		fmt.Println("错误:", err)
// 	} else {
// 		// 如果 err 为 nil，说明计算成功，打印结果
// 		fmt.Println("结果:", result)
// 	}

// 	// 调用 multiply 函数，演示命名返回值和裸返回
// 	// 计算 4 * 6 = 24
// 	fmt.Println(multiply(4, 6))

// 	// 调用 sum 函数，演示可变参数的使用
// 	// 传入 5 个参数，函数内部会将它们作为切片处理
// 	// 计算 1 + 2 + 3 + 4 + 5 = 15
// 	fmt.Println(sum(1, 2, 3, 4, 5))

// 	// 调用 add 函数，演示匿名函数/闭包的使用
// 	// 计算 5 + 3 = 8
// 	fmt.Println(add(5, 3))
// }

// Go: 数据类型和结构
package main

import "fmt"

func main() {
	// 基本类型
	var number int = 42
	var text string = "Hello"
	var boolean bool = true

	// 数组（固定大小）
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	// [4]interface{} 表示一个长度固定为 4 的数组，元素类型为 interface{}（可以容纳任意类型）。
	var mixedArray [4]interface{} = [4]interface{}{1, "hello", true, nil}

	// 切片（动态数组）
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6) // 添加元素

	// Map
	// map[string]interface{} // Go 的内建映射类型（类似于 JS 的对象或 Map）。
	// string // 表示 key 必须是字符串类型。
	// interface{} // 表示 value 可以是任意类型（因为 interface{} 是空接口，任何类型都实现了它）。
	object := map[string]interface{}{
		"name":     "Alice",
		"age":      30,
		"isActive": true,
	}

	// 结构体（类似对象但有定义的结构）
	type Person struct {
		Name     string
		Age      int
		IsActive bool
	}

	person := Person{
		Name:     "Bob",
		Age:      25,
		IsActive: true,
	}

	fmt.Printf("数组长度: %d\n", len(array))
	fmt.Printf("切片长度: %d\n", len(slice))
	fmt.Printf("Map 键: %v\n", getMapKeys(object))
	fmt.Printf("Person: %+v\n", person)
}

// 辅助函数获取 map 键
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
