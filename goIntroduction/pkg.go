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
package goIntroduction

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

// 错误处理的惯用模式
// ========================================
// 在 Go 中，处理错误的正确方式是立即检查返回的 error 对象。
// 这种模式通常被称为 "Guard Clause"（卫语句）或 "Early Return"（尽早返回）。
// 它可以避免嵌套层级过深（避免 "箭头型" 代码），使正常逻辑流保持在最左侧。
func safeDivide(a, b float64) {
	// 接收多返回值：result 接收计算结果，err 接收可能的错误
	result, err := divide(a, b)

	// 检查错误：如果 err 不为 nil，说明发生了错误
	if err != nil {
		// 处理错误：这里只是简单打印，实际项目中可能需要记录日志、重试或向上传递错误
		fmt.Printf("错误: %v\n", err)
		// 发生错误后通常需要中断当前流程（return）
		return
	}

	// 如果代码执行到这里，说明 err 为 nil，操作成功
	// 这被称为 "Happy Path"（快乐路径）
	fmt.Printf("结果: %v\n", result)
}

// 自定义错误类型
// ========================================
// 结构体 ValidationError 用于表示数据验证失败的情况。
// 它不仅包含错误信息，还记录了具体是哪个字段验证失败，
// 展示了如何通过结构体携带更丰富的错误上下文。
type ValidationError struct {
	Field   string // 验证失败的字段名称（例如 "username" 或 "age"）
	Message string // 具体的错误描述（例如 "不能为空" 或 "必须大于18"）
}

// Error 实现 error 接口
// ----------------------------------------
// 任何实现了 Error() string 方法的类型都可以被视为 error 类型。
// 这里我们将 ValidationError 结构体适配为标准的 error 接口。
//
// 方法接收者 (Method Receiver) 语法说明：
// ========================================
// 这里的 (e ValidationError) 表示 Error() 方法属于 ValidationError 类型。
//
//  1. 语法解释
//     func (e ValidationError) Error() string { ... }
//     - e : 接收者的变量名（类似于其他语言中的 this 或 self）
//     - ValidationError : 接收者的类型
//
//  2. 与普通函数对比
//     | 类型     | 语法定义                                 | 调用方式      |
//     | :------- | :--------------------------------------- | :------------ |
//     | 普通函数 | func Error() string                      | Error()       |
//     | 方法     | func (e ValidationError) Error() string  | e.Error()     |
func (e ValidationError) Error() string {
	// 格式化并返回易读的错误字符串
	return fmt.Sprintf("验证错误 [字段: %s]: %s", e.Field, e.Message)
}

// 带自定义错误的函数
// ========================================
// validateAge 演示如何使用自定义错误类型进行业务逻辑验证。
// 它展示了 Go 中将领域规则封装在验证函数中的惯用模式。
//
// 函数签名解读：
//   - validateAge(age int) : 接收一个整数类型的年龄参数
//   - error               : 返回标准的 error 接口类型
//
// 返回值说明：
//   - 验证通过时返回 nil（表示无错误）
//   - 验证失败时返回 ValidationError 结构体（实现了 error 接口）
//
// 这种设计的优势：
//  1. 类型安全：调用者可以通过类型断言获取详细的错误信息
//  2. 可扩展性：可以轻松添加更多验证规则
//  3. 可测试性：每个验证规则都可以独立测试
func validateAge(age int) error {
	// 验证规则 1：年龄不能为负数
	if age < 0 {
		// 返回自定义错误，携带字段名和具体错误描述
		// 这样调用者既可以使用 err.Error() 获取格式化的错误信息，
		// 也可以通过类型断言访问 Field 和 Message 字段
		return ValidationError{Field: "age", Message: "年龄不能为负数"}
	}
	// 验证规则 2：年龄不能超过合理范围
	if age > 150 {
		return ValidationError{Field: "age", Message: "年龄似乎不现实"}
	}
	// 所有验证通过，返回 nil 表示成功
	// 这里的 nil 会被解释为"一个没有具体值的 error 接口"，表示没有错误发生
	return nil
}

