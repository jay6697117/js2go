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

// // 4.Go: 数据类型和结构
// package main

// import "fmt"

// func main() {
// 	// 基本类型
// 	var number int = 42
// 	var text string = "Hello"
// 	var boolean bool = true
// 	fmt.Printf("基本类型 - number: %v, text: %v, boolean: %v\n", number, text, boolean)

// 	// 数组（固定大小）
// 	var array [5]int = [5]int{1, 2, 3, 4, 5}
// 	// [4]interface{} 表示一个长度固定为 4 的数组，元素类型为 interface{}（可以容纳任意类型）。
// 	var mixedArray [4]interface{} = [4]interface{}{1, "hello", true, nil}
// 	fmt.Printf("混合数组: %v\n", mixedArray)

// 	// 切片（动态数组）
// 	slice := []int{1, 2, 3, 4, 5}
// 	slice = append(slice, 6) // 添加元素

// 	// Map
// 	// map[string]interface{} // Go 的内建映射类型（类似于 JS 的对象或 Map）。
// 	// string // 表示 key 必须是字符串类型。
// 	// interface{} // 表示 value 可以是任意类型（因为 interface{} 是空接口，任何类型都实现了它）。
// 	object := map[string]interface{}{
// 		"name":     "Alice",
// 		"age":      30,
// 		"isActive": true,
// 	}

// 	// 结构体（类似对象但有定义的结构）
// 	type Person struct {
// 		Name     string
// 		Age      int
// 		IsActive bool
// 	}

// 	person := Person{
// 		Name:     "Bob",
// 		Age:      25,
// 		IsActive: true,
// 	}

// 	fmt.Printf("数组长度: %d\n", len(array))
// 	fmt.Printf("数组: %v\n", array)
// 	fmt.Printf("切片长度: %d\n", len(slice))
// 	fmt.Printf("切片: %d\n", slice)
// 	fmt.Printf("Map 键: %v\n", getMapKeys(object))
// 	fmt.Printf("Map: %v\n", object)
// 	fmt.Printf("Person: %+v\n", person)
// }

// // getMapKeys 从 map 中提取所有的键并返回一个字符串切片
// // 功能说明：
// //
// //	该函数遍历输入的 map，提取所有键并以字符串切片的形式返回
// //	适用于需要获取 map 所有键进行进一步处理的场景
// //
// // 参数:
// //
// //	m: 需要提取键的 map，键类型为 string，值类型为 interface{}（可以是任意类型）
// //	   - 如果传入 nil map，函数会返回空切片
// //	   - 如果传入空 map（长度为0），函数也会返回空切片
// //
// // 返回值:
// //
// //	[]string: 包含 map 中所有键的字符串切片
// //	         - 键的顺序不保证，因为 Go 的 map 遍历顺序是随机的
// //	         - 返回的切片长度等于输入 map 的长度
// //
// // 时间复杂度: O(n)，其中 n 是 map 中键值对的数量
// // 空间复杂度: O(n)，需要创建一个新的切片来存储所有键
// func getMapKeys(m map[string]interface{}) []string {
// 	// 使用 make 创建切片，第二个参数 0 表示初始长度为 0
// 	// 第三个参数 len(m) 表示预分配容量为 map 的长度
// 	// 性能优化：预分配容量避免了切片在 append 过程中的多次内存重新分配
// 	// 如果不预分配容量，切片会在容量不足时自动扩容（通常是翻倍），可能导致多次内存拷贝
// 	keys := make([]string, 0, len(m))

// 	// 使用 range 遍历 map，语法：for key, value := range map
// 	// 这里只需要键，所以省略了值的接收变量（Go 允许这种简化写法）
// 	// 注意：Go 的 map 遍历顺序是随机的，每次运行可能得到不同的键顺序
// 	for k := range m {
// 		// append 函数将键 k 追加到 keys 切片的末尾
// 		// 由于预分配了足够的容量，这里不会触发切片扩容，性能较好
// 		keys = append(keys, k)
// 	}

// 	// 返回包含所有键的切片
// 	// 调用者可以对返回的切片进行排序、过滤等进一步操作
// 	return keys
// }

// // 5.Go: 作用域和生命周期
// package main

// import "fmt"

// func example() {
// 	x := 10 // 函数作用域

// 	if true {
// 		y := 20        // 块作用域
// 		fmt.Println(x) // 10
// 		fmt.Println(y) // 20
// 	}

