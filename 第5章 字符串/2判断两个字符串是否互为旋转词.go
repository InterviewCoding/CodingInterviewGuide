package main

/**
 * @Author: yirufeng
 * @Date: 2020/11/22 6:27 下午
 * @Desc: 判断两个字符串是否互为旋转词
 **/


//A的旋转操作就是将A 最左边的字符移动到最右边。例如, 若A = 'abcde'，在移动一次之后结果就是'bcdea'。如果在若干次旋转操作之后，A能变成B，那么返回True。
//自己的想法：因为旋转次可以旋转前k个字符，其中k为0-长度
//思路：拼接两个A为一个新字符串newStr，之后判断B是否在newStr里面即可
func rotateString(A string, B string) bool {

	//判断B是否为A+A的子串，如果是的话直接返回开始位置的下标，否则返回-1
	ret := isSubString(A+A, B)
	if ret >= 0 {
		return true
	}

	//说明B不是A+A的子串
	return false
}

//判断str2是否在str1里面，如果在就直接返回开始位置的下标，否则返回-1
func isSubString(str1 string, str2 string) int {

}

func main() {

}
