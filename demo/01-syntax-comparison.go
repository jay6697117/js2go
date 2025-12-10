// Go: 运算符
package demo

import (
	"fmt"
	"math"
)

func syntaxComparison() {
	a, b := 10, 3

	// 算术
	sum := a + b                              // 13
	difference := a - b                       // 7
	product := a * b                          // 30
	quotient := float64(a) / float64(b)       // 3.333... (需要浮点转换)
	remainder := a % b                        // 1
	power := math.Pow(float64(a), float64(b)) // 1000 (需要 math 包)

	// 比较
	isEqual := (a == b)  // false
	isGreater := (a > b) // true
	isLess := (a < b)    // false

	// 逻辑
	and := true && false // false
	or := true || false  // true
	not := !true         // false

	// 位运算
	bitwiseAnd := a & b  // 2
	bitwiseOr := a | b   // 11
	bitwiseXor := a ^ b  // 9
	leftShift := a << 1  // 20
	rightShift := a >> 1 // 5

	fmt.Printf("算术: %d, %d, %d, %.2f, %d, %.0f\n",
		sum, difference, product, quotient, remainder, power)
	fmt.Printf("比较: %t, %t, %t\n", isEqual, isGreater, isLess)
	fmt.Printf("逻辑: %t, %t, %t\n", and, or, not)
	fmt.Printf("位运算: %d, %d, %d, %d, %d\n",
		bitwiseAnd, bitwiseOr, bitwiseXor, leftShift, rightShift)
}
