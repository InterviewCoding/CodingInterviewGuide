package main

import "fmt"

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/11/5 10:05 上午
 * @Desc: 在其他数都出现k次的数组中找到只出现1次的数

解题思路：
1. 遍历arr，将遍历到的每一个数转换为k进制数，我们通过一个[]int存储转换后的k进制数的每一位，之后我们分别对每一位进行运算并得到结果
2. 所有数进行k进制无进位相加，得到ret
3. 将ret转换为十进制数并进行返回
*/

//在arr中只有1个数出现1次，其他数都出现k次，找到只出现1次的数并进行返回
func onceNum(arr []int, k int) int {

	//1. 建立一个结果对应的位数组，用于存放数组中所有数对应的k进制数无进位相加的结果
	retBitArr := make([]int, 32)
	//2. 将数组中每个数分别转换为k进制并按位和retBitArr进行无进位相加
	for i := 0; i < len(arr); i++ {
		tempNum := convertToKSys(arr[i], k)

		if i == 27 {
			fmt.Println("27之前：", retBitArr)
		}

		fmt.Println("转换：----", tempNum)
		for j := 0; j < 32; j++ {
			retBitArr[j] = (retBitArr[j] + tempNum[j]) % k
		}

		if i == 27 {
			fmt.Println("27之后：", retBitArr)
		}
	}

	fmt.Println("31：", retBitArr)
	//3. 将异或之后的结果k进制数组转换为我们的10进制数并返回
	ret := convertKSysToNum(retBitArr, k)

	//4. 返回结果
	return ret
}

//将十进制的num转换为k进制的[]int并进行返回
func convertToKSys(num int, k int) []int {
	ret := make([]int, 32)
	index := 0
	for num != 0 {
		ret[index] = num % k
		num /= k
		index++
	}

	return ret
}

func convertKSysToNum(retBitArr []int, k int) int {
	ret := 0
	temp := 1
	for i := 0; i < 32; i++ {
		ret += temp * retBitArr[i]
		temp *= k
	}
	return ret
}

func convertKSysToNum2(retBitArr []int, k int) int {
	ret := 0
	for i := 31; i != -1; i-- {
		ret = ret*k + retBitArr[i]
	}
	return ret
}

func main() {
	//以上代码会在这一行测试用例报错
	//fmt.Println(onceNum([]int{-401451, -177656, -2147483646, -473874, -814645, -2147483646, -852036, -457533,
	//	-401451, -473874, -401451, -216555, -917279, -457533, -852036, -457533, -177656, -2147483646, -177656,
	//	-917279, -473874, -852036, -917279, -216555, -814645, 2147483645, -2147483648, 2147483645, -814645,
	//	2147483645, -216555}, 3))

	//fmt.Println(onceNum([]int{-1, -1, -1, 2, 2, 2, 23233232}, 3))
	fmt.Println(convertToKSys(-2147483648, 3))
	fmt.Println(convertKSysToNum([]int{-2, 0, -1, -2, 0, -2, 0, -1, -1, -2, -1, -2, -2, -2, -1, -2, -1, -1, -2, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 3))

	fmt.Println(onceNum([]int{-401451, -177656, -2147483646, -473874, -814645, -2147483646, -852036, -457533,
		-401451, -473874, -401451, -216555, -917279, -457533, -852036, -457533, -177656, -2147483646, -177656,
		-917279, -473874, -852036, -917279, -216555, -814645, 2147483645, -2147483648, 2147483645, -814645,
		2147483645, -216555}, 3))

	//-2147483648
	//fmt.Println(onceNum([]int{-401451, -177656, -2147483646, -473874, -814645, -852036, -216555, -917279, -457533, 2147483645}, 3))

	//nums := []int{-401451, -177656, -2147483646, -473874, -814645, -852036, -216555, -917279, -457533, 2147483645}
	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(convertToKSys(nums[i], 3))
	//}
	//fmt.Println()
	//fmt.Println(convertToKSys(2147483648, 3))
	fmt.Println(convertKSysToNum2([]int{-2, 0, -1, -2, 0, -2, 0, -1, -1, -2, -1, -2, -2, -2, -1, -2, -1, -1, -2, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 3))
}
