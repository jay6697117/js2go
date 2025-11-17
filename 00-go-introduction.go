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
func multiply(a, b int) (result int) {
	result = a * b
	return // 裸返回
}

// 带可变参数的函数（类似剩余参数）
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
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
