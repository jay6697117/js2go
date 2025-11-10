package main

import "fmt"

func main() {
	temperature := 25

	if temperature > 30 {
		fmt.Println("天气很热！")
	} else if temperature > 20 {
		fmt.Println("天气温暖。")
	} else {
		fmt.Println("天气寒冷。")
	}

	// Go 需要显式的布尔表达式
	if temperature > 0 {
		fmt.Println("温度是正数")
	}
}
