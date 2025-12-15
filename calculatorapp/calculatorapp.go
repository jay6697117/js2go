package calculatorapp

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	fmt.Println("简易计算器 (支持格式: 10+2, 10 + 2, 10.5 * 2 等)")
	fmt.Print("请输入: ")

	// 1. 读取用户输入 (整行读取)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入失败:", err)
		os.Exit(1)
	}

	// 去除首尾空格
	input = strings.TrimSpace(input)

	// 2. 使用正则表达式解析输入
	// 正则说明:
	// ^\s*             : 开头可能有空格
	// (-?\d+(?:\.\d+)?) : 第一个数字 (支持负数和小数)
	// \s*              : 可能有空格
	// ([\+\-\*\/])     : 运算符 (+, -, *, /)
	// \s*              : 可能有空格
	// (-?\d+(?:\.\d+)?) : 第二个数字 (支持负数和小数)
	// \s*$             : 结尾可能有空格
	re := regexp.MustCompile(`^\s*(-?\d+(?:\.\d+)?)\s*([\+\-\*\/])\s*(-?\d+(?:\.\d+)?)\s*$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		fmt.Println("输入格式错误: 请输入类似 10+2 或 10 + 2 的格式")
		os.Exit(1)
	}

	// 解析数字和运算符
	num1Str := matches[1]
	operator := matches[2]
	num2Str := matches[3]

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("数字解析错误")
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
