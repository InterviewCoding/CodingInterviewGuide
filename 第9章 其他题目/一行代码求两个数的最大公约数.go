package main

import "fmt"

/**
 * @Author: yirufeng
 * @Date: 2020/11/21 11:26 上午
 * @Desc: 一行代码求两个数的最大公约数

思路：使用辗转相除法求两个数的最大公约数，例如 a ➗ b = m ..... n。那么a和b的最大公约数将会是b 与 n的最大公约数
 **/
func getGcd(a, b int) int {
	if b == 0 {
		return a
	}

	return getGcd(b, a % b)
}

func main() {
	fmt.Println(getGcd(10, 3))
	fmt.Println(getGcd(9, 3))
}