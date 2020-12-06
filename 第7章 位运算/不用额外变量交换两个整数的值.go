package main

import "fmt"

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/11/6 7:09 下午
 * @Desc:


两种交换的方式：
1. 使用异或
2. 使用语法特性
3. 其他可以交换的方式，如加减乘除需要注意溢出，因此不建议使用
 */

//第一种方式：直接进行交换
func swapNum(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

//第二种方式：使用异或进行交换可以有效避免溢出
func swapNum2(a int, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	fmt.Println(swapNum(10, 2))
	fmt.Println(swapNum2(10, 2))
}
