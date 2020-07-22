
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--> Hello, World!");

	tips := "这里是my类型转换?"
	fmt.Println("--> tips: " + tips);

	// int 类型转成 string 类型
	fmt.Println("---- Itoa -----")
	num := 3311
	str := strconv.Itoa(num)
	fmt.Printf("--> 类型: %T, 值: %v \n", num, num)
	// 输出结果: “--> 类型: string, 值: 3311”
	fmt.Printf("--> 类型: %T, 值: %v \n", str, str)
	// 输出结果: “--> 类型: string, 值: 3311”

	// Atoi: 将 string 类型转成 int 类型
	fmt.Println("---- Atoi -----")
	num2 := "1012"
	if s, err := strconv.Atoi(num2); err == nil {
		fmt.Printf("--> 类型: %T, 值: %v \n", s, s)
		// 输出结果: “--> 类型: int, 值: 1012”
	}

	// ParseFloat: 将字符串转换成浮点数
	fmt.Println("---- ParseFloat -----")
	v3 := "3.1415926535"
	if s, err := strconv.ParseFloat(v3, 32); err == nil {
		// 输出结果: “--> 类型: float64, 值: 3.1415927410125732”
		fmt.Printf("--> 类型: %T, 值: %v \n", s, s)
	}
	if s, err := strconv.ParseFloat(v3, 64); err == nil {
		// 输出结果: “--> 类型: float64, 值: 3.1415926535”
		fmt.Printf("--> 类型: %T, 值: %v \n", s, s)
	}

	// 特殊字符转义
	fmt.Println("---- Quote -----")
	// there is a tab character inside the string literal
	s := strconv.Quote(`"Fran & \n \t Freddie's Diner   ☺"`)
	// 输出结果: “"\"Fran & \\n \\t Freddie's Diner\t☺\""”
	fmt.Println(s)
	// QuoteToASCII 将字符串转换为“双引号”引起来的 ASCII 字符串
	fmt.Println(strconv.QuoteToASCII("to quote Shakespeare 引用莎士比亚的话"))

	// 特殊字符取消转义
	fmt.Println("---- Unquote -----")
	s1 := "`Hello   世界！`"             // 解析反引号字符串
	s2 := `"Hello\t\u4e16\u754c\uff01"` // 解析双引号字符串
	s3 := `"to quote Shakespeare \u5f15\u7528\u838e\u58eb\u6bd4\u4e9a\u7684\u8bdd"`
	// 解析双引号字符串
	fmt.Println(strconv.Unquote(s1))    // Hello    世界！ <nil>
	fmt.Println(strconv.Unquote(s2))    // Hello    世界！ <nil>
	fmt.Println(strconv.Unquote(s3))    // to quote Shakespeare 引用莎士比亚的话 <nil>
}