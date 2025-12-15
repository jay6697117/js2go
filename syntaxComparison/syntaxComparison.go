// Go: 运算符
package syntaxcomparison

import (
	"fmt"
	"math"
)

func SyntaxComparison() {
	a, b := 10, 3

	// 1. 算术运算符 (Arithmetic Operators)
	// ==========================================
	// Go 支持标准的算术运算。
	// 注意：除法运算 `/` 对于整数会进行截断（向下取整）。
	sum := a + b        // 加法: 13
	difference := a - b // 减法: 7
	product := a * b    // 乘法: 30

	// 类型转换 (Type Conversion)
	// ------------------------------------------
	// Go 是强类型语言，不支持隐式类型转换。
	// int 和 float64 不能直接运算，必须显式转换。
	// 这里将 int 转换为 float64 以获得精确的除法结果。
	quotient := float64(a) / float64(b) // 除法: 3.333...

	remainder := a % b // 取模 (余数): 1

	// Math 包的使用
	// ------------------------------------------
	// 幂运算没有内置操作符（如 **），必须使用 math.Pow 函数。
	// math.Pow 接收 float64 参数并返回 float64。
	power := math.Pow(float64(a), float64(b)) // 幂运算: 10^3 = 1000

	// 2. 比较运算符 (Comparison Operators)
	// ==========================================
	// 比较结果是布尔值 (bool)：true 或 false
	isEqual := (a == b)  // 等于: false
	isGreater := (a > b) // 大于: true
	isLess := (a < b)    // 小于: false

	// 3. 逻辑运算符 (Logical Operators)
	// ==========================================
	// 用于组合布尔值。
	and := true && false // 逻辑与 (AND): 只有两边都为 true 时才为 true
	or := true || false  // 逻辑或 (OR): 只要有一边为 true 就为 true
	not := !true         // 逻辑非 (NOT): 取反

	// 4. 位运算符 (Bitwise Operators)
	// ==========================================
	// 针对二进制位进行操作。
	// a = 10 (二进制 1010)
	// b = 3  (二进制 0011)
	bitwiseAnd := a & b  // 按位与 (AND): 0010 (十进制 2)
	bitwiseOr := a | b   // 按位或 (OR):  1011 (十进制 11)
	bitwiseXor := a ^ b  // 按位异或 (XOR): 1001 (十进制 9) - 相同为0，不同为1
	leftShift := a << 1  // 左移 1 位: 10100 (十进制 20) - 相当于乘以 2
	rightShift := a >> 1 // 右移 1 位: 0101 (十进制 5) - 相当于除以 2

	// 格式化输出 (Formatting)
	// ==========================================
	// %d   : 十进制整数
	// %.2f : 浮点数，保留两位小数
	// %.0f : 浮点数，不保留小数 (四舍五入)
	// %t   : 布尔值 (true/false)
	fmt.Printf("算术: %d, %d, %d, %.2f, %d, %.0f\n",
		sum, difference, product, quotient, remainder, power)
	fmt.Printf("比较: %t, %t, %t\n", isEqual, isGreater, isLess)
	fmt.Printf("逻辑: %t, %t, %t\n", and, or, not)
	fmt.Printf("位运算: %d, %d, %d, %d, %d\n",
		bitwiseAnd, bitwiseOr, bitwiseXor, leftShift, rightShift)
}
