// Go: 静态类型与类型推断
package main

import "fmt"

func main() {
	age := 30                                                  // int (类型推断)
	name := "Alice"                                            // string (类型推断)
	isActive := true                                           // bool (类型推断)
	scores := []int{85, 92, 78}                                // int 切片
	person := map[string]interface{}{"name": "Bob", "age": 25} // map

	// Go 没有像 JavaScript 那样的 typeof，但可以使用反射
	fmt.Printf("age 类型: %T\n", age)       // int
	fmt.Printf("name 类型: %T\n", name)     // string
	fmt.Printf("scores 类型: %T\n", scores) // []int
}