// 带错误处理的 HTTP 请求
// ========================================
// fetchData 演示了 Go 中进行 HTTP 请求的标准模式，包含完整的错误处理。
// 这是一个同步请求函数，会阻塞直到请求完成或发生错误。
//
// 函数签名解读：
//   - fetchData(url string)  : 接收一个 URL 字符串作为参数
//   - ([]byte, error)        : 返回响应体的字节切片和可能的错误
//
// 核心知识点：
//  1. http.Get()     : 标准库提供的便捷 HTTP GET 请求方法
//  2. defer 语句     : 延迟执行，确保资源正确释放
//  3. 错误包装 (%w)  : Go 1.13+ 引入的错误链机制
//  4. io.ReadAll()   : 读取整个响应体到内存
func fetchData(url string) ([]byte, error) {
	// http.Get 是 net/http 包提供的便捷函数
	// ----------------------------------------
	// 它是 http.DefaultClient.Get(url) 的简写形式
	// 返回值：
	//   - resp *http.Response : HTTP 响应对象，包含状态码、响应头、响应体等
	//   - err error           : 网络错误（如 DNS 解析失败、连接超时、TLS 握手失败等）
	//
	// 注意：即使服务器返回 4xx 或 5xx 状态码，只要网络通信成功，err 仍为 nil
	// 因此需要单独检查 resp.StatusCode
	resp, err := http.Get(url)

	// 错误检查：网络层面的错误
	// ----------------------------------------
	// 这里捕获的是网络传输层面的错误，例如：
	//   - 无法解析域名（DNS 错误）
	//   - 无法建立 TCP 连接（目标主机不可达）
	//   - TLS/SSL 证书验证失败
	//   - 请求超时
	if err != nil {
		// fmt.Errorf 配合 %w 动词进行错误包装 (Error Wrapping)
		// ----------------------------------------
		// %w 是 Go 1.13 引入的特殊格式化动词，用于包装错误。
		// 它的作用：
		//   1. 保留原始错误信息（可通过 errors.Unwrap() 获取）
		//   2. 添加上下文描述，使错误信息更具可读性
		//   3. 支持 errors.Is() 和 errors.As() 进行错误链检查
		//
		// 对比其他方式：
		//   - %v : 只保留错误字符串，丢失原始错误类型
		//   - %w : 保留完整错误链，可追溯根因
		return nil, fmt.Errorf("获取数据失败: %w", err)
	}

	// defer 延迟执行：资源清理的最佳实践
	// ========================================
	// defer 关键字确保 resp.Body.Close() 在函数返回前执行，无论是正常返回还是发生错误。
	//
	// 为什么必须关闭 resp.Body？
	//   1. resp.Body 是一个 io.ReadCloser，底层持有网络连接资源
	//   2. 不关闭会导致连接泄漏，最终耗尽系统的文件描述符
	//   3. http.Transport 默认维护连接池，关闭 Body 后连接可被复用
	//
	// defer 的执行时机：
	//   - 在包含它的函数返回时执行（不是当前代码块）
	//   - 多个 defer 按 LIFO（后进先出）顺序执行
	//   - 即使发生 panic，defer 仍会执行（用于资源清理）
	//
	// 最佳实践：在获取需要关闭的资源后，立即使用 defer 声明关闭操作
	defer resp.Body.Close()

	// HTTP 状态码检查
	// ----------------------------------------
	// http.StatusOK 是常量 200，表示请求成功。
	// 常见的 HTTP 状态码常量：
	//   - http.StatusOK           (200) : 成功
	//   - http.StatusCreated      (201) : 资源创建成功
	//   - http.StatusBadRequest   (400) : 客户端请求错误
	//   - http.StatusUnauthorized (401) : 未认证
	//   - http.StatusForbidden    (403) : 禁止访问
	//   - http.StatusNotFound     (404) : 资源不存在
	//   - http.StatusInternalServerError (500) : 服务器内部错误
	//
	// 注意：3xx 重定向状态码通常由 http.Client 自动处理（默认最多跟随 10 次重定向）
	if resp.StatusCode != http.StatusOK {
		// %d 格式化动词用于输出十进制整数
		return nil, fmt.Errorf("HTTP 错误! 状态: %d", resp.StatusCode)
	}

	// 读取响应体
	// ----------------------------------------
	// io.ReadAll (Go 1.16+) 从 Reader 中读取所有数据直到 EOF。
	// 在 Go 1.15 及更早版本中，使用 ioutil.ReadAll（现已弃用）。
	//
	// 返回值：
	//   - body []byte : 响应体的完整内容
	//   - err error   : 读取过程中发生的错误（如连接中断）
	//
	// 注意事项：
	//   1. ReadAll 会将整个响应体加载到内存，对于大文件可能导致 OOM
	//   2. 对于大文件，应使用流式处理（io.Copy 或分块读取）
	//   3. 响应体只能读取一次，因为它是一个流（io.Reader）
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 成功：返回响应体数据和 nil 错误
	return body, nil
}