// 	fmt.Println(x) // 10
// 	// fmt.Println(y) // Error: undefined: y
// }

// // Go 中的闭包 (Closure)
// // ========================================
// // 闭包是一个函数值（function value），它引用了其函数体之外的变量。
// // 该函数可以访问并修改这些被引用的变量，换句话说，该函数被"绑定"在了这些变量上。
// //
// // createCounter 是一个"工厂函数"，它返回一个闭包。
// // 函数签名解读：
// //   - func createCounter()  : 函数名为 createCounter，不接受任何参数
// //   - func() int            : 返回值类型是"一个不接受参数、返回 int 的函数"
// func createCounter() func() int {
// 	// count 是外部函数 createCounter 的局部变量。
// 	// 正常情况下，当 createCounter 执行完毕后，count 应该被销毁。
// 	count := 0

// 	// 返回一个匿名函数（Anonymous Function）。
// 	// 这个匿名函数"捕获"了外部变量 count，形成了一个闭包。
// 	// 由于 count 被闭包引用，Go 的逃逸分析（Escape Analysis）会将 count 分配到堆上，
// 	// 使其生命周期延长，即使 createCounter 函数返回后，count 依然存活。
// 	return func() int {
// 		// 每次调用这个闭包，都在修改和访问同一个 count 变量。
// 		// 这就是闭包"保留状态"的核心机制。
// 		count++
// 		return count
// 	}
// }

// // 使用示例：
// //   counter1 := createCounter()  // 创建一个计数器，拥有独立的 count (初始为0)
// //   counter1()                   // 返回 1
// //   counter1()                   // 返回 2
// //
// //   counter2 := createCounter()  // 创建另一个计数器，拥有自己独立的 count
// //   counter2()                   // 返回 1 (与 counter1 互不影响)

// func main() {
// 	example()

// 	counter := createCounter()
// 	fmt.Println(counter()) // 1
// 	fmt.Println(counter()) // 2
// }

// 6. Go: 错误处理
package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// 返回错误的函数
// ========================================
// Go 的错误处理采用"显式返回 error"的设计哲学，而非 try/catch 异常机制。
// 这使得错误处理路径清晰可见，强制开发者思考每个可能出错的地方。
//
// 函数签名解读：
//   - divide(a, b float64)  : 接收两个 float64 类型的参数
//   - (float64, error)      : 返回两个值 —— 计算结果和错误信息
//
// Go 的 error 是一个内置接口，定义如下：
//
//	type error interface {
//	    Error() string
//	}
//
// 任何实现了 Error() string 方法的类型都可以作为 error 使用。
func divide(a, b float64) (float64, error) {
	// 边界条件检查：除数不能为零
	if b == 0 {
		// errors.New() 创建一个简单的错误对象
		// 返回值约定：当发生错误时，通常返回零值（0）和非 nil 的 error
		return 0, errors.New("除零错误")
	}
	// 正常情况：返回计算结果和 nil
	// nil 表示"没有错误"，这是 Go 错误处理的核心约定
	return a / b, nil
}

// 处理错误的函数
func safeDivide(a, b float64) {
	result, err := divide(a, b)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("结果: %v\n", result)
}

// 自定义错误类型
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("验证错误 %s: %s", e.Field, e.Message)
}

// 带自定义错误的函数
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{Field: "age", Message: "年龄不能为负数"}
	}
	if age > 150 {
		return ValidationError{Field: "age", Message: "年龄似乎不现实"}
	}
	return nil
}

// 带错误处理的 HTTP 请求
func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("获取数据失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP 错误! 状态: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	return body, nil
}

func main() {
	safeDivide(10, 2) // 结果: 5
	safeDivide(10, 0) // 错误: 除零错误

	if err := validateAge(-5); err != nil {
		fmt.Printf("验证错误: %v\n", err)
	}

	if err := validateAge(200); err != nil {
		fmt.Printf("验证错误: %v\n", err)
	}

	// 测试 fetchData：访问 GitHub API
	fmt.Println("\n--- 测试 fetchData ---")
	data, err := fetchData("https://api.github.com/users/octocat")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
	} else {
		// 打印前 200 个字符（避免输出过长）
		output := string(data)
		if len(output) > 200 {
			output = output[:200] + "..."
		}
		fmt.Printf("GitHub API 响应:\n%s\n", output)
	}
}
