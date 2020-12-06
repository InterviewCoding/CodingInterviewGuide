package main

import "fmt"

//递归传入切片作为参数，递归回来的时候切片不会造成任何改变
func recursion(nums []int) {
	if len(nums) <= 0 {
		return
	}

	fmt.Println(nums[0])
	nums = nums[1:]
	recursion(nums)
	fmt.Println(nums)
}

//测试是否切片切去元素也会造成地址的改变，按照自己的理解是会改变的
//测试结果：如果切片移除元素，切片对应的地址也会改变
func test(nums []int) {
	fmt.Printf("address of slice %p\n", &nums)
	if len(nums) <= 0 {
		return
	}
	fmt.Println(&nums)
	nums = nums[:len(nums)-1]
	test(nums)
}

func main() {
	nums := []int{1, 2, 3}
	test(nums)
	fmt.Println(nums)
}
