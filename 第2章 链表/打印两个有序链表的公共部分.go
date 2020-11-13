package main

import "fmt"

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/11/9 9:09 上午
 * @Desc: 打印两个有序链表的公共部分
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func printCommonComponent(head1 *ListNode, head2 *ListNode) {
	if head1 == nil || head2 == nil {
		return
	}

	cur1, cur2 := head1, head2
	for cur1 != nil && cur2 != nil {
		if cur1.Val == cur2.Val {
			fmt.Printf("%d", cur1.Val)
			cur1, cur2 = cur1.Next, cur2.Next
		} else if cur1.Val > cur2.Val {
			cur2 = cur2.Next
		} else {
			cur1 = cur1.Next
		}
	}
}

func main() {

}