func GoIntroduction() {
	// 示例 1: 基本错误处理模式
	// ----------------------------------------
	// 调用 safeDivide 函数，它内部演示了"卫语句"（Guard Clause）模式
	// 这种模式在 Go 中非常常见：立即检查错误并处理，保持主逻辑的简洁
	safeDivide(10, 2) // 正常情况：输出结果 5
	safeDivide(10, 0) // 错误情况：输出错误信息

	// 示例 2: if 语句中的简短变量声明
	// ----------------------------------------
	// 语法特点：if <初始化语句>; <条件表达式> { ... }
	//
	// 这里 err := validateAge(-5) 是初始化语句：
	//   1. 先执行 validateAge(-5)
	//   2. 将返回值赋给新变量 err
	//   3. err 的作用域仅限于这个 if-else 代码块内
	//
	// 这种写法的优势：
	//   - 限制了 err 变量的作用域，避免污染外部作用域
	//   - 代码更加紧凑，逻辑更加连贯
	if err := validateAge(-5); err != nil {
		// 这里 err 不为 nil，说明验证失败
		// %v 会打印错误的默认格式
		fmt.Printf("验证错误: %v\n", err)
	}

	// 再次演示简短变量声明，这次使用不同的变量名或复用（但在不同作用域）
	// 注意：这里的 err 是一个新的变量，与上面的 err 无关
	if err := validateAge(200); err != nil {
		fmt.Printf("验证错误: %v\n", err)
	}

	// 示例 3: 综合应用 - HTTP 请求与数据处理
	// ----------------------------------------
	// 演示真实场景下的错误处理流程：
	//   1. 发起请求
	//   2. 检查错误
	//   3. 处理结果
	fmt.Println("\n--- 测试 fetchData ---")

	// 调用 fetchData 获取 GitHub 用户信息
	// data: 响应内容的字节切片 ([]byte)
	// err : 可能发生的错误
	data, err := fetchData("https://api.github.com/users/jay6697117")

	// 检查请求是否成功
	if err != nil {
		// 如果发生错误（如断网、域名解析失败、404等），打印错误并结束当前逻辑
		fmt.Printf("请求失败: %v\n", err)
	} else {
		// 请求成功，处理返回的数据
		// ----------------------------------------------------
		// 1. 将字节切片转换为字符串，以便打印
		output := string(data)

		// 2. 简单的截断处理：避免控制台输出过长内容
		//    由于 GitHub API 返回的 JSON 可能很长，这里只展示前 200 个字符
		if len(output) > 2000 {
			output = output[:2000] + "..."
		}

		// 3. 打印最终结果
		fmt.Printf("GitHub API 响应:\n%s\n", output)
	}
}
