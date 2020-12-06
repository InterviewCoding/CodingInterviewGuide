package main

import "fmt"

/**
 * @Author: yirufeng
 * @Date: 2020/11/22 6:15 下午
 * @Desc: 判断两个字符串是否互为变形词
 **/

//变形词指的是出现字符串长度相同且出现的字符种类和数目相同的情况
func isDeformation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	//如果长度相同则使用一个map来统计字符出现的次数
	strMap := make(map[uint8]int)
	//遍历str1统计字符出现的次数
	for i := 0; i < len(str1); i++ {
		strMap[str1[i]] += 1
	}
	//之后使用对应的map遍历str2
	for i := 0; i < len(str2); i++ {
		strMap[str2[i]] -= 1
		//如果出现之后，减去次数为-1，那么说明不对，就返回false
		if strMap[str2[i]] < 0 {
			return false
		}
	}

	//否则到最后我们直接返回true
	return true
}

func main() {
	fmt.Println(isDeformation("123", "345"))
	fmt.Println(isDeformation("123123", "321123"))
	fmt.Println(isDeformation("123123", "322123"))

}
