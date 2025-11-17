// // Go: If/Else
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

// // Go: For 循环
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

// Go: 函数定义
package main

import "fmt"

// 函数声明
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// 带多个返回值的函数
// divide 执行除法运算
// 参数: a - 被除数, b - 除数
// 返回: 商(float64) 和 错误(error)
// 当除数为0时返回错误，否则返回nil
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除零错误") // 返回零值和错误
	}
	return a / b, nil // 返回计算结果和nil表示无错误
}

// 带命名返回值的函数
// 命名返回值特点：
// 1. 在函数签名中声明返回值的变量名 (result int)
// 2. 该变量会在函数开始时自动初始化为类型的零值（int 类型的零值为 0）
// 3. 可以在函数体内直接使用和修改这个变量
// 4. 支持裸返回（naked return）：return 语句后不需要显式指定返回值
// 参数说明：a, b 为两个整数输入参数
// 返回值：result 为两数相乘的结果
func multiply(a, b int) (result int) {
	// 直接对命名返回值 result 赋值
	// 这里 result 相当于函数内的一个局部变量
	result = a * b
	// 裸返回：return 后不需要指定返回值，会自动返回 result 的当前值
	// 等价于：return result
	// 注意：裸返回虽然简洁，但在复杂函数中可能降低可读性，建议谨慎使用
	return
}

// 带可变参数的函数（类似 JavaScript 的剩余参数 ...args）
// 可变参数特点：
// 1. 使用 ...type 语法声明，必须是函数的最后一个参数
// 2. 在函数内部，可变参数被当作 []int 切片（数组）使用
// 3. 调用时可以传入任意数量（0个或多个）的参数，如：sum(1, 2, 3) 或 sum()
// 4. 也可以使用切片展开语法：nums := []int{1,2,3}; sum(nums...)
// 参数说明：numbers 接收任意数量的整数参数
// 返回值：所有数字的总和
func sum(numbers ...int) int {
	// 初始化累加变量
	total := 0
	// 使用 range 遍历可变参数切片
	// range 返回两个值：索引和元素值
	// _ 表示忽略索引，只使用元素值 num
	for _, num := range numbers {
		// 累加每个数字
		total += num
	}
	// 返回累加结果
	return total
}

// 函数作为变量（类似函数表达式）
var add = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greet("Bob"))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	fmt.Println(multiply(4, 6))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(add(5, 3))
}
