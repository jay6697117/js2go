package calculatorapp

import (
	"errors"
	"fmt"
	"os"
)

// calculate 执行运算
// Go 特性：函数可以返回多个值 (result, error)
func CalculateFunc(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			// Go 特性：使用 errors 包创建错误对象
			return 0, errors.New("除数不能为零")
		}
		return a / b, nil
	default:
		return 0, errors.New("无效的运算符")
	}
}

func CalculateExecute() {
	var num1, num2 float64
	var operator string

	fmt.Println("简易计算器 (格式: 数字 运算符 数字，例如: 10 + 2)")
	fmt.Print("请输入: ")

	// 1. 处理用户输入
	// fmt.Scan 会自动根据空格分割输入并解析到对应的变量地址中
	_, err := fmt.Scan(&num1, &operator, &num2)

	// 2. 输入错误处理
	if err != nil {
		fmt.Println("输入错误: 请确保格式正确 (数字 空格 运算符 空格 数字)")
		// 非零状态码退出表示程序异常结束
		os.Exit(1)
	}

	// 3. 执行计算
	result, err := CalculateFunc(num1, num2, operator)

	// 4. 逻辑错误处理
	if err != nil {
		fmt.Printf("计算错误: %v\n", err)
	} else {
		// %g 格式化会自动省略不必要的浮点数小数位
		fmt.Printf("结果: %g\n", result)
	}
}
